多項式回歸 (Polynomial Regression) 是一個利用多項式進行回歸分析的方法。在進行線性回歸時，常常會遇到一些資料不符合線性關係的情況，這時就可以採用多項式回歸來進行分析。

多項式回歸的基本概念是將自變量 $x$ 進行多次方項的拓展，例如 $x$ 的平方、立方等，然後將這些項作為獨立變量進行回歸分析。通常，多項式回歸所使用的項數可以由資料集的形狀和特徵來決定。

例如，我們有一組數據，其中自變量 $x$ 和應變量 $y$ 的關係不符合線性關係，如下圖所示：

![image.png](attachment:image.png)

在這種情況下，我們可以使用多項式回歸來建立一個多項式方程，以逼近資料的分佈情況。在本例中，我們選擇二次多項式，使用以下方程來進行回歸分析：

$$
y = \beta_0 + \beta_1 x + \beta_2 x^2
$$

其中 $\beta_0$、$\beta_1$ 和 $\beta_2$ 分別是模型的截距項和斜率項，以及二次項的係數。

透過這個模型，我們可以得到以下拟合圖形：

![image-2.png](attachment:image-2.png)

從圖中可以看出，多項式回歸模型可以很好地拟合數據集，有效地反映出資料的分佈情況。