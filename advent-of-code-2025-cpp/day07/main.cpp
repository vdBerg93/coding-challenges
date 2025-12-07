#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <algorithm>
#include <bits/stdc++.h>

using ll = long long;
using namespace std;

using Grid = vector<vector<char>>;
using Memoize = vector<vector<ll>>;

Grid loadGrid(string filePath){
    Grid g;
    ifstream file(filePath);
    string row;
    while (file >> row) {
        vector<char> rowVec(row.begin(), row.end());
        g.push_back(rowVec);
    }

    return g;
}

void mark(Grid& g, const ll& x, const ll& y){
    if (g[y][x] == '.') g[y][x] = '|';
}

void checkResult(ll want, ll got){
    if (want > 0 && want != got){
        cerr << "want: "<<want<<", got: "<<got<<endl;
        exit(1);
    }
}

void part1_countSplits(const string& filePath, ll want = 0){
    Grid g = loadGrid(filePath);
    ll splits = 0;

    for (ll y=0; y<g.size()-1; y++){
        for(ll x=0; x<g[0].size(); x++){
            const char c = g[y][x];

            if ( c == '|' || c == 'S'){
                mark(g, x, y+1);

            }else if (c == '^' && (g[y-1][x] == '|') ){
                mark(g, x-1, y);
                mark(g, x+1, y);
                mark(g, x-1, y+1); // Already visited so check again.
                ++splits;
            }
        }
    }
    checkResult(want,splits);
    cout << "Part1("<<filePath<<")= "<<splits<<endl;
}

ll countPaths(Grid g, Memoize& m, ll x, ll y){
    if (y >= g.size()) return 0;       // Done
    if (m[y][x] != -1) return m[y][x]; // Memoized

    const char c = g[y][x];
    ll count = 0;

    switch (c){
        case 'S': 
            count++; // Falthrough intended
        case '.': {
            count += countPaths(g, m, x, y+1);
            break;
        } 
        default: {
            count += 1+ countPaths(g, m, x-1, y) + countPaths(g, m, x+1, y);
        }
    }
    m[y][x] = count; // Memoize

    return count;
}

void part2_countPaths(const string& filePath, ll xStart, ll yStart, ll want = 0){
    Grid g = loadGrid(filePath);

    Memoize m(
        g.size(), 
        vector<ll>(g[0].size(),-1)
    );
    
    ll splits = countPaths(g, m, xStart, yStart);

    checkResult(want,splits);
    cout << "Part2("<<filePath<<")= "<<splits<<endl;
}


int main(int argc, char** argv){
    part1_countSplits("example.txt",21);
    part1_countSplits("input.txt");
    part2_countPaths("example.txt",7,0,40);
    part2_countPaths("input.txt",70,0);
}