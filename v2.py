import math
def lcm(a, b):
    return (a * b) // math.gcd(a, b)


ll = [x for x in open("input.txt").read().strip().split('\n\n')]


inst = list(ll[0])
conn = {}
for l in ll[1].split("\n"):
    a = l.split(" ")[0]
    b = l.split("(")[1].split(",")[0]
    c = l.split(" ")[3].split(")")[0]
    conn[a] = (b, c)
pos = 'AAA'
idx = 0
while pos != 'ZZZ':
    d = inst[idx % len(inst)]
    pos = conn[pos][0 if d == 'L' else 1]
    idx += 1
print("p1", idx)


def solvesteps(start_pos):
    pos = start_pos
    idx = 0
    while not pos.endswith('Z'):
        d = inst[idx % len(inst)]
        pos = conn[pos][0 if d == 'L' else 1]
        idx += 1
    return idx


ret = 1
for start in conn:
    if start.endswith('A'):
        ret = lcm(ret, solvesteps(start))
print("p2", ret)
