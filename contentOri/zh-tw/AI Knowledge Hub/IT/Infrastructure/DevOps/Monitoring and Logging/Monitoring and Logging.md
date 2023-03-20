+++
title = "監控和日誌記錄"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Monitoring and Logging的問題
{{< /ask_chatgpt >}}



1. 什麼是監控（monitoring）？在系統維護中的作用是什麼？
2. 為什麼日誌（logging）是關鍵的系統維護工具？日誌能提供哪些信息？
3. 如何在AWS（Amazon Web Services）中建立一個完整的監控和日誌系統？有哪些可以使用的服務？
4. 若是要進行故障排除，哪些監控和日誌工具可以幫助我們找到問題的源頭？
5. 如何保護監控和日誌資料的敏感性？有哪些保護方法和措施？   

## 說明知識
{{< ask_chatgpt >}}
說明Monitoring and Logging並舉例
{{< /ask_chatgpt >}}



Monitoring和Logging是系統管理的相關項目，用於監控和記錄應用程式或系統運行的狀態和數據。 

Monitoring （監控）是持續監視系統或應用程式的運行狀態，以確保它們在生產環境中持續運行和提供符合預期的性能。透過設置警報和通知，監控可以用來批判性地識別故障，以及即時響應事件以保證服務的可用性。

舉例來說，Web應用程式可以進行監控，以確保其在高流量時沒有崩潰。系統管理員可以實時監看CPU、記憶體和網路使用率，也可以監測多個伺服器上的訪問數據，並設置警報以在發生故障時及時通知相關人員。

Logging （記錄）是記錄應用程式或系統的活動詳情，可以用於後續排除錯誤和維護。記錄可以捕獲有意義的事件和資訊，如使用者行為、API請求、鏈接和代碼錯誤。在這些記錄中，還可以存儲與每個事件相關的細節，如數據、時間和敘述。

舉例來說，Web應用程式可以記錄所有訪問數據、頁面加載時間，以及使用者資訊等等。系統管理員可以檢查這些紀錄以分析出應用程式遇到的故障或優化應用程式效能。

總而言之，監控和日誌記錄是主要的系統管理工具，可幫助管理員更好地運行和維護系統，以確保其在生產環境中提供穩定和可靠的性能。   

## 彙總重點
{{< ask_chatgpt >}}
條列Monitoring and Logging的重點
{{< /ask_chatgpt >}}



1. 監控（Monitoring）： 監控是一種監視系統或應用程序的活動的過程。通過監控，您可以瞭解系統的整體運行狀況並快速診斷問題。

2. 日誌記錄（Logging）： 日誌記錄是一種記錄系統或應用程序活動的方法。它可以提供詳細的事件記錄以用於問題診斷和問題的研究。

3. 實時性（Real-time）：監控和日誌記錄需要在實時性方面具備高效性和準確性。

4. 可擴展性（Scalability）：監控和日誌記錄需要在處理大量數據時具備可擴展性。

5. 安全性（Security）：監控和日誌記錄需要確保數據的機密性、完整性和可用性。可以使用加密、訪問控制和審計日誌等安全性措施。

6. 自動化（Automation）：監控和日誌記錄需要自動化以提高效率並最大程度減少人為錯誤。

7. 可視化（Visualization）：通過儀表板和可視化工具可以輕鬆展示監控和日誌記錄的數據以查看趨勢、問題和故障。

8. 告警和通知（Alerting and Notification）：對於問題和故障，需要能夠在適當的時間通知相關人員，並自動報警以加快故障排查和解決。

9. 分析和調試（Analysis and Debugging）：監控和日誌記錄提供了必要的數據進行分析和調試以解決問題和故障。

10. 系統性能（System Performance）：監控和日誌記錄需要評估系統性能以識別瓶頸並進行優化。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Monitoring and Logging的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 當一個應用程式在運行時出現錯誤，請說明如何從日誌檔案中搜尋問題所在。

答案：搜尋與時間相關的關鍵字或錯誤訊息，並尋找包含詳細錯誤訊息的日誌檔案。

2. 請說明有哪些方法可以監測伺服器的資源使用情況。

答案：監測CPU、記憶體、磁碟和網路使用情況、使用系統監控軟體、使用Log分析工具監測CPU負載和記憶體使用情況、使用系統日誌記錄磁碟和網路使用狀況。

3. 請簡述何謂日誌旋轉（Log rotation），以及其用途。

答案：日誌旋轉是一個過程，它將歷史日誌檔案從系統中移除並轉存到特定位置，以便更暢順地管理儲存空間、儲存更多日誌、以及確保日誌檔案不會過大而影響效能。

4. 如何防止Logging出現的資訊遭到竄改？

答案：使用加密方式將日誌資訊儲存、限制日誌檔案的讀取權限、使用數位簽章驗證、定期備份資料以及將日誌存放在安全的網路存取區域等。

5. 請說明何謂模式（pattern）在Logging中的作用。

答案：模式在Logging中用於指定文字的格式和內容，可以非常清晰地說明訊息，引導開發者快速解決問題。例如，它可能包括打印日期、時間、事件優先級、動作、最終修改人員等資訊。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Monitoring and Logging的網路資料
{{< /ask_chatgpt >}}



1. "Introduction to Monitoring and Logging" by Datadog: https://www.datadoghq.com/blog/introduction-to-monitoring-and-logging/

This article provides an overview of the concepts behind monitoring and logging, and why they are important for IT operations. It covers the differences between the two, common tools that are used for both, and best practices for effective monitoring and logging.

2. "Monitor and log Kubernetes clusters with Prometheus and Grafana" by IBM Developer: https://developer.ibm.com/technologies/containers/articles/monitor-and-log-kubernetes-clusters-with-prometheus-and-grafana/

This guide explores monitoring and logging in Kubernetes, a popular container orchestration system. The article introduces Prometheus and Grafana, tools that are used for monitoring and visualizing metrics, and explains how to set them up in a Kubernetes environment.

3. "What is Log Management? How It Works, Best Practices, and More" by Loggly: https://www.loggly.com/ultimate-guide/what-is-log-management/

Loggly's Ultimate Guide to Log Management provides a comprehensive explanation of logging and its importance in IT operations. It covers common uses of logs, how to collect and store them, and best practices for analyzing and acting on log data.

4. "10 best practices for effective logging" by Dynatrace: https://www.dynatrace.com/news/blog/10-best-practices-for-effective-logging/

This article outlines ten best practices for logging in IT operations. It covers topics such as log formatting, filtering, and retention, and provides tips for troubleshooting and debugging with log data.

5. "Monitoring and Logging Best Practices for AWS" by AWS: https://aws.amazon.com/blogs/architecture/monitoring-and-logging-best-practices-for-aws-workloads/

This blog post from AWS provides guidance on monitoring and logging in the context of cloud-based workloads. It covers strategies for collecting and analyzing metrics and logs in AWS, and offers recommendations for tools and services that can help.   
