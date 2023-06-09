

- Prim's Algorithm是一種用來找尋最小生成樹的演算法。
- 演算法通常需要一個起始點，從該起始點開始建立最小生成樹。
- 演算法維護兩個數據結構，分別是集合S和邊集合T。S表示已經在最小生成樹上的點，而T則表示已經在最小生成樹上的邊。
- 首先將起始點加入集合S中。
- 從集合S中的所有點開始找尋最小生成樹的下一個邊，找到其中最小權重的邊並加入邊集合T中。
- 將邊的兩個端點中未被加入集合S中的點加入集合S中，重複上述過程，直到所有點都在集合S中。
- Prim's Algorithm的時間複雜度為O(E log V)，其中E是邊的數量，V是點的數量。