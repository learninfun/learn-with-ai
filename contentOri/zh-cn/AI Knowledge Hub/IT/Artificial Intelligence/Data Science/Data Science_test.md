1. 試著用 Python 利用 KMeans 分群演算法對鳶尾花資料集進行分群，並計算出每個群組的中心點座標。

答案：下方為 Python 程式碼示範：

```python
from sklearn.cluster import KMeans
from sklearn.datasets import load_iris

#讀入鳶尾花資料集
X = load_iris().data

#使用 KMeans 演算法進行分群（4 組）
kmeans = KMeans(n_clusters=4, random_state=0).fit(X)

#列印出每個群組的中心點座標
print(kmeans.cluster_centers_)
```

2. 假設你已經使用前述的 KMeans 演算法進行分群，並且有了每個資料點所屬的群組標籤，試著使用 Python 計算出每個群組的平均值、中位數和標準差。

答案：下方為 Python 程式碼示範：

```python
import numpy as np
from sklearn.datasets import load_iris
from sklearn.cluster import KMeans

#讀入鳶尾花資料集
X = load_iris().data

#使用 KMeans 演算法進行分群（4 組）
kmeans = KMeans(n_clusters=4, random_state=0).fit(X)

#計算每個群組的平均值
means = [np.mean(X[kmeans.labels_ == i], axis=0) for i in range(kmeans.n_clusters)]

#計算每個群組的中位數
medians = [np.median(X[kmeans.labels_ == i], axis=0) for i in range(kmeans.n_clusters)]

#計算每個群組的標準差
std_dev = [np.std(X[kmeans.labels_ == i], axis=0) for i in range(kmeans.n_clusters)]

#列印出計算結果
print("群組平均值：\n", means)
print("群組中位數：\n", medians)
print("群組標準差：\n", std_dev)
```

3. 試著使用 Python 計算出鐵達尼號資料集中，男性與女性的生存率（以百分比呈現）。

答案：下方為 Python 程式碼示範：

```python
import pandas as pd

#讀入鐵達尼號資料集
titanic = pd.read_csv('titanic.csv')

#計算男性生存率
male_survival = titanic.loc[titanic['Sex'] == 'male', 'Survived'].mean()

#計算女性生存率
female_survival = titanic.loc[titanic['Sex'] == 'female', 'Survived'].mean()

#將結果以百分比呈現
print("男性生存率：{:.2f}%".format(male_survival*100))
print("女性生存率：{:.2f}%".format(female_survival*100))
```

4. 試著使用 Python 計算出鐵達尼號資料集中，船票價格的平均值、中位數、標準差以及最大、最小值。

答案：下方為 Python 程式碼示範：

```python
import pandas as pd

#讀入鐵達尼號資料集
titanic = pd.read_csv('titanic.csv')

#計算票價平均值
mean_fare = titanic['Fare'].mean()

#計算票價中位數
median_fare = titanic['Fare'].median()

#計算票價標準差
std_dev_fare = titanic['Fare'].std()

#取得票價的最大和最小值
min_fare = titanic['Fare'].min()
max_fare = titanic['Fare'].max()

#列印出計算結果
print("票價平均值：", mean_fare)
print("票價中位數：", median_fare)
print("票價標準差：", std_dev_fare)
print("票價最大值：", max_fare)
print("票價最小值：", min_fare)
```

5. 試著使用 Python 計算出美國高中生測驗數據集中，數學分數最高的 10 個男生和 10 個女生的平均分數。

答案：下方為 Python 程式碼示範：

```python
import pandas as pd

#讀入測驗數據集
scores = pd.read_csv('scores.csv')

#選出男生中數學分數最高的 10 名
male_top_10 = scores.loc[scores['gender'] == 'male', 'math score'].nlargest(10)

#選出女生中數學分數最高的 10 名
female_top_10 = scores.loc[scores['gender'] == 'female', 'math score'].nlargest(10)

#計算男生的平均分數
male_mean = male_top_10.mean()

#計算女生的平均分數
female_mean = female_top_10.mean()

#列印出計算結果
print("男生數學分數前十名的平均分數：", male_mean)
print("女生數學分數前十名的平均分數：", female_mean)
```

以上皆為示範題目，請讀者自行思考和練習其他題目。