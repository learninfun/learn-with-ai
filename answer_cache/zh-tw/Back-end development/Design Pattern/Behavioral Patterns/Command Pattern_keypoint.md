

1. Command Pattern是一種行為型設計模式。
2. Command Pattern讓你能夠將特定操作的信息從其執行中分離出來，並封裝成一個獨立的物件中。
3. Command Pattern讓你可以將特定的操作序列化、日誌記錄、取消或延遲其執行。
4. Command Pattern中的關鍵角色有Command、Invoker、Receiver和Client。
5. Command是行為請求的選擇接口，Invoker引用並調用命令，Receiver實現命令和最終操作，Client則創建一個具體的Command對象並將其傳遞給Invoker。
6. Command Pattern的優點包括解耦程式碼、易於修改、簡化操作層級和支援撤銷和恢復功能等。
7. Command Pattern的缺點包括生成大量命令物件可能會影響效能、需要額外實現的模式如果實現不好可能會產生更多的問題。