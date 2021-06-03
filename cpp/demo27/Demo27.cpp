//
// Created by Administrator on 2021/6/1.
//

#include "Demo27.h"

class Demo27 {
public:
    int removeElement(vector<int>& nums, int val) {
        int n = nums.size();
        int left = 0;

        for (int right = 0; right < n; right++) {
            if (nums[right] != val) {
                nums[left] = nums[right];
                left++;
            }
        }
        return left;
    }
};