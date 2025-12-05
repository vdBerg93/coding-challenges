#include <fstream>
#include <vector>
#include <string>
#include <iostream>
#include <functional>

using namespace std;

using ull = unsigned long long;

vector<pair<ull,ull>> loadData(string path){
    ifstream file(path);
    vector<pair<ull,ull>> data;

    ull start, end;
    while (file >> start){
        file.ignore(1,'-');
        file >> end;
        file.ignore(1,',');
        data.push_back({start,end});
    }

    return data;
}

ull sumInvalidProducts(const vector<pair<ull,ull>>& data, function<bool(ull)> isValid){
    ull invalidSum = 0;
    for (const auto& range : data){
        for (ull id = range.first; id <= range.second; id++){
            if (!isValid(id)){
                invalidSum += id;
            }
        }
    }
    return invalidSum;
}

bool validProduct1(ull id){
    string str = to_string(id);
    int len = str.length();
    return len % 2 != 0 || str.substr(0, len/2) != str.substr(len/2);
}

bool repeatingPattern(const string& text, int sz){
    if (text.length() % sz != 0) return false;

    string pattern = text.substr(0, sz);
    for (size_t i = sz; i < text.length(); i += sz){
        if (text.substr(i, sz) != pattern) return false;
    }
    return true;
}

bool validProduct2(ull id){
    string text = to_string(id);
    for (int sz = 1; sz <= text.length() / 2; sz++){
        if (repeatingPattern(text, sz)) return false;
    }
    return true;
}

ull part1(const vector<pair<ull,ull>>& data){
    return sumInvalidProducts(data, validProduct1);
}

ull part2(const vector<pair<ull,ull>>& data){
    return sumInvalidProducts(data, validProduct2);
}

void runPart1(string filePath){
    ull result = part1(loadData(filePath));
    cout << "Part1("<<filePath<<") = "<<result<<endl;
}

void runPart2(string filePath){
    ull result = part2(loadData(filePath));
    cout << "Part2("<<filePath<<") = "<<result<<endl;
}

int main(int argc, char** argv){
    runPart1("input.txt");
    runPart2("input.txt");
}