class Aircraft(object):
    def __init__(
            self,
            type: str,
            reg_number: str,
            wear: int,
            a_check_hours: int
            ) -> None:
        self.type = type
        self.reg_number = reg_number
        self.wear = wear
        self.a_check_hours = a_check_hours