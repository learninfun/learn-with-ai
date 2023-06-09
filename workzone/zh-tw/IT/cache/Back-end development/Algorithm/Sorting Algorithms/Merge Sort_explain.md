

Merge Sort是一種分治法排序演算法，它的基本概念是將待排序數列不斷拆分為較小的子序列，直到每個子序列裡面只有一個元素為止，再將相鄰的子序列合併為更大的有序子序列，直到最後只剩下一個有序的序列。

以下是Merge Sort的步驟：

1. 分割：把待排序的序列拆分成兩個子序列，每個子序列包含約一半的元素。 

2. 對每個子序列遞歸排序，直到每個子序列裡面只剩下一個元素。 

3. 合併：把排序好的子序列合併成一個大的有序序列。 

舉例來說，如果我們要對以下數列進行排序：

{38, 27, 43, 3, 9, 82, 10}

分割後會得到兩個子序列：

{38, 27, 43, 3}

{9, 82, 10}

對每個子序列遞歸排序，最終得到排序好的子序列：

{3, 27, 38, 43}

{9, 10, 82}

最後將排序好的子序列合併，得到最終的有序序列：

{3, 9, 10, 27, 38, 43, 82}