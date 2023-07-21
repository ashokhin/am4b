import datetime
import logging
import time

import selenium

from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait

from . import func


logging.getLogger(__name__).addHandler(logging.NullHandler())


class AM4BaseClass(object):
    ### Login page
    #
    # xpath for button 'login' on main page
    xbtn_login = '//div[@class="login-frontpage"]//button[contains(@onclick, "#login")]'
    # xpath for text field 'username' in popup window
    xtf_username = '//form[@id="loginForm"]//input[@id="lEmail"]'
    # xpath for text field 'password' in popup window
    xtf_password = '//form[@id="loginForm"]//input[@id="lPass"]'
    # xpath for check box 'remember me' in popup window
    xcb_remember_me = '//form[@id="loginForm"]//input[@id="remember"]'
    # xpath for button 'login' in popup window
    xbtn_auth = '//form[@id="loginForm"]//button[@id="btnLogin"]'

    # Loading overlay
    #
    #
    xelem_load_overlay = '/html/body/div[@class="preloader exo xl-text" and contains(@style, "AM_loading.jpg") and contains(@style, "display: none")]'
    # Main game page
    #
    # xpath for text 'Account' on main game page, which contains available money
    xtxt_money = '//nav[@id="topMenu"]//span[@id="headerAccount"]'

    ### 'Flight Info' menu
    #
    # xpath for button 'landed' in side menu
    xbtn_landed = '//div[@id="flightInfo"]//button[@id="flightStatusLanded"]'
    # xpath for list of landed aircraft in 'landed' list
    xelem_list_landed = '//div[@id="landedList"]/div[contains(@id, "flightStatus") and contains(@onclick, "showFlightInfo")]'
    # xpath for 'depart' button
    xbtn_depart = '//div[@id="flightInfo"]//button[contains(@onclick, "route_depart.php?mode=all&ids=x")]'

    ### 'Finance, Marketing & Stock' menu
    #
    # xpath for button 'Finance, Marketing & Stock' in main window
    xbtn_finance = '//div[@id="mapMaint" and @data-original-title="Finance, Marketing & Stock"]'
    # xpath for tab 'Marketing' in popup window 'finance'
    xbtn_marketing_tab = '//div[@id="popup"]//button[@id="popBtn2"]'
    # xpath for button 'New campaign' in popup window 'finance'
    xbtn_mrktn_new_campaign = '//div[@id="financeAction"]//button[@id="newCampaign"]'
    # xpath for campaign row 'Airline reputation' in popup window 'finance'
    xelem_mrktn_company_n1 = '//div[@id="campaign-1"]//tr[contains(@onclick, "marketing_new.php?type=1")]'
    # xpath for button 'Start campaign' in popup window 'finance'
    xbtn_mrktn_company_n1_do = '//div[@id="campaign-2"]//button[@id="c4Btn"]'
    # 
    xelem_mrktn_company_n2 = '//div[@id="campaign-1"]//tr[contains(@onclick, "marketing_new.php?type=5")]'
    #
    xbtn_mrktn_company_n2_do = '//div[@id="campaign-2"]//button[contains(@onclick, "marketing_new.php?type=5&mode=do&c=1")]'
    #
    xelem_list_mrktn_companies = '//div[@id="active-campaigns"]//td[@class="hasCountdown"]'

    ### 'Fuel & CO2' popup
    #
    #
    fuel_elements_xpath = {
        # xpath for button 'Fuel' in main window
        'btn_menu': '//div[@id="mapMaint" and @data-original-title="Fuel & co2"]',
        # xpath for text field 'Amount to purchase' for fuel and CO2 in popup window 'fuel'
        'tf_amount': '//div[@id="fuelMain"]//input[@id="amountInput"]',
        # xpath for tab 'co2' in popup window 'fuel'
        'btn_co2_tab': '//div[@id="popup"]//button[@id="popBtn2"]',
        'fuel': {
            # xpath for text 'current price' for fuel in popup window 'fuel'
            'txt_price': '//div[@id="fuelMain"]/div/div[1]/span[2]/b',
            # xpath for text 'max.capacity' for fuel in popup window 'fuel'
            'txt_max_cap': '//div[@id="fuelMain"]//span[@class="s-text" and contains(text(), "Lbs")]',
            # xpath for text 'current capacity' for fuel in popup window 'fuel'
            'txt_cur_cap': '//div[@id="fuelMain"]//span[@id="remCapacity"]',
            # xpath for button 'Purchase' for fuel in popup window 'fuel'
            'btn_purchase': '//div[@id="fuelMain"]//button[contains(@onclick, "fuel.php?mode=do&amount=")]',
        },
        'co2': {
            # xpath for text 'current price' for CO2 in popup window 'fuel'
            'txt_price': '//div[@id="co2Main"]/div/div[2]/span[2]/b',
            # xpath for text 'max.capacity' for CO2 in popup window 'fuel'
            'txt_max_cap': '//div[@id="co2Main"]//span[@class="s-text" and contains(text(), "Quotas")]',
            # xpath for text 'current capacity' for CO2 in popup window 'fuel'
            'txt_cur_cap': '//div[@id="co2Main"]//span[@id="remCapacity"]',
            # xpath for button 'Purchase' for CO2 in popup window 'fuel'
            'btn_purchase': '//div[@id="co2Main"]//button[contains(@onclick, "co2.php?mode=do&amount=")]',
        },
    }

    # xpath for button 'close' in all popup windows
    xbtn_popup_close = '//div[@id="popup"]//span[@onclick="closePop();"]'
    
    ### 'Maintenance' popup
    #
    # xpath for button 'Maintenance' in main window
    xbtn_maintenance = '//div[@id="mapMaint" and @data-original-title="Maintenance"]'
    # xpath for button 'Plan' in popup window 'maintenance'
    xbtn_mnt_plan = '//div[@id="popup"]//button[@id="popBtn2"]'
    # 
    xbtn_mnt_sort_by_wear = """//div[@id="maintView"]//button[@onclick="sortMaint();"]"""
    #
    xbtn_mnt_sort_by_acheck = """//div[@id="maintView"]//button[@onclick="sortMaint('check');"]"""
    #
    xelem_list_mnt_to_base = '//div[@id="acListView"]/div[@data-base="1"]'
    #
    xbtn_mnt_repair_plan = './/button[contains(@onclick, "maint_plan_do.php?type=repair&id=")]'
    #
    xtxt_mnt_repair_cost = '//div[@id="typeRepair"]//div[contains(text(), "$ ")]'
    #
    xbtn_mnt_repair_do = '//div[@id="typeRepair"]//button[contains(@onclick, "maint_plan_do.php?mode=do&type=repair&id=")]'
    #
    xbtn_mnt_acheck_plan = '//div[@id="acListView"]//button[contains(@onclick, "maint_plan_do.php?type=check&id=")]'
    #
    xtxt_mnt_acheck_cost = '//div[@id="typeCheck"]//div[contains(text(), "$ ")]'
    #
    xbtn_mnt_acheck_do = '//div[@id="typeCheck"]//button[contains(@onclick, "maint_plan_do.php?mode=do&type=check&id=")]'
    #
    xbtn_mnt_modify_plan = './/button[contains(@onclick, "maint_plan_do.php?type=modify&id=")]'
    #
    xelem_list_mnt_modify_checkbox = '//div[@id="typeModify"]//label[@class="check-container"]'
    #
    xcb_mnt_modify_reduced_co2 = './/input[@id="mod1"]'
    #
    xcb_mnt_modify_speed_increase = './/input[@id="mod2"]'
    #
    xcb_mnt_modify_reduced_fuel = './/input[@id="mod3"]'
    #
    xtxt_mnt_modify_cost = '//div[@id="typeModify"]//span[@id="acCost"]'
    #
    xbtn_mnt_modify_do = '//div[@id="typeModify"]//button[contains(@onclick, "modifyAction")]'

    ### 'Fleet & routes' popup
    #
    # xpath for button 'Fleet & routes' in main window
    xbtn_fleetandroutes = '//div[@id="mapRoutes" and @data-original-title="Fleet & routes"]'
    # xpath for button '+ Order' in popup window 'Fleet & routes'
    xbtn_fl_order = '//div[@id="popup"]//button[@id="popBtn2"]'
    #
    xbtn_fl_list_sort = '//div[@id="routeAction"]//button[@id="listSort"]'
    #
    xelem_list_fl_ac = '//div[@id="sortManu"]//div[@id="acListDetail"]//div[contains(@onclick, "#modelSelection") and contains(@onclick, "#acModel")]'
    #
    xtxt_fl_ac_model_name = './/div[@class="col-6"]/b'
    #
    xtxt_fl_ac_speed = '//div[@id="acModel"]//div[@id="acModelList"]//span[@id="acSpeed"]'
    #
    xtxt_fl_ac_capacity = '//div[@id="acModel"]//div[@id="acModelList"]//span[@id="acCapacity"]'
    #
    xtxt_fl_ac_price = '//div[@id="acModel"]//div[@id="acModelList"]//span[@id="acCost"]'
    #
    xtxt_fl_ac_capacity = '//div[@id="acModel"]//div[@id="acModelList"]//span[@id="acCapacity"]'
    #
    xtxt_fl_ac_range = '//div[@id="acModel"]//div[@id="acModelList"]//div[@class="col-sm-6"]//table[@class="table m-text"]//b[contains(text(), " km")]'
    #
    xtxt_fl_ac_runway = '//div[@id="acModel"]//div[@id="acModelList"]//div[@class="col-sm-6"]//table[@class="table m-text"]//td[@style="width:25%;"]/b[contains(text(), " ft")]'

    def __init__(self) -> None:
        logging.debug("Init driver")
        self._driver = webdriver.Chrome(options=self._set_chrome_options())
        self._am4_base_url = "https://www.airlinemanager.com/"
        self._am4_credentials = {
            'username': '',
            'password': '',
        }
        ### Class inner variables
        self._login_attempts = 0
        self._login_last_attempt = datetime.datetime.now()
        self._logged_in = False

    @property
    def am4_base_url(self) -> str:
        return self._am4_base_url
    
    @am4_base_url.setter
    def am4_base_url(self, value: str):
        self._am4_base_url = value

    @property
    def username(self) -> str:
        return self._am4_credentials['username']
    
    @username.setter
    def username(self, value: str):
        self._am4_credentials['username'] = value
    
    @property
    def password(self) -> str:
        return self._am4_credentials['password']
    
    @password.setter
    def password(self, value: str):
        self._am4_credentials['password'] = value

    def _set_chrome_options(self) -> Options:
        chrome_options = webdriver.ChromeOptions()
        chrome_options.add_argument("--headless")
        chrome_options.add_argument("--no-sandbox")
        chrome_options.add_argument("--disable-dev-shm-usage")
        # chrome_options.binary_location = "./chromedriver"
        
        return chrome_options
    
    def _refresh_page(self):
        self._driver.refresh()
        time.sleep(5)
    
    def _click_button(self, button: any):
        try:
            if isinstance(button, str):
                logging.debug("Find button '{}'".format(button))
                btn = self._driver.find_element('xpath', button)
            else:
                btn = button
        
            logging.debug("Click button '{}'".format(btn))
            if btn.is_displayed():
                btn.click()
                time.sleep(2)
            else:
                logging.warning("Button element '{}' isn't clickable".format(button))
        except selenium.common.exceptions.NoSuchElementException as nselx:
            logging.error("No such element exception. Unable to locate element: '{}'".format(button))
            logging.exception("Exception: \n{}".format(nselx))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return
        except selenium.common.exceptions.ElementClickInterceptedException as ecie:
            logging.error("Button '{}' not available for click".format(button))
            logging.exception("Exception: \n{}".format(ecie))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return
        except selenium.common.exceptions.StaleElementReferenceException as sere:
            logging.error("Button '{}' not available for click".format(button))
            logging.exception("Exception: \n{}".format(sere))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return
    
    def _type_text_in_field(self, element_xpath: str, input_text: str):
        try:
            logging.debug("Write field '{}'".format(element_xpath))
            text_field = self._driver.find_element('xpath', element_xpath)
            text_field.clear()
            text_field.send_keys(input_text)
        except selenium.common.exceptions.NoSuchElementException as nselex:
            logging.error("No such element exception. Unable to locate element: '{}'".format(element_xpath))
            logging.exception("Exception: \n{}".format(nselex))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return
    
    def _get_text_from_element(self, element: any) -> str:
        try:
            logging.debug("Get text '{}'".format(element))
            if isinstance(element, str):
                return self._driver.find_element('xpath', element).text
            return element.text
        except selenium.common.exceptions.InvalidArgumentException as iae:
            logging.error("Error with element '{}'".format(element))
            logging.exception("Exception: \n{}".format(iae))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return ""
        except selenium.common.exceptions.NoSuchElementException as nsee:
            logging.error("Error with element '{}'".format(element))
            logging.exception("Exception: \n{}".format(nsee))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return ""
    
    def _get_int_from_element(self, element: any):
        return func.extract_int_from_string(self._get_text_from_element(element))

    def _login(self):
        logging.info("Login...")
        logging.info("Login attempts: {}".format(self._login_attempts))
        logging.info("Last login attempt: {}".format(self._login_last_attempt.isoformat()))
        self._driver.delete_all_cookies()
        self._logged_in = False
        if self._login_attempts > 5:
            time_delta_sec = (datetime.datetime.now() - self._login_last_attempt).seconds
            if time_delta_sec < 60:
                logging.error("Maximum (5) login attempts reached.")
                raise UserWarning("Maximum (5) login attempts reached.")
            else:
                self._login_attempts = 0

        self._login_attempts += 1
        self._login_last_attempt = datetime.datetime.now()
        self._driver.get(self._am4_base_url)
        self._click_button(self.xbtn_login)

        self._type_text_in_field(self.xtf_username, self._am4_credentials['username'])
        self._type_text_in_field(self.xtf_password, self._am4_credentials['password'])
        self._click_button(self.xcb_remember_me)
        self._click_button(self.xbtn_auth)

        try:
            logging.info("Wait loading page after authentication...")
            _ = WebDriverWait(self._driver, 30).until(
                EC.presence_of_element_located(('xpath', self.xelem_load_overlay))
            )
            self._logged_in = True
        except Exception as ex:
            logging.exception("Login exception:\n{}".format(ex))
            raise ex
    
    def login(self):
        self._login()