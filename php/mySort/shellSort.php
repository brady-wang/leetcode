<?php


function shellSort($arr)
{
    $len = count($arr);
    $temp = 0;
    $gap = 1;
    while($gap < $len / 3) {
        $gap = $gap * 3 + 1;
    }
    for ($gap; $gap > 0; $gap = floor($gap / 3)) {
        for ($i = $gap; $i < $len; $i++) {
            $temp = $arr[$i];
            for ($j = $i - $gap; $j >= 0 && $arr[$j] > $temp; $j -= $gap) {
                $arr[$j+$gap] = $arr[$j];
            }
            $arr[$j+$gap] = $temp;
        }
    }
    return $arr;
}


$arr = [3, 4, 234, 345, 63, 22, 35, 2, 1, 35];
echo join(',', $arr);
echo "<br />";
$res = selectSortV3($arr);
echo join(',', $res);