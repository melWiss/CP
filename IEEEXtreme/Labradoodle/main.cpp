#include <iostream>
#include <map>
#include <list>

int main() {
    int n, m;
    std::cin >> n >> m;

    std::map<std::string, std::list<std::string>*> suffix;
    std::map<std::string, std::list<std::string>*> prefix;

    for (int i = 0; i < n; i++) {
        std::string key;
        std::cin >> key;

        std::string suffixStr = key.substr(0, 3);
        if (suffix[suffixStr] == nullptr) {
            suffix[suffixStr] = new std::list<std::string>();
        }
        suffix[suffixStr]->push_back(key);

        std::string prefixStr = key.substr(key.length() - 3);
        if (prefix[prefixStr] == nullptr) {
            prefix[prefixStr] = new std::list<std::string>();
        }
        prefix[prefixStr]->push_back(key);
    }

    for (int i = 0; i < m; i++) {
        std::string word;
        std::cin >> word;

        std::string prefixStr = word.substr(word.length() - 3);
        std::string suffixStr = word.substr(0, 3);

        if (suffix[suffixStr] == nullptr || prefix[prefixStr] == nullptr) {
            std::cout << "none" << std::endl;
        } else if (suffix[suffixStr]->size() > 1 || prefix[prefixStr]->size() > 1) {
            std::cout << "ambiguous" << std::endl;
        } else {
            std::cout << *(suffix[suffixStr]->begin()) << " " << *(prefix[prefixStr]->begin()) << std::endl;
        }
    }

    // Clean up memory
    for (auto it = suffix.begin(); it != suffix.end(); ++it) {
        delete it->second;
    }

    for (auto it = prefix.begin(); it != prefix.end(); ++it) {
        delete it->second;
    }

    return 0;
}
