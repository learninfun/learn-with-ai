+++
title = "F1得分"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题F1 Score的中文问题
{{< /ask_chatgpt >}}

1. F1 Score是什么，有什么作用？
2. F1 Score的计算公式是什么，如何解释？
3. F1 Score和Accuracy的区别是什么？
4. F1 Score在解决什么样的问题中特别重要？
5. 如何利用F1 Score来评估机器学习模型的表现？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明F1 Score并举例
{{< /ask_chatgpt >}}

F1 Score是一个综合评估模型准确率和召回率的指标，它是精度和召回率的调和平均值。F1 Score值介于0和1之间，值越接近1表示模型的准确性和召回率越高。

例如，假设有一个二元分类问题，假阳性(FP)代表预测某个样本为阳性，但实际上是阴性；假阴性(FN)代表预测某个样本为阴性，但实际上是阳性；真阴性(TN)代表预测某个样本为阴性，实际上也是阴性；真阳性(TP)代表预测某个样本为阳性，实际上也是阳性。假设一个分类器对一个资料集进行预测，得到以下的confusion matrix：

|                      | Predicted Positive | Predicted Negative |
| -------------------- | ------------------ | ------------------ |
| **Actual Positive**  | TP = 80            | FN = 20            |
| **Actual Negative**  | FP = 10            | TN = 90            |

根据上述的confusion matrix，可以计算出模型的精度和召回率，其中精度= TP / (TP+FP)，而召回率= TP / (TP+FN)。F1 Score则可以通过以下公式计算：F1 Score = 2 x (精度x召回率)/(精度+召回率)。在这个例子中，精度=80 / (80+10)=0.89，召回率=80 / (80+20)=0.80。因此，F1 Score = 2 x (0.89x0.80) / (0.89+0.80) = 0.84。这意味着模型的准确度和召回率都很高，整体表现较为良好。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列F1 Score的重点
{{< /ask_chatgpt >}}

1. F1 Score是評估模型預測結果準確性的指標。
2. F1 Score是Precision和Recall的調和平均數，綜合考慮了真陽性、假陽性和假陰性。
3. F1 Score的取值範圍在0和1之間，值越高表示模型的預測結果越準確。
4. F1 Score在不同的問題和資料集中具有不同的意義和應用。
5. F1 Score需要根據實際問題和模型的需求來進行調參。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题F1 Score的中等难度问题，并在后面列出答案
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

