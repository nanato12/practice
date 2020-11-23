import datetime
import workdays
import jpholiday

# 2019/01/01 ~ 2021/12/31の祝日取得
holidays = [
    holiday_info[0]
    for holiday_info in jpholiday.between(
        start_date=datetime.datetime(2019, 1, 1),
        end_date=datetime.datetime(2021, 12, 31),
    )
]
print("---------- 祝日 ----------\n")
# 取得した祝日を表示するよ
for holiday in holidays:
    print(holiday)

# 2020/11/24から2営業日前、祝日はさっき取ったやつ使うよ
print("---------- 2020/11/24から2営業日前 ----------")
start_date = datetime.datetime(2020, 11, 24)
print(workdays.workday(start_date, days=-2, holidays=holidays))
