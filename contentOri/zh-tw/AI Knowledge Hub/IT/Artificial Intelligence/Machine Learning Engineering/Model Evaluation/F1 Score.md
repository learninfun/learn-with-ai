+++
title = "F1得分"
weight = "3"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題F1 Score的中文問題
{{< /ask_chatgpt >}}

1. F1 Score是什麼，有什麼作用？
2. F1 Score的計算公式是什麼，如何解釋？
3. F1 Score和Accuracy的區別是什麼？
4. F1 Score在解決什麼樣的問題中特別重要？
5. 如何利用F1 Score來評估機器學習模型的表現？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明F1 Score並舉例
{{< /ask_chatgpt >}}

F1 Score是一個綜合評估模型準確率和召回率的指標，它是精度和召回率的調和平均值。F1 Score值介於0和1之間，值越接近1表示模型的準確性和召回率越高。

例如，假設有一個二元分類問題，假陽性(FP)代表預測某個樣本為陽性，但實際上是陰性；假陰性(FN)代表預測某個樣本為陰性，但實際上是陽性；真陰性(TN)代表預測某個樣本為陰性，實際上也是陰性；真陽性(TP)代表預測某個樣本為陽性，實際上也是陽性。假設一個分類器對一個資料集進行預測，得到以下的confusion matrix：

|                      | Predicted Positive | Predicted Negative |
| -------------------- | ------------------ | ------------------ |
| **Actual Positive**  | TP = 80            | FN = 20            |
| **Actual Negative**  | FP = 10            | TN = 90            |

根據上述的confusion matrix，可以計算出模型的精度和召回率，其中精度= TP / (TP+FP)，而召回率= TP / (TP+FN)。F1 Score則可以通過以下公式計算：F1 Score = 2 x (精度x召回率)/(精度+召回率)。在這個例子中，精度=80 / (80+10)=0.89，召回率=80 / (80+20)=0.80。因此，F1 Score = 2 x (0.89x0.80) / (0.89+0.80) = 0.84。這意味著模型的準確度和召回率都很高，整體表現較為良好。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列F1 Score的重點
{{< /ask_chatgpt >}}

1. F1 Score是評估模型預測結果準確性的指標。
2. F1 Score是Precision和Recall的調和平均數，綜合考慮了真陽性、假陽性和假陰性。
3. F1 Score的取值範圍在0和1之間，值越高表示模型的預測結果越準確。
4. F1 Score在不同的問題和資料集中具有不同的意義和應用。
5. F1 Score需要根據實際問題和模型的需求來進行調參。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題F1 Score的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}

1. 計算二元分類問題中的F1 Score，若正例數為80，負例數為20，True Positive為60，False Positive為10，False Negative為20，則F1 Score為多少？
答案：F1 Score = 2 x (Precision x Recall) / (Precision + Recall) = 2 x (60 / 70 x 60 / 80) / (60 / 70 + 60 / 80) ≈ 0.7895

2. 計算多元分類問題中的weighted F1 Score，若有5個分類，每個分類的預測數量和實際數量如下表所示，則weighted F1 Score為多少？

|           | Class 1 | Class 2 | Class 3 | Class 4 | Class 5 |
|:---------:|:-------:|:-------:|:-------:|:-------:|:-------:|
| 預測 Class 1 |    25   |    5    |    5    |    0    |    0    |
| 預測 Class 2 |    10   |    35   |    5    |    5    |    0    |
| 預測 Class 3 |    5    |    5    |    50   |    0    |    5    |
| 預測 Class 4 |    0    |    5    |    0    |    40   |    5    |
| 預測 Class 5 |    0    |    0    |    5    |    5    |    35   |

答案：對每個類別計算Precision、Recall和F1 Score，並加權平均。例如，對於Class 1，Precision = 25 / (25 + 10 + 5 + 0 + 0) ≈ 0.625，Recall = 25 / (25 + 5 + 5 + 0 + 0) ≈ 0.714，F1 Score = 2 * (0.625 * 0.714) / (0.625 + 0.714) ≈ 0.667。經過計算，weighted F1 Score ≈ 0.685

3. 若有一個二元分類模型，其中True Positive Rate (TPR) = 0.8，False Positive Rate (FPR) = 0.1，則該模型的F1 Score為多少？
答案：由TPR和FPR可求出Precision和Recall，Precision = TP / (TP + FP) = TPR / (TPR + FPR - 1) ≈ 0.8 / 0.7 ≈ 1.143，Recall = TPR = 0.8，F1 Score = 2 * (Precision * Recall) / (Precision + Recall) ≈ 0.848

4. 若一個模型的Precision和Recall都等於0.9，則其F1 Score為多少？
答案：F1 Score = 2 * (Precision * Recall) / (Precision + Recall) = 2 * (0.9 * 0.9) / (0.9 + 0.9) ≈ 0.9

5. 若有一個二元分類問題，其中正例數量為70，負例數量為30，若該問題有兩個不同的模型，其中模型A的Precision為0.8，Recall為0.7，模型B的Precision為0.7，Recall為0.8，則哪個模型的F1 Score更高？
答案：模型A的F1 Score = 2 * (0.8 * 0.7) / (0.8 + 0.7) ≈ 0.7619，模型B的F1 Score = 2 * (0.7 * 0.8) / (0.7 + 0.8) ≈ 0.7619。由計算可知，兩個模型的F1 Score相同。   

