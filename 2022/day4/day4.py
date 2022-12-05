import fileinput
from collections import defaultdict

def contains(r1, r2):
    return r1[0] <= r2[0] and r1[1] >= r2[1]

def containsAny(r1, r2):
    return (r1[0] <= r2[0] and r1[1] >= r2[0]) or (r1[0] <= r2[1] and r1[1] >= r2[1])

total = 0
total2 = 0
for line in fileinput.input():
    # 2-4,6-8
    range1, range2 = line.strip().split(',')
    r1 = list(map(int, range1.split('-')))
    r2 = list(map(int, range2.split('-')))
    print(r1)
    print(r2)
    total += 1 if (contains(r1, r2) or contains(r2, r1)) else 0
    total2 += 1 if (containsAny(r1, r2) or containsAny(r2, r1)) else 0

print(total)
print(total2)

