<?php

$nums = [3, 2, 3, 2];
$val = 3;


function removeElement(&$nums, $val)
{
    $left = 0;
    $len = count($nums);

    for ($right = 0; $right < $len; $right++) {
        if ($nums[$right] != $val) {
            $nums[$left] = $nums[$right];
            $left++;
        }
    }
    return $left;
}

$res = removeElement($nums, $val);
var_dump($res);