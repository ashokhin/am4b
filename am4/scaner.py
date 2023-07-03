import logging
import time

from .am4 import AM4BaseClass


logging.getLogger(__name__).addHandler(logging.NullHandler())


class AM4Scaner(AM4BaseClass):
    def __init__(self) -> None:
        super().__init__()
        self._planes_dict = {}
    
    def _get_all_planes(self):
        plane_dict = {}
        self._click_button(self.xbtn_fleetandroutes)
        self._click_button(self.xbtn_fl_order)
        sort_button_text = self._get_text_from_element(self.xbtn_fl_list_sort)
        if sort_button_text == 'Show list':
            self._click_button(self.xbtn_fl_list_sort)
        
        aircrafts_webelem_list = self._driver.find_elements('xpath', self.xelem_list_fl_ac)
        logging.info("Print first: '{}'".format(aircrafts_webelem_list[0]))
        logging.info("Click first")
        self._click_button(aircrafts_webelem_list[0])
        logging.info("Model: '{}', Speed: '{}'".format(
            self._get_text_from_element(self.xtxt_fl_ac_model_name),
            self._get_text_from_element(self.xtxt_fl_ac_speed)
        ))
        self._click_button(self.xbtn_fl_order)
        aircrafts_webelem_list2 = self._driver.find_elements('xpath', self.xelem_list_fl_ac)
        logging.info("Print first: '{}'".format(aircrafts_webelem_list2[0]))
        logging.info("Click second")
        self._click_button(aircrafts_webelem_list[1])
        logging.info("Model: '{}', Speed: '{}'".format(
            self._get_text_from_element(self.xtxt_fl_ac_model_name),
            self._get_text_from_element(self.xtxt_fl_ac_speed)
        ))
        _ = input()
        

    def _scan(self):
        self._login()
        self._get_all_planes()
    
    def scan(self):
        self._scan()