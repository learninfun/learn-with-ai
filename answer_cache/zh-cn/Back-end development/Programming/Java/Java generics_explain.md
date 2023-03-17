

Java generics是一個Java語言的特性，它允許使用者在編譯時定義類型通用代碼，而不是將泛型化代碼應用於編譯時期。

舉例來說，如果一個程式需要一個集合來存儲整數，程式可以使用Integer類型的ArrayList，如下所示：

```java
ArrayList<Integer> ints = new ArrayList<>();
ints.add(1);
ints.add(2);
ints.add(3);
```

這裡，ArrayList的類型引數是Integer。這樣一來，不僅程式可以證明ArrayList只包含Integer物件，也可以確保如果試圖添加其他類型的元素，編譯器會給出一個錯誤。

Java generics還允許創建泛型方法。以下是一個返回任意型別陣列中最大元素的方法：

```java
public static <T extends Comparable<T>> T max(T[] arr) {
    T max = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i].compareTo(max) > 0) {
            max = arr[i];
        }
    }
    return max;
}
```

在這裡，<T extends Comparable<T>>是類型參數，它通過限制類型T為Comparable<T>的子類來定義可接受的類型。此方法可以接受任何Comparable<T>類型的陣列，例如Integer，Double或String。