1. 什麼是Confusion matrix？如何使用Confusion matrix評估模型效能？
2. 什麼是ROC曲線？如何使用ROC曲線評估二元分類模型效能？
3. 過擬合與欠擬合分別代表什麼？如何使用Validation curve判斷模型是否過擬合或欠擬合？
4. 什麼是Bias-Variance trade-off？如何使用Learning curve找出模型的最佳平衡點？
5. 什麼是Cross-validation？如何使用Cross-validation找出最佳的模型超參數？

答案：
1. Confusion matrix是用來評估二元分類模型效能的方法，將實際類別與預測類別每個部分分別放於四個方格中，可計算出準確率、精確率、召回率、F1 score等指標。
2. ROC曲線是以偽陽性率(FPR)為X軸、真陽性率(TPR)為Y軸所繪製的曲線，可以用來評估二元分類模型的效能。曲線下方面積(AUC)越高，模型效能越好。
3. 過擬合指模型在訓練集上表現過於優秀，但在測試集上表現較差；欠擬合指模型無法在訓練集上表現良好，因此在測試集上也無法表現良好。使用Validation curve可以繪製出不同超參數下的訓練集與測試集的得分對比圖，以判斷模型是否過擬合或欠擬合。
4. Bias-Variance trade-off指的是在模型表現優化過程中，Bias誤差與Variance誤差之間需要平衡。Learning curve可以繪製出不同訓練集大小下的訓練集與測試集的得分對比圖，以判斷模型是否需要更多的數據來降低Bias誤差或更換更複雜的模型來減少Variance誤差。
5. Cross-validation是一種交叉驗證的方法，可幫助找出最佳的模型超參數。使用K-Fold交叉驗證，將數據集分為K個部分，模型訓練時取K-1份作為訓練集，取1份作為驗證集，重複K次，求得K次得分的平均值。最佳的模型超參數為能夠使得得分最高或者方差最小的超參數。