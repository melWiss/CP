#include <iostream>
#include <map>

long long calc(long long n) {
    long long k = n;
    long long sum = 0;
    while (k > 0 || n > 0) {
        k = n % 10;
        sum += k * k;
        n /= 10;
    }
    return sum;
}

int main() {
    std::ios::sync_with_stdio(0);std::cin.tie(0);
    long long a, b, totalHappyNumbers = 0;
    std::cin >> a >> b;
    std::map<long long, int> numbers;

    for (long long n = a; n <= b; n++) {
        long long i = calc(n);
        while (true) {
            numbers[i] += 1;
            if (numbers[i] >= 2) {
                if (i == 1) {
                    totalHappyNumbers++;
                }
                break;
            }
            i = calc(i);
        }

        numbers.clear();
    }

    std::cout << totalHappyNumbers << std::endl;
    return 0;
}
