import fileinput
from collections import defaultdict

def find_match(str1, str2):
    hm = defaultdict(lambda: False)
    for ch in str1:
        hm[ch] = True
    for ch in str2:
        if hm[ch]:
            return ch

def find_badge(str1, str2, str3):
    tot = defaultdict(lambda: 0)
    for stri in [str1, str2, str3]:
        hm = defaultdict(lambda: False)
        for ch in stri:
            hm[ch] = True
        for ch in hm:
            tot[ch] = tot[ch] + 1
    for ch in tot:
        if tot[ch] == 3:
            return ch
    raise Exception('nope')


def find_value(ch):
    # a-z is 1-26
    # A-Z is 27-52
    if ch.islower():
        return ord(ch) - 96
    else:
        return ord(ch) - 38

# part 1
# sumt = 0
# for line in fileinput.input():
#     line = line.strip()
#     print(len(line.strip()))
#     ch = find_match(line[:int(len(line)/2)], line[int(len(line)/2):])
#     val = find_value(ch)
#     print(ch, val)
#     sumt = sumt + val

# print(sumt)

# part 2
sumt = 0
lines = []
for line in fileinput.input():
    lines.append(line.strip())
    if len(lines) < 3:
        continue
    print(lines)
    ch = find_badge(lines[0], lines[1], lines[2])
    val = find_value(ch)
    print(ch, val)
    sumt = sumt + val
    lines = []

print(sumt)
