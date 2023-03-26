1. Cross-Validation的目的在于评估机器学习模型的性能，并尽可能地减少过拟合或欠拟合的情况。

2. Cross-Validation通常包含以下步骤：将资料集分为训练集和测试集、使用训练集建立模型、使用测试集对模型进行评估。

3. 传统的Cross-Validation方法包括K-Fold Cross-Validation和Leave-One-Out Cross-Validation。K-Fold Cross-Validation将资料集分为K份，每次选取其中一份作为测试集，其余K-1份作为训练集；Leave-One-Out Cross-Validation则是每次只选择一笔资料作为测试集，其余资料作为训练集。

4. 在Cross-Validation中，需注意测试集和训练集应该是相互独立的，而且测试集应尽可能地代表整个资料集。

5. Cross-Validation的评估指标通常包括准确率、F1-score、ROC曲线等。

6. Cross-Validation的一个重要应用是选择最佳模型，即通过比较不同模型的Cross-Validation结果，选择性能最好的模型作为最终模型。

7. Cross-Validation的限制包括时间、计算资源等成本，以及对于资料集的大小和分布等限制。因此，在实际应用中，需根据实际情况选择适当的Cross-Validation方法。