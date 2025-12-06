#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;

using ull = unsigned long long;

vector<vector<char>> loadGrid(string filePath){
    ifstream file(filePath);
    if (!file.is_open()) exit(1);

    vector<vector<char>> grid;
    string row;
    while ( file >> row ) {
        vector<char> rowVec(row.begin(),row.end());
        grid.push_back(rowVec);
    }

    return grid;
}

// value - Convert from ASCI char to int
int value(char c){
    return int(c)-'0';
}

struct battery{
    char val;
    int pos;
};

battery findMax(vector<char> b, int first, int last){
    int max = -1;
    char ch;
    int pos = -1;
    for(int i=first; i<=last;i++){
        int val = value(b[i]);
        if (val > max ) {
            max=val;
            ch=b[i];
            pos=i;
        }
    }

    return battery{.val=ch, .pos=pos};
}

ull selectMaxDigits(vector<char> row, int N){
    vector<battery> selection;
    int last = -1;
    for (int i=0; i<N
    ; i++){
        auto b = findMax(row,last+1,row.size()-N
    +i);
        last = b.pos;
        selection.push_back(b);
    }
    // Combine results
    string s;
    s.reserve(N);
    for (const auto& b : selection){
        s += b.val;
    }

    return stoull(s);
}

void solve(string filePath, string name, int N){
    vector<vector<char>> data = loadGrid(filePath);
    ull totalPwr = 0;
    for (auto r : data){
        totalPwr += selectMaxDigits(r,N);
    }

    cout << name<<"("<<filePath<<") = "<<totalPwr<<endl;
}

int main(int argc, char** argv){

    solve("example.txt","part1",2);
    solve("input.txt","part1",2);
    solve("example.txt","part2",12);
    solve("input.txt","part2",12);
}