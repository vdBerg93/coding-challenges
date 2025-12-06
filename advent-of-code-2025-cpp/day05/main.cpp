#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <algorithm>


using namespace std;

using id = unsigned long long;
using range = pair<id,id>;

void loadData(
    const string& filePath,
    vector<range>& fresh,
    vector<id>& available
){
    ifstream file(filePath);

    bool parsingRanges = true;

    string line;
    while(getline(file,line)){
        if (line.empty()){
            parsingRanges=false;
            continue;
        }

        if (parsingRanges){
            stringstream ss(line);
            id first, second;
            char dash;
            ss >> first >> dash >> second;
            fresh.emplace_back(first,second);
        }else{
            available.emplace_back(stoull(line));
        }
    }
}

int countFreshAvailable(
    vector<range>& fresh,
    vector<id>& available
){
    int result = 0;
    for (auto i : available){
        cout << i<<endl;
        for (auto [f,l] : fresh){
            if (i >= f && i <= l){
                result++;
                break;
            }
        }
    }
    return result;
}

id countFresh(vector<range> available){
    id result = 0;
    for (auto p : available){
        result += p.second-p.first + 1;
    }
    return result;
}

// Merge overlapping or adjacent ranges (non-recursive, O(n log n))
vector<range> mergeOverlappingRanges(vector<range> ranges) {
    if (ranges.empty()) return ranges;

    // Sort ranges by start position
    sort(ranges.begin(), ranges.end());

    vector<range> merged;
    merged.push_back(ranges[0]);

    for (size_t i = 1; i < ranges.size(); i++) {
        range& last = merged.back();
        const range& current = ranges[i];

        // Check if current overlaps or is adjacent to last
        if (current.first <= last.second + 1) {
            // Merge by extending the end
            last.second = max(last.second, current.second);
        } else {
            // No overlap, add as new range
            merged.push_back(current);
        }
    }

    return merged;
}


void mergeRanges(vector<pair<id,id>>& ranges){
    for (int i = 0; i<ranges.size(); i++){
        for (int j = 0; j<ranges.size(); j++){
            if ( i==j ) continue;
            if (ranges[i].first <= ranges[j].second &&
                 ranges[i].second >= ranges[j].second){

                // Merge
                ranges[j].second = ranges[i].second;
                ranges[j].first = min(ranges[j].first,ranges[i].first);
                ranges.erase(ranges.begin()+i);

                mergeRanges(ranges);
                
                return;
            }
        }
    }
}

int part1(const string& filePath){
    vector<pair<id,id>> fresh;
    vector<id> available;
    loadData(filePath,fresh,available);
    auto result = countFreshAvailable(fresh, available);
    cout << "Part1("<<filePath<<") = "<<result<<endl;
}

int part2(const string& filePath){
    vector<pair<id,id>> fresh;
    vector<id> available;
    loadData(filePath,fresh,available);
    mergeRanges(fresh);
    id result = countFresh(fresh);
    cout << "Part1("<<filePath<<") = "<<result<<endl;
}

int main(int argc, char** argv){
    part1("example.txt");
    part1("input.txt"); // 737
    part2("example.txt");
    part2("input.txt");
}