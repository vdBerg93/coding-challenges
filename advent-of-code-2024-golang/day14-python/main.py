import time


class Robot:
    def __init__(self,_x, _y, _vx, v_y):
        self.x,self.y,self.vx,self.vy = _x,_y,_vx,v_y
    def move(self,N,xMax,ymax):
        self.x += N*self.vx
        self.y += N*self.vy
        self.x = self.x % xMax
        self.y = self.y % ymax


def parse_input(path:str)->list[Robot]:
    with open(path) as f:
        robots = []
        for l in f.read().splitlines():
            f = l.replace("p=","").replace("v=","").split()
            x,y = map(int,f[0].split(","))
            vx,vy = map(int,f[1].split(","))
            robots.append(Robot(x,y,vx,vy))
    return robots

def solve(path:str,N:int,xmax:int,ymax:int)->int:
    robots = parse_input(path)
    for i in range(N):
        for r in robots:
            r.move(1,xmax,ymax)
        map_print(i, robots, xmax, ymax)
    return safety_factor(map_get(robots, xmax, ymax))


def map_get(robots:list[Robot], xmax:int, ymax:int)->list[list[int]]:
    m = []
    for y in range(ymax):
        row = [0 for x in range(xmax)]
        m.append(row)

    for r in robots:
        m[r.y][r.x] = m[r.y][r.x]+1

    return m


def map_print(i:int, robots:list[Robot], xmax:int, ymax:int):
    m = map_get(robots, xmax, ymax)
    prob = get_tree_likelihood(m)

    TREE_THRESHOLD=200
    if prob < TREE_THRESHOLD:
       return

    print("i:{} tree probability: {}".format(i + 1, prob))
    print("Map at i={}s".format(i + 1))
    [print(str(row).replace(", ","").replace("0"," ")) for row in m]
    print("Part2:", i + 1)
    exit(0)


def get_tree_likelihood(m:list[list[int]])->int:
    ymax, xmax = len(m), len(m[0])
    prob = 0
    for y in range(ymax):
        for x in range(xmax):
            if detect_slope(m, x, y):
                prob+=1

    return prob


def detect_slope(m:list[list[int]], x:int, y:int)->bool:
    #   .1
    #   1p
    if x > 0 and y > 0 and m[y][x-1] >=1 and m[y-1][x] >= 1:
        return True
    #   1.
    #   p1
    if x < len(m[0])-1 and y > 0 and m[y][x+1] >=1 and m[y-1][x] >= 1:
        return True

    return False


def safety_factor(m:list[list[int]])->int:
    xmax, ymax = len(m[0]), len(m)

    Q1,Q2,Q3,Q4 = 0,0,0,0
    for x in range(xmax):
        for y in range(ymax):
            v = m[y][x]
            if x < int(xmax/2) and y < int(ymax/2):
                Q1+=v
            if x > int(xmax/2) and y < int(ymax/2):
                Q2+=v
            if x < int(xmax/2) and y > int(ymax/2):
                Q3+=v
            if x > int(xmax/2) and y > int(ymax/2):
                Q4+=v

    SF = 1
    for i in [Q1,Q2,Q3,Q4]:
        SF *= i

    return SF


if __name__ == "__main__":
    p1e = solve('example.txt',100,11,7)
    p1 = solve('input.txt',100,101,103)
    print("Part1:",p1)
    solve('input.txt',50000,101,103)