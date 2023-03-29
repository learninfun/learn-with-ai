1. 如何在Kubernetes中實現水平擴展？
2. 如何設置Kubernetes中的滾動升級（Rolling Updates）？
3. 在Kubernetes Pod中，如何使用init容器（init containers）？
4. 如何在Kubernetes中設置自動水平擴展（Auto Scaling）？
5. 在Kubernetes中如何設置服務儲存器（Service Discovery）？

答案：
1. 在Kubernetes中實現水平擴展可通過設置相應的Pod Replica Set規則，如標籤選擇器（Label Selectors）和容器健康檢查（Liveness Probes）等。
2. 可以使用Kubernetes Deployment對象進行滾動升級，並設置相關參數，如更新策略（update strategy）和最大失敗次數（max unavailable）等。
3. 在Kubernetes Pod中，通過設置init容器來實現初始化操作，如下載資源、設置環境變數和執行命令等。
4. 在Kubernetes中設置自動水平擴展通過配置Pod Auto Scaler對象，藉助CPU利用率和RAM利用率等指標來調整Pod實例數量。
5. 服務儲存器（Service Discovery）在Kubernetes中可使用Kubernetes Service對象來實現，通過設置相應的標籤選擇器和端口設置等參數，來實現服務發現和通訊。