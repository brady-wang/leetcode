<?php



//选择排序也是一种简单直观的排序算法。它的工作原理很容易理解：初始时在序列中找到最小（大）元素，放到序列的起始位置作为已排序序列；
//然后，再从剩余未排序元素中继续寻找最小（大）元素，放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。

//注意选择排序与冒泡排序的区别：冒泡排序通过依次交换相邻两个顺序不合法的元素位置，从而将当前最小（大）元素放到合适的位置；
//而选择排序每遍历一次都记住了当前最小（大）元素的位置，最后仅需一次交换操作即可将其放到合适的位置。

function selectSort($arr)
{
// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n^2)
// 最优时间复杂度 ---- O(n^2)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 不稳定

    for ($i = 0; $i < count($arr) - 1; $i++) {
        $min = $i;
        for ($j = $i + 1; $j < count($arr); $j++) {
            if ($arr[$j] < $arr[$min]) {
                $min = $j;
            }
        }
        var_dump("找到最小的为" . $arr[$min]);
        if ($min != $i) {
            $tmp = $arr[$i];
            $arr[$i] = $arr[$min];
            $arr[$min] = $tmp;
        }
        echo "排序后" . join(',', $arr) . "<br />";
    }
    return $arr;
}

function selectSortV2($arr)
{
    $n = count($arr);

    for ($i = 0; $i < $n - 1; $i++) {
        $min = $i; // 定义个最小下标，最开始为当前的第一个 后面依次与后面的对比，记录下 最后交换
        for ($j = $i + 1; $j < $n; $j++) { // 从下一个开始，依次和前一个对比 i的位置最终要被最小的占用
            if ($arr[$j] < $arr[$min]) {
                $min = $j; // 如果找到更小的 或者更大的，把下标给min
            }
        }
        // 循环结束，如果 min ！= i需要交换
        if ($min != $i) {
            $tmp = $arr[$i];
            $arr[$i] = $arr[$min];
            $arr[$min] = $tmp;
        }
    }
    return $arr;
}

// 外层循环n-1次 内层每次都让当前元素和没有排序的i-1 到n-1之间的数比，如果找到更小的，就存入min下标 最终判断i下标和min是否一样，如果不一样，说明需要交换
function selectSortV3($arr)
{
    $n = count($arr);

    for ($i = 0; $i < $n - 1; $i++) {
        $min = $i;
        for ($j = $i + 1; $j < $n; $j++) {
            if ($arr[$j] < $arr[$min]) {
                $min = $j;
            }
        }

        if($i != $min){
            $tmp = $arr[$i];
            $arr[$i] = $arr[$min];
            $arr[$min] = $tmp;
        }
    }
    return $arr;
}

$arr = [3, 4, 234, 345, 63, 22, 35, 2, 1, 35];
echo join(',', $arr);
echo "<br />";
$res = selectSortV3($arr);
echo join(',', $res);