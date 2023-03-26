1. Gated Recurrent Unit (GRU)是一種循環神經網絡 (Recurrent Neural Network, RNN) 的一種變種，用於處理序列式資料。

2. 透過 Gate Mechanism 控制每個單元 (unit) 的訊息流動，可以有效地解決梯度消失 (vanishing gradient) 和梯度爆炸 (exploding gradient) 的問題。

3. GRU 的單元包含了更新門 (update gate)、重置門 (reset gate) 和新的候選隱藏狀態 (candidate hidden state)。

4. 更新門控制了過去隱藏狀態的影響程度，重置門控制了新輸入和過去隱藏狀態的交互作用，候選隱藏狀態決定了當前時刻的隱藏狀態。

5. GRU 通常用於語音識別、機器翻譯、文本生成等自然語言處理任務中。

6. 可以通過調整 GRU 單元的參數和架構的設計，提高模型的性能。