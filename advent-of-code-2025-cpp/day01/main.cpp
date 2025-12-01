#include <vector>
#include <fstream>
#include <iostream>

using namespace std;

vector<pair<char,int>> loadData(string path){

    ifstream file(path);
    vector<pair<char,int>> data;

    char letter;
    int num;

    while (file >> letter >> num){
        data.push_back({letter,num});
    }

    return data;
};

int getStartScore(int pos){
    int result = 0;
    if (pos == 0) {
        result++;
    }
    return result;
}

void move(int& pos, pair<char,int> action){
    int dir = (action.first == 'R') ? 1 : -1;
    pos += dir*action.second;
    pos = ((pos % 100) + 100) % 100;
};

int part1(int pos,vector<pair<char,int>> data){
    int result = getStartScore(pos);
    for (const auto& d : data){
        move(pos,d);
        if (pos == 0){
            result++;
        }
    }

    return result;
}


int part2(int pos,vector<pair<char,int>> data){
    int result = getStartScore(pos);
    for (const auto& d : data){
        // Count how many times we cross through 0 during this rotation
        int distance_to_zero = (d.first == 'R') ? (100-pos):pos;
        int crossings = 0;
        if (d.second >= distance_to_zero && pos != 0){
            crossings = (d.second - distance_to_zero) / 100 + 1;
        } else if (pos == 0){
            crossings = d.second / 100;        
        }
        result += crossings;

        move(pos,d);
    }

    return result;
}

void runPart1(int start, string filepath){
    vector<pair<char,int>> data = loadData(filepath);
    int result = part1(start,data);
    cout<<"Part1("<<start<<", "<<filepath<<")="<<result<<endl;
}

void runPart2(int start, string filepath){
    vector<pair<char,int>> data = loadData(filepath);
    int result = part2(start,data);
    cout<<"Part2("<<start<<", "<<filepath<<")="<<result<<endl;
}


int main(int argc, char** argv){

    runPart1(50,"example.txt"); 
    runPart1(50,"input.txt");
   
    runPart2(50,"example.txt");
    runPart2(50,"input.txt");
    // 4324 too low
    // 4730 too low
    // 5118 too low
}
