

State Pattern的重點如下：

1. 狀態對象：定義不同狀態下的行為和屬性。

2. 狀態接口：定義狀態的行為方法。

3. 上下文對象：持有狀態對象，根據不同狀態調用對應的方法。

4. 狀態轉換：每個狀態都有可能轉換到其他狀態。

5. 繼承：使用繼承可以減少重複代碼並提高代碼複用性。

6. 聚合：使用聚合可以實現更靈活的狀態轉換。

7. 可拓展性：可以輕鬆增加新的狀態和行為，而不需要修改現有的代碼。

8. 測試和維護：狀態模式讓代碼更具可讀性和可維護性，並使測試更容易。