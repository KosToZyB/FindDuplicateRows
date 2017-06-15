#include <iostream>
#include <fstream>
#include <map>
#include <vector>
#include <thread>
#include <mutex>
#include <unordered_map>

std::unordered_map <std::string, uint8_t> data;
std::mutex lock;

void searchRepetition(std::string fileName) {
    auto startTime = std::chrono::steady_clock::now();
    std::string line;
    std::ifstream fileData(fileName);

    if (fileData.is_open()) {
        while (fileData.good()) {
            std::getline(fileData, line);
            std::unique_lock<std::mutex> lk(lock);
            data[line]++;
        }

        fileData.close();
    } else {
        std::cout << "Unable to open file: " << fileName;
    }
    auto endTime = std::chrono::steady_clock::now();
    auto elapsedSeconds = std::chrono::duration_cast<std::chrono::seconds>(endTime - startTime);
    std::cout << fileName << " processing at " << elapsedSeconds.count() << " s\n";
}

int main() {
    auto startTime = std::chrono::steady_clock::now();
    std::vector<std::thread> threads;
    for (int i = 0; i < 5; ++i) {
        threads.push_back(std::thread(searchRepetition, "data" + std::to_string(i) + ".txt"));
    }

    for (auto &t : threads) {
        t.join();
    }

    uint64_t repeatedLine(0);
    for (const auto &p : data) {
        if (p.second > 1) {
            std::cout << p.first << ": " << unsigned(p.second) << std::endl;
            repeatedLine += p.second;
        }
    }

    std::cout << "repeated lines: " << repeatedLine << std::endl;

    auto endTime = std::chrono::steady_clock::now();
    auto elapsedSeconds = std::chrono::duration_cast<std::chrono::seconds>(endTime - startTime);
    std::cout << "Programm work at " << elapsedSeconds.count() << " s\n";
    return 0;
}