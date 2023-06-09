1. 狀態(State)：系統可能處於的不同狀態
2. 行動(Action)：系統可執行的各種行動
3. 狀態轉移函數(Transition Function)：描述系統如何由一個狀態轉移到另一個狀態，並受行動影響
4. 即時獎勵函數(Reward Function)：描述當系統處於某個狀態且採取某個行動時，獲得的即時獎勵值
5. 持續時間收益(Return)：在一個序列操作中，按照當前策略採取一系列行動後，獲得的總收益
6. 策略(Policy)：系統在各個狀態下，選擇執行哪個行動的策略
7. 優化問題(Optimization Problem)：在已知系統的狀態轉移函數、即時獎勵函數和策略的情況下，如何設計最優策略使得持續時間收益最大化
8. 廣義策略迭代(Generalized Policy Iteration)：以價值函數(Value Function)和策略迭代(PI)交替進行的一種模型學習算法
9. Q-Learning算法：一種不需要事先知道狀態轉移函數的基於Q值學習法則的強化學習算法
10. 蒙特卡洛方法(Monte Carlo Method)：直接從經驗中學習策略的一種方法，通過對環境進行多次仿真實驗，獲得期望收益和狀態轉移函數等信息，進而學習最優策略