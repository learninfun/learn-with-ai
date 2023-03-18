

Layouts（佈局）是指在Android中用來定義UI界面的框架或模板。它定義了UI元素的位置、大小、顏色等屬性，可以幫助開發者更方便地設計和實現應用程序的UI界面。

Android中常用的布局有幾種，主要包括：

1. LinearLayout：水平或垂直排列子視圖，適用於簡單的UI佈局。

2. RelativeLayout：基於視圖之間相對位置定位子視圖。

3. FrameLayout：用於將子視圖疊加在一起，常用於實現動畫效果。

4. ConstraintLayout：通過將視圖之間的關係定義為約束來佈局。

以下是LinearLayout和RelativeLayout的示例代碼：

```xml
<!--LinearLayout示例-->
<LinearLayout
    android:layout_width="match_parent"
    android:layout_height="wrap_content"
    android:orientation="horizontal"
    android:gravity="center">

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="Hello,"
        android:textSize="18sp"/>

    <TextView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="World!"
        android:textSize="18sp"/>

</LinearLayout>

<!--RelativeLayout示例-->
<RelativeLayout
    android:layout_width="match_parent"
    android:layout_height="wrap_content">

    <Button
        android:id="@+id/btn_ok"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:text="OK"
        android:layout_alignParentRight="true"/>

    <EditText
        android:id="@+id/et_input"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:layout_toLeftOf="@id/btn_ok"/>

</RelativeLayout>
```

上面的LinearLayout示例展示了如何將兩個TextView水平排列在一起，並使用gravity屬性將它們置於中心。RelativeLayout示例展示了如何將一個Button放置於EditText的右側。

以上就是Layouts及其示例的簡單介紹。在實際開發中，開發者應根據實際需求來選擇適當的佈局方式。