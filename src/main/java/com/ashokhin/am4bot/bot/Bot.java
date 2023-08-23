package com.ashokhin.am4bot.bot;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.openqa.selenium.By;
import org.openqa.selenium.WebElement;

import com.ashokhin.am4bot.model.APIXpath;
import com.ashokhin.am4bot.model.Aircraft;
import com.ashokhin.am4bot.model.AirplaneFuel;
import com.ashokhin.am4bot.model.FuelType;
import com.ashokhin.am4bot.model.Maintenance;
import com.ashokhin.am4bot.model.MaintenanceOperation;
import com.ashokhin.am4bot.model.Marketing;
import com.ashokhin.am4bot.model.MarketingCompany;
import com.ashokhin.am4bot.model.MarketingCompanyType;
import com.ashokhin.am4bot.utils.AircraftSortingComparator;

public final class Bot extends BotBase {
    private static final Logger logger = LogManager.getLogger(Bot.class);
    private static int accountMoney;
    private BotMode botMode;
    private int fuelBudgetPercent;
    private int maintenanceBudgetPercent;
    private int marketingBudgetPercent;
    private int wearPercent;
    private int maximumHoursBeforeACheck;
    private Map<FuelType, Integer> fuelPricesMap;
    private Map<FuelType, AirplaneFuel> fuelDataMap;
    private ArrayList<MarketingCompany> marketingDataList = new ArrayList<MarketingCompany>();

    public Bot(String baseUrl, String login, String password) {
        super(baseUrl, login, password);
        fuelDataMap = new HashMap<FuelType, AirplaneFuel>();
    }

    public Bot(
            String baseUrl,
            String login,
            String password,
            BotMode botMode,
            int fuelGoodPrice,
            int co2GoodPrice,
            int criticalFuelLevelPercent,
            int fuelBudgetPercent,
            int maintenanceBudgetPercent,
            int marketingBudgetPercent,
            int aircraftWearPercent,
            int aircraftMaximumHoursBeforeACheck) {
        super(baseUrl, login, password);
        this.fuelPricesMap = new HashMap<FuelType, Integer>() {
            {
                put(FuelType.FUEL, fuelGoodPrice);
                put(FuelType.CO2, co2GoodPrice);
            }
        };
        AirplaneFuel.setCriticalLevelPercent(criticalFuelLevelPercent);
        this.fuelBudgetPercent = fuelBudgetPercent;
        this.maintenanceBudgetPercent = maintenanceBudgetPercent;
        this.marketingBudgetPercent = marketingBudgetPercent;
        this.wearPercent = aircraftWearPercent;
        this.maximumHoursBeforeACheck = aircraftMaximumHoursBeforeACheck;
        this.botMode = botMode;
        fuelDataMap = new HashMap<FuelType, AirplaneFuel>();
    }

    public final void setSettings(
            int fuelGoodPrice,
            int co2GoodPrice,
            int criticalFuelLevelPercent,
            int fuelBudgetPercent,
            int maintenanceBudgetPercent,
            int marketingBudgetPercent,
            int aircraftWearPercent,
            int aircraftMaximumHoursBeforeACheck) {
        this.fuelPricesMap = new HashMap<FuelType, Integer>() {
            {
                put(FuelType.FUEL, fuelGoodPrice);
                put(FuelType.CO2, co2GoodPrice);
            }
        };
        AirplaneFuel.setCriticalLevelPercent(criticalFuelLevelPercent);
        this.fuelBudgetPercent = fuelBudgetPercent;
        this.maintenanceBudgetPercent = maintenanceBudgetPercent;
        this.marketingBudgetPercent = marketingBudgetPercent;
        this.wearPercent = aircraftWearPercent;
        this.maximumHoursBeforeACheck = aircraftMaximumHoursBeforeACheck;
    }

    @Override
    public void run() {
        switch (this.botMode) {
            case ALL:
                this.startOnce();
                this.quit();
                break;
            case BUY_FUEL:
                super.startBot();
                this.buyFuel();
                this.quit();
                break;
            case DEPART:
                super.startBot();
                this.departAllAircraft();
                this.quit();
                break;
            case MAINTENANCE:
                super.startBot();
                this.maintenanceAircraft();
                this.quit();
                break;
            case MARKETING:
                super.startBot();
                this.startMarketingCompanies();
                this.quit();
                break;
        }
    }

