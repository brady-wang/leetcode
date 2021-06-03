#include <stdio.h>

int pivotIndex(int *nums, int numsSize) {
    int total = 0;
    int sum = 0;
    for (int i = 0; i < numsSize; i++) {
        total += nums[i];
    }
    for (int j = 0; j < numsSize; j++) {
        if (2 * sum + nums[j] == total) {
            return j;
        }
        sum += nums[j];
    }
    return -1;
}

int main(int argc, char const *argv[]) {
    int nums[] = {1, 7, 3, 6, 5, 6};
    int size = sizeof(nums) / sizeof(nums[0]);
    int index;

    index = pivotIndex(nums, size);
    printf("%d", index);

}




