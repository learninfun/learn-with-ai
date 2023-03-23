+++
title = "廣度優先搜尋"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Breadth-First Search的中文問題
{{< /ask_chatgpt >}}



1. 在一張無向圖中，如何找出兩個給定節點之間的最短路徑？
2. 在一個網格中，如何找到從起點到終點的最短路徑，其中障礙物可能會阻擋移動？
3. 如何在一個社交網絡中找到關於某個人的直接或間接聯繫，並計算出與他們之間的距離？
4. 如何從起點到盡可能多的目標節點，而不必找到最短路徑？
5. 如何在一個巨大的圖形資料庫中找到與搜索條件最相符的所有圖形，並以最短路徑顯示它們之間的關係？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Breadth-First Search並舉例
{{< /ask_chatgpt >}}



Breadth-First Search（BFS）是一種搜尋演算法，以廣度優先的方式來遍歷一個圖形或樹狀結構。

其遍歷的順序是逐層往下，也就是先遍歷所有的同一深度節點，再遍歷下一深度的節點。在BFS遍歷中，使用一個隊列來維護已經被廣度遍歷的節點，以便按層訪問下一階段節點時使用。

舉個例子：假設我們有一個有向圖如下圖所示：

<img src="https://i.imgur.com/Vhh9XfO.png" width="200"/>

我們從節點1開始進行BFS遍歷，首先將節點1加入隊列中。接下來，按照節點編號的大小順序，先遍歷節點2和節點3。

<img src="https://i.imgur.com/s8BDGdc.png" width="200"/>

然後，把節點2的相鄰節點4，7加入隊列中，把節點3的相鄰節點5，6加入隊列中。

<img src="https://i.imgur.com/gPXs4L4.png" width="200"/>

再遍歷節點4和節點7，因為它們沒有相鄰節點可加入隊列中，所以直接跳過。

<img src="https://i.imgur.com/yujfiFB.png" width="200"/>

最後，遍歷節點5和節點6，發現節點5有一個相鄰節點8，所以把節點8加入隊列中。遍歷完節點5和節點6，隊列已經空了，此時遍歷結束。

<img src="https://i.imgur.com/iCCvSTN.png" width="200"/>

這個例子中，我們首先訪問距離起點節點最近的節點，然後是次近的節點，然後是更遠的節點。這種BFS的訪問順序可以確保我們在最短時間內訪問到所有節點。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Breadth-First Search的重點
{{< /ask_chatgpt >}}



1. 廣度優先搜尋演算法（BFS）是一種用來搜索圖形或樹形結構的技術。

2. BFS從起始節點開始搜尋，然後逐層擴展，直到達到終點或所有節點都被訪問為止。

3. BFS適用於找到最短路徑和最少操作的問題，因為它保證了先找到的路徑長度最短或者操作最少。

4. BFS使用FIFO（先進先出）佇列來保存待處理的節點，這有助於記錄搜尋順序和計算層次。

5. BFS通常需要使用標記訪問過的節點，以避免重複訪問和死循環。

6. BFS可以用來應對未知的圖形和樹形結構，並且可以與其他搜尋演算法結合使用。

7. BFS的時間複雜度為O(V+E)，其中V是圖形的節點數，E是圖形的邊數。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Breadth-First Search的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 題目： 給定一個連通無向圖，起始節點為S，每個節點都是一個整數，節點間的距離為1。你的任務是從起始節點開始按字典序打印所有節點。輸入的圖是用鄰接矩陣表示的，值為1表示兩個節點相連，0表示不連通。 同時輸入的起始節點的下標，輸出樣例中「#」表示空格。

答案：

from queue import Queue

def bfs(n, start_node):
    visited = [0] * n
    queue = Queue()
    queue.put(start_node)
    visited[start_node] = 1
    while not queue.empty():
        node = queue.get()
        print(node, end=' ')
        for i in range(n):
            if adj[node][i] == 1 and visited[i] == 0:
                visited[i] = 1
                queue.put(i)

n = int(input())
adj = []
for i in range(n):
    adj.append(list(map(int, input().split())))

start_node = int(input())

bfs(n, start_node)


2. 題目：有n個物品，每種物品都有自己的重量和價格，在限定的總重量內，選出若幹件物品使得物品的總重量不超過總重量且總價格最大。假設背包的容量為C（Capacity），每個物品的重量為w（weight），價值為v（value）。使用BFS求背包問題的最優解。

答案：

from queue import Queue

class Node:
	def __init__(self, level, weight, value, bound):
		self.level = level
		self.weight = weight
		self.value = value
		self.bound = bound

def bound(node, max_weight, n, values, weights):
	if node.weight >= max_weight:
		return 0
	result = node.value
	j = node.level + 1
	total_weight = node.weight
	while j < n and total_weight + weights[j] <= max_weight:
		total_weight += weights[j]
		result += values[j]
		j += 1
	if j < n:
		result += (max_weight - total_weight) * (values[j] / weights[j])
	return result

