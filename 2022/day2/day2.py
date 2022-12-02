import fileinput

value_scores = {
    ('A', 'X'): 1 + 3,
    ('A', 'Y'): 2 + 6,
    ('A', 'Z'): 3 + 0,
    ('B', 'X'): 1 + 0,
    ('B', 'Y'): 2 + 3,
    ('B', 'Z'): 3 + 6,
    ('C', 'X'): 1 + 6,
    ('C', 'Y'): 2 + 0,
    ('C', 'Z'): 3 + 3,
}
value2_scores = {
    ('A', 'X'): 3 + 0,
    ('A', 'Y'): 1 + 3,
    ('A', 'Z'): 2 + 6,
    ('B', 'X'): 1 + 0,
    ('B', 'Y'): 2 + 3,
    ('B', 'Z'): 3 + 6,
    ('C', 'X'): 2 + 0,
    ('C', 'Y'): 3 + 3,
    ('C', 'Z'): 1 + 6,
}


total = 0
total2 = 0
for line in fileinput.input():
    score = value_scores[tuple(line.split())]
    total += score

    score2 = value2_scores[tuple(line.split())]
    total2 += score2

print(total)
print(total2)