    private final synchronized void checkMoney() {
        logger.debug("Check money");

        this.refreshPage();

        Bot.accountMoney = this.getIntFromElement(APIXpath.xpathTextAccount);

        logger.debug(String.format("Account money $%d", Bot.accountMoney));
    }

    private static final synchronized void decreaseMoney(int moneySpent) {
        Bot.accountMoney -= moneySpent;
    }

    private final void checkFuelType(FuelType fuelType) {
        logger.info(String.format("Check '%s' price and capacity", fuelType.getTitle()));

        switch (fuelType) {
            case FUEL:
                this.clickButton(APIXpath.xpathAllFuelElementsMap.get("common").get("xpathButtonFuelTab"));
                break;

            case CO2:
                this.clickButton(APIXpath.xpathAllFuelElementsMap.get("common").get("xpathButtonCO2Tab"));
                break;
        }

        int fuelPrice = this.getIntFromElement(APIXpath.xpathAllFuelElementsMap
                .get(fuelType.getTitle()).get("xpathTextPrice"));

        logger.debug(String.format("Price for '%s' is $%d", fuelType.getTitle(), fuelPrice));

        int fuelCurrentCapacity = this.getIntFromElement(APIXpath.xpathAllFuelElementsMap
                .get(fuelType.getTitle()).get("xpathTextCurrentCapacity"));

        logger.debug(
                String.format("Current capacity for '%s' is %d %s", fuelType.getTitle(), fuelCurrentCapacity,
                        fuelType.getUnit()));

        int fuelMaxCapacity = this.getIntFromElement(APIXpath.xpathAllFuelElementsMap
                .get(fuelType.getTitle()).get("xpathTextMaxCapacity"));

        logger.debug(String.format("Maximum capacity for '%s' is %d %s", fuelType.getTitle(), fuelMaxCapacity,
                fuelType.getUnit()));

        if (this.fuelDataMap.containsKey(fuelType)) {
            this.fuelDataMap.get(fuelType).update(fuelPrice, fuelCurrentCapacity);
        } else {
            this.fuelDataMap.put(fuelType, new AirplaneFuel(fuelType, fuelPrice, this.fuelPricesMap.get(fuelType),
                    this.fuelBudgetPercent, fuelCurrentCapacity, fuelMaxCapacity));
        }

        logger.info(
                String.format("'%s' price: $%d, holding capacity: %d / %d %s", fuelType.getTitle(),
                        this.fuelDataMap.get(fuelType).getPrice(), this.fuelDataMap.get(fuelType).getHoldingCapacity(),
                        this.fuelDataMap.get(fuelType).getMaximumCapacity(), fuelType.getUnit()));
    }

    private final void buyFuelAmount(AirplaneFuel airplaneFuel) {
        int needFuelAmount = airplaneFuel.getNeedAmount(Bot.accountMoney);

        if (needFuelAmount == 0) {

            return;
        }

        logger.debug(String.format("%s buy: %d", airplaneFuel.getType(), needFuelAmount));

        this.typeTextInField(APIXpath.xpathAllFuelElementsMap.get("common").get("xpathTextFieldAmount"),
                String.valueOf(needFuelAmount));

        this.clickButton(
                APIXpath.xpathAllFuelElementsMap.get(airplaneFuel.getType()).get("xpathButtonPurchase"));
        airplaneFuel.buyFuelAmount(needFuelAmount);
    }

    private final void buyFuelType(FuelType fuelType) {
        this.checkFuelType(fuelType);
        AirplaneFuel currentFuel = this.fuelDataMap.get(fuelType);

        if (currentFuel.isFull()) {
            logger.info(String.format("We already have enough %s", fuelType.getTitle()));

            return;
        }

        if (currentFuel.notEnoughFuel()) {
            logger.warn(String.format("We haven't enough %s", fuelType.getTitle()));
        }

        logger.debug(String.format("Fuel data info: %s", currentFuel));

        this.buyFuelAmount(currentFuel);
    }

