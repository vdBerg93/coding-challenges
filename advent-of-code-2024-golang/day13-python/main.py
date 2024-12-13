
def load_games(filePath:str):
    with open(filePath) as f:
        games = []
        A,B,P = [],[],[]
        for i,L in enumerate(f.read().splitlines()):
            filter = [" X", " Y", "+","="]
            for f in filter:
                L = L.replace(f,"")
            if "Button A:"in L:
                L = L.lstrip("Button A:").split(",")
                A = [int(L[0]),int(L[1])]
            elif "Button B:" in L:
                L = L.lstrip("Button B:").split(",")
                B = [int(L[0]), int(L[1])]
            elif "Prize:" in L:
                L = L.lstrip("Prize:").split(",")
                P = [int(L[0]),int(L[1])]
            else:
                games.append((A, B, P))
                continue
        games.append((A, B, P))
    return games

def solveGames(games:list, MAX_PRESSES:int, P) ->(int,int):
    COST_A, COST_B = 3, 1
    wins,tokens = 0, 0

    for game in games:
        A,B,ok = solve4(game,P)
        if ok == False:
            continue
        C = COST_A * A + COST_B * B

        if MAX_PRESSES>0 and (A > MAX_PRESSES or B > MAX_PRESSES):
            continue
        wins += 1
        tokens += C
    return wins, tokens

def solve4(game:tuple,P)->(int,int,bool):
    x1, x2 = game[0][0], game[1][0]
    y1, y2 = game[0][1], game[1][1]
    Gx, Gy = game[2][0] + P, game[2][1]+ P

    a = (Gx * y2 - Gy * x2) / (x1 * y2 - y1 * x2)
    b = (Gy * x1 - Gx * y1) / (x1 * y2 - y1 * x2)
    if a == int(a) and b == int(b):
        return a,b,True

    return 0,0,False

if __name__ == '__main__':
    games = load_games("input.txt")
    wins, tokens = solveGames(games,100,0)
    print("Part1: Wins={}, tokens={}".format(wins, int(tokens)))

    wins, tokens = solveGames(games,-1,10000000000000)
    print("Part2: Wins={}, tokens={}".format(wins, int(tokens)))
