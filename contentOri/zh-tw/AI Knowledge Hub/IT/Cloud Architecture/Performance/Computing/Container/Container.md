+++
title = "容器"
weight = "3"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Container的中文問題
{{< /ask_chatgpt >}}

1. 什麼是貨櫃運輸？
2. 如何確保貨櫃在運輸過程中不受損傷？
3. 貨櫃租賃和買賣的差異是什麼？
4. 如何選擇最適合自己需求的貨櫃尺寸？
5. 貨櫃運輸的優缺點是什麼？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Container並舉例
{{< /ask_chatgpt >}}

Container是一種封裝資源的技術，可以將應用程式、庫以及所有相關的配置文件、程式碼等資源打包到同一個環境中，實現一次性配置，避免了不同環境下的配置變化和應用程式的不相容性問題，使應用程序在不同環境中運行更加穩定和可靠。

舉例來說，Docker就是一種常見的Container技術。使用Docker，我們可以創建一個包含應用程序、庫、配置文件等所有相關資源的映像，然後將其部署到不同的主機、雲平台等環境中，實現應用程式的無縫運行，同時也方便了維護和升級。此外，Docker還可以實現應用程序的快速移植和多版本共存，為軟體開發、部署等帶來了更多便利。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Container的重點
{{< /ask_chatgpt >}}

1. 容器指的是一個能夠盛放物品的物體或空間。
2. 容器可以分為固體容器和液體容器。
3. 常見的固體容器包括箱子、袋子、罐子、瓶子等。
4. 常見的液體容器包括杯子、碗、壺、瓶子、桶等。
5. 容器的種類很多，不同種類的容器具有不同的特點和用途。
6. 容器需要具有承載能力和防漏性能，以確保內部物品的安全。
7. 現代化的容器還要符合環保要求，可以回收再利用。
8. 容器廣泛應用於生活、工業、農業、旅遊等各個領域。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Container的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}

1. 设计一个栈，支持push、pop、top和min函数，min函数能够返回栈中元素最小值，时间复杂度要求O(1)。
答案：使用辅助栈，每次将当前栈中最小值入栈到辅助栈中。

2. 设计一个队列，支持push、pop、peek和empty函数，其中peek函数能够返回队列顶部的元素（不删除），empty函数返回队列是否为空。
答案：使用两个栈，将元素push到栈1中，pop时先检查栈2，如果不为空，则直接pop，否则将栈1中的所有元素依次pop并push到栈2中，再从栈2中pop。

3. 设计一个可以在常数时间内获取中位数的数据结构，支持addNum和findMedian两个函数。
答案：使用两个堆，一个最大堆和一个最小堆，最大堆中存放小于中位数的元素，最小堆中存放大于中位数的元素，每次加入元素时，根据元素大小插入到相应的堆中，然后修正两个堆中的元素个数。

4. 设计一个可以存储任意数据类型的容器，支持push、pop和getRandom函数，并且时间复杂度均为O(1)。
答案：使用一个哈希表存储元素和相应的位置，在存储时保持元素和位置的一致性，删除时将该元素和相应位置交换到数组末尾并删除，然后更新哈希表。getRandom函数随机取一个索引并返回该索引位置的元素。

5. 设计一个可以支持添加、删除、查找、随机访问等操作的容器，且操作效率均为O(1)。
答案：使用哈希表存储元素和相应的位置，使用ArrayList顺序存储元素，添加元素时将元素存入ArrayList末尾，并在哈希表中记录相应的位置，删除元素时先在哈希表中查找元素位置，然后将其与数组末尾进行交换再删除，最后更新哈希表。查找、随机访问等操作均直接在ArrayList中进行。   