    private final void buyFuel() {
        this.checkMoney();
        logger.info("Buy fuel");

        this.clickButton(APIXpath.xpathAllFuelElementsMap.get("common").get("xpathButtonFuelMenu"));

        for (FuelType fuelType : FuelType.values()) {
            this.buyFuelType(fuelType);
        }

        this.clickButton(APIXpath.xpathButtonPopupClose);
    }

    private final List<Aircraft> findAllAircraftForMaintenance(MaintenanceOperation maintenanceOperation) {
        List<Aircraft> aircraftForMaintenance = new ArrayList<Aircraft>();

        this.clickButton(APIXpath.xpathButtonMaintenancePlan);

        switch (maintenanceOperation) {
            case A_CHECK:
                this.clickButton(APIXpath.xpathButtonMaintenanceSortByACheck);
                break;
            case REPAIR:
                this.clickButton(APIXpath.xpathButtonMaintenanceSortByWear);
                break;
            case MODIFY:
                break;
            default:
                break;
        }

        for (WebElement aircraftWebElement : this.getElements(APIXpath.xpathElementListMaintenanceToBase)) {
            aircraftForMaintenance.add(new Aircraft(
                    aircraftWebElement.getAttribute(Aircraft.TYPE),
                    aircraftWebElement.getAttribute(Aircraft.REG_NUMBER),
                    Integer.valueOf(aircraftWebElement.getAttribute(Aircraft.A_CHECK)),
                    Float.valueOf(aircraftWebElement.getAttribute(Aircraft.WEAR)).intValue()));
        }

        return aircraftForMaintenance;
    }

    private final WebElement findChildButton(
            MaintenanceOperation maintenanceOperation,
            Aircraft aircraftForMaintenance) {

        logger.debug(String.format("Find child button for '%s' operation for '%s'", maintenanceOperation.getTitle(),
                aircraftForMaintenance.getRegNumber()));

        String xpathButtonSort = null;
        String xpathButtonForSearchChild = null;

        switch (maintenanceOperation) {
            case A_CHECK:
                xpathButtonSort = APIXpath.xpathButtonMaintenanceSortByACheck;
                xpathButtonForSearchChild = APIXpath.xpathButtonMaintenanceACheckPlan;
                break;
            case REPAIR:
                xpathButtonSort = APIXpath.xpathButtonMaintenanceSortByWear;
                xpathButtonForSearchChild = APIXpath.xpathButtonMaintenanceRepairPlan;
                break;

            case MODIFY:
                xpathButtonForSearchChild = APIXpath.xpathButtonMaintenanceModifyPlan;
                break;

            default:
                break;
        }

        this.clickButton(APIXpath.xpathButtonMaintenancePlan);
        if (xpathButtonSort != null) {
            this.clickButton(xpathButtonSort);
        }

        WebElement childButton = null;

        for (WebElement aircraftFromList : this.getElements(APIXpath.xpathElementListMaintenanceToBase)) {
            if (aircraftFromList.getAttribute(Aircraft.REG_NUMBER).equals(aircraftForMaintenance.getRegNumber())) {
                // Find child button
                childButton = aircraftFromList.findElement(By.xpath(xpathButtonForSearchChild));
                return childButton;
            }
        }

        return childButton;
    }

    private final boolean clickMaintenanceButton(WebElement maintenanceButton) {
        if (maintenanceButton == null) {
            logger.warn("Child maintenance button not found");

            return false;
        }

        try {
            this.clickButton(maintenanceButton);
            return true;
        } catch (Exception e) {
            e.printStackTrace();
            return false;
        }

    }

    private final boolean aCheckAircraft(Aircraft aircraftForACheck) {
        logger.debug(String.format("Try to A-Check '%s'", aircraftForACheck));

        if (!this.clickMaintenanceButton(
                this.findChildButton(MaintenanceOperation.A_CHECK, aircraftForACheck))) {
            logger.warn(String.format("Button for aircraft '%s' not found", aircraftForACheck.getRegNumber()));

            return false;
        }

        int aCheckPrice = this.getIntFromElement(APIXpath.xpathTextMaintenanceACheckPrice);
        int availableMoney = (int) Math.round((Bot.accountMoney * (this.maintenanceBudgetPercent * 0.01)));

        if (aCheckPrice > availableMoney) {
            logger.warn(
                    String.format("A-Check is too expensive. A-Check cost: $%d, available money for maintenance: $%d",
                            aCheckPrice, availableMoney));

            return false;
        }

        this.clickButton(APIXpath.xpathButtonMaintenanceACheckDo);

        logger.info(String.format("Aircraft '%s' planed for A-Check for $%d", aircraftForACheck.getRegNumber(),
                aCheckPrice));

        Bot.decreaseMoney(aCheckPrice);

        return true;
    }

