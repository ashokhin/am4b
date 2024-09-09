package com.ashokhin.am4bot.model;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public final class APIXpath {
    /*
     * Login page
     */
    // xpath for the button 'play-now' on the main page
    public static final String xpathButtonPlayNow = "//div[@class='button-container']//button[@class='play-now']";
    // xpath for the button 'login' on the main page
    public static final String xpathButtonLogin = "//div[@class='login-container']//button[@onclick=\"login('show');\"]";
    // xpath for a text field 'username' in the popup window
    public static final String xpathTextFieldUsername = "//form[@id='loginForm']//input[@id='lEmail']";
    // xpath for a text field 'password' in the popup window
    public static final String xpathTextFieldPassword = "//form[@id='loginForm']//input[@id='lPass']";
    // xpath for a check box 'remember me' in the popup window
    public static final String xpathCheckboxRememberMe = "//form[@id='loginForm']//label[@for='remember']";
    // xpath for a button 'login' in the popup window
    public static final String xpathButtonAuth = "//form[@id='loginForm']//button[@id='btnLogin']";

    /*
     * Loading overlay
     */
    public static final String xpathElementLoadingOverlay = "/html/body/div[@class='preloader exo xl-text' and contains(@style, 'AM_loading.jpg') and contains(@style, 'display: none')]";

    /*
     * Game's main page
     */
    // xpath for a text 'Account' on the main game page, which contains available
    // money
    public static final String xpathTextAccount = "//nav[@id='topMenu']//span[@id='headerAccount']";
    // xpath for the button 'Account' on the main game page
    public static final String xpathButtonAccount = "//nav[@id='topMenu']//li[contains(@onclick,'banking.php')]";
    //
    public static final String xpathElementListAccountLines = "//div[@id='popup']//div[@id='popContent']//div[@id='bankingAction']//tbody//tr[contains(@onclick, 'banking_account.php?id=')]";
    //
    public static final String xpathElementAccountLine = "./td";

    /*
     * Menu: 'Flight Info'
     */
    // xpath for the button 'Hubs' in the side menu
    public static final String xpathButtonHubs = "//div[@id='flightInfo']//button[contains(@onclick, 'hubs.php') and contains(@onclick, 'Hubs')]";
    // xpath for the lounges alert icon element
    public static final String xpathElementLoungeAlertIcon = "//div[@id='flightInfo']//span[@id='loungeAlertIcon' and @class='glyphicons glyphicons-exclamation-sign text-warning']";
    //
    public static final String xpathButtonLounges = "//div[@id='popup']//button[@id='loungeBtn']";
    //
    public static final String xpathElementListLoungeLine = "//div[@id='popup']//div[@id='popContent']//table[@class='table table-sm m-text']/tbody/tr[contains(@id, 'lList')]";
    //
    public static final String xpathElementLoungeLine = "./td";
    // xpath for the button 'landed' in the side menu
    public static final String xpathButtonLanded = "//div[@id='flightInfo']//button[@id='flightStatusLanded']";
    // xpath for a list of landed aircraft in the 'landed' list
    public static final String xpathElementListLanded = "//div[@id='landedList']/div[contains(@id, 'flightStatus') and contains(@onclick, 'showFlightInfo') and contains(@data-grounded, '0')]";
    // xpath for the button 'depart'
    public static final String xpathButtonDepart = "//div[@id='flightInfo']//button[contains(@onclick, 'route_depart.php?mode=all&ids=x')]";

    /*
     * Menu: 'Company'
     */
    // xpath for the button 'Company' in the main window
    public static final String xpathButtonCompanyMenu = "//div[@id='mapAcList' and @data-original-title='Company, staff & highscore']";
    // xpath for tab 'Stuff' in the popup window 'Company'
    public static final String xpathButtonCompanyStuffTab = "//div[@id='popup']//button[@id='popBtn2']";
    // hashmap for pilots morale elements
    private static final Map<String, String> xpathStuffPilotsMoraleElements = new HashMap<String, String>() {
        {
            // xpath for the pilots salary text
            put("xpathTextSalary", "//div[@id='staffMain']//b[@id='pilotSalary']");
            // xpath for the pilots morale text
            put("xpathTextMorale", "//div[@id='staffMain']//span[@id='pilotMorale']");
            // xpath for the pilots salary increase
            put("xpathButtonSalaryRaise",
                    "//div[@id='staffMain']//button[contains(@onclick, 'staff_action.php?type=pilot&mode=raise')]");
            // xpath for the pilots salary decrease
            put("xpathButtonSalaryPaycut",
                    "//div[@id='staffMain']//button[contains(@onclick, 'staff_action.php?type=pilot&mode=cut')]");

        }
    };
    // hashmap for crew morale elements
    private static final Map<String, String> xpathStuffCrewMoraleElements = new HashMap<String, String>() {
        {
            // xpath for the crew salary text
            put("xpathTextSalary", "//div[@id='staffMain']//b[@id='crewSalary']");
            // xpath for the crew morale text
            put("xpathTextMorale", "//div[@id='staffMain']//span[@id='crewMorale']");
            // xpath for the crew salary increase
            put("xpathButtonSalaryRaise",
                    "//div[@id='staffMain']//button[contains(@onclick, 'staff_action.php?type=crew&mode=raise')]");
            // xpath for the crew salary decrease
            put("xpathButtonSalaryPaycut",
                    "//div[@id='staffMain']//button[contains(@onclick, 'staff_action.php?type=crew&mode=cut')]");

        }
    };
    // hashmap for engineers morale elements
    private static final Map<String, String> xpathStuffEngineersMoraleElements = new HashMap<String, String>() {
        {
            // xpath for the engineers salary text
            put("xpathTextSalary", "//div[@id='staffMain']//b[@id='engineerSalary']");
            // xpath for the engineers morale text
            put("xpathTextMorale", "//div[@id='staffMain']//span[@id='engineerMorale']");
            // xpath for the engineers salary increase
            put("xpathButtonSalaryRaise",
                    "//div[@id='staffMain']//button[contains(@onclick, 'staff_action.php?type=engineer&mode=raise')]");
            // xpath for the engineers salary decrease
            put("xpathButtonSalaryPaycut",
                    "//div[@id='staffMain']//button[contains(@onclick, 'staff_action.php?type=engineer&mode=cut')]");

        }
    };
    // hashmap for technicians morale elements
    private static final Map<String, String> xpathStuffTechniciansMoraleElements = new HashMap<String, String>() {
        {
            // xpath for the technicians salary text
            put("xpathTextSalary", "//div[@id='staffMain']//b[@id='techSalary']");
            // xpath for the technicians morale text
            put("xpathTextMorale", "//div[@id='staffMain']//span[@id='techMorale']");
            // xpath for the technicians salary increase
            put("xpathButtonSalaryRaise",
                    "//div[@id='staffMain']//button[contains(@onclick, 'staff_action.php?type=tech&mode=raise')]");
            // xpath for the technicians salary decrease
            put("xpathButtonSalaryPaycut",
                    "//div[@id='staffMain']//button[contains(@onclick, 'staff_action.php?type=tech&mode=cut')]");

        }
    };
    // hashmap for joining all stuff's morale elements together
    public static final Map<String, Map<String, String>> xpathAllStuffMoraleElementsMap = new HashMap<String, Map<String, String>>() {
        {
            put("pilots", xpathStuffPilotsMoraleElements);
            put("crew", xpathStuffCrewMoraleElements);
            put("engineers", xpathStuffEngineersMoraleElements);
            put("technicians", xpathStuffTechniciansMoraleElements);
        }
    };

    // xpath for button 'close' in all popup windows
    public static final String xpathButtonPopupClose = "//div[@id='popup']//span[@onclick='closePop();']";

    /*
     * Menu: 'Finance, Marketing & Stock'
     */
    // xpath for button 'Finance, Marketing & Stock' in the main window
    public static final String xpathButtonFinanceMarketingMenu = "//div[@id='mapMaint' and @data-original-title='Finance, Marketing & Stock']";
    // xpath for tab 'Marketing' in the popup window 'finance'
    public static final String xpathButtonFinanceMarketingTab = "//div[@id='popup']//button[@id='popBtn2']";
    // xpath for button 'New campaign' in the popup window 'finance'
    public static final String xpathButtonFinanceMarketingNewCampaign = "//div[@id='financeAction']//button[@id='newCampaign']";
    // xpath for campaign row 'Airline reputation' in the popup window 'finance'
    public static final String xpathElementFinanceMarketingCompany1 = "//div[@id='campaign-1']//tr[contains(@onclick, 'marketing_new.php?type=1')]";
    // xpath for campaign duration in dropdown menu
    public static final String xpathElementFinanceMarketingCompany1Select = "//div[@id='campaign-2']//select[@id='dSelector']";
    // xpath for button 'Start campaign' in the popup window 'finance'
    public static final String xpathButtonFinanceMarketingCompany1Do = "//div[@id='campaign-2']//button[@id='c4Btn']";
    //
    public static final String xpathElementFinanceMarketingCompany2 = "//div[@id='campaign-1']//tr[contains(@onclick, 'marketing_new.php?type=5')]";
    //
    public static final String xpathButtonFinanceMarketingCompany2Do = "//div[@id='campaign-2']//button[contains(@onclick, 'marketing_new.php?type=5&mode=do&c=1')]";
    //
    public static final String xpathElementListFinanceMarketingCompanies = "//div[@id='active-campaigns']//td[@class='hasCountdown']";
    //
    public static final String xpathElementFinanceMarketingCompany3 = "//div[@id='campaign-1']//tr[contains(@onclick, 'marketing_new.php?type=2')]";
    //
    public static final String xpathElementFinanceMarketingCompany3Select = "//div[@id='campaign-2']//select[@id='dSelector']";
    //
    public static final String xpathButtonFinanceMarketingCompany3Do = "//div[@id='campaign-2']//button[contains(@onclick, 'marketing_new.php?type=2&c=4&mode=do')]";

    /*
     * Menu: 'Fuel & CO2'
     */
    // hashmap for common elements in 'Fuel & CO2' menu
    private static final Map<String, String> xpathCommonElementsMap = new HashMap<String, String>() {
        {
            // xpath for button 'Fuel' in the main window
            put("xpathButtonFuelMenu", "//div[@id='mapMaint' and @data-original-title='Fuel & co2']");
            // xpath for text field 'Amount to purchase' for fuel and CO2 in the popup
            // window 'fuel'
            put("xpathTextFieldAmount", "//div[@id='fuelMain']//input[@id='amountInput']");
            // xpath for tab 'fuel' in the popup window 'fuel'
            put("xpathButtonFuelTab", "//div[@id='popup']//button[@id='popBtn1']");
            // xpath for tab 'co2' in the popup window 'fuel'
            put("xpathButtonCO2Tab", "//div[@id='popup']//button[@id='popBtn2']");
        }
    };

    // hashmap for 'fuel' elements
    private static final Map<String, String> xpathFuelElementsMap = new HashMap<String, String>() {
        {
            // xpath for text 'current price' for fuel in the popup window 'fuel'
            put("xpathTextPrice", "//div[@id='fuelMain']/div/div[1]/span[2]/b");
            // xpath for text 'max.capacity' for fuel in the popup window 'fuel'
            put("xpathTextMaxCapacity", "//div[@id='fuelMain']//span[@class='s-text' and contains(text(), 'Lbs')]");
            // xpath for text 'current capacity' for fuel in the popup window 'fuel'
            put("xpathTextCurrentCapacity", "//div[@id='fuelMain']//span[@id='remCapacity']");
            // xpath for button 'Purchase' for fuel in the popup window 'fuel'
            put("xpathButtonPurchase", "//div[@id='fuelMain']//button[contains(@onclick, 'fuel.php?mode=do&amount=')]");
        }
    };

    // hashmap for 'CO2' elements
    private static final Map<String, String> xpathCO2ElementsMap = new HashMap<String, String>() {
        {
            // xpath for text 'current price' for CO2 in the popup window 'fuel'
            put("xpathTextPrice", "//div[@id='co2Main']/div/div[2]/span[2]/b");
            // xpath for text 'max.capacity' for CO2 in the popup window 'fuel'
            put("xpathTextMaxCapacity", "//div[@id='co2Main']//span[@class='s-text' and contains(text(), 'Quotas')]");
            // xpath for text 'current capacity' for CO2 in the popup window 'fuel'
            put("xpathTextCurrentCapacity", "//div[@id='co2Main']//span[@id='remCapacity']");
            // xpath for button 'Purchase' for CO2 in the popup window 'fuel'
            put("xpathButtonPurchase", "//div[@id='co2Main']//button[contains(@onclick, 'co2.php?mode=do&amount=')]");
        }
    };

    // hashmap for joining all 'fuel' and 'CO2' hashmaps together
    public static final Map<String, Map<String, String>> xpathAllFuelElementsMap = new HashMap<String, Map<String, String>>() {
        {
            put("common", xpathCommonElementsMap);
            put(FuelType.FUEL.getTitle(), xpathFuelElementsMap);
            put(FuelType.CO2.getTitle(), xpathCO2ElementsMap);
        }
    };

    /*
     * Menu: 'Maintenance'
     */
    // xpath for button 'Maintenance' in the main window
    public static final String xpathButtonMaintenanceMenu = "//div[@id='mapMaint' and @data-original-title='Maintenance']";
    // xpath for button 'Plan' in the popup window 'maintenance'
    public static final String xpathButtonMaintenancePlan = "//div[@id='popup']//button[@id='popBtn2']";
    //
    public static final String xpathButtonMaintenanceSortByWear = "//div[@id='maintView']//button[@onclick='sortMaint();']";
    //
    public static final String xpathButtonMaintenanceSortByACheck = "//div[@id='maintView']//button[@onclick=\"sortMaint('check');\"]";
    //
    public static final String xpathElementListMaintenanceToBase = "//div[@id='acListView']/div[@data-base='1']";
    //
    public static final String xpathButtonMaintenanceRepairPlan = ".//button[contains(@onclick, 'maint_plan_do.php?type=repair&id=')]";
    //
    public static final String xpathTextMaintenanceRepairPrice = "//div[@id='typeRepair']//div[contains(text(), '$ ')]";
    //
    public static final String xpathButtonMaintenanceRepairDo = "//div[@id='typeRepair']//button[contains(@onclick, 'maint_plan_do.php?mode=do&type=repair&id=')]";
    //
    public static final String xpathButtonMaintenanceACheckPlan = ".//button[contains(@onclick, 'maint_plan_do.php?type=check&id=')]";
    //
    public static final String xpathTextMaintenanceACheckPrice = "//div[@id='typeCheck']//div[contains(text(), '$ ')]";
    //
    public static final String xpathButtonMaintenanceACheckDo = "//div[@id='typeCheck']//button[contains(@onclick, 'maint_plan_do.php?mode=do&type=check&id=')]";
    //
    public static final String xpathButtonMaintenanceModifyPlan = ".//button[contains(@onclick, 'maint_plan_do.php?type=modify&id=')]";
    //
    public static final String xpathElementListMaintenanceModifyCheckbox = "//div[@id='typeModify']//label[@class='check-container']";
    //
    private static final String xpathCheckboxMaintenanceModifyReduceCo2 = ".//input[@id='mod1']";
    //
    private static final String xpathCheckboxMaintenanceModifySpeedIncrease = ".//input[@id='mod2']";
    //
    private static final String xpathCheckboxMaintenanceModifyReduceFuel = ".//input[@id='mod3']";
    //
    public static final List<String> xpathCheckboxMaintenanceModifyList = new ArrayList<String>() {
        {
            add(xpathCheckboxMaintenanceModifyReduceCo2);
            add(xpathCheckboxMaintenanceModifySpeedIncrease);
            add(xpathCheckboxMaintenanceModifyReduceFuel);
        }
    };
    //
    public static final String xpathTextMaintenanceModifyPrice = "//div[@id='typeModify']//span[@id='acCost']";
    //
    public static final String xpathTextMaintenanceModifyCargoPrice = "//div[@id='typeModify']//span[@id='totalCost']";
    //
    public static final String xpathButtonMaintenanceModifyDo = "//div[@id='typeModify']//button[contains(@onclick, 'modifyAction')]";

    /*
     * Menu: 'Fleet & routes'
     */
    // xpath for button 'Fleet & routes' in the main window
    public static final String xpathButtonFleetAndRoutesMenu = "//div[@id='mapRoutes' and @data-original-title='Fleet & routes']";
    // xpath for button '+ Order' in the popup window 'Fleet & routes'
    public static final String xpathButtonFleetAndRoutes_order = "//div[@id='popup']//button[@id='popBtn2']";
    //
    public static final String xpathButtonFleetAndRoutes_list_sort = "//div[@id='routeAction']//button[@id='listSort']";

    /*
     * Pop-up window 'Overview'
     */
    // xpath for button 'Overview' in the main window
    public static final String xpathButtonOverviewMenu = "//div[@id='flightInfo']//div[@id='flightInfoSecContainer']//button[contains(@onclick, 'overview.php')]";
    // xpath for text element 'Airline Reputation'
    public static final String xpathTextOverviewAirlineReputation = "//div[@id='popup']//div[@id='popContent']//div[contains(text(),'Airline Reputation')]";
    // xpath for text element 'Cargo reputation'
    public static final String xpathTextOverviewCargoReputation = "//div[@id='popup']//div[@id='popContent']//div[contains(text(),'Cargo reputation')]";
    // xpath for text element 'Fuel cost'
    public static final String xpathTextOverviewFuelCost = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[1]/td[2]/span[1]";
    // xpath for text element 'Co2 quota cost'
    public static final String xpathTextOverviewCo2Cost = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[2]/td[2]/span[1]";
    // xpath for text element 'Fleet size'
    public static final String xpathTextOverviewFleetSize = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[3]/td[2]";
    // xpath for text element 'Pending delivery'
    public static final String xpathTextOverviewPendingDelivery = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[4]/td[2]";
    // xpath for text element 'Routes'
    public static final String xpathTextOverviewRoutes = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[5]/td[2]";
    // xpath for text element 'Hubs'
    public static final String xpathTextOverviewHubs = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[6]/td[2]";
    // xpath for text element 'Pending maintenance'
    public static final String xpathTextOverviewPendingMaintenance = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[7]/td[2]";
    // xpath for text element 'Fuel holding'
    public static final String xpathTextOverviewFuelHolding = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[8]/td[2]";
    // xpath for text element 'Hangar capacity'
    public static final String xpathTextOverviewHangarCapacity = "//div[@id='popup']//div[@id='popContent']/div[1]/div[6]/table/tbody/tr[9]/td[2]";
    // xpath for text element 'Co2 quotas'
    public static final String xpathTextOverviewCo2Quotas = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[1]/td[2]";
    // xpath for text element 'A/C Inflight'
    public static final String xpathTextOverviewACInflight = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[2]/td[2]";
    // xpath for text element 'Share value'
    public static final String xpathTextOverviewShareValue = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[3]/td[2]/div";
    // xpath for text element 'Flights operated'
    public static final String xpathTextOverviewFlightsOperated = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[4]/td[2]";
    // xpath for text element 'Y class pax'
    public static final String xpathTextOverviewYClassPax = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[5]/td[2]";
    // xpath for text element 'J class pax'
    public static final String xpathTextOverviewJClassPax = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[6]/td[2]";
    // xpath for text element 'F class pax'
    public static final String xpathTextOverviewFClassPax = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[7]/td[2]";
    // xpath for text element 'Large load'
    public static final String xpathTextOverviewLargeLoad = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[8]/td[2]";
    // xpath for text element 'Heavy load'
    public static final String xpathTextOverviewHeavyLoad = "//div[@id='popup']//div[@id='popContent']/div[1]/div[7]/table/tbody/tr[9]/td[2]";
}
