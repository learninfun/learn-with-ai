梯度提升（Gradient Boosting）是一種集成學習法，它通過集成多個模型來提升模型性能。Gradient Boosting通常基於決策樹模型，同時也能應用於其他模型上。

Gradient Boosting的主要思想是通過序列化地擬合模型，將前一個模型失效的樣本加權以提高後續模型的性能。這樣一來，後續模型會更加關注前一個模型中錯誤的樣本，儘量更好地捕捉這些樣本的特徵。

舉例來說，假設我們要預測房價，我們可以建立一個基礎模型，比如簡單的線性回歸。接下來，我們可以基於殘差（即真實值與預測值之差）學習一個決策樹模型，以提高我們的預測能力。然後，我們可以基於新的殘差再次學習一個決策樹模型，以進一步提高性能。最終，我們可以將所有的模型集成起來，得到一個更強大的預測模型。這個模型會考慮所有模型的預測結果，並使用加權平均的方式得出最終的預測。

Gradient Boosting在許多領域中取得了驚人的成功，比如網絡廣告和推薦系統。在這些應用中，Gradient Boosting的主要優勢在於它能夠處理大量的非線性特徵，並產生高精度的預測結果。