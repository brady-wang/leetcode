<?php
// 两数字之和


// 循环方式 外层从2开始到11  依次与内层的当前数的下一位到最后一位相加  相同就找到,否则未找到
// 第一次外层循环 2 +  7 11 15
// 第二次外层循环 7 + 11 15
// 第三次外层循环 11 + 15


function twoSum($nums, $target)
{
    $len = count($nums);
    for ($i = 0; $i < $len-1; $i++) {
        for ($j = $i + 1; $j < $len; $j++) {
            if ($nums[$i] + $nums[$j] == $target) {
                return [$i, $j];
            }
        }

    }
    return [];
}

$arr = [2, 7, 11, 15];
$target = 26;

$res = twoSum($arr, $target);
var_dump($res);

// 结果

/*array(2) {
[0]=>
int(2)
[1]=>
int(3)
}*/