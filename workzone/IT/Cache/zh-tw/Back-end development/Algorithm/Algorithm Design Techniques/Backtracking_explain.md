

背朝追蹤是解決問題的一種演算法，其中通過在解決方案的過程中回溯並反悔，以避免訪問無用的狀態和降低演算法的複雜性。這個演算法通常在求解組合問題（如最短路徑、旅行商問題、數獨）和搜索問題（如填字、八皇后）時使用。

在背朝追蹤中，我們開始尋找問題的解決方案，並假定解決方案是正確的。如果發現某個部分無法通過限制條件，則回溯到先前考慮的狀態並嘗試其他解決方案。在這個過程中，我們保留了先前的計算和結果，以節省計算成本。

舉一個背朝追蹤演算法的例子，考慮以下問題：從一個起點開始，找到一條到達終點的最短路徑。我們可以使用深度優先搜索來找到路徑，並通過回溯來避免重複搜索相同的狀態。以下是一個簡單的背向搜索算法，在網格上搜索從起點到終點的最短路徑：

```python
def backtrack(start, end, grid, visited):
    # Base case: we have reached the end
    if start == end:
        return [end]
    
    # Check if we already visited this state
    if start in visited:
        return []
    
    visited.add(start)
    
    # Try moving in all directions
    rows, cols = len(grid), len(grid[0])
    x, y = start
    candidates = []
    
    for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
        nx, ny = x + dx, y + dy
        if 0 <= nx < rows and 0 <= ny < cols and grid[nx][ny]:
            candidates.append((nx, ny))
    
    # Try to backtrack using each candidate
    for next_pos in candidates:
        path = backtrack(next_pos, end, grid, visited)
        if path:
            return [start] + path
    
    # No valid path found, so backtrack
    return []
```

使用背向搜索算法，我們可以以線性時間和空間複雜度找到從一個起點到另一個終點的最短路徑。