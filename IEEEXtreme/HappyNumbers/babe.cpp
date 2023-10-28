#include <iostream>
#include <map>

int calc(int n) {
    int k = n;
    int sum = 0;
    while (k > 0 || n > 0) {
        k = n % 10;
        sum += k * k;
        n /= 10;
    }
    return sum;
}

int main() {
    std::ios::sync_with_stdio(0);std::cin.tie(0);
    int a, b, totalHappyNumbers = 0;
    bool isHappy = true;
    std::cin >> a >> b;
    std::map<int, int> numbers;

    for (int n = a; n <= b; n++) {
        int i = calc(n);
        while (isHappy) {
            numbers[i]++;
            if (numbers[i] >= 2) {
                if (i == 1) {
                    break;
                } else {
                    isHappy = false;
                    break;
                }
            }
            i = calc(i);
        }

        numbers.clear();

        if (isHappy) {
            totalHappyNumbers++;
        } else {
            isHappy = true;
        }
    }

    std::cout << totalHappyNumbers << std::endl;

    return 0;
}
