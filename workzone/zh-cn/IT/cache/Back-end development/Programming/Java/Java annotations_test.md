

1. 給定以下的Java annotation，該注釋用於方法上，它的功能是什麼？
```
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface MyAnnotation {
   String value();
}
```
答案：該注釋是一個自定義注釋，用於方法上，可以指定一個字符串值。

2. 給定以下的Java annotation，該注釋用於類上，它的功能是什麼？
```
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.TYPE)
public @interface MyAnnotation {
   String author() default "unknown";
   String date();
}
```
答案：該注釋是一個自定義注釋，用於類上，可以指定類的作者和日期。

3. 定義一個Java注釋，描述當前的方法是否為只讀方法（即不允許對數據進行更改）。
答案：
```
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface ReadOnly {}
```

4. 定義一個Java注釋，描述當前的類是一個單例模式的類。該注釋不能用在任何接口或抽象類上。
答案：
```
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.TYPE)
public @interface Singleton {}
```

5. 定義一個Java注釋，描述當前的方法被調用時，必須在指定時間段內完成，否則將拋出異常。
答案：
```
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface TimeLimit {
   int seconds() default 5;
}
```