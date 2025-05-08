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
	ICON_FI_LOUNGE_ALERT string = `div[id="flightInfo"] > span[id="loungeAlertIcon"]`
	BUTTON_FI_OVERVIEW   string = `div#flightInfo div#flightInfoSecContainer button[onclick="popup('overview.php','Overview');"]`

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

	// "Hubs" pop-up
	LIST_HUBS_HUBS                   string = "div#hubList > div.row.mt-1.opa.rounded"
	ELEMENT_HUB                      string = "div.row.mt-1.opa.rounded > div:nth-child(3) > div:nth-child(1)"
	TEXT_HUBS_HUB_NAME               string = "div.p-2.col-9.exo.m-text > b"
	TEXT_HUBS_HUB_DEPARTURES         string = "div.row.mt-1.opa.rounded > div:nth-child(3) > div:nth-child(1) > div:nth-child(1) > span:nth-child(3)"
	TEXT_HUBS_HUB_ARRIVALS           string = "div.row.mt-1.opa.rounded > div:nth-child(3) > div:nth-child(1) > div:nth-child(2) > span:nth-child(3)"
	TEXT_HUBS_HUB_PAX_DEPARTED       string = "div.row.mt-1.opa.rounded > div:nth-child(4) > div:nth-child(1) > div:nth-child(1) > span:nth-child(3)"
	TEXT_HUBS_HUB_PAX_ARRIVED        string = "div.row.mt-1.opa.rounded > div:nth-child(4) > div:nth-child(1) > div:nth-child(2) > span:nth-child(3)"
	BUTTON_HUBS_HUB_MANAGE           string = "div#hubDetail button.btn.btn-danger.btn-xs-real"
	TEXT_HUBS_HUB_MANAGE_REPAIR_COST string = "#loungeRepairCost"
	BUTTON_HUBS_HUB_MANAGE_REPAIR    string = "div#hubDetail.hidden button#loungeRepairBtn"
	BUTTON_HUBS_HUB_MANAGE_BACK      string = "#hubReturnBtn > button:nth-child(1)"

	// "Company" pop-up
	TEXT_COMPANY_RANK                  string = "div.text-secondary"
	TEXT_COMPANY_STAFF_TRAINING_POINTS string = "span#tPoints"

	// "Fuel" pop-up
	TEXT_FUEL_FUEL_PRICE    string = "div#fuelMain span#sumCost"
	TEXT_FUEL_FUEL_HOLDING  string = "div#fuelMain #holding"
	TEXT_FUEL_FUEL_CAPACITY string = "div#fuelMain span.s-text:nth-child(4)"
	TEXT_FIELD_FUEL_AMOUNT  string = "input#amountInput"
	BUTTON_FUEL_BUY         string = "div#fuelMain button.btn-block:nth-child(2)"
)
