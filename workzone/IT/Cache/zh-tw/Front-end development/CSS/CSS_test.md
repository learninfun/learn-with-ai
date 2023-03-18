

1. 如何讓一個div置中又靠下，且不知道高度?

```
答案： 
div {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, 50%);
}
```

2. 如何讓一個背景顏色區塊產生光圈效果?

```
答案：
div {
  background-color: #333;
  box-shadow: 0 0 20px #fff;
}
```

3. 如何讓一個圓圈鏡像成為一個半圓?

```
答案：
div {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  overflow: hidden;
}
div:after {
  content: "";
  display: block;
  width: 100px;
  height: 50px;
  background-color: #333;
  border-radius: 0 0 50% 50%;
  transform: translateX(-50%);
}
```

4. 如何讓一個圖片hover放大?

```
答案：
div:hover img {
  transform: scale(1.2);
  transition: all 0.2s ease-in-out;
}
```

5. 如何讓一個文字保持在固定位置，而不隨著滾動而上下偏移?

```
答案：
div {
  position: fixed;
  top: 50px;
  left: 50px;
}
```