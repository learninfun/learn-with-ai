

- Saga Pattern可以用來協調分散的事務，保證系統在失敗的情況下能夠恢復正常運作
- 單獨的每個事務都是幂等的，即在重複執行時不會改變系統的狀態
- Saga是一個長時間的交互過程，它管理一系列相關的事務
- Saga Pattern具有不可撤銷性，一旦一個事務完成，它就無法被撤回
- Saga模式的實現方式有兩種： Choreography-based saga和Orchestration-based saga