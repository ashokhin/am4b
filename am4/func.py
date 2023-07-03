def extract_int_from_string(data_string: str) -> int:
    if len(data_string) > 0:
        return int(''.join(filter(str.isdigit, data_string)))
    else:
        return 0