from datetime import timedelta,datetime

def add(moment: datetime) -> datetime:
    gigasecond = 10**9
    return moment + timedelta(seconds=gigasecond)
