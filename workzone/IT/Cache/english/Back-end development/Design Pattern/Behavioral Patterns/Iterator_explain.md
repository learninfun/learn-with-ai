

Iterator is a design pattern that provides a way to traverse a collection of objects sequentially without exposing its underlying representation. It allows you to access objects without knowing how they are stored, which is useful when you need to iterate over complex data structures (e.g., trees or graphs).

In Java, an iterator is an interface that defines methods for accessing elements of a collection. It provides three main methods:

- hasNext(): Returns true if there are more elements in the collection.
- next(): Returns the next element in the collection.
- remove(): Removes the last element returned by the iterator (optional operation).

Here is an example of how to use an iterator to iterate over an ArrayList:

```
import java.util.ArrayList;
import java.util.Iterator;

public class IteratorExample {
    public static void main(String[] args) {
        ArrayList<String> fruits = new ArrayList<>();
        fruits.add("apple");
        fruits.add("banana");
        fruits.add("cherry");
        
        Iterator<String> it = fruits.iterator();
        while (it.hasNext()) {
            String fruit = it.next();
            System.out.println(fruit);
        }
    }
}
```

This code creates an ArrayList of fruits and then creates an iterator for the collection using the `iterator()` method. It then uses a `while` loop to iterate over the collection and print each element using the `next()` method. The loop exits when there are no more elements in the collection (determined by `hasNext()` returning false).