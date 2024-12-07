from itertools import product

def readdata(path:str)->map:
    data = {}
    with open(path) as file:
        for line in file:
            result, values = line.split(':')
            result = int(result.strip())
            values = list(map(int,values.strip().split()))
            data[result]=values
    return data

def solve(N):  
    solution = 0    
    for result, values in data.items():
        solution += checknumber(N, result, values)
    return solution
    
def checknumber(N:[int], result:int, values:[int]):
    configs = list(product(N, repeat=len(values)-1))
    for config in configs:
        if checkconfiguration(config, result, values):
            return result
    return 0
  
def checkconfiguration(config, result, values)->bool:    
    value = 0
    for idx, x in enumerate(values):
        if idx==0:
            value = x
        else:
            c = config[idx-1]
            if c==1:
                value *= x
            elif c==2:
                value += x
            elif c==3:
                value = int(str(value)+str(x))
            else:
                print("invalid: {}".format(c))
                exit
    if value == result:
        return True  



data = readdata('input.txt')
print("Part 1: {}".format(solve([1,2])))
print("Part 2: {}".format(solve([1,2,3])))