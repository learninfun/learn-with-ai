Recurrent Neural Networks (RNNs)是一種神經網路的架構，最初被創造用來處理序列的資料，例如語言、音訊及影像等。

RNNs的特點是有一個額外的網路層，被稱為hidden layer。這個hidden layer，跟一般神經網路的hidden layer不一樣的地方在於：每個hidden node都會接收上一個時間點的hidden node的訊息。換句話說，hidden layer具有時序性，並且可以在處理序列資料時記憶先前預測的結果。這個過程稱為記憶體(memory)或循環性(recurrence)。

以下是一個簡單的RNNs例子：我們要用RNNs創建一個模型，可以翻譯英文為法文的問候語。假設我們有以下的英語問候語：

"Hello, how are you?"
"Good morning, how's your day going?"
"Hi there, what's up?"

RNNs模型會被訓練，輸入每個單詞的資料，並且在最後輸出法文的翻譯。在這過程中，RNNs會從前一個word node儲存目前的資訊，並且隨著時間記錄預測翻譯的過程。這樣，RNNs可以在翻譯每個單詞時，考慮到整個語句上下文，並進行更準確的翻譯。