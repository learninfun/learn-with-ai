

Heap Sort是一种基于二元树（Binary Tree）的排序演算法，具体来说就是基于完全二元树（Complete Binary Tree）来实现的，且可看成是选择排序的一种改进版本。

在Heap Sort中，先建立一个“最大堆（Max Heap）”或“最小堆（Min Heap）”，然后将堆顶元素与堆底元素互换，接着维护堆的性质，再对剩下的元素进行堆排序，重复上述步骤直到所有元素都排好序。在最大堆的情况下，堆顶元素为最大值，堆底元素为最小值；在最小堆的情况下，堆顶元素为最小值，堆底元素为最大值。

举例来说，对一个数列{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}进行Heap Sort，具体步骤如下：

1. 建立最大堆

将数列转化成最大堆并维护其性质：

```
                        3
                     /     \
                    1       4
                  /  \     /  \
                 1    5   9    2
                / \
               6   5
```

2. 交换堆顶元素与堆底元素

将堆顶元素3与堆底元素3互换：

```
                        3
                     /     \
                    1       4
                  /  \     /  \
                 1    5   9    2
                / \
               6   5
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        3
                     /     \
                    1       4
                  /  \     /  \
                 1    5   9    2
                / \
               6   5
                        ↑
                        3
```

3. 维护堆的性质

因为堆顶元素发生了变化，所以需要重新维护堆的性质：

```
                        1
                     /     \
                    1       4
                  /  \     /  \
                 3    5   9    2
                / \
               6   5
```

4. 依次重复步骤2、3

重复步骤2和3，直到所有元素都排好序：

```
                        1
                     /     \
                    1       2
                  /  \     /  \
                 3    5   9    4
                / \
               6   5
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    1       2
                  /  \     /  \
                 5    5   9    4
                / \
               6   3
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    3       2
                  /  \     /  \
                 5    5   9    4
                / \
               6   1
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        2
                     /     \
                    3       1
                  /  \     /  \
                 5    5   9    4
                / \
               6   1
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    3       1
                  /  \     /  \
                 5    5   9    4
                / \
               6   2
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    3       1
                  /  \     /  \
                 5    5   4    9
                / \
               6   2
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    3       2
                  /  \     /  \
                 5    5   4    9
                / \
               6   1
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        2
                     /     \
                    3       1
                  /  \     /  \
                 5    5   4    9
                / \
               6   1
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    3       1
                  /  \     /  \
                 5    5   4    6
                        \
                         9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    3       1
                  /  \     /  \
                 5    5   4    2
                        \
                         6
                          \
                           9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    2       1
                  /  \     /  \
                 5    5   4    3
                        \
                         6
                          \
                           9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    1       2
                  /  \     /  \
                 5    5   4    3
                        \
                         6
                          \
                           9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    1       3
                  /  \     /  \
                 5    5   4    2
                        \
                         6
                          \
                           9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        2
                     /     \
                    1       3
                  /  \     /  \
                 5    5   4    1
                        \
                         6
                          \
                           9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    1       3
                  /  \     /  \
                 5    5   4    2
                        \
                         6
                          \
                           9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    1       2
                  /  \     /  \
                 5    5   4    3
                        \
                         6
                          \
                           9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    1       2
                  /  \     /  \
                 5    4   5    3
                        \
                         6
                          \
                           9
↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
                        1
                     /     \
                    1       2
                  /  \     /  \
                 4    3   5    5
                        \
                         6
                          \
                           9
```

最终得到一个排好序的数列：{1, 1, 2, 3, 4, 5, 5, 6, 9}。

优点是平均与最坏情况下，时间复杂度均为O(nlogn)，且适用于大资料排序。但因为常数因素较大，相对于快速排序、合并排序等排序方法，速度相对较慢。