def bfs(n, max_weight, values, weights):
	queue = Queue()
	root = Node(-1, 0, 0, 0)
	queue.put(root)
	max_value = 0
	while not queue.empty():
		node = queue.get()
		if node.level == -1:
			left = Node(0, 0, 0, 0)
		elif node.level == n - 1:
			continue
		else:
			left = Node(node.level + 1, node.weight + weights[node.level + 1], node.value + values[node.level + 1], 0)
			left.bound = bound(left, max_weight, n, values, weights)
			if left.weight <= max_weight and left.value > max_value:
				max_value = left.value
			if left.bound > max_value:
				queue.put(left)
		right = Node(node.level + 1, node.weight, node.value, 0)
		right.bound = bound(right, max_weight, n, values, weights)
		if right.weight <= max_weight and right.value > max_value:
			max_value = right.value
		if right.bound > max_value:
			queue.put(right)
	return max_value

n = int(input())
max_weight = int(input())
values = list(map(int, input().split()))
weights = list(map(int, input().split()))

max_value = bfs(n, max_weight, values, weights)

print(max_value)


3. 題目：在一個NxN的網格中，1代表通路，0代表牆，從左上方的格子開始走，每次只能向右或向下走，走到右下角的格子後結束，尋找一條最短的路徑。

答案：

from queue import Queue

class Node:
	def __init__(self, row, col, steps):
		self.row = row
		self.col = col
		self.steps = steps

def bfs(n, grid):
	directions = [(1, 0), (0, 1)]
	visited = [[False for _ in range(n)] for _ in range(n)]
	queue = Queue()
	start = Node(0, 0, 0)
	queue.put(start)
	visited[0][0] = True
	while not queue.empty():
		currentNode = queue.get()
		if currentNode.row == n - 1 and currentNode.col == n - 1:
			return currentNode.steps
		for dir in directions:
			newRow = currentNode.row + dir[0]
			newCol = currentNode.col + dir[1]
			
			if newRow >= 0 and newRow < n and newCol >= 0 and newCol < n and grid[newRow][newCol] == 1 and not visited[newRow][newCol]:
				queue.put(Node(newRow, newCol, currentNode.steps + 1))
				visited[newRow][newCol] = True
	return -1

n = int(input())
grid = []
for i in range(n):
	grid.append(list(map(int, input().split())))

print(bfs(n, grid))


4. 題目： 給定一個大小為n的矩陣，每個位置的值代表該位置的魔法值，每次可以進行一次魔法轉換，將以該位置為起始點的行和列值全部加1。求將整個矩陣的魔法值加到X需要進行最少多少次魔法轉換。

答案：

from queue import Queue

class Node:
	def __init__(self, i, j, value, steps):
		self.i = i
		self.j = j
		self.value = value
		self.steps = steps

def bfs(n, grid, x):
	visited_row = [False for _ in range(n)]
	visited_col = [False for _ in range(n)]
	queue = Queue()
	start = Node(0, 0, grid[0][0], 0)
	visited_row[0] = True
	visited_col[0] = True
	queue.put(start)
	count = 0
	while not queue.empty():
		node = queue.get()
		if node.value >= x:
			count = node.steps
			break
		if not visited_row[node.i]:
			for j in range(n):
				newValue = node.value + grid[node.i][j]
				queue.put(Node(node.i, j, newValue, node.steps + 1))
			visited_row[node.i] = True
		if not visited_col[node.j]:
			for i in range(n):
				newValue = node.value + grid[i][node.j]
				queue.put(Node(i, node.j, newValue, node.steps + 1))
			visited_col[node.j] = True
	return count

n = int(input())
grid = []
for i in range(n):
	grid.append(list(map(int, input().split())))
	
x = int(input())

print(bfs(n, grid, x))


5. 題目：將一個由字符串s1轉換為字符串s2，可以進行三種操作：插入一個字符、刪除一個字符、替換一個字符。求最小的操作次數。

答案：

from queue import Queue

class Node:
	def __init__(self, s, steps):
		self.s = s
		self.steps = steps

def bfs(s1, s2):
	if s1 == s2:
		return 0
	queue = Queue()
	visited = set()
	queue.put(Node(s1, 0))
	visited.add(s1)
	while not queue.empty():
		node = queue.get()
		for i in range(len(s1)):
			for j in range(26):
				newChar = chr(ord('a') + j)
				if newChar != node.s[i]:
					newStr = node.s[:i] + newChar + node.s[i+1:]
					if newStr == s2:
						return node.steps + 1
					if newStr not in visited:
						queue.put(Node(newStr, node.steps + 1))
						visited.add(newStr)
		if len(node.s) < len(s2):
			newStr = node.s + 'a'
			if newStr == s2:
				return node.steps + 1
			if newStr not in visited:
				queue.put(Node(newStr, node.steps + 1))
				visited.add(newStr)
			
			newStr = 'a' + node.s
			if newStr == s2:
				return node.steps + 1
			if newStr not in visited:
				queue.put(Node(newStr, node.steps + 1))
				visited.add(newStr)
			
		elif len(node.s) > len(s2):
			for i in range(len(node.s)):
				newStr = node.s[:i] + node.s[i+1:]
				if newStr == s2:
					return node.steps + 1
				if newStr not in visited:
					queue.put(Node(newStr, node.steps + 1))
					visited.add(newStr)
	return -1

s1 = input()
s2 = input()

print(bfs(s1, s2))   

