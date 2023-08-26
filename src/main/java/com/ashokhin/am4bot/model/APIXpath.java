package com.ashokhin.am4bot.model;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public final class APIXpath {
    // Login page
    //
    // xpath for button 'login' on main page
    public static final String xpathButtonLogin = "//div[@class='login-frontpage']//button[contains(@onclick, '#login')]";
    // xpath for text field 'username' in popup window
    public static final String xpathTextFieldUsername = "//form[@id='loginForm']//input[@id='lEmail']";
    // xpath for text field 'password' in popup window
    public static final String xpathTextFieldPassword = "//form[@id='loginForm']//input[@id='lPass']";
    // xpath for check box 'remember me' in popup window
    public static final String xpathCheckboxRememberMe = "//form[@id='loginForm']//input[@id='remember']";
    // xpath for button 'login' in popup window
    public static final String xpathButtonAuth = "//form[@id='loginForm']//button[@id='btnLogin']";
    // Loading overlay
    //
    //
    public static final String xpathElementLoadingOverlay = "/html/body/div[@class='preloader exo xl-text' and contains(@style, 'AM_loading.jpg') and contains(@style, 'display: none')]";
    // Main game page
    //
    // xpath for text 'Account' on main game page, which contains available money
    public static final String xpathTextAccount = "//nav[@id='topMenu']//span[@id='headerAccount']";
    // 'Flight Info' menu
    //
    // xpath for button 'landed' in side menu
    public static final String xpathButtonLanded = "//div[@id='flightInfo']//button[@id='flightStatusLanded']";
    // xpath for list of landed aircraft in 'landed' list
    public static final String xpathElementListLanded = "//div[@id='landedList']/div[contains(@id, 'flightStatus') and contains(@onclick, 'showFlightInfo') and contains(@data-grounded, '0')]";
    // xpath for 'depart' button
    public static final String xpathButtonDepart = "//div[@id='flightInfo']//button[contains(@onclick, 'route_depart.php?mode=all&ids=x')]";
    // 'Finance, Marketing & Stock' menu
    //
    // xpath for button 'Finance, Marketing & Stock' in main window
    public static final String xpathButtonFinanceMarketingMenu = "//div[@id='mapMaint' and @data-original-title='Finance, Marketing & Stock']";
    // xpath for tab 'Marketing' in popup window 'finance'
    public static final String xpathButtonFinanceMarketingTab = "//div[@id='popup']//button[@id='popBtn2']";
    // xpath for button 'New campaign' in popup window 'finance'
    public static final String xpathButtonFinanceMarketingNewCampaign = "//div[@id='financeAction']//button[@id='newCampaign']";
    // xpath for campaign row 'Airline reputation' in popup window 'finance'
    public static final String xpathElementFinanceMarketingCompany1 = "//div[@id='campaign-1']//tr[contains(@onclick, 'marketing_new.php?type=1')]";
    //
    public static final String xpathElementFinanceMarketingCompany1Select = "//div[@id='campaign-2']//select[@id='dSelector']";
    // xpath for button 'Start campaign' in popup window 'finance'
    public static final String xpathButtonFinanceMarketingCompany1Do = "//div[@id='campaign-2']//button[@id='c4Btn']";
    //
    public static final String xpathElementFinanceMarketingCompany2 = "//div[@id='campaign-1']//tr[contains(@onclick, 'marketing_new.php?type=5')]";
    //
    public static final String xpathButtonFinanceMarketingCompany2Do = "//div[@id='campaign-2']//button[contains(@onclick, 'marketing_new.php?type=5&mode=do&c=1')]";
    //
    public static final String xpathElementListFinanceMarketingCompanies = "//div[@id='active-campaigns']//td[@class='hasCountdown']";
    // 'Fuel & CO2' popup
    //
    //
    private static final Map<String, String> xpathCommonElementsMap = new HashMap<String, String>() {
        {
            // xpath for button 'Fuel' in main window
            put("xpathButtonFuelMenu", "//div[@id='mapMaint' and @data-original-title='Fuel & co2']");
            // xpath for text field 'Amount to purchase' for fuel and CO2 in popup window
            // 'fuel'
            put("xpathTextFieldAmount", "//div[@id='fuelMain']//input[@id='amountInput']");
            // xpath for tab 'fuel' in popup window 'fuel'
            put("xpathButtonFuelTab", "//div[@id='popup']//button[@id='popBtn1']");
            // xpath for tab 'co2' in popup window 'fuel'
            put("xpathButtonCO2Tab", "//div[@id='popup']//button[@id='popBtn2']");
        }
    };
    private static final Map<String, String> xpathFuelElementsMap = new HashMap<String, String>() {
        {
            // xpath for text 'current price' for fuel in popup window 'fuel'
            put("xpathTextPrice", "//div[@id='fuelMain']/div/div[1]/span[2]/b");
            // xpath for text 'max.capacity' for fuel in popup window 'fuel'
            put("xpathTextMaxCapacity", "//div[@id='fuelMain']//span[@class='s-text' and contains(text(), 'Lbs')]");
            // xpath for text 'current capacity' for fuel in popup window 'fuel'
            put("xpathTextCurrentCapacity", "//div[@id='fuelMain']//span[@id='remCapacity']");
            // xpath for button 'Purchase' for fuel in popup window 'fuel'
            put("xpathButtonPurchase", "//div[@id='fuelMain']//button[contains(@onclick, 'fuel.php?mode=do&amount=')]");
        }
    };
    private static final Map<String, String> xpathCO2ElementsMap = new HashMap<String, String>() {
        {
            // xpath for text 'current price' for CO2 in popup window 'fuel'
            put("xpathTextPrice", "//div[@id='co2Main']/div/div[2]/span[2]/b");
            // xpath for text 'max.capacity' for CO2 in popup window 'fuel'
            put("xpathTextMaxCapacity", "//div[@id='co2Main']//span[@class='s-text' and contains(text(), 'Quotas')]");
            // xpath for text 'current capacity' for CO2 in popup window 'fuel'
            put("xpathTextCurrentCapacity", "//div[@id='co2Main']//span[@id='remCapacity']");
            // xpath for button 'Purchase' for CO2 in popup window 'fuel'
            put("xpathButtonPurchase", "//div[@id='co2Main']//button[contains(@onclick, 'co2.php?mode=do&amount=')]");
        }
    };

    public static final Map<String, Map<String, String>> xpathAllFuelElementsMap = new HashMap<String, Map<String, String>>() {
        {
            put("common", xpathCommonElementsMap);
            put(FuelType.FUEL.getTitle(), xpathFuelElementsMap);
            put(FuelType.CO2.getTitle(), xpathCO2ElementsMap);
        }
    };

    // xpath for button 'close' in all popup windows
    public static final String xpathButtonPopupClose = "//div[@id='popup']//span[@onclick='closePop();']";
    // 'Maintenance' popup
    //
    // xpath for button 'Maintenance' in main window
    public static final String xpathButtonMaintenanceMenu = "//div[@id='mapMaint' and @data-original-title='Maintenance']";
    // xpath for button 'Plan' in popup window 'maintenance'
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
    public static final String xpathButtonMaintenanceACheckPlan = "//div[@id='acListView']//button[contains(@onclick, 'maint_plan_do.php?type=check&id=')]";
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
    public static final String xpathButtonMaintenanceModifyDo = "//div[@id='typeModify']//button[contains(@onclick, 'modifyAction')]";
    // 'Fleet & routes' popup
    //
    // xpath for button 'Fleet & routes' in main window
    public static final String xpathButtonFleetAndRoutesMenu = "//div[@id='mapRoutes' and @data-original-title='Fleet & routes']";
    // xpath for button '+ Order' in popup window 'Fleet & routes'
    public static final String xpathButtonFleetAndRoutes_order = "//div[@id='popup']//button[@id='popBtn2']";
    //
    public static final String xpathButtonFleetAndRoutes_list_sort = "//div[@id='routeAction']//button[@id='listSort']";

    /**
     * public static final String xelem_list_fl_ac =
     * "//div[@id='sortManu']//div[@id='acListDetail']//div[contains(@onclick,
     * '#modelSelection') and contains(@onclick, '#acModel')]";
     * //
     * public static final String xtxt_fl_ac_model_name =
     * ".//div[@class='col-6']/b";
     * //
     * public static final String xtxt_fl_ac_speed =
     * "//div[@id='acModel']//div[@id='acModelList']//span[@id='acSpeed']";
     * //
     * public static final String xtxt_fl_ac_capacity =
     * "//div[@id='acModel']//div[@id='acModelList']//span[@id='acCapacity']";
     * //
     * public static final String xtxt_fl_ac_price =
     * "//div[@id='acModel']//div[@id='acModelList']//span[@id='acCost']";
     * //
     * public static final String xtxt_fl_ac_range =
     * "//div[@id='acModel']//div[@id='acModelList']//div[@class='col-sm-6']//table[@class='table
     * m-text']//b[contains(text(), ' km')]";
     * //
     * public static final String xtxt_fl_ac_runway =
     * "//div[@id='acModel']//div[@id='acModelList']//div[@class='col-sm-6']//table[@class='table
     * m-text']//td[@style='width:25%;']/b[contains(text(), ' ft')]";
     */
}
