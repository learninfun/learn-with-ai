

1. 如何使用Java reflection取得某個class的所有方法名稱？

```
Class<?> clazz = MyClass.class;
Method[] methods = clazz.getDeclaredMethods();
for (Method method : methods) {
    System.out.println(method.getName());
}
```

2. 如何使用Java reflection取得某個class的所有變數名稱？

```
Class<?> clazz = MyClass.class;
Field[] fields = clazz.getDeclaredFields();
for (Field field : fields) {
    System.out.println(field.getName());
}
```

3. 如何使用Java reflection動態新增一個物件實例？

```
Class<?> clazz = MyClass.class;
Object newInstance = clazz.getDeclaredConstructor().newInstance();
```

4. 如何使用Java reflection設定某個物件實例的屬性值？

```
MyClass obj = new MyClass();
Class<?> clazz = MyClass.class;
Field field = clazz.getDeclaredField("fieldName");
field.setAccessible(true);
field.set(obj, "newValue");
```

5. 如何使用Java reflection呼叫某個物件實例的方法？

```
MyClass obj = new MyClass();
Class<?> clazz = MyClass.class;
Method method = clazz.getDeclaredMethod("methodName", String.class);
method.invoke(obj, "parameterValue");
```

答案僅供參考，實際實現方式可能有所不同。