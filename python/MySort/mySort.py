def bubbleSort(list):
    n = len(list)
    for i in range(0, n - 1):
        print(i)
        for j in range(0, n - i - 1):
            if list[j] > list[j + 1]:
                list[j], list[j + 1] = list[j + 1], list[j]

    return list


# 外层循环n-1次 内层每次都让当前元素和没有排序的i-1 到n-1之间的数比，如果找到更小的，就存入min下标 最终判断i下标和min是否一样，如果不一样，说明需要交换
def selectSort(list):
    n = len(list)
    for i in range(n - 1):
        minIndex = i
        for j in range(i + 1, n):
            if list[j] < list[minIndex]:
                minIndex = j

        if minIndex != i:
            list[i], list[minIndex] = list[minIndex], list[i]
    return list


arr = [1, 3, 4, 2, 5, 6, 9, 8, 0]
res = selectSort(arr)
print(res)
