import logging


class AirplaneFuel(object):
    fuel_types = ["fuel", "co2"]

    def __init__(self, 
                 fuel_type: str = "", 
                 fuel_price: int = 0,
                 current_capacity: int = 0,
                 maximum_capacity: int = 0) -> None:
        self.fuel_type = fuel_type.lower()
        self.fuel_price = fuel_price
        self.current_capacity = current_capacity
        self.maximum_capacity = maximum_capacity
        self.holding_capacity = (maximum_capacity - current_capacity)
        if (self.holding_capacity <= int(self.maximum_capacity * 0.2)):
            logging.warning("You are holding less than 20% ({} / {}) of {}".format(self.holding_capacity, 
                                                                                   self.maximum_capacity, 
                                                                                   self.fuel_type))
            self.not_enough_fuel = True
        else:
            self.not_enough_fuel = False
    

    def __str__(self) -> str:
        return f"Airline fuel type '{self.fuel_type}'"
    
    
    def __repr__(self) -> str:
        return f"{self.__class__}('{self.fuel_type}', {self.fuel_price}, {self.holding_capacity}, {self.current_capacity}, {self.maximum_capacity})"