    private final void aCheckAllAircraft() {
        logger.info("Search aircraft which need A-Check");

        List<Aircraft> aircraftNeedACheck = new ArrayList<Aircraft>();

        for (Aircraft aircraftForMaintenance : this.findAllAircraftForMaintenance(MaintenanceOperation.A_CHECK)) {
            if (aircraftForMaintenance.getACheckHours() < this.maximumHoursBeforeACheck) {
                aircraftNeedACheck.add(aircraftForMaintenance);
            }
        }

        if (aircraftNeedACheck.isEmpty()) {
            logger.info("No aircraft need A-Check");

            return;
        }

        int aCheckedAircraftCount = 0;
        for (Aircraft aircraftForACheck : aircraftNeedACheck) {
            if (this.aCheckAircraft(aircraftForACheck)) {
                aCheckedAircraftCount++;
            }
        }

        logger.info(String.format("Aircraft planed for A-Check: %d", aCheckedAircraftCount));
    }

    private final boolean repairAircraft(Aircraft aircraftForRepair) {
        logger.debug(String.format("Try to repair '%s'", aircraftForRepair));

        if (!this.clickMaintenanceButton(
                this.findChildButton(MaintenanceOperation.REPAIR, aircraftForRepair))) {
            logger.warn(String.format("Button for aircraft '%s' not found", aircraftForRepair.getRegNumber()));

            return false;
        }

        int repairPrice = this.getIntFromElement(APIXpath.xpathTextMaintenanceRepairPrice);
        int availableMoney = (int) Math.round((Bot.accountMoney * (this.maintenanceBudgetPercent * 0.01)));

        if (repairPrice > availableMoney) {
            logger.warn(
                    String.format("Repair is too expensive. Repair cost: $%d, available money for maintenance: $%d",
                            repairPrice, availableMoney));

            return false;
        }

        this.clickButton(APIXpath.xpathButtonMaintenanceRepairDo);

        logger.info(String.format("Aircraft '%s' planed to repair for $%d", aircraftForRepair, repairPrice));

        Bot.decreaseMoney(repairPrice);

        return true;
    }

    private final void repairAllAircraft() {
        logger.info("Search aircraft which need repair");

        List<Aircraft> aircraftNeedRepair = new ArrayList<Aircraft>();
        for (Aircraft aircraftForMaintenance : this.findAllAircraftForMaintenance(MaintenanceOperation.REPAIR)) {
            if (aircraftForMaintenance.getWearPercent() >= this.wearPercent) {
                aircraftNeedRepair.add(aircraftForMaintenance);
            }
        }

        if (aircraftNeedRepair.isEmpty()) {
            logger.info("No aircraft need repair");

            return;
        }

        int repairedAircraftCount = 0;
        for (Aircraft aircraftForRepair : aircraftNeedRepair) {
            if (this.repairAircraft(aircraftForRepair)) {
                repairedAircraftCount++;
            }
        }

        logger.info(String.format("Aircraft planed for repair: %d", repairedAircraftCount));
    }

