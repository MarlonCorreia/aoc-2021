
def part1(steps):
    l = []
    for x in set(steps):
        f = 0
        for v in steps:
            f += abs(x - v)
        l.append(f)
    
    return min(l)

with open('/home/marlinho/projects/aoc-2021/07/input.txt') as f:
    steps = [int(x) for x in f.read().split(',')]

    p1 = part1(steps)
    print(f"Part1 {p1}")
