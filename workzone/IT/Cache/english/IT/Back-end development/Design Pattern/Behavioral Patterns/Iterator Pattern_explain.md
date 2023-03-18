

Iterator pattern is a design pattern that provides a way to traverse the elements of a collection without exposing its underlying implementation. It provides two important interfaces - an Iterator interface that defines the operations to iterate over the collection, and a Collection interface that defines the contract for a collection.

For example, let's say we have an array of integers and we want to iterate over it. One way to do it is like this:

```
int[] nums = {1, 2, 3, 4};
for(int i=0; i<nums.length; i++) {
    System.out.println(nums[i]);
}
```

But what if the collection is not an array, but a linked list or a tree? The implementation of the iterator would be different for each collection, but the client code that uses the iterator would remain the same. That's where the Iterator pattern comes in.

Let's define an interface for our iterator:

```
public interface Iterator<T> {
    boolean hasNext();
    T next();
}
```

And an interface for our collection:

```
public interface Collection<T> {
    Iterator<T> iterator();
}
```

Now we can implement our array iterator like this:

```
public class ArrayIterator<T> implements Iterator<T> {
    private T[] array;
    private int index = 0;
    
    public ArrayIterator(T[] array) {
        this.array = array;
    }
    
    public boolean hasNext() {
        return index < array.length;
    }
    
    public T next() {
        return array[index++];
    }
}
```

And our client code would look like this:

```
int[] nums = {1, 2, 3, 4};
Iterator<Integer> it = new ArrayIterator<Integer>(nums);
while(it.hasNext()) {
    System.out.println(it.next());
}
```

Now if we have a linked list or a tree collection, we can implement their respective iterators using the same `Iterator` interface and our client code would remain the same.