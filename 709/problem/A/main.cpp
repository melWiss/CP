#include <stdio.h>

int main()
{
    long long int n, b, d, s, a, j;
    scanf("%lld %lld %lld\n", &n, &b, &d);
    s = 0;
    j = 0;

    for (int i = 0; i < n; i++)
    {
        scanf("%lld", &a);
        if (a <= b)
        {
            s += a;
        }
        if (s > d) {
            s = 0;
            j++;
        }
    }

    printf("%lld\n", j);
    return 0;
}
