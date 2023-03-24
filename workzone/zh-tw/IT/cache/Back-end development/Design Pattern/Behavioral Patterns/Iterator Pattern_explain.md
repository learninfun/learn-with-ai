

Iterator Pattern是一種設計模式，它提供了一種方法來遍歷集合物件的元素，而不需要暴露物件的內部實現細節。

在Iterator Pattern中，我們定義了一個Iterator介面，它定義了迭代器的方法，包括獲取下一個元素、是否還有下一個元素等。接著，我們在集合物件中實現一個Iterator，用於獲取集合中的元素。

舉例來說，假設我們有一個列表List，我們希望能夠瀏覽列表中的元素，這時我們可以使用Iterator Pattern。首先，我們定義一個Iterator介面如下：

```
public interface Iterator {
    public boolean hasNext();
    public Object next();
}
```

接著，我們在List中實現Iterator如下：

```
public class ListIterator implements Iterator {
    private List list;
    private int index;
    
    public ListIterator(List list) {
        this.list = list;
        this.index = 0;
    }
    
    public boolean hasNext() {
        return index < list.size();
    }
    
    public Object next() {
        Object obj = list.get(index);
        index++;
        return obj;
    }
}
```

最後，我們可以使用Iterator來遍歷List中的元素：

```
List list = new ArrayList();
list.add("A");
list.add("B");
list.add("C");

Iterator iterator = new ListIterator(list);
while(iterator.hasNext()) {
    System.out.println(iterator.next());
}
```

這樣就能夠瀏覽List中的所有元素，同時不需要暴露List的內部實現細節。