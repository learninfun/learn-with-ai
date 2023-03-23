+++
title = "Java泛型"
weight = "9"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Java generics的问题
{{< /ask_chatgpt >}}



1. 什麼是Java的泛型？

2. 泛型的好處是什麼？

3. 泛型的限制是什麼？

4. 如何在Java中使用泛型？

5. 使用泛型時，如何避免運行時出現ClassCastException？   

## 说明知识
{{< ask_chatgpt >}}
说明Java generics并举例
{{< /ask_chatgpt >}}



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

## 汇总重点
{{< ask_chatgpt >}}
条列Java generics的重点
{{< /ask_chatgpt >}}



1. 泛型的概念：泛型是一种让类或方法能够处理一组不同类型的数据的机制。

2. 泛型类：使用泛型的类，可以根据需要存储不同类型的对象。

3. 泛型方法：使用泛型的方法，可以在方法调用时指定参数类型，也可以在方法返回值中使用泛型。

4. 类型参数：使用类型参数可以声明泛型类或方法中使用的类型。

5. 通配符：在泛型中使用的通配符有两种，一种是 "?" 表示未知类型参数，一种是 "? extends T" 表示类型参数必须是 T 或 T 的子类。

6. 类型擦除：Java 泛型实际上是一种编译期特性，编译器会在编译时将泛型类型擦除成相应的实际类型。

7. 桥接方法：类型擦除会导致泛型继承中出现一些问题，为了解决这些问题，编译器会自动生成桥接方法。

8. 通配符捕获：使用通配符捕获可以在泛型方法中使用泛型类型，而不需要在方法声明中指定具体的泛型类型。

9. 泛型边界：使用泛型边界可以限制泛型类型的范围，例如可以指定类型参数必须是某个类的子类。

10. 泛型通信：Java 泛型允许在泛型类或方法中通过泛型类型参数进行通信，而不需要进行硬编码。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Java generics的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



問題1:

請寫一個泛型方法，在給定的List中，找到第二大的元素。

答案:

```java
public static <T extends Comparable<? super T>> T findSecondLargest(List<T> list) {
    if (list == null || list.size() < 2) {
        return null;
    }

    T max = list.get(0);
    T secondMax = null;

    for (int i = 1; i < list.size(); i++) {
        T current = list.get(i);
        if (current.compareTo(max) > 0) {
            secondMax = max;
            max = current;
        } else if (secondMax == null || current.compareTo(secondMax) > 0) {
            secondMax = current;
        }
    }

    return secondMax;
}
```

問題2:

請寫一個泛型方法，在給定的Map中，查找value最小的key。

答案:

```java
public static <K, V extends Comparable<? super V>> K findKeyWithMinValue(Map<K, V> map) {
    if (map == null || map.isEmpty()) {
        return null;
    }

    K result = null;
    V minValue = null;

    for (Map.Entry<K, V> entry : map.entrySet()) {
        if (minValue == null || entry.getValue().compareTo(minValue) < 0) {
            minValue = entry.getValue();
            result = entry.getKey();
        }
    }

    return result;
}
```

問題3:

請寫一個泛型方法，在給定的List中，找到出現次數最多的元素。

答案:

```java
public static <T> T findMostFrequentValue(List<T> list) {
    if (list == null || list.isEmpty()) {
        return null;
    }

    Map<T, Integer> countMap = new HashMap<>();
    T mostFrequentValue = null;
    int maxCount = 0;

    for (T value : list) {
        int count = countMap.getOrDefault(value, 0) + 1;
        countMap.put(value, count);
        if (count > maxCount) {
            mostFrequentValue = value;
            maxCount = count;
        }
    }

    return mostFrequentValue;
}
```

問題4:

請寫一個泛型方法，將給定的List中的所有元素轉換為String後，連接起來返回。

答案:

```java
public static <T> String joinToString(List<T> list, String delimiter) {
    if (list == null || list.isEmpty()) {
        return "";
    }

    StringBuilder sb = new StringBuilder();
    sb.append(list.get(0));

    for (int i = 1; i < list.size(); i++) {
        sb.append(delimiter).append(list.get(i));
    }

    return sb.toString();
}
```

問題5:

请编写一个泛型方法，接受一个由一些表示姓名的字符串和整数数组对组成的序列，并以每个姓名为键，对应年龄为值，返回一个Map<String, Integer>。

答案:

```java
public static <T> Map<String, Integer> createMapFromPairs(List<Pair<String, Integer>> pairs) {
    if (pairs == null || pairs.isEmpty()) {
        return Collections.emptyMap();
    }

    Map<String, Integer> result = new HashMap<>();

    for (Pair<String, Integer> pair : pairs) {
        result.put(pair.getKey(), pair.getValue());
    }

    return result;
}
```   