    private final boolean modifyAircraft(Aircraft aircraftForModify) {
        logger.debug(String.format("Try to modify '%s'", aircraftForModify));

        if (!this.clickMaintenanceButton(
                this.findChildButton(MaintenanceOperation.MODIFY, aircraftForModify))) {
            logger.warn(String.format("Button for aircraft '%s' not found", aircraftForModify.getRegNumber()));

            return false;
        }

        for (WebElement modifyCheckboxRow : this.getElements(APIXpath.xpathElementListMaintenanceModifyCheckbox)) {
            for (String modifyCheckboxXPath : APIXpath.xpathCheckboxMaintenanceModifyList) {
                WebElement checkboxWebElem = null;
                try {
                    checkboxWebElem = modifyCheckboxRow.findElement(By.xpath(modifyCheckboxXPath));
                } catch (org.openqa.selenium.NoSuchElementException e) {
                    // go to try find next element
                    continue;
                }

                if (checkboxWebElem.getAttribute("checked") != null) {
                    // checkbox already checked
                    logger.trace(String.format("The checkbox '%s' is already checked", checkboxWebElem));

                    break;
                }

                WebElement checkboxCheckmarkWebElem = modifyCheckboxRow
                        .findElement(By.xpath(".//span[@class=\"checkmark\"]"));

                if (checkboxCheckmarkWebElem != null) {
                    this.clickButton(checkboxCheckmarkWebElem);
                    break;
                }
            }
        }

        int modifyPrice = this.getIntFromElement(APIXpath.xpathTextMaintenanceModifyPrice);

        if (modifyPrice == 0) {
            logger.debug(String.format("The aircraft '%s' already fully modified", aircraftForModify));

            return false;
        }
        int availableMoney = (int) Math.round((Bot.accountMoney * (this.maintenanceBudgetPercent * 0.01)));

        if (modifyPrice > availableMoney) {
            logger.warn(
                    String.format(
                            "Modification is too expensive. Modification cost: $%d, available money for maintenance: $%d",
                            modifyPrice, availableMoney));

            return false;
        }

        this.clickButton(APIXpath.xpathButtonMaintenanceModifyDo);

        logger.info(String.format("Aircraft '%s' planed for modification for $%d", aircraftForModify, modifyPrice));

        Bot.decreaseMoney(modifyPrice);

        return true;
    }

    private final void modifyAllAircraft() {
        logger.info("Search aircraft which need modification");

        List<Aircraft> aircraftNeedModify = this.findAllAircraftForMaintenance(MaintenanceOperation.MODIFY);
        // Sort Aircraft list by 'aircraftRegNumber'
        Collections.sort(aircraftNeedModify, new AircraftSortingComparator());
        // Get only last N (Maintenance.MODIFY_AIRCRAFT_NUMBER) aircraft for modify
        logger.debug(String.format("Check only last %d aircraft", Maintenance.MODIFY_AIRCRAFT_NUMBER));

        List<Aircraft> aircraftListForModify = aircraftNeedModify.subList(
                Math.max((aircraftNeedModify.size() - Maintenance.MODIFY_AIRCRAFT_NUMBER), 0),
                aircraftNeedModify.size());

        if (aircraftListForModify.isEmpty()) {
            logger.info(String.format("No aircraft for modification"));

            return;
        }

        int modifiedAircraftCount = 0;
        for (Aircraft aircraftForModify : aircraftListForModify) {

            if (this.modifyAircraft(aircraftForModify)) {
                modifiedAircraftCount++;
            }
        }

        if (modifiedAircraftCount > 0) {
            logger.info(String.format("Aircraft planed for modification: %d", modifiedAircraftCount));
        } else {
            logger.info(String.format("No aircraft modified"));
        }
    }

    private final void maintenanceAircraft() {
        this.checkMoney();
        logger.info("Maintenance aircraft");
        this.clickButton(APIXpath.xpathButtonMaintenanceMenu);

        this.aCheckAllAircraft();
        this.repairAllAircraft();
        this.modifyAllAircraft();

        this.clickButton(APIXpath.xpathButtonPopupClose);
    }

