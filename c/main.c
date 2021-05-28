#include <stdio.h>

int main(int argc, char const *argv[])
{
    // 将数组按照从小到大排序
    int a[] = {3, 5, 7, 4, 2, 9};
    int i, j, n, temp;
    n =
    for(i = 1; i < 6; i++) {
        for(j = 0; j < i; j++) {
            // 当插入元素较小时，进行插入操作
            if (a[i] < a[j]) {
                temp = a[i];
                // 数组的元素依次移动
                for(n = i; n > j; n--) {
                    a[n] = a[n - 1];
                }
                a[j] = temp;
            }
        }
    }


    int init;
    for(init = 0; init < 6; init++) {
        printf("%d", a[init] );
    }
    return 0;
}