

1. 如何使用Java reflection取得某个class的所有方法名称？

```
Class<?> clazz = MyClass.class;
Method[] methods = clazz.getDeclaredMethods();
for (Method method : methods) {
    System.out.println(method.getName());
}
```

2. 如何使用Java reflection取得某个class的所有变数名称？

```
Class<?> clazz = MyClass.class;
Field[] fields = clazz.getDeclaredFields();
for (Field field : fields) {
    System.out.println(field.getName());
}
```

3. 如何使用Java reflection动态新增一个物件实例？

```
Class<?> clazz = MyClass.class;
Object newInstance = clazz.getDeclaredConstructor().newInstance();
```

4. 如何使用Java reflection设定某个物件实例的属性值？

```
MyClass obj = new MyClass();
Class<?> clazz = MyClass.class;
Field field = clazz.getDeclaredField("fieldName");
field.setAccessible(true);
field.set(obj, "newValue");
```

5. 如何使用Java reflection呼叫某个物件实例的方法？

```
MyClass obj = new MyClass();
Class<?> clazz = MyClass.class;
Method method = clazz.getDeclaredMethod("methodName", String.class);
method.invoke(obj, "parameterValue");
```

答案仅供参考，实际实现方式可能有所不同。