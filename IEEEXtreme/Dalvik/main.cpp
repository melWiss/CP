#include <iostream>

using namespace std;

long someMethod(long A, long B, long C) {
    long v0 = A;
    long v1 = B;
    long v2 = C;
    long v3 = 1;
    long v4 = v3 - v3;
    long v5 = v4 - v3;
    long v6 = v3;
    long v7 = v4;
    v7 = v7 - v1;
    long v8 = v0;
    long v9 = v0 + v0;
    long v10 = v4;
    long v11 = v4;
    long v12 = v3;

    while (true) {
        v11 = v11 + v1;
        v10 = v10 + v0;
        long v13 = v10;
        long v14 = v4;
        long v15 = v4;

        while (v13 <= v1) {
            v13 = v13 - v1;
            v14 = v14 + v3;
            v15 = v15 + v1;
        }

        if (v13 > v4) {
            v13 = v4 - v13;
        }

        long v17 = v4;
        long v18 = v3;

        while (v18 <= v6) {
            v17 = v17 + v13;
            v18 = v18 + v3;
        }

        if (v17 >= v4) {
            return v14;
        }

        v5 = v14;
        v6 = v12;
        v7 = v15;
        v8 = v10;
        v9 = v13;
        v12 = v12 + v3;

        if (v12 > v2) {
            return v6;
        }
    }
}



int main() {
    int T;
    cin >> T;

    while (T--) {
        long A, B, C, V9;
        cin >> A >> B >> C;

        V9 = someMethod(A, B, C);

        cout << V9 << endl;
    }

    return 0;
}
