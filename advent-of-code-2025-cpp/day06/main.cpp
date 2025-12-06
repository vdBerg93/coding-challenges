#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <algorithm>


using namespace std;

using ull = unsigned long long;

struct operation{
    char op;
    vector<ull> values;
     operation(char o, vector<ull> v) : op(o), values(v) {}
};

vector<operation> loadDataPart2(
    const string& filePath,
    int operatorRow
){
    ifstream file(filePath);
    vector<string> rowStr;
    vector<char> operators;
    string line;
    int lineNr = 0;

    while(getline(file, line)){
        lineNr++;
        reverse(line.begin(), line.end());

        if (lineNr < operatorRow){
            rowStr.push_back(line);
        }else{
            stringstream ss(line);
            char op;
            while(ss >> op){
                operators.push_back(op);
            }
        }
    }

    vector<vector<ull>> columns(operators.size());
    int nCol = rowStr[0].size();
    int lastSpaceCol = -1;
    int currentCol = 0;

    for (int col = 0; col <= nCol; col++){
        // Check if this column is empty (all spaces) or out of range
        bool isEmptyCol = true;
        if (col < nCol){
            for (const auto& r : rowStr){
                if (r[col] != ' '){
                    isEmptyCol = false;
                    break;
                }
            }
        }

        if (!isEmptyCol) continue;

        for (int x = lastSpaceCol + 1; x < col; x++){
            string s;
            for (const auto& r : rowStr){
                if (r[x] != ' ') s += r[x];
            }
            ull val = stoull(s);
            columns[currentCol].push_back(val);
        }
        lastSpaceCol = col;
        currentCol++;
    }

    vector<operation> ops;
    for (size_t i = 0; i < operators.size(); i++){
        ops.emplace_back(operators[i], columns[i]);
    }

    return ops;
}

vector<operation> loadDataPart1(
    const string& filePath,
    int operatorRow
){
    ifstream file(filePath);
    vector<vector<ull>> columns;
    vector<char> operators;
    string line;
    int lineNr = 0;

    while(getline(file, line)){
        lineNr++;
        stringstream ss(line);

        if (lineNr < operatorRow){
            ull val;
            int i = 0;
            while (ss >> val){
                if (i >= columns.size()){
                    columns.push_back({});
                }
                columns[i++].push_back(val);
            }
        }else{
            char op;
            while(ss >> op){
                operators.push_back(op);
            }
        }
    }

    vector<operation> ops;
    for (size_t i = 0; i < columns.size(); i++){
        char op = (i < operators.size()) ? operators[i] : ' ';
        ops.emplace_back(op, columns[i]);
    }

    return ops;
}

ull doMath(const vector<operation>& ops){
    ull total = 0;
    for (const auto& o : ops){
        ull result = (o.op == '*') ? 1 : 0;
        for(auto v : o.values){
            result = (o.op == '*') ? result * v : result + v;
        }
        total += result;
    }
    return total;
}

void validateResult(ull want, ull got){
    if (want > 0 && want != got){
        cerr << "want: "<<want<<", got: "<<got<<endl;
        exit(1);
    }
}

void part1(const string& filePath, int opRow, ull want = 0){
    auto ops = loadDataPart1(filePath,opRow);
    ull total = doMath(ops);
    validateResult(want, total);
    cout << "Part1("<<filePath<<") = "<<total<<endl;
}


void part2(const string& filePath, int opRow, ull want = 0){
    auto ops = loadDataPart2(filePath,opRow);
    ull total = doMath(ops);
    validateResult(want, total);
    cout << "Part2("<<filePath<<") = "<<total<<endl;
}


int main(int argc, char** argv){
    part1("example.txt",4,4277556);
    part1("input.txt",5);
    part2("example.txt",4,3263827);
    part2("input.txt",5); 
}