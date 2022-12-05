import fileinput
from collections import defaultdict
import re

def contains(r1, r2):
    return r1[0] <= r2[0] and r1[1] >= r2[1]

def containsAny(r1, r2):
    return (r1[0] <= r2[0] and r1[1] >= r2[0]) or (r1[0] <= r2[1] and r1[1] >= r2[1])

stacks = [' ']
input = fileinput.input()
for line in input:
    if not line.strip():
        break
    stacks.append(line.strip())
print(stacks)

def moveStacks(amount, fromStack, toStack):
    copyStr = stacks[fromStack][-amount:]
    stacks[toStack] = stacks[toStack] + copyStr# ''.join(reversed(copyStr))
    stacks[fromStack] = stacks[fromStack][:-amount]

for line in input:
    m = re.match(r'move (\d+) from (\d+) to (\d+)', line.strip())
    moveStacks(int(m[1]), int(m[2]), int(m[3]))
    print(stacks)

print(''.join(map(lambda x: x[-1], stacks)))
