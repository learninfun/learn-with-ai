1. Generator是Python中的一種特殊類型，可以用來創建一個可以迭代的對象。
2. 使用yield關鍵字可以使函數返回一個Generator對象。
3. Generator對象可以在迭代時生成一系列的值，而不是一次性返回所有的值。
4. 與列表相比，Generator的主要優勢是可以減少內存的使用，因為只有在需要時才會生成值。
5. 在使用Generator時，可以使用for循環或者next()函數來遍歷生成的值。
6. Generator可以用來處理大型數據集，或者需要逐步計算的問題，例如計算斐波那契數列。
7. Generator的yield關鍵字可以有多個值，每次產生的值可以是一個元組、列表或字典等集合型別。
8. 一個Generator函數可以調用另一個Generator函數，並使用yield from語法遍歷產生的值。
9. 在使用Generator時，注意避免重復使用同一個Generator對象，否則會導致值的重複生成。
10. Generator是Python中較為高級的概念，但使用起來非常靈活和方便，可以大大提升代碼的可讀性和效率。