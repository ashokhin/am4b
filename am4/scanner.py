import csv
import logging
import time

from selenium.webdriver.remote.webelement import WebElement

from .am4 import AM4BaseClass


logging.getLogger(__name__).addHandler(logging.NullHandler())


class AM4Scanner(AM4BaseClass):
    def __init__(self) -> None:
        super().__init__()
        self._csv_file = ""
        self._csv_fieldnames = []
    
    def _open_ac_order(self) -> None:
        self._click_button(self.xbtn_fleetandroutes)
        self._click_button(self.xbtn_fl_order)
        sort_button_text = self._get_text_from_element(self.xbtn_fl_list_sort)
        if sort_button_text == 'Show list':
            self._click_button(self.xbtn_fl_list_sort)
        
        return
    
    def _get_ac_model(self, ac_elem: WebElement) -> str:
        ac_model_child_webelem = ac_elem.find_element('xpath', self.xtxt_fl_ac_model_name)
        
        return self._get_text_from_element(ac_model_child_webelem)
    
    def _get_ac_details(self, ac_model_name: str) -> dict:
        logging.info("Get details about AC '{}'".format(ac_model_name))
        self._open_ac_order()
        ac_dict = {}
        ac_model: str

        for ac in self._driver.find_elements('xpath', self.xelem_list_fl_ac):
            ac_model = self._get_ac_model(ac)
            if ac_model_name == ac_model:
                self._click_button(ac)
                break
        
        ac_dict = {
            'model': ac_model,
            'capacity': self._get_int_from_element(self.xtxt_fl_ac_capacity),
            'price': self._get_int_from_element(self.xtxt_fl_ac_price),
            'range': self._get_int_from_element(self.xtxt_fl_ac_range),
            'runway': self._get_int_from_element(self.xtxt_fl_ac_runway),
            'speed': self._get_int_from_element(self.xtxt_fl_ac_speed),
        }
        
        self._click_button(self.xbtn_popup_close)

        return ac_dict

    def _get_all_aircraft(self) -> list:
        aircraft_data = []
        self._open_ac_order()
        aircraft_webelem_list = self._driver.find_elements('xpath', self.xelem_list_fl_ac)
        ac_models_list = []
        for ac in aircraft_webelem_list:
            ac_models_list.append(self._get_ac_model(ac))
        
        self._click_button(self.xbtn_popup_close)

        logging.info("Total AC: {}".format(len(ac_models_list)))

        for ac in ac_models_list[0:5]:
            aircraft_data.append(self._get_ac_details(ac))
        
        return aircraft_data
    
    def _write_to_csv(self, csv_data: list):
        with open(self._csv_file, mode='w') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=self._csv_fieldnames, delimiter=';', quoting=csv.QUOTE_ALL)
            writer.writeheader()
            writer.writerows(csv_data)

    def _scan(self):
        self._login()
        self._csv_file = "./am4scanner.csv"
        self._csv_fieldnames = ["model", "capacity", "price", "range", "runway", "speed"]
        self._write_to_csv(self._get_all_aircraft())
    
    def scan(self):
        self._scan()