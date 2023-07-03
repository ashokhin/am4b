import datetime
import logging
import time
import selenium

from selenium import webdriver
from selenium.webdriver.remote.webelement import WebElement
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


logging.getLogger(__name__).addHandler(logging.NullHandler())


def extract_int_from_string(data_string: str) -> int:
        return int(''.join(filter(str.isdigit, data_string)))


class AirlineManager4Bot(object):
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
    # xpath for list of landed aircrafts in 'landed' list
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
    # xpath for button 'Fuel' in main window
    xbtn_fuel = '//div[@id="mapMaint" and @data-original-title="Fuel & co2"]'
    # xpath for text 'current price' for fuel in popup window 'fuel'
    xtxt_fuel_price = '//div[@id="fuelMain"]/div/div[1]/span[2]/b'
    # xpath for text 'max.capacity' for fuel in popup window 'fuel'
    xtxt_fuel_max_cap = '//div[@id="fuelMain"]//span[@class="s-text" and contains(text(), "Lbs")]'
    # xpath for text 'current capacity' for fuel in popup window 'fuel'
    xtxt_fuel_cur_cap = '//div[@id="fuelMain"]//span[@id="remCapacity"]'
    # xpath for text field 'Amount to purchase' for fuel and CO2 in popup window 'fuel'
    xtf_fuel_and_co2_amount = '//div[@id="fuelMain"]//input[@id="amountInput"]'
    # xpath for button 'Purchase' for fuel in popup window 'fuel'
    xbtn_fuel_purchase = '//div[@id="fuelMain"]//button[contains(@onclick, "fuel.php?mode=do&amount=")]'
    # xpath for tab 'co2' in popup window 'fuel'
    xbtn_co2_tab = '//div[@id="popup"]//button[@id="popBtn2"]'
    # xpath for text 'current price' for CO2 in popup window 'fuel'
    xtxt_co2_price = '//div[@id="co2Main"]/div/div[2]/span[2]/b'
    # xpath for text 'max.capacity' for CO2 in popup window 'fuel'
    xtxt_co2_max_cap = '//div[@id="co2Main"]//span[@class="s-text" and contains(text(), "Quotas")]'
    # xpath for text 'current capacity' for CO2 in popup window 'fuel'
    xtxt_co2_cur_cap = '//div[@id="co2Main"]//span[@id="remCapacity"]'
    # xpath for button 'Purchase' for CO2 in popup window 'fuel'
    xbtn_co2_purchase = '//div[@id="co2Main"]//button[contains(@onclick, "co2.php?mode=do&amount=")]'

    # xpath for button 'close' in all popup windows
    xbtn_popup_close = '//div[@id="popup"]//span[@onclick="closePop();"]'
    
    ### 'Maintanance' popup
    # xpath for button 'Maintanance' in main window
    xbtn_maintanance = '//div[@id="mapMaint" and @data-original-title="Maintenance"]'
    # xpath for button 'Plan' in popup window 'maintanance'
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

    def __init__(self) -> None:
        logging.debug("Init driver")
        self._driver = webdriver.Chrome(options=self._set_chrome_options())
        self._am4_base_url = "https://www.airlinemanager.com/"
        self._am4_credentials = {
            'username': '',
            'password': '',
        }
        ### Set good prices for fuel and co2
        self._fuel_good_price = 350
        self._co2_good_price = 120
        ### Set available percent of budget for fuel, maintanance and marketing
        self._fuel_budget_percent = 30
        self._maintanance_budget_percent = 30
        self._marketing_budget_percent = 30
        ### Set borders for maintanance (Repairs and A-Checks)
        self._aircraft_wear_percent = 20
        self._aircraft_max_hours_to_acheck = 24
        ### Class inner variables
        self._login_attempts = 0
        self._login_last_attempt = datetime.datetime.now()
        self._loged_in = False
        self._account_money = 0
        self._fuel_data = {
            'holding': 0,
            'price': 0,
            'current_capacity': 0,
            'maximum_capacity': 0,
        }
        self._co2_data = {
            'holding': 0,
            'price': 0,
            'current_capacity': 0,
            'maximum_capacity': 0,
        }
        self._not_enough_fuel = True
        self._not_enough_co2 = True
    
    def __exit__(self):
        self._driver.close()
    
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
    
    @property
    def fuel_good_price(self) -> int:
        return self._fuel_good_price
    
    @fuel_good_price.setter
    def fuel_good_price(self, value: int):
        self._fuel_good_price = int(value)
    
    @property
    def co2_good_price(self) -> int:
        return self._co2_good_price
    
    @co2_good_price.setter
    def co2_good_price(self, value: int):
        self._co2_good_price = int(value)
    
    @property
    def fuel_budget_percent(self) -> int:
        return int(self._fuel_budget_percent)
    
    @fuel_budget_percent.setter
    def fuel_budget_percent(self, value: int):
        self._fuel_budget_percent = int(value)
    
    @property
    def maintanance_budget_percent(self) -> int:
        return int(self._maintanance_budget_percent)
    
    @maintanance_budget_percent.setter
    def maintanance_budget_percent(self, value: int):
        self._maintanance_budget_percent = int(value)

    @property
    def marketing_budget_percent(self) -> int:
        return int(self._marketing_budget_percent)
    
    @marketing_budget_percent.setter
    def marketing_budget_percent(self, value: int):
        self._marketing_budget_percent = int(value)

    @property
    def aircraft_wear_percent(self) -> int:
        return self._aircraft_wear_percent
    
    @aircraft_wear_percent.setter
    def aircraft_wear_percent(self, value: int):
        self._aircraft_wear_percent = int(value)
    
    @property
    def aircraft_max_hours_to_acheck(self) -> int:
        return self._aircraft_max_hours_to_acheck
    
    @aircraft_max_hours_to_acheck.setter
    def aircraft_max_hours_to_acheck(self, value: int):
        self._aircraft_max_hours_to_acheck = int(value)

    def _set_chrome_options(self) -> selenium.webdriver.chrome.options.Options:
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
            logging.error("Button '{}' not avaiable for click".format(button))
            logging.exception("Exception: \n{}".format(ecie))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return
        except selenium.common.exceptions.StaleElementReferenceException as sere:
            logging.error("Button '{}' not avaiable for click".format(button))
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
    
    def _get_text_from_element(self, element_xpath: str) -> str:
        try:
            logging.debug("Get text '{}'".format(element_xpath))
            return self._driver.find_element('xpath', element_xpath).text
        except selenium.common.exceptions.InvalidArgumentException as iae:
            logging.error("Error with element '{}'".format(element_xpath))
            logging.exception("Exception: \n{}".format(iae))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return ""
        except selenium.common.exceptions.NoSuchElementException as nsee:
            logging.error("Error with element '{}'".format(element_xpath))
            logging.exception("Exception: \n{}".format(nsee))
            logging.debug("Page source: {}".format(self._driver.page_source))

            return ""

    def _login(self):
        logging.info("Login...")
        logging.info("Login attempts: {}".format(self._login_attempts))
        logging.info("Last login attempt: {}".format(self._login_last_attempt.isoformat()))
        self._driver.delete_all_cookies()
        self._loged_in = False
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
            logging.info("Wait loading page after authentification...")
            _ = WebDriverWait(self._driver, 30).until(
                EC.presence_of_element_located(('xpath', self.xelem_load_overlay))
            )
            self._loged_in = True
        except Exception as ex:
            logging.exception("Login exception:\n{}".format(ex))
            raise ex
    
    def login(self):
        self._login()
    
    def _company_active(self, company_xpath: str) -> bool:
        company_web_elem = self._driver.find_element('xpath', company_xpath)
        if company_web_elem.get_attribute('class') == "not-active":
            return False
        
        return True

    def _enable_marketing_company(self, marketing_company: dict):
        logging.info("Check marketing company '{}'...".format(marketing_company['name']))

        self._click_button(self.xbtn_finance)
        self._click_button(self.xbtn_marketing_tab)
        self._click_button(self.xbtn_mrktn_new_campaign)
        if not self._company_active(marketing_company['row_xpath']):
            logging.info("Marketing company '{}' already active".format(marketing_company['name']))
            self._click_button(self.xbtn_popup_close)
            return

        self._click_button(marketing_company['row_xpath'])

        company_cost_xpath = ""

        if marketing_company['name'] == 'Airline reputation':
            company_cost_xpath = '{}//span[@id="c4"]'.format(marketing_company['button_xpath'])
        
        if marketing_company['name'] == 'Eco friendly':
            company_cost_xpath = marketing_company['button_xpath']

        company_cost = extract_int_from_string(self._get_text_from_element(company_cost_xpath))
        available_money = int(self._account_money * (self._marketing_budget_percent * 0.01))

        if company_cost > available_money:
            logging.warning("Not enough money for marketing company. Available money for marketing company: ${}. Marketing company price: ${}".format(available_money,
                                                                                                                                                      company_cost))
            self._click_button(self.xbtn_popup_close)
            return

        logging.info("Activate marketing company '{}' for ${}".format(marketing_company['name'], company_cost))
        self._click_button(marketing_company['button_xpath'])
        self._click_button(self.xbtn_popup_close)
        self._account_money -= company_cost
    
    def _enable_marketing_companies(self):
        marketing_companies = [
            {
                'name': 'Airline reputation',
                'row_xpath': self.xelem_mrktn_company_n1,
                'button_xpath': self.xbtn_mrktn_company_n1_do,
            },
            {
                'name': 'Eco friendly',
                'row_xpath': self.xelem_mrktn_company_n2,
                'button_xpath': self.xbtn_mrktn_company_n2_do,
            },
        ]

        self._check_money()
        for marketing_company in marketing_companies:
            self._enable_marketing_company(marketing_company)
    
    def _check_marketing_companies(self):
        self._click_button(self.xbtn_finance)
        self._click_button(self.xbtn_marketing_tab)
        active_marketing_companies = self._driver.find_elements('xpath', self.xelem_list_mrktn_companies)

        self._click_button(self.xbtn_popup_close)
        return len(active_marketing_companies)

    def _marketing_companies(self):
        self._get_info()
        if self._not_enough_fuel:
            logging.warning("Not enough fuel ({} / {}). Skip marketing companies.".format(self._fuel_data['holding'], 
                                                                                         self._fuel_data['maximum_capacity']))
            return
        
        if self._not_enough_co2:
            logging.warning("Not enough CO2 ({} / {}). Skip marketing companies.".format(self._co2_data['holding'], 
                                                                                        self._co2_data['maximum_capacity']))
            return
        
        logging.info("Search marketing companies to enable...")
        
        mcc = self._check_marketing_companies()

        if mcc == 2:
            logging.info("All marketing companies enabled: {}".format(mcc))
            return
        
        logging.info("{} marketing company(s) enabled".format(mcc))
        
        self._enable_marketing_companies()
    
    def marketing_companies(self):
        self._marketing_companies()

    def _check_ready_for_depart(self) -> int:
        self._click_button(self.xbtn_landed)
        elems = self._driver.find_elements('xpath', self.xelem_list_landed)
        
        return len(elems)
    
    def _depart(self):
        logging.info("Depart all available planes...")
        ready_for_depart_ac = self._check_ready_for_depart()

        if ready_for_depart_ac > 0:
            logging.info("Aircrafts ready for depart: {}".format(ready_for_depart_ac))
            self._click_button(self.xbtn_depart)
            departed_ac = (ready_for_depart_ac - self._check_ready_for_depart())
            logging.info("Aircrafts departed: {}".format(departed_ac))

            return
        
        logging.info("No aircrafts ready to depart")
    
    def depart(self):
        self._depart()

    def _check_money(self):
        logging.info("Check account money...")
        self._refresh_page()
        account_money = self._get_text_from_element(self.xtxt_money)
        if account_money == '':
            logging.error("Amount of money not found")

            return
        
        self._account_money = extract_int_from_string(account_money)
    
    def check_money(self):
        self._check_money()

    def _check_fuel(self):
        logging.info("Check fuel/CO2 prices and capacity...")
        # Open popup window 'fuel'
        self._click_button(self.xbtn_fuel)

        # Get info about fuel
        fuel_price = self._get_text_from_element(self.xtxt_fuel_price)
        if fuel_price == '':
            logging.error("Fuel price not found")
            return

        self._fuel_data['price'] = extract_int_from_string(fuel_price)

        fuel_cur_cap = self._get_text_from_element(self.xtxt_fuel_cur_cap)
        if fuel_cur_cap == '':
            logging.error("Fuel current capacity not found")
            return
        
        self._fuel_data['current_capacity'] = extract_int_from_string(fuel_cur_cap)

        fuel_max_cap = self._get_text_from_element(self.xtxt_fuel_max_cap)
        if fuel_max_cap == '':
            logging.error("Fuel maximum capacity not found")
            return
        
        self._fuel_data['maximum_capacity'] = extract_int_from_string(fuel_max_cap)

        self._fuel_data['holding'] = int(self._fuel_data['maximum_capacity'] - self._fuel_data['current_capacity'])

        if self._fuel_data['holding'] <= int(self._fuel_data['maximum_capacity'] * 0.2):
            logging.warning("You are holding less than 20% ({} / {}) of fuel".format(self._fuel_data['holding'], 
                                                                                     self._fuel_data['maximum_capacity']))
            self._not_enough_fuel = True
        else:
            self._not_enough_fuel = False
        
        # Get info about CO2
        self._click_button(self.xbtn_co2_tab)

        # Get CO2 price
        co2_price = self._get_text_from_element(self.xtxt_co2_price)
        if co2_price == '':
            logging.error("CO2 price not found")
            return

        self._co2_data['price'] = extract_int_from_string(co2_price)

        # Get CO2 current capacity
        co2_cur_cap = self._get_text_from_element(self.xtxt_co2_cur_cap)
        if co2_cur_cap == '':
            logging.error("CO2 current capacity not found")
            return
        
        self._co2_data['current_capacity'] = extract_int_from_string(co2_cur_cap)

        # Get CO2 maximum capacity
        co2_max_cap = self._get_text_from_element(self.xtxt_co2_max_cap)
        if co2_max_cap == '':
            logging.error("Fuel maximum capacity not found")
            return
        
        self._co2_data['maximum_capacity'] = extract_int_from_string(co2_max_cap)

        self._co2_data['holding'] = int(self._co2_data['maximum_capacity'] - self._co2_data['current_capacity'])

        if self._co2_data['holding'] <= int(self._co2_data['maximum_capacity'] * 0.2):
            logging.warning("You are holding less than 20% ({} / {}) of CO2".format(self._co2_data['holding'], 
                                                                                     self._co2_data['maximum_capacity']))
            self._not_enough_co2 = True
        else:
            self._not_enough_co2 = False

        # Close popup window 'fuel'
        self._click_button(self.xbtn_popup_close)        

    def check_fuel(self):
        self._check_fuel()
    
    def _get_info(self):
        self._check_money()
        self._check_fuel()

    def get_info(self):
        self._get_info()
        print("""
===Airline info===
Account:\t$ {}
Fuel price:\t$ {}
Fuel capacity:\t{:.2f} %
CO2 price:\t$ {}
CO2 capacity:\t{:.2f} %
===================""".format(self._account_money,
                              self._fuel_data['price'],
                              100 * float(self._fuel_data['current_capacity'])/float(self._fuel_data['maximum_capacity']),
                              self._co2_data['price'],
                              100 * float(self._co2_data['current_capacity'])/float(self._co2_data['maximum_capacity'])))

    def _buy_fuel_amount(self, amount: int):
        logging.info("Buy fuel. {} Lbs for ${}".format(amount, int((self._fuel_data['price'] * amount)/1000)))
        # Open popup window 'fuel'
        self._click_button(self.xbtn_fuel)
        # Enter fuel amount
        self._type_text_in_field(self.xtf_fuel_and_co2_amount, str(amount))
        # Click 'purchase' button
        self._click_button(self.xbtn_fuel_purchase)
        # Close popup window 'fuel'
        self._click_button(self.xbtn_popup_close)
        self._check_money()

    def _buy_fuel_percent(self):
        # self._refresh_page()
        if self._fuel_data['price'] > self._fuel_good_price:
            logging.info("Fuel price is too high. Current: ${}, recommended: ${}".format(self._fuel_data['price'], self._fuel_good_price))
            return
        
        logging.info("Buy fuel for good price: ${}...".format(self._fuel_data['price']))
        available_money = int(self._account_money * (self._fuel_budget_percent * 0.01))
        fuel_total_price = int((self._fuel_data['price'] * self._fuel_data['current_capacity'])/1000)
        logging.info("Available money for fuel: ${}, fuel total price: ${}, available capacity: {}".format(available_money, 
                                                                                                           fuel_total_price,
                                                                                                           self._fuel_data['current_capacity']))
        if fuel_total_price <= 0:
            logging.info("No need to buy more fuel")
            return

        if fuel_total_price <= available_money:
            self._buy_fuel_amount(self._fuel_data['current_capacity'])
            return
        
        avaiable_amount = int(available_money / self._fuel_data['price']) * 1000
        self._buy_fuel_amount(avaiable_amount)

    def _buy_co2_amount(self, amount: int):
        logging.info("Buy CO2. {} Quotas for ${}".format(amount, int((self._co2_data['price'] * amount)/1000)))
        # Open popup window 'fuel'
        self._click_button(self.xbtn_fuel)
        # Go to tab 'CO2' in popup window 'fuel'
        self._click_button(self.xbtn_co2_tab)
        # Enter fuel amount
        self._type_text_in_field(self.xtf_fuel_and_co2_amount, str(amount))
        # Click 'purchase' button
        self._click_button(self.xbtn_co2_purchase)
        # Close popup window 'fuel'
        self._click_button(self.xbtn_popup_close)
        self._check_money()

    def _buy_co2_percent(self):
        if self._co2_data['price'] > self._co2_good_price:
            logging.info("CO2 price is too high. Current: ${}, recommended: ${}".format(self._co2_data['price'], self._co2_good_price))
            return

        logging.info("Buy CO2 quotas for good price: ${}...".format(self._co2_data['price']))
        
        available_money = int(self._account_money * (self._fuel_budget_percent * 0.01))
        co2_total_price = int((self._co2_data['price'] * self._co2_data['current_capacity'])/1000)
        logging.info("Available money for CO2: ${}, CO2 total price: ${}".format(available_money, co2_total_price))
        if co2_total_price <= 0:
            logging.info("No need to buy more CO2")
            return

        if co2_total_price <= available_money:
            self._buy_co2_amount(self._co2_data['current_capacity'])
            return
        
        avaiable_amount = int(available_money / self._co2_data['price']) * 1000
        self._buy_co2_amount(avaiable_amount)

    def _buy_fuel(self):
        logging.info("Try to buy fuel...")
        self._get_info()
        self._buy_fuel_percent()
        self._buy_co2_percent()            

    def buy_fuel(self):
        self._buy_fuel()
    
    def _find_all_for_maintanance(self) -> list[WebElement]:
        self._click_button(self.xbtn_maintanance)
        self._click_button(self.xbtn_mnt_plan)
        aircrafts_on_base = self._driver.find_elements('xpath', self.xelem_list_mnt_to_base)
        # Close popup window 'maintanance'
        self._click_button(self.xbtn_popup_close)

        return aircrafts_on_base
    
    def _repair_aircraft(self, aircraft_reg: str) -> bool:
        logging.info("Repair aircraft...")

        self._click_button(self.xbtn_maintanance)
        self._click_button(self.xbtn_mnt_plan)
        self._click_button(self.xbtn_mnt_sort_by_wear)
        ac_data_type: str
        ac_data_reg: str
        ac_data_wear: str
        child_element_repair_button: WebElement

        for ac in self._driver.find_elements('xpath', self.xelem_list_mnt_to_base):
            ac_data_reg = str(ac.get_attribute('data-reg'))
            if ac_data_reg == aircraft_reg:
                ac_data_type = str(ac.get_attribute('data-type'))
                ac_data_wear = str(ac.get_attribute('data-wear'))
                # Find 'Repair' button
                child_element_repair_button = ac.find_element('xpath', self.xbtn_mnt_repair_plan)
                break

        
        logging.info("AC type: {}, AC reg: {}, AC wear: {}".format(ac_data_type,
                                                                   ac_data_reg,
                                                                   ac_data_wear))
        
        # Click 'Repair' button
        self._click_button(child_element_repair_button)

        repair_cost = extract_int_from_string(self._get_text_from_element(self.xtxt_mnt_repair_cost))
        available_money = int(self._account_money * (self._maintanance_budget_percent * 0.01))

        if repair_cost > available_money:
            logging.warn("Repair is too expensive. Repair cost: ${}, available money for repair: ${}".format(repair_cost, 
                                                                                                             available_money))
            # Close popup window 'maintanance'
            self._click_button(self.xbtn_popup_close)
            return False
        
        self._click_button(self.xbtn_mnt_repair_do)
        logging.info("Aircraft '{}' planed to repair for ${}".format(ac_data_reg, repair_cost))
        # Close popup window 'maintanance'
        self._click_button(self.xbtn_popup_close)
        self._account_money -= repair_cost

        return True

    def _repair_all_aircrafts(self):
        logging.info("Search aircrafts which need repair")
        acs_need_repair = []
        aircrafts_on_base = self._find_all_for_maintanance()
        for ac in aircrafts_on_base:
            ac_wear = int(float(ac.get_attribute('data-wear')))
            if ac_wear >= self._aircraft_wear_percent:
                acs_need_repair.append(str(ac.get_attribute('data-reg')))
        
        if len(acs_need_repair) == 0:
            logging.info("No aircrafts need repair")
            return
        
        self._check_money()
        repaired_acs = 0
        for ac in acs_need_repair:
            if self._repair_aircraft(ac):
                repaired_acs += 1
        
        logging.info("Aircrafts repaired: {}".format(repaired_acs))
    
    def _acheck_aircraft(self, aircraft_reg) -> bool:
        logging.info("A-Check aircraft...")
        
        self._click_button(self.xbtn_maintanance)
        self._click_button(self.xbtn_mnt_plan)
        self._click_button(self.xbtn_mnt_sort_by_acheck)
        ac_data_type: str
        ac_data_reg: str
        ac_data_hours: str
        child_element_acheck_button: WebElement

        for ac in self._driver.find_elements('xpath', self.xelem_list_mnt_to_base):
            ac_data_reg = str(ac.get_attribute('data-reg'))
            if ac_data_reg == aircraft_reg:
                ac_data_type = str(ac.get_attribute('data-type'))
                ac_data_hours = str(ac.get_attribute('data-hours'))
                # Find 'A-Check' button
                child_element_acheck_button = ac.find_element('xpath', self.xbtn_mnt_acheck_plan)
                break
        
        logging.info("AC type: {}, AC reg: {}, AC hours to check: {}".format(ac_data_type,
                                                                             ac_data_reg,
                                                                             ac_data_hours))
        
        # Click 'A-Check' button
        self._click_button(child_element_acheck_button)

        acheck_cost = extract_int_from_string(self._get_text_from_element(self.xtxt_mnt_acheck_cost))
        available_money = int(self._account_money * (self._maintanance_budget_percent * 0.01))

        if acheck_cost > available_money:
            logging.warn("A-Check is too expensive. A-Check cost: ${}, available money for A-Check: ${}".format(acheck_cost,
                                                                                                                available_money))
            # Close popup window 'maintanance'
            self._click_button(self.xbtn_popup_close)

            return False
        
        self._click_button(self.xbtn_mnt_acheck_do)
        logging.info("Aircraft '{}' planed to A-Check for ${}".format(ac_data_reg, acheck_cost))
        # Close popup window 'maintanance'
        self._click_button(self.xbtn_popup_close)
        self._account_money -= acheck_cost

        return True
    
    def _acheck_all_aircrafts(self):
        logging.info("Search aircrafts which need A-Check")
        acs_need_acheck = []
        aircrafts_on_base = self._find_all_for_maintanance()
        for ac in aircrafts_on_base:
            ac_hours_to_acheck = int(ac.get_attribute('data-hours'))
            if ac_hours_to_acheck < self._aircraft_max_hours_to_acheck:
                acs_need_acheck.append(str(ac.get_attribute('data-reg')))
        
        if len(acs_need_acheck) == 0:
            logging.info("No aircrafts needs A-Check")
            return
        
        self._check_money()
        achecked_acs = 0
        for ac in acs_need_acheck:
            if self._acheck_aircraft(ac):
                achecked_acs += 1
        
        logging.info("Aircrafts planed for A-Check: {}".format(achecked_acs))
    
    def _modify_aircraft(self, aircraft_reg: str) -> bool:
        self._click_button(self.xbtn_maintanance)
        self._click_button(self.xbtn_mnt_plan)
        ac_data_type: str
        ac_data_reg: str
        child_element_modify_button: WebElement

        for ac in self._driver.find_elements('xpath', self.xelem_list_mnt_to_base):
            ac_data_reg = str(ac.get_attribute('data-reg'))
            if ac_data_reg == aircraft_reg:
                ac_data_type = str(ac.get_attribute('data-type'))
                # Find 'Modify' button
                child_element_modify_button = ac.find_element('xpath', self.xbtn_mnt_modify_plan)
                break

            
        # Click 'Modify' button
        self._click_button(child_element_modify_button)

        for modification_checkbox_row in self._driver.find_elements('xpath', self.xelem_list_mnt_modify_checkbox):
            for checkbox_xpath in [self.xcb_mnt_modify_reduced_co2, self.xcb_mnt_modify_speed_increase, self.xcb_mnt_modify_reduced_fuel]:
                try:
                    checkbox_web_elem = modification_checkbox_row.find_element('xpath', checkbox_xpath)

                    if bool(checkbox_web_elem.get_attribute('checked')):
                        break
                    
                    checkbox_span = modification_checkbox_row.find_element('xpath', './/span[@class="checkmark"]')
                    self._click_button(checkbox_span)
                    break

                except:
                    continue
        
        modification_cost = extract_int_from_string(self._get_text_from_element(self.xtxt_mnt_modify_cost))
        if modification_cost == 0:
            # Close popup window 'maintanance'
            self._click_button(self.xbtn_popup_close)
            return False
        logging.info("AC type: {}, AC reg: {}".format(ac_data_type,
                                                      aircraft_reg))
        available_money = int(self._account_money * (self._maintanance_budget_percent * 0.01))
        if modification_cost > available_money:
            logging.warn("Modification is too expensive. Modification cost: ${}, available money for modification: ${}".format(modification_cost,
                                                                                                                               available_money))
            # Close popup window 'maintanance'
            self._click_button(self.xbtn_popup_close)
            return False

        self._click_button(self.xbtn_mnt_modify_do)
        logging.info("Aircraft '{}' planed for modification for ${}".format(ac_data_reg,
                                                                           modification_cost))
        # Close popup window 'maintanance'
        self._click_button(self.xbtn_popup_close)
        self._account_money -= modification_cost
        
        return True

    def _modify_all_aircrafts(self):
        logging.info("Search aircrafts which need modification")
        modified_acs = []
        aircrafts_on_base = self._find_all_for_maintanance()
        acs_regs = []
        for ac in aircrafts_on_base:
            acs_regs.append(str(ac.get_attribute('data-reg')))

        if len(acs_regs) > 0:
            self._check_money()
            if len(acs_regs) > 5:
                # Check only last 5 aircrafts
                logging.info("Check only last 5 aircrafts for modification need...")
                acs_regs.sort()
                acs_regs = acs_regs[-5:]
            else:
                logging.info("Check {} aircrafts for modification need...".format(len(acs_regs)))
        
        for aircraft_reg in acs_regs:
            if self._modify_aircraft(aircraft_reg):
                modified_acs.append(aircraft_reg)
        
        logging.info("Aircrafts planed for modification: {}".format(len(modified_acs)))
        logging.debug("Modification planed for aircrafts: '{}'".format(modified_acs))

    def _do_maintanance(self):
        logging.info("Check aircrafts maintanance needs...")
        # A-Check
        self._acheck_all_aircrafts()
        # Repair
        self._repair_all_aircrafts()
        # Modification
        self._modify_all_aircrafts()

    def do_maintanance(self):
        self._do_maintanance()
    
    def run_once(self):
        logging.info("Run all actions")
        self._login()
        self._marketing_companies()
        self._do_maintanance()
        self._depart()
        self._buy_fuel()
        self._driver.close()
    
    def _run_service(self, seconds_to_sleep: int=300):
        def start_service(seconds_to_sleep):
            logging.debug("Start service")
            self.login()
            while self._loged_in:
                self._marketing_companies()
                self._depart()
                self._do_maintanance()
                self._buy_fuel()
                time.sleep(seconds_to_sleep)
        
        while True:
            try:
                logging.info("Start AM4Bot as a service")
                start_service(seconds_to_sleep)
            except KeyboardInterrupt:
                return
            except Exception as ex:
                raise ex

    def run_service(self, seconds_to_sleep: int=300):
        try:
            self._run_service(seconds_to_sleep)
        finally:
            self._driver.close()
