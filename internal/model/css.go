package model

const (
	// Login screen
	BUTTON_PLAY_NOW     string = "button.play-now"
	BUTTON_LOGIN        string = `button[onclick="login('show');"]`
	TEXT_FIELD_LOGIN    string = "input#lEmail"
	TEXT_FIELD_PASSWORD string = "input#lPass"
	BUTTON_AUTH         string = "button#btnLogin"
	OVERLAY_LOADING     string = "div.preloader.exo.xl-text"

	// Main screen
	BUTTON_MAIN_HUBS        string = `button[onclick="popup('hubs.php','Hubs');"]`
	BUTTON_MAIN_ACCOUNT     string = `li.text-center[onclick="popup('banking.php','Banking');"]`
	BUTTON_MAIN_COMPANY     string = "div#smallMainMenu div#mapAcList"
	BUTTON_MAIN_FLEET       string = "div#smallMainMenu div#mapRoutes"
	BUTTON_MAIN_FUEL        string = `div#smallMainMenu div#mapMaint[data-original-title="Fuel & co2"]`
	BUTTON_MAIN_MAINTENANCE string = `div#smallMainMenu div#mapMaint[data-original-title="Maintenance"]`
	BUTTON_MAIN_FINANCE     string = `div#smallMainMenu div#mapMaint[data-original-title="Finance, Marketing & Stock"]`
	BUTTON_MAIN_BONUS       string = `div#smallMainMenu div#mapMaint[data-original-title="Bonus & Increase"]`
	ICON_FREE_REWARDS       string = "div#smallMainMenu img#bonusDutyFreeIconAlert"

	// Account pop-up
	LIST_ACCOUNT_ACCOUNTS        string = "div#bankingAction > table > tbody > tr"
	TEXT_ACCOUNT_ACCOUNT_NAME    string = "tr > td:nth-child(1)"
	TEXT_ACCOUNT_ACCOUNT_BALANCE string = "tr > td:nth-child(2)"

	// Common elements
	BUTTON_COMMON_TAB1        string = "#popBtn1"
	BUTTON_COMMON_TAB2        string = "#popBtn2"
	BUTTON_COMMON_TAB3        string = "#popBtn3"
	BUTTON_COMMON_CLOSE_POPUP string = `span[onclick="closePop();"]`

	// "Flight info" elements
	ICON_FI_LOUNGE_ALERT  string = "div#flightInfo span#loungeAlertIcon"
	BUTTON_ALLIANCE_INFO  string = `div#flightInfo span[onclick="popup('alliance.php','Alliance');"]`
	BUTTON_FI_OVERVIEW    string = `div#flightInfo div#flightInfoSecContainer button[onclick="popup('overview.php','Overview');"]`
	BUTTON_FI_DEPART_ALL  string = "div#flightInfo button.btn-xs:nth-child(2)"
	TEXT_FI_DEPART_AMOUNT string = "div#flightInfo span#listDepartAmount"

	// "Overview" pop-up
	TEXT_OVERVIEW_AIRLINE_REPUTATION              string = "div#popup div#popContent div.col-6:nth-child(4)"
	TEXT_OVERVIEW_CARGO_REPUTATION                string = "div#popup div#popContent div.col-6:nth-child(5)"
	TEXT_OVERVIEW_FLEET_SIZE                      string = "div#popup div#popContent div.col-sm-6:nth-child(7) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(2)"
	TEXT_OVERVIEW_AC_PENDING_DELIVERY             string = "div#popup div#popContent div.col-sm-6:nth-child(7) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(4) > td:nth-child(2)"
	TEXT_OVERVIEW_ROUTES                          string = "div#popup div#popContent div.col-sm-6:nth-child(7) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(5) > td:nth-child(2)"
	TEXT_OVERVIEW_HUBS                            string = "div#popup div#popContent div.col-sm-6:nth-child(7) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(6) > td:nth-child(2)"
	TEXT_OVERVIEW_AC_PENDING_MAINTENANCE          string = "div#popup div#popContent div.col-sm-6:nth-child(7) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(7) > td:nth-child(2)"
	TEXT_OVERVIEW_HANGAR_CAPACITY                 string = "div#popup div#popContent div.col-sm-6:nth-child(7) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(9) > td:nth-child(2)"
	TEXT_OVERVIEW_AC_INFLIGHT                     string = "div#popup div#popContent div.col-sm-6:nth-child(8) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(2)"
	TEXT_OVERVIEW_SHARE_PRICE                     string = "div#popup div#popContent div.col-sm-6:nth-child(8) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(2)"
	TEXT_OVERVIEW_FLIGHTS_OPERATED                string = "div#popup div#popContent div.col-sm-6:nth-child(8) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(4) > td:nth-child(2)"
	TEXT_OVERVIEW_PASSENGERS_ECONOMY_TRANSPORTED  string = "div#popup div#popContent div.col-sm-6:nth-child(8) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(5) > td:nth-child(2)"
	TEXT_OVERVIEW_PASSENGERS_BUSINESS_TRANSPORTED string = "div#popup div#popContent div.col-sm-6:nth-child(8) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(6) > td:nth-child(2)"
	TEXT_OVERVIEW_PASSENGERS_FIRST_TRANSPORTED    string = "div#popup div#popContent div.col-sm-6:nth-child(8) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(7) > td:nth-child(2)"
	TEXT_OVERVIEW_CARGO_TRANSPORTED_LARGE         string = "div#popup div#popContent div.col-sm-6:nth-child(8) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(8) > td:nth-child(2)"
	TEXT_OVERVIEW_CARGO_TRANSPORTED_HEAVY         string = "div#popup div#popContent div.col-sm-6:nth-child(8) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(9) > td:nth-child(2)"

	// "Alliance" pop-up
	TEXT_ALLIANCE_CONTRIBUTED_TOTAL   string = "div#popup div#popContent div#member-container tr.td-sort.bg-light > td:nth-child(3)"
	TEXT_ALLIANCE_CONTRIBUTED_PER_DAY string = "div#popup div#popContent div#member-container tr.td-sort.bg-light > td:nth-child(4)"
	TEXT_ALLIANCE_FLIGHTS             string = "div#popup div#popContent div#member-container tr.td-sort.bg-light > td:nth-child(6)"
	TEXT_ALLIANCE_SEASON_MONEY        string = "div#popup div#popContent div#member-container tr.td-sort.bg-light > td:nth-child(8)"

	// "Hubs" pop-up
	BUTTON_HUBS_LOUNGES_MAINTENANCE       string = "div#popContent button#loungeBtn"
	LIST_HUBS_LOUNGES                     string = "div#popContent table.table.table-sm.m-text > tbody > tr"
	TEXT_HUBS_LOUNGES_LOUNGE_NAME         string = "tr > td:nth-child(1)"
	TEXT_HUBS_LOUNGES_LOUNGE_WEAR_PERCENT string = "tr > td:nth-child(2) > b:nth-child(1)"
	TEXT_HUBS_LOUNGES_LOUNGE_REPAIR_COST  string = "tr > td:nth-child(2) > span:nth-child(3)"
	BUTTON_HUBS_LOUNGES_LOUNGE_REPAIR     string = "tr > td:nth-child(3) > button:nth-child(1)"
	BUTTON_HUBS_LOUNGES_BACK_TO_HUBS      string = `div#popContent button[onclick="popup('hubs.php','Hubs');"]`
	LIST_HUBS_HUBS                        string = "div#hubList > div.row.mt-1.opa.rounded"
	ELEMENT_HUB                           string = "div.row.mt-1.opa.rounded > div:nth-child(3) > div:nth-child(1)"
	TEXT_HUBS_HUB_NAME                    string = "div.p-2.col-9.exo.m-text > b"
	TEXT_HUBS_HUB_DEPARTURES              string = "div.row.mt-1.opa.rounded > div:nth-child(3) > div:nth-child(1) > div:nth-child(1) > span:nth-child(3)"
	TEXT_HUBS_HUB_ARRIVALS                string = "div.row.mt-1.opa.rounded > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > span:nth-child(3)"
	TEXT_HUBS_HUB_PAX_DEPARTED            string = "div.row.mt-1.opa.rounded > div:nth-child(4) > div:nth-child(1) > div:nth-child(1) > span:nth-child(3)"
	TEXT_HUBS_HUB_PAX_ARRIVED             string = "div.row.mt-1.opa.rounded > div:nth-child(4) > div:nth-child(1) > div:nth-child(2) > span:nth-child(3)"
	BUTTON_HUBS_HUB_MANAGE                string = "div#hubDetail button.btn.btn-danger.btn-xs-real"
	TEXT_HUBS_HUB_MANAGE_REPAIR_COST      string = "#loungeRepairCost"
	BUTTON_HUBS_HUB_MANAGE_REPAIR         string = "div#hubDetail.hidden button#loungeRepairBtn"
	BUTTON_HUBS_HUB_MANAGE_BACK           string = "#hubReturnBtn > button:nth-child(1)"
	ICON_HUBS_CATERING                    string = "div.row.mt-1.opa.rounded span.glyphicons-fast-food"
	BUTTON_HUBS_ADD_CATERING              string = "div#hubDetail button.btn-success:nth-child(1)"
	ELEM_HUBS_CATERING_OPTION_3           string = "div#caterMain div.col-4:nth-child(4)"
	SELECT_HUBS_CATERING_DURATION         string = "div#caterMain select#durationSelector"
	SELECT_HUBS_CATERING_AMOUNT           string = "div#caterMain select#caterAmount"
	TEXT_HUBS_CATERING_COST               string = "div#caterMain span#sumCost"
	BUTTON_HUBS_CATERING_BUY              string = "div#caterMain button#btnCaterDo"

	// "Company" pop-up
	TEXT_COMPANY_RANK                           string = "div.text-secondary"
	TEXT_COMPANY_STAFF_TRAINING_POINTS          string = "span#tPoints"
	TEXT_COMPANY_STAFF_PILOT_SALARY             string = "#pilotSalary"
	TEXT_COMPANY_STAFF_PILOT_MORALE             string = "#pilotMorale"
	BUTTON_COMPANY_STAFF_PILOT_SALARY_UP        string = "#pilot_main > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(1) > button:nth-child(1)"
	BUTTON_COMPANY_STAFF_PILOT_SALARY_DOWN      string = "#pilot_main > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(2) > button:nth-child(1)"
	TEXT_COMPANY_STAFF_CREW_SALARY              string = "#crewSalary"
	TEXT_COMPANY_STAFF_CREW_MORALE              string = "#crewMorale"
	BUTTON_COMPANY_STAFF_CREW_SALARY_UP         string = "#crew_main > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(1) > button:nth-child(1)"
	BUTTON_COMPANY_STAFF_CREW_SALARY_DOWN       string = "#crew_main > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(2) > button:nth-child(1)"
	TEXT_COMPANY_STAFF_ENGINEER_SALARY          string = "#engineerSalary"
	TEXT_COMPANY_STAFF_ENGINEER_MORALE          string = "#engineerMorale"
	BUTTON_COMPANY_STAFF_ENGINEER_SALARY_UP     string = "#engineer_main > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(1) > button:nth-child(1)"
	BUTTON_COMPANY_STAFF_ENGINEER_SALARY_DOWN   string = "#engineer_main > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(2) > button:nth-child(1)"
	TEXT_COMPANY_STAFF_TECHNICIAN_SALARY        string = "#techSalary"
	TEXT_COMPANY_STAFF_TECHNICIAN_MORALE        string = "#techMorale"
	BUTTON_COMPANY_STAFF_TECHNICIAN_SALARY_UP   string = "#tech_main > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(1) > button:nth-child(1)"
	BUTTON_COMPANY_STAFF_TECHNICIAN_SALARY_DOWN string = "#tech_main > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(2) > button:nth-child(1)"

	// "Fuel" pop-up
	TEXT_FUEL_FUEL_PRICE    string = "div#fuelMain span#sumCost"
	TEXT_FUEL_FUEL_HOLDING  string = "div#fuelMain #holding"
	TEXT_FUEL_FUEL_CAPACITY string = "div#fuelMain span.s-text:nth-child(4)"
	TEXT_FIELD_FUEL_AMOUNT  string = "input#amountInput"
	BUTTON_FUEL_BUY         string = "div#fuelMain button.btn-block:nth-child(2)"

	// "Maintenance" pop-up
	BUTTON_MAINTENANCE_BASE_ONLY        string = "div#maintAction button#baseOnly"
	BUTTON_MAINTENANCE_SORT_BY_CHECK    string = `div#maintAction button[onclick="sortMaint('check');"]`
	BUTTON_MAINTENANCE_SORT_BY_WEAR     string = `div#maintAction button[onclick="sortMaint();"]`
	LIST_MAINTENANCE_AC_LIST            string = "div#maintAction div#acListView > div.at-base"
	TEXT_MAINTENANCE_AC_A_CHECK_HOURS   string = "data-hours"
	TEXT_MAINTENANCE_AC_WEAR_PERCENT    string = "data-wear"
	TEXT_MAINTENANCE_AC_REG_NUMBER      string = "data-reg"
	TEXT_MAINTENANCE_AC_TYPE            string = "data-type"
	BUTTON_MAINTENANCE_A_CHECK          string = `div[role="group"] button:nth-child(1)`
	BUTTON_MAINTENANCE_REPAIR           string = `div[role="group"] button:nth-child(2)`
	BUTTON_MAINTENANCE_MODIFY           string = `div[role="group"] button:nth-child(3)`
	CHECKBOX_MAINTENANCE_MODIFY_MOD1    string = "div#typeModify table.table.table-sm.exo > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(1) > label:nth-child(1) > span:nth-child(2)"
	CHECKBOX_MAINTENANCE_MODIFY_MOD2    string = "div#typeModify table.table.table-sm.exo > tbody:nth-child(1) > tr:nth-child(2) > td:nth-child(1) > label:nth-child(1) > span:nth-child(2)"
	CHECKBOX_MAINTENANCE_MODIFY_MOD3    string = "div#typeModify table.table.table-sm.exo > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(1) > label:nth-child(1) > span:nth-child(2)"
	TEXT_MAINTENANCE_A_CHECK_TOTAL_COST string = "div#typeCheck div.col-6:nth-child(6) > div:nth-child(3)"
	TEXT_MAINTENANCE_REPAIR_TOTAL_COST  string = "div#typeRepair div:nth-child(4) > div:nth-child(3)"
	TEXT_MAINTENANCE_MODIFY_TOTAL_COST  string = "div#typeModify div.row > div.col-6.text-center > span.text-danger.font-weight-bold"
	BUTTON_MAINTENANCE_PLAN_CHECK       string = "div#typeCheck button.btn.btn-xs-real.btn-danger"
	BUTTON_MAINTENANCE_PLAN_REPAIR      string = "div#typeRepair button.btn.btn-xs-real.btn-danger"
	BUTTON_MAINTENANCE_PLAN_MODIFY      string = "div#typeModify button.btn-danger:nth-child(1)"

	// "Finance" pop-up
	BUTTON_FINANCE_MARKETING_NEW_COMPANY               string = "div#financeAction button#newCampaign"
	ELEM_FINANCE_MARKETING_LIST                        string = "div#financeAction #active-campaigns > table:nth-child(1) > tbody:nth-child(1)"
	ELEM_FINANCE_MARKETING_INC_AIRLINE_REP             string = "div#financeAction table.table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(1)"
	ELEM_FINANCE_MARKETING_INC_CARGO_REP               string = "div#financeAction table.table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(2)"
	ELEM_FINANCE_MARKETING_ECO_FRIENDLY                string = "div#financeAction table.table:nth-child(2) > tbody:nth-child(1) > tr:nth-child(3)"
	SELECT_FINANCE_MARKETING_COMPANY_DURATION          string = "div#financeAction select#dSelector"
	OPTION_FINANCE_MARKETING_INC_AIRLINE_REP_24H_VALUE string = "6"
	OPTION_FINANCE_MARKETING_INC_CARGO_REP_24H_VALUE   string = "6"
	TEXT_FINANCE_MARKETING_INC_AIRLINE_REP_COST        string = "div#financeAction span#c4"
	TEXT_FINANCE_MARKETING_INC_CARGO_REP_COST          string = "div#financeAction span#c4"
	TEXT_FINANCE_MARKETING_ECO_FRIENDLY_COST           string = "div#financeAction button.btn-danger:nth-child(1)"
	BUTTON_FINANCE_MARKETING_INC_AIRLINE_REP_BUY       string = "div#financeAction button#c4Btn"
	BUTTON_FINANCE_MARKETING_INC_CARGO_REP_BUY         string = "div#financeAction button#c4Btn"
	BUTTON_FINANCE_MARKETING_ECO_FRIENDLY_BUY          string = TEXT_FINANCE_MARKETING_ECO_FRIENDLY_COST

	// "Bonus" pop-up
	BUTTON_BONUS_DUTY_FREE_TAB string = "div#popContent button#dutyFree"
	BUTTON_BONUS_CLAIM_GIFT    string = "div#popContent div#dutyFree button#claim_gift"
)
