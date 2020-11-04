import csv
from enum import Enum

class Month(Enum):
    Never = 0
    Jan = 1
    Feb = 2
    Mar = 3
    Apr = 4
    May = 5
    June = 6
    July = 7
    Aug = 8
    Sep = 9
    Oct = 10
    Nov = 11
    Dec = 12

"""
Fields are
id,
area,
water,
months,
is_never,
Jan,Feb,Mar,Apr,May,June,July,Aug,Sep,Oct,Nov,Dec,
crop,
land_percent,
yield_measure,
normal_yield,
rwf_measure,
rwf_rent,
seed_chem_rwf,
labor_rwf,
irr_rwf
"""
def RepresentsInt(s):
    try:
        int(s)
        return True
    except ValueError:
        return False

class Farmer:
    def __init__(self, args):
        for key, val in args.items():
            if RepresentsInt(val):
                val = int(val)
            setattr(self, key, val)

    def income(self):
        return (self.normal_yield * self.rwf_measure)

    def costs(self):
        return (self.rwf_rent + self.seed_chem_rwf + self.labor_rwf + self.irr_rwf)

    def get_profits(self):
        return self.income() - self.costs()

if __name__ == "__main__":
    farmers = []
    with open('farmers.csv') as csv_file:
        csv_reader = csv.reader(csv_file, delimiter=',')
        keys =[]
        row_num = 0
        for row in csv_reader:
            if row_num == 0:
                keys = row
            else:
                if row[0] == '':
                    continue
                farmers.append(Farmer({keys[i]:row[i] for i in range(len(row))}))
            row_num += 1

    #Example of getting proftis from all farmers
    # for farmer in farmers:
    #     print(farmer.get_profits())

    #Example of increasing the yield for all farmers
    for farmer in farmers:
        x = farmer.normal_yield
        farmer.normal_yield = int(1.5 * farmer.normal_yield)
        print(f'Yield was increased from {x} to {farmer.normal_yield}')
