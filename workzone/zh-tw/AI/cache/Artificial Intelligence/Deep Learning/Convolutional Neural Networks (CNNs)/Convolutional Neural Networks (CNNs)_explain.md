卷積神經網絡（CNNs）是一種神經網絡架構，特別設計用於處理具有類似網格結構的數據，例如圖像和聲音。CNNs可以學習識別圖像中的對象、人臉、文本等。它們利用卷積層來檢測特徵，然後使用池化層來縮小特徵映射的大小，最終使用全連接層來生成輸出。

舉一個簡單的例子，假設我們有一張手寫數字的圖像，我們想使用CNNs來識別數字。我們可以將圖像的像素表示為輸入層，然後將它們送入卷積層。卷積層將包含許多卷積核，每個卷積核都尋找可以識別的特定模式。例如，卷積層中的某個卷積核可能尋找相鄰像素之間的邊緣或者線條。池化層將減小特徵映射的尺寸，防止 CNN 對噪聲或不相關的數據過分關注。最終的全連接層將輸出數字的概率分佈，然後我們可以使用argmax函數來找到最有可能的數字。