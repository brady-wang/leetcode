<?php



//　对于未排序数据(右手抓到的牌)，在已排序序列(左手已经排好序的手牌)中从后向前扫描，找到相应位置并插入。
//
//　　插入排序在实现上，通常采用in-place排序（即只需用到O(1)的额外空间的排序），因而在从后向前扫描过程中，需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。
//
//　　具体算法描述如下：
//
//1. 从第一个元素开始，该元素可以认为已经被排序
//2. 取出下一个元素，在已经排序的元素序列中从后向前扫描
//3. 如果该元素（已排序）大于新元素，将该元素移到下一位置
//4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
//5. 将新元素插入到该位置后
//6. 重复步骤2~5
function insertSort($arr)
{
   $count = count($arr);

   for($i=1;$i<$count;$i++){
       $tmp = $arr[$i];
       for($j = $i -1 ;$j>=0;$j--){
           if($arr[$j] < $tmp){
               $arr[$j+1] = $arr[$j];
               $arr[$j] = $tmp;
           } else {
               break;
           }
       }
   }

    return $arr;
}

function insertSortV2($arr)
{
    $count = count($arr);

    for($i=1;$i<$count;$i++){
        $tmp = $arr[$i];

        for($j=$i-1;$j>=0;$j--){
            if($arr[$j] > $tmp){
                $arr[$j+1] = $arr[$j];
                $arr[$j] = $tmp;
            }
        }
    }
    return $arr;
}

$arr = [234,4,453,5,22,234,345,45,665,78];

echo join(',',$arr);
echo "<br />";
$res = insertSortV2($arr);
echo join(',',$res);