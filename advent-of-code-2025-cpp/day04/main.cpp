#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>

using namespace std;

vector<vector<char>> loadGrid(string filePath){
    vector<vector<char>> grid;
    ifstream file(filePath);
    string row;
    while (file >> row) {
        vector<char> rowVec(row.begin(), row.end());
        grid.push_back(rowVec);
    }

    return grid;
}

bool reachable(vector<vector<char>>& grid, int x, int y){
    int adjacent = 0;
    for ( int dy = -1; dy<=1; dy++){
        for (int dx = -1; dx <= 1; dx++){
            if (dy == 0 && dx == 0){
                continue; // Being evaluated
            }
            int xn = x+dx, yn = y+dy;
            if (xn < 0 || yn < 0 || xn >= grid[0].size() || yn >= grid.size()){
                continue; // Out of range
            }
            if ( grid[yn][xn] == '@' ){
                if (++adjacent >= 4) return false;
            }
        }
    }
    return true;
}

vector<pair<int,int>> findReachable(vector<vector<char>> grid){
    vector<pair<int,int>> reachableTiles;
    for (int y = 0; y < grid.size(); y++){
        for (int x = 0; x < grid[0].size(); x++){
            if (grid[y][x] == '@' && reachable(grid, x, y)){
                reachableTiles.push_back(pair(y,x));
            }
        }
    }
    return reachableTiles;
}

void runPart1(string filePath){
    vector<vector<char>> grid = loadGrid(filePath);
    auto reachableTiles = findReachable(grid);
    cout << "Part1("<<filePath<<") = "<<reachableTiles.size()<<endl;
}

void runPart2(string filePath){
    vector<vector<char>> grid = loadGrid(filePath);
    int totalReachable = 0;
    for (auto r = findReachable(grid); !r.empty(); r = findReachable(grid)){
        totalReachable += r.size();
        for (auto [y,x] : r){
            grid[y][x] = '.';
        }
        if (r.size()==0){
            break;
        }
    }
    
    cout << "Part2("<<filePath<<") = "<<totalReachable<<endl;
}

int main(int argc, char** argv){
    runPart1("example.txt");
    runPart1("input.txt");
    runPart2("example.txt");
    runPart2("input.txt");
}