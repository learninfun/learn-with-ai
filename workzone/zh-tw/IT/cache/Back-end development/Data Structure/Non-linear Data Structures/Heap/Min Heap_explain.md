

Min Heap是一種資料結構，它是一棵完全二元樹，即除了最後一層之外，每一層都是滿的，並且最後一層的節點都從左邊開始填入。在Min Heap中，每一個節點的值都小於或等於它的兩個子節點的值，因此根節點保存的是整個Heap中的最小值。

舉例來說，以下是一個Min Heap：

```
          3
       /     \
     12       5
    /  \     / \
   20  11   7   14
```

在這個Min Heap中，根節點的值為3，並且根節點的兩個子節點的值都大於等於3，即12和5。同樣地，12和5這兩個節點也滿足節點值小於或等於它們的子節點的值。

另外，Min Heap的一個重要特點是，當我們從Min Heap中刪除根節點時，Heap結構會自動調整，使得新的根節點還是整個Heap中的最小值。具體來說，我們可以先將最後一個節點複製到根節點的位置，然後不斷地比較新的根節點和它的兩個子節點的值，如果它的值大於其中某一個子節點的值，就交換它們的位置，直到找到了合適的位置為止。這樣一來，我們就可以在O(log n)的時間內刪除Min Heap中的最小值了。