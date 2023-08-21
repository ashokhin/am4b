import logging
import time
from am4.aircraft import Aircraft

from selenium.webdriver.remote.webelement import WebElement

from . import func
from .am4 import AM4BaseClass
from .fuel import AirplaneFuel


logging.getLogger(__name__).addHandler(logging.NullHandler())


class AirlineManager4Bot(AM4BaseClass):
    def __init__(self) -> None:
        super().__init__()
        ### Set good prices for fuel and co2
        self._fuel_good_price = 350
        self._co2_good_price = 120
        self._good_price = {
            'fuel': 350,
            'co2': 120,
        }
        ### Set available percent of budget for fuel, maintenance and marketing
        self._fuel_budget_percent = 30
        self._maintenance_budget_percent = 30
        self._marketing_budget_percent = 30
        ### Set borders for maintenance (Repairs and A-Checks)
        self._aircraft_wear_percent = 20
        self._aircraft_max_hours_to_a_check = 24
        ### Class inner variables
        self._account_money = 0
        self._fuel_data = {}
    
    
    def __exit__(self):
        self._driver.close()
    
    @property
    def fuel_good_price(self) -> int:
        return self._good_price['fuel']
    
    @fuel_good_price.setter
    def fuel_good_price(self, value: int):
        self._good_price['fuel'] = int(value)
    
    @property
    def co2_good_price(self) -> int:
        return self._good_price['co2']
    
    @co2_good_price.setter
    def co2_good_price(self, value: int):
        self._good_price['co2'] = int(value)
    
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
    def aircraft_max_hours_to_a_check(self) -> int:
        return self._aircraft_max_hours_to_a_check
    
    @aircraft_max_hours_to_a_check.setter
    def aircraft_max_hours_to_a_check(self, value: int):
        self._aircraft_max_hours_to_a_check = int(value)
    

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
        for fuel_type in AirplaneFuel.fuel_types:
            if self._fuel_data[fuel_type].not_enough_fuel:
                logging.warning("Not enough {} ({} / {}). Skip marketing companies.".format(fuel_type, 
                                                                                            self._fuel_data[fuel_type].holding_capacity, 
                                                                                            self._fuel_data[fuel_type].maximum_capacity))
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


    def _check_fuel_type(self, fuel_type: str) -> None:
        logging.info("Check '{}' price and capacity...".format(fuel_type))
        # Open popup window 'fuel'
        self._click_button(self.fuel_elements_xpath['btn_menu'])

        if fuel_type == 'co2':
            # Open 'CO2' tab
            self._click_button(self.fuel_elements_xpath['btn_co2_tab'])

        # Get info about fuel
        fuel_price = self._get_text_from_element(self.fuel_elements_xpath[fuel_type]['txt_price'])
        if fuel_price == '':
            logging.error("'{}' price not found".format(fuel_type))
            return

        fuel_price = func.extract_int_from_string(fuel_price)

        fuel_cur_cap = self._get_text_from_element(self.fuel_elements_xpath[fuel_type]['txt_cur_cap'])
        if fuel_cur_cap == '':
            logging.error("'{}' current capacity not found".format(fuel_type))
            return
        
        fuel_cur_cap = func.extract_int_from_string(fuel_cur_cap)

        fuel_max_cap = self._get_text_from_element(self.fuel_elements_xpath[fuel_type]['txt_max_cap'])
        if fuel_max_cap == '':
            logging.error("'{}' maximum capacity not found".format(fuel_type))
            return
        
        fuel_max_cap = func.extract_int_from_string(fuel_max_cap)

        self._fuel_data[fuel_type] = AirplaneFuel(fuel_type=fuel_type,
                                                  fuel_price=fuel_price,
                                                  current_capacity=fuel_cur_cap,
                                                  maximum_capacity=fuel_max_cap)
        
        # Close popup window 'fuel'
        self._click_button(self.xbtn_popup_close)
    

    def check_fuel(self):
        for fuel_type in AirplaneFuel.fuel_types:
            self._check_fuel_type(fuel_type)
    

    def _get_info(self):
        self._check_money()
        for fuel_type in AirplaneFuel.fuel_types:
            self._check_fuel_type(fuel_type)


    def _buy_fuel_type_amount(self, amount: int, fuel_type: str):
        logging.info("Buy '{}'. {} for ${}".format(fuel_type, amount, int((self._fuel_data[fuel_type].fuel_price * amount)/1000)))
        # Open popup window 'fuel'
        self._click_button(self.fuel_elements_xpath['btn_menu'])

        if fuel_type == 'co2':
            # Open 'CO2' tab
            self._click_button(self.fuel_elements_xpath['btn_co2_tab'])
        
        # Enter fuel amount
        self._type_text_in_field(self.fuel_elements_xpath['tf_amount'], str(amount))
        # Click 'purchase' button
        self._click_button(self.fuel_elements_xpath[fuel_type]['btn_purchase'])
        # Close popup window 'fuel'
        self._click_button(self.xbtn_popup_close)
    

    def _buy_fuel_type(self, fuel_type: str):
        self.check_money()
        if fuel_type not in self._fuel_data:
            self._check_fuel_type(fuel_type)

        if self._fuel_data[fuel_type].fuel_price > self._good_price[fuel_type]:
            logging.info("'{}' price is too high. Current: ${}, recommended: ${}".format(fuel_type, 
                                                                                         self._fuel_data[fuel_type].fuel_price, 
                                                                                         self._good_price[fuel_type]))
            if not self._fuel_data[fuel_type].not_enough_fuel:
                return
        
        logging.info("Buy '{}' for ${}...".format(fuel_type, self._fuel_data[fuel_type].fuel_price))
        available_money = int(self._account_money * (self._fuel_budget_percent * 0.01))
        fuel_total_price = int((self._fuel_data[fuel_type].fuel_price * self._fuel_data[fuel_type].current_capacity)/1000)
        logging.info("Available money for {fuel_type}: ${available_money}, {fuel_type} total price: ${fuel_total_price}, available {fuel_type} capacity: {available_capacity}".format(available_money=available_money,
                                                                                                                                                                                      fuel_type=fuel_type,
                                                                                                                                                                                      fuel_total_price=fuel_total_price,
                                                                                                                                                                                      available_capacity=self._fuel_data[fuel_type].current_capacity))
        
        if fuel_total_price <= 0:
            logging.info("No need to buy more {}".format(fuel_type))
            return

        if available_money >= fuel_total_price:
            self._buy_fuel_type_amount(self._fuel_data[fuel_type].current_capacity, fuel_type)
            return
        
        available_amount = int(available_money / self._fuel_data[fuel_type].fuel_price) * 1000
        self._buy_fuel_type_amount(available_amount, fuel_type)


    def _check_fuel(self):
        for fuel_type in AirplaneFuel.fuel_types:
            self._check_fuel_type(fuel_type)
    

    def _buy_fuel(self):
        logging.info("Try to buy fuel...")
        for fuel_type in AirplaneFuel.fuel_types:
            self._buy_fuel_type(fuel_type)


    def buy_fuel(self):
        self._buy_fuel()
    

    def _find_all_for_maintenance(self) -> list[Aircraft]:
        self._click_button(self.xbtn_maintenance)
        self._click_button(self.xbtn_mnt_plan)
        ac_data = []
        for ac in self._driver.find_elements('xpath', self.xelem_list_mnt_to_base):
            ac_data.append(
                Aircraft(
                    type = str(ac.get_attribute('data-type')),
                    reg_number = str(ac.get_attribute('data-reg')),
                    wear = int(float(ac.get_attribute('data-wear'))),
                    a_check_hours = int(float(ac.get_attribute('data-hours')))
                )
            )
        # Close popup window 'maintenance'
        self._click_button(self.xbtn_popup_close)

        return ac_data
    

    def _repair_aircraft(self, aircraft_for_repair: Aircraft) -> bool:
        logging.info("Repair aircraft...")

        self._click_button(self.xbtn_maintenance)
        self._click_button(self.xbtn_mnt_plan)
        self._click_button(self.xbtn_mnt_sort_by_wear)

        child_element_repair_button = None

        for ac in self._driver.find_elements('xpath', self.xelem_list_mnt_to_base):
            ac_data_reg = str(ac.get_attribute('data-reg'))
            if aircraft_for_repair.reg_number == str(ac.get_attribute('data-reg')):
                # Find 'Repair' button
                child_element_repair_button = ac.find_element('xpath', self.xbtn_mnt_repair_plan)
                break

        
        logging.info("AC type: {}, AC reg: {}, AC wear: {}".format(aircraft_for_repair.type,
                                                                   aircraft_for_repair.reg_number,
                                                                   aircraft_for_repair.wear))
        
        if not child_element_repair_button:
            logging.warning("Button for repair not found")

            return False
        
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
        logging.info("Aircraft '{}' planed to repair for ${}".format(aircraft_for_repair.reg_number, repair_cost))
        # Close popup window 'maintenance'
        self._click_button(self.xbtn_popup_close)
        self._account_money -= repair_cost

        return True


    def _repair_all_aircraft(self):
        logging.info("Search aircraft which need repair")
        acs_need_repair = []
        aircraft_on_base = self._find_all_for_maintenance()

        for ac in aircraft_on_base:
            if ac.wear >= self._aircraft_wear_percent:
                acs_need_repair.append(ac)
        
        if len(acs_need_repair) == 0:
            logging.info("No aircraft need repair")
            return
        
        self._check_money()
        repaired_acs = 0

        for ac in acs_need_repair:
            if self._repair_aircraft(ac):
                repaired_acs += 1
        
        logging.info("Aircraft repaired: {}".format(repaired_acs))
    

    def _a_check_aircraft(self, aircraft_for_a_check: Aircraft) -> bool:
        logging.info("A-Check aircraft...")
        
        self._click_button(self.xbtn_maintenance)
        self._click_button(self.xbtn_mnt_plan)
        self._click_button(self.xbtn_mnt_sort_by_a_check)
        child_element_a_check_button = None

        for ac in self._driver.find_elements('xpath', self.xelem_list_mnt_to_base):
            if aircraft_for_a_check.reg_number == str(ac.get_attribute('data-reg')):
                # Find 'A-Check' button
                child_element_a_check_button = ac.find_element('xpath', self.xbtn_mnt_a_check_plan)
                break
        
        logging.info("AC type: {}, AC reg: {}, AC hours to check: {}".format(aircraft_for_a_check.type,
                                                                             aircraft_for_a_check.reg_number,
                                                                             aircraft_for_a_check.a_check_hours))
        
        if not child_element_a_check_button:
            logging.warning("Button for A-Check not found")

            return False

        # Click 'A-Check' button
        self._click_button(child_element_a_check_button)

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
            if ac_hours_to_acheck < self._aircraft_max_hours_to_a_check:
                acs_need_acheck.append(str(ac.get_attribute('data-reg')))
        
        if len(acs_need_acheck) == 0:
            logging.info("No aircraft need A-Check")
            return
        
        self._check_money()
        achecked_acs = 0
        for ac in acs_need_acheck:
            if self._a_check_aircraft(ac):
                achecked_acs += 1
        
        logging.info("Aircraft planed for A-Check: {}".format(achecked_acs))
    

    def _modify_aircraft(self, aircraft_reg: str) -> bool:
        self._click_button(self.xbtn_maintenance)
        self._click_button(self.xbtn_mnt_plan)
        ac_data_type = ""
        ac_data_reg = ""
        child_element_modify_button = ""

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
        logging.info("Check aircraft maintenance need...")
        # A-Check
        self._acheck_all_aircraft()
        # Repair
        self._repair_all_aircraft()
        # Modification
        self._modify_all_aircraft()


    def do_maintenance(self):
        self._do_maintenance()
    

    def run_once(self):
        logging.info("Run all actions")
        self._login()
        self._buy_fuel()
        self._marketing_companies()
        self._do_maintenance()
        self._depart()
        self._check_fuel()
        self._buy_fuel()
        self._driver.close()
    

    def _run_service(self, seconds_to_sleep: int=300):
        def start_service(seconds_to_sleep):
            logging.debug("Start service")
            self.login()
            while self._logged_in:
                self._buy_fuel()
                self._marketing_companies()
                self._depart()
                self._do_maintenance()
                self._check_fuel()
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
