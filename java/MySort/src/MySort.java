import sort.Sort;

import java.util.Arrays;
public class MySort {
    public static void main(String[] args) {
        int[] myList = {2, 3, 1, 4, 6, 7, 9, 8, 0};

        Sort mysort = new Sort();
        myList = mysort.quickSort(myList);
        System.out.println(Arrays.toString(myList));
    }


}