    private final void startMarketingCompanies() {
        logger.debug("Try to start marketing companies");

        for (FuelType fuelType : FuelType.values()) {
            if (this.fuelDataMap.get(fuelType).notEnoughFuel()) {
                logger.warn(String.format("Not enough %s (%d / %d). Skip marketing companies.",
                        fuelType.getTitle(), this.fuelDataMap.get(fuelType).getHoldingCapacity(),
                        this.fuelDataMap.get(fuelType).getMaximumCapacity()));

                return;
            }
        }

        this.clickButton(APIXpath.xpathButtonFinanceMarketingMenu);
        this.clickButton(APIXpath.xpathButtonFinanceMarketingTab);
        logger.info("Search marketing companies for enabling...");

        this.checkMarketingCompanies();

        int activeMarketingCompaniesCount = 0;

        for (MarketingCompany marketingCompany : marketingDataList) {
            if (marketingCompany.isActive()) {
                activeMarketingCompaniesCount++;
                continue;
            } else {
                if (this.enableMarketingCompany(marketingCompany)) {
                    activeMarketingCompaniesCount++;
                    this.clickButton(APIXpath.xpathButtonFinanceMarketingNewCampaign);
                }
            }
        }

        switch (activeMarketingCompaniesCount) {
            case 0:
                logger.info("No active marketing companies");
                break;
            default:
                logger.info(String.format("%d marketing companies active", activeMarketingCompaniesCount));
                break;
        }

        this.clickButton(APIXpath.xpathButtonPopupClose);
    }

    private void checkMarketingCompanies() {
        logger.info("Check marketing companies");

        this.clickButton(APIXpath.xpathButtonFinanceMarketingNewCampaign);

        for (MarketingCompanyType marketingCompanyType : MarketingCompanyType.values()) {
            logger.debug(String.format("Try to check company '%s'", marketingCompanyType));
            this.marketingDataList.add(
                    new MarketingCompany(
                            marketingCompanyType,
                            this.getAttribute(marketingCompanyType.getRowXpath(), "class")
                                    .equals("not-active")));
        }
    }

    private final boolean enableMarketingCompany(MarketingCompany marketingCompany) {
        logger.debug("Try to activate marketing company '%s'", marketingCompany.getName());

        this.clickButton(marketingCompany.getRowXpath());

        String marketingCompanyCostFullXpath = "";

        if (marketingCompany.getType() == MarketingCompanyType.AIRLINE_REPUTATION) {
            this.selectFromDropdown(APIXpath.xpathElementFinanceMarketingCompany1Select,
                    Marketing.MARKETING_COMPANY_REPUTATION_DURATION);

            marketingCompanyCostFullXpath = String.format("%s//span[@id='c4']",
                    marketingCompany.getButtonXpath());
        }

        if (marketingCompany.getType() == MarketingCompanyType.ECO_FRIENDLY) {
            marketingCompanyCostFullXpath = String.format("%s",
                    marketingCompany.getButtonXpath());
        }

        int marketingCompanyPrice = this.getIntFromElement(marketingCompanyCostFullXpath);
        int availableMoney = (int) Math.round((Bot.accountMoney * (this.marketingBudgetPercent * 0.01)));

        if (marketingCompanyPrice > availableMoney) {
            logger.warn(String.format(
                    "Not enough money for marketing company '%s'. Available money for marketing company: $%d. Marketing company price: $%s",
                    marketingCompany.getName(),
                    availableMoney,
                    marketingCompanyPrice));

            return false;
        }

        logger.info(String.format("Activate marketing company '%s' for $%d",
                marketingCompany.getName(),
                marketingCompanyPrice));

        this.clickButton(marketingCompany.getButtonXpath());

        return true;
    }

    private final int getReadyForDepartCount() {
        this.clickButton(APIXpath.xpathButtonLanded);

        return this.getElements(APIXpath.xpathElementListLanded).size();
    }

    private final void departAllAircraft() {
        logger.info("Depart all available aircraft...");
        int readyForDepartCount = this.getReadyForDepartCount();

        if (readyForDepartCount == 0) {
            logger.info("No aircraft ready for depart");

            return;
        }

        logger.info(String.format("Aircraft ready for depart: %d", readyForDepartCount));

        int aircraftDeparted = this.getReadyForDepartCount();

        while (readyForDepartCount > 0) {
            logger.debug("Depart available aircraft");
            this.clickButton(APIXpath.xpathButtonDepart);
            readyForDepartCount = this.getReadyForDepartCount();
            // Buy fuel after each depart
            this.buyFuel();
        }

        logger.info(String.format("Aircraft departed: %d", aircraftDeparted));
    }

    public final void startOnce() {
        logger.info("Start Bot");
        super.startBot();
        this.buyFuel();
        this.startMarketingCompanies();
        this.maintenanceAircraft();
        this.departAllAircraft();
    }

    @Override
    public final void quit() {
        super.quit();
    }
}
