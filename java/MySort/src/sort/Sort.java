package sort;

//https://www.cnblogs.com/flyingdreams/p/11161157.html
public class Sort {

    /**
     * 冒泡排序 思路：外层循环从1到n-1，内循环从当前外层的元素的下一个位置开始，依次和外层的元素比较，出现逆序就交换，通过与相邻元素的比较和交换来把小的数交换到最前面。
     *
     * @param arr
     * @return
     */
    public int[] bubbleSort(int[] arr) {
        int length = arr.length;
        // 外层控制趟数，没趟找到一个最大的，一共要走n-1趟
        for (int i = 0; i < length - 1; i++) {
            // 内层每次找到一个最大的，冒泡到最后面 每次循环次数为 n-1-i 次 内层除了总循环次数和外层有关，其他无关，都是自己和自己下一个比较
            for (int j = 0; j < length - 1 - i; j++) { // 从第一个数往后，如果发现相邻的比较，自己比后面的大就交换
                if (arr[j] > arr[j + 1]) {
                    int tmp;
                    tmp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = tmp;
                }
            }
        }
        return arr;
    }

    //选择排序也是一种简单直观的排序算法。它的工作原理很容易理解：初始时在序列中找到最小（大）元素，放到序列的起始位置作为已排序序列；
    //然后，再从剩余未排序元素中继续寻找最小（大）元素，放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。

    //注意选择排序与冒泡排序的区别：冒泡排序通过依次交换相邻两个顺序不合法的元素位置，从而将当前最小（大）元素放到合适的位置；
    //而选择排序每遍历一次都记住了当前最小（大）元素的位置，最后仅需一次交换操作即可将其放到合适的位置。
    public int[] selectSort(int[] arr) {
        int n = arr.length;
        int min;
        int tmp;
        for (int i = 0; i < n - 1; i++) {
            min = i;
            for (int j = i + 1; j < n; j++) {
                if (arr[j] < arr[min]) {
                    min = j;
                }
            }
            if (min != i) {
                tmp = arr[i];
                arr[i] = arr[min];
                arr[min] = tmp;
            }
        }
        return arr;
    }

    // 思路：通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。可以理解为玩扑克牌时的理牌；
    public int[] insertSort(int[] arr) {
        int n = arr.length;
        int tmp;
        for (int i = 1; i < n; i++) {
            tmp = arr[i]; // 等待插入的数字

            for (int j = i - 1; j >= 0; j--) {
                if (tmp < arr[j]) {
                    arr[j + 1] = arr[j];
                    arr[j] = tmp;
                } else {
                    break;
                }
            }
        }
        return arr;
    }

    /**
     * 快速排序
     * @param array
     */
    public  int[]  quickSort(int[] array) {
        int len;
        if(array == null
                || (len = array.length) == 0
                || len == 1) {
            return array;
        }
        sort(array, 0, len - 1);
        return array;
    }

    /**
     * 快排核心算法，递归实现
     * @param array
     * @param left
     * @param right
     */
    public static void   sort(int[] array, int left, int right) {
        if(left > right) {
            return;
        }
        // base中存放基准数
        int base = array[left];
        int i = left, j = right;
        while(i != j) {
            // 顺序很重要，先从右边开始往左找，直到找到比base值小的数
            while(array[j] >= base && i < j) {
                j--;
            }

            // 再从左往右边找，直到找到比base值大的数
            while(array[i] <= base && i < j) {
                i++;
            }

            // 上面的循环结束表示找到了位置或者(i>=j)了，交换两个数在数组中的位置
            if(i < j) {
                int tmp = array[i];
                array[i] = array[j];
                array[j] = tmp;
            }
        }

        // 将基准数放到中间的位置（基准数归位）
        array[left] = array[i];
        array[i] = base;

        // 递归，继续向基准的左右两边执行和上面同样的操作
        // i的索引处为上面已确定好的基准值的位置，无需再处理
        sort(array, left, i - 1);
        sort(array, i + 1, right);
    }

}
