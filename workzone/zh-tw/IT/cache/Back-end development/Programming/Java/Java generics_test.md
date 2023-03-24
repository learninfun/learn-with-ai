

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