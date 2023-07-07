import logging
import time

from selenium.webdriver.remote.webelement import WebElement

from . import func
from .am4 import AM4BaseClass


logging.getLogger(__name__).addHandler(logging.NullHandler())


class AirlineManager4Bot(AM4BaseClass):
    def __init__(self) -> None:
        super().__init__()
        ### Set good prices for fuel and co2
        self._fuel_good_price = 350
        self._co2_good_price = 120
        ### Set available percent of budget for fuel, maintenance and marketing
        self._fuel_budget_percent = 30
        self._maintenance_budget_percent = 30
        self._marketing_budget_percent = 30
        ### Set borders for maintenance (Repairs and A-Checks)
        self._aircraft_wear_percent = 20
        self._aircraft_max_hours_to_acheck = 24
        ### Class inner variables
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
    def maintenance_budget_percent(self) -> int:
        return int(self._maintenance_budget_percent)
    
    @maintenance_budget_percent.setter
    def maintenance_budget_percent(self, value: int):
        self._maintenance_budget_percent = int(value)

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

        company_cost = func.extract_int_from_string(self._get_text_from_element(company_cost_xpath))
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
        
        logging.info("Search marketing companies for enabling...")
        
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
            logging.info("Aircraft ready for depart: {}".format(ready_for_depart_ac))
            self._click_button(self.xbtn_depart)
            departed_ac = (ready_for_depart_ac - self._check_ready_for_depart())
            logging.info("Aircraft departed: {}".format(departed_ac))

            return
        
        logging.info("No aircraft ready to depart")
    
    def depart(self):
        self._depart()

    def _check_money(self):
        logging.info("Check account money...")
        self._refresh_page()
        account_money = self._get_text_from_element(self.xtxt_money)
        if account_money == '':
            logging.error("Amount of money not found")

            return
        
        self._account_money = func.extract_int_from_string(account_money)
    
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

        self._fuel_data['price'] = func.extract_int_from_string(fuel_price)

        fuel_cur_cap = self._get_text_from_element(self.xtxt_fuel_cur_cap)
        if fuel_cur_cap == '':
            logging.error("Fuel current capacity not found")
            return
        
        self._fuel_data['current_capacity'] = func.extract_int_from_string(fuel_cur_cap)

        fuel_max_cap = self._get_text_from_element(self.xtxt_fuel_max_cap)
        if fuel_max_cap == '':
            logging.error("Fuel maximum capacity not found")
            return
        
        self._fuel_data['maximum_capacity'] = func.extract_int_from_string(fuel_max_cap)

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

        self._co2_data['price'] = func.extract_int_from_string(co2_price)

        # Get CO2 current capacity
        co2_cur_cap = self._get_text_from_element(self.xtxt_co2_cur_cap)
        if co2_cur_cap == '':
            logging.error("CO2 current capacity not found")
            return
        
        self._co2_data['current_capacity'] = func.extract_int_from_string(co2_cur_cap)

        # Get CO2 maximum capacity
        co2_max_cap = self._get_text_from_element(self.xtxt_co2_max_cap)
        if co2_max_cap == '':
            logging.error("Fuel maximum capacity not found")
            return
        
        self._co2_data['maximum_capacity'] = func.extract_int_from_string(co2_max_cap)

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
        
        available_amount = int(available_money / self._fuel_data['price']) * 1000
        self._buy_fuel_amount(available_amount)

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
        
        available_amount = int(available_money / self._co2_data['price']) * 1000
        self._buy_co2_amount(available_amount)

    def _buy_fuel(self):
        logging.info("Try to buy fuel...")
        self._get_info()
        self._buy_fuel_percent()
        self._buy_co2_percent()            

    def buy_fuel(self):
        self._buy_fuel()
    
    def _find_all_for_maintenance(self) -> list[WebElement]:
        self._click_button(self.xbtn_maintenance)
        self._click_button(self.xbtn_mnt_plan)
        aircraft_on_base = self._driver.find_elements('xpath', self.xelem_list_mnt_to_base)
        # Close popup window 'maintenance'
        self._click_button(self.xbtn_popup_close)

        return aircraft_on_base
    
    def _repair_aircraft(self, aircraft_reg: str) -> bool:
        logging.info("Repair aircraft...")

        self._click_button(self.xbtn_maintenance)
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

        repair_cost = func.extract_int_from_string(self._get_text_from_element(self.xtxt_mnt_repair_cost))
        available_money = int(self._account_money * (self._maintenance_budget_percent * 0.01))

        if repair_cost > available_money:
            logging.warn("Repair is too expensive. Repair cost: ${}, available money for repair: ${}".format(repair_cost, 
                                                                                                             available_money))
            # Close popup window 'maintenance'
            self._click_button(self.xbtn_popup_close)
            return False
        
        self._click_button(self.xbtn_mnt_repair_do)
        logging.info("Aircraft '{}' planed to repair for ${}".format(ac_data_reg, repair_cost))
        # Close popup window 'maintenance'
        self._click_button(self.xbtn_popup_close)
        self._account_money -= repair_cost

        return True

    def _repair_all_aircraft(self):
        logging.info("Search aircraft which need repair")
        acs_need_repair = []
        aircraft_on_base = self._find_all_for_maintenance()
        for ac in aircraft_on_base:
            ac_wear = int(float(ac.get_attribute('data-wear')))
            if ac_wear >= self._aircraft_wear_percent:
                acs_need_repair.append(str(ac.get_attribute('data-reg')))
        
        if len(acs_need_repair) == 0:
            logging.info("No aircraft need repair")
            return
        
        self._check_money()
        repaired_acs = 0
        for ac in acs_need_repair:
            if self._repair_aircraft(ac):
                repaired_acs += 1
        
        logging.info("Aircraft repaired: {}".format(repaired_acs))
    
    def _acheck_aircraft(self, aircraft_reg) -> bool:
        logging.info("A-Check aircraft...")
        
        self._click_button(self.xbtn_maintenance)
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

        acheck_cost = func.extract_int_from_string(self._get_text_from_element(self.xtxt_mnt_acheck_cost))
        available_money = int(self._account_money * (self._maintenance_budget_percent * 0.01))

        if acheck_cost > available_money:
            logging.warn("A-Check is too expensive. A-Check cost: ${}, available money for A-Check: ${}".format(acheck_cost,
                                                                                                                available_money))
            # Close popup window 'maintenance'
            self._click_button(self.xbtn_popup_close)

            return False
        
        self._click_button(self.xbtn_mnt_acheck_do)
        logging.info("Aircraft '{}' planed to A-Check for ${}".format(ac_data_reg, acheck_cost))
        # Close popup window 'maintenance'
        self._click_button(self.xbtn_popup_close)
        self._account_money -= acheck_cost

        return True
    
    def _acheck_all_aircraft(self):
        logging.info("Search aircraft which need A-Check")
        acs_need_acheck = []
        aircraft_on_base = self._find_all_for_maintenance()
        for ac in aircraft_on_base:
            ac_hours_to_acheck = int(ac.get_attribute('data-hours'))
            if ac_hours_to_acheck < self._aircraft_max_hours_to_acheck:
                acs_need_acheck.append(str(ac.get_attribute('data-reg')))
        
        if len(acs_need_acheck) == 0:
            logging.info("No aircraft needs A-Check")
            return
        
        self._check_money()
        achecked_acs = 0
        for ac in acs_need_acheck:
            if self._acheck_aircraft(ac):
                achecked_acs += 1
        
        logging.info("Aircraft planed for A-Check: {}".format(achecked_acs))
    
    def _modify_aircraft(self, aircraft_reg: str) -> bool:
        self._click_button(self.xbtn_maintenance)
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
        
        modification_cost = func.extract_int_from_string(self._get_text_from_element(self.xtxt_mnt_modify_cost))
        if modification_cost == 0:
            # Close popup window 'maintenance'
            self._click_button(self.xbtn_popup_close)
            return False
        logging.info("AC type: {}, AC reg: {}".format(ac_data_type,
                                                      aircraft_reg))
        available_money = int(self._account_money * (self._maintenance_budget_percent * 0.01))
        if modification_cost > available_money:
            logging.warn("Modification is too expensive. Modification cost: ${}, available money for modification: ${}".format(modification_cost,
                                                                                                                               available_money))
            # Close popup window 'maintenance'
            self._click_button(self.xbtn_popup_close)
            return False

        self._click_button(self.xbtn_mnt_modify_do)
        logging.info("Aircraft '{}' planed for modification for ${}".format(ac_data_reg,
                                                                           modification_cost))
        # Close popup window 'maintenance'
        self._click_button(self.xbtn_popup_close)
        self._account_money -= modification_cost
        
        return True

    def _modify_all_aircraft(self):
        logging.info("Search aircraft which need modification")
        modified_acs = []
        aircraft_on_base = self._find_all_for_maintenance()
        acs_regs = []
        for ac in aircraft_on_base:
            acs_regs.append(str(ac.get_attribute('data-reg')))

        if len(acs_regs) > 0:
            self._check_money()
            if len(acs_regs) > 5:
                # Check only last 5 aircraft
                logging.info("Check only last 5 aircraft for modification need...")
                acs_regs.sort()
                acs_regs = acs_regs[-5:]
            else:
                logging.info("Check {} aircraft for modification need...".format(len(acs_regs)))
        
        for aircraft_reg in acs_regs:
            if self._modify_aircraft(aircraft_reg):
                modified_acs.append(aircraft_reg)
        
        logging.info("Aircraft planed for modification: {}".format(len(modified_acs)))
        logging.debug("Modification planed for aircraft: '{}'".format(modified_acs))

    def _do_maintenance(self):
        logging.info("Check aircraft maintenance needs...")
        # A-Check
        self._acheck_all_aircraft()
        # Repair
        # self._repair_all_aircraft()
        # Modification
        self._modify_all_aircraft()

    def do_maintenance(self):
        self._do_maintenance()
    
    def run_once(self):
        logging.info("Run all actions")
        self._login()
        self._marketing_companies()
        self._do_maintenance()
        self._depart()
        self._buy_fuel()
        self._driver.close()
    
    def _run_service(self, seconds_to_sleep: int=300):
        def start_service(seconds_to_sleep):
            logging.debug("Start service")
            self.login()
            while self._logged_in:
                self._marketing_companies()
                self._depart()
                self._do_maintenance()
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
