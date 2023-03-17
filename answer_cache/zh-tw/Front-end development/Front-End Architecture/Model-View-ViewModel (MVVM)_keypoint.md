

以下是MVVM架構重點：

1. Model：應用程式的資料模型，處理資料邏輯與存取資料的方法，並且不會與View或ViewModel互動。

2. View：使用者介面，可看成是呈現Model的內容，不直接向Model或任何物件發出指令。

3. ViewModel：連結View與Model的重要角色，擔負著兩方面的轉換：資料轉換(activity to model)及操作轉換(view to model)。例如，當View有任何事件(按鈕按下、輸入資料等)並通知ViewModel時，ViewModel便可依照事件種類將請求轉換為Model可處理的資訊，並返回處理結果到View上，以此牽引整個MVVM的工作流程。

4. 兩者的綁定：綁定是讓View和ViewModel之間傳遞資料的方法之一，它是一種方式，讓變數狀態隨時同步，而不需撰寫其他邏輯程式。例如，當model資料修改時，View會被通知資料變動，並自動更新自己所顯示的內容。

5. 重用性：因為View和ViewModel解耦且彼此分離，所以顯示模版和顯示資訊都容易復用。

6. 測試性：雖然View和ViewModel都是需要人工測試，但是Model則是可以使用單元測試學派的測試方式，讓整體架構更為健全。

7. 編寫方式：在MVVM中，View主要是需要XAML語言進行拖拉式設計，也可直接刻畫視界設計，ViewModel更多是以程式碼的方式撰寫，讓封裝、邏輯分層與非同步處理更為容易。