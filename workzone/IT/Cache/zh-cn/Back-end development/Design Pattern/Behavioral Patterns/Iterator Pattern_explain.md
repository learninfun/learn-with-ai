

Iterator Pattern是一种设计模式，它提供了一种方法来遍历集合物件的元素，而不需要暴露物件的内部实现细节。

在Iterator Pattern中，我们定义了一个Iterator介面，它定义了迭代器的方法，包括获取下一个元素、是否还有下一个元素等。接着，我们在集合物件中实现一个Iterator，用于获取集合中的元素。

举例来说，假设我们有一个列表List，我们希望能够浏览列表中的元素，这时我们可以使用Iterator Pattern。首先，我们定义一个Iterator介面如下：

```
public interface Iterator {
    public boolean hasNext();
    public Object next();
}
```

接着，我们在List中实现Iterator如下：

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

最后，我们可以使用Iterator来遍历List中的元素：

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

这样就能够浏览List中的所有元素，同时不需要暴露List的内部实现细节。