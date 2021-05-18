<?php

$arr = [1, 3, 34, 2, 32, 2, 78, -43, 53, -35, 0];

//鸡尾酒排序，也叫定向冒泡排序，是冒泡排序的一种改进。此算法与冒泡排序的不同处在于从低到高然后从高到低，而冒泡排序则仅从低到高去比较序列里的每个元素。他可以得到比冒泡排序稍微好一点的效能
function CocktailSort($arr)
{
    $left = 0;// 定义左边界
    $right = count($arr) - 1;// 定义右边界
    $count = 0;
    while ($left < $right) {
        for ($i = $left; $i < $right; $i++) {// 从左向右，将最大元素放在后面
            if ($arr[$i] > $arr[$i + 1]) {
                $t = $arr[$i];
                $arr[$i] = $arr[$i + 1];
                $arr[$i + 1] = $t;
            }
        }
        $right--; // 右边界缩小
        for ($i = $right; $i > $left; $i--) {// 从右向左，将最小元素放在前面
            if ($arr[$i-1] > $arr[$i]) {
                $t = $arr[$i -1];
                $arr[$i - 1] = $arr[$i];
                $arr[$i] = $t;
            }
        }
        $left++;// 左边界增大

        echo ("第".($count+1)."次排序后".join(',',$arr).'<br />');
    }
    return $arr;
}
echo join(',',$arr);
echo "<br />";
$res = CocktailSort($arr);
echo join(',',$res);
