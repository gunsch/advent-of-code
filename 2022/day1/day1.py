import fileinput


elves = []
current_elf = []

for line in fileinput.input():
    if line.strip():
        calories = int(line)
        print(calories)
        current_elf.append(calories)
    else:
        elves.append(current_elf)
        print(current_elf)
        current_elf = []

max_count = 0
counts = []
for elf in elves:
    max_count = max(max_count, sum(elf))
    counts.append(sum(elf))


print(max_count)


print(sum(sorted(counts)[-3:]))
