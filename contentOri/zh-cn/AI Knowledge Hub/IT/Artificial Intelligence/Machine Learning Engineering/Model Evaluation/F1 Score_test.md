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