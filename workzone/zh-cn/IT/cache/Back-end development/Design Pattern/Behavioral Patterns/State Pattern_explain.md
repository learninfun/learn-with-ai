

State Pattern是一種行為型設計模式，它允許物件在其內部狀態改變時動態改變其行為，並且不需要大幅修改原有物件的結構、方法組成等。這個模式的核心理念是將物件的狀態轉換成一個獨立的物件，並且不斷的調用該物件對應的方法實現狀態所對應的行為。

下面是一個以簡單宝可梦战斗为例的狀態模式：

假設有三個寶可夢，開戰前和戰鬥中的狀態是不同的：

```python
# 開戰前狀態
class PokemonState:
    def execute(self):
        pass

class Ready(PokemonState):
    def execute(self):
        return '預備'

# 戰鬥狀態
class Battle(PokemonState):
    def execute(self):
        return '戰鬥'
```

然後建立主角的狀態：

```python
class Pokemon:
    def __init__(self):
        self.current_state = Ready()

    # 切換狀態
    def switch_state(self, state):
        self.current_state = state

    def do_action(self):
        return self.current_state.execute()

```

實例化三只隨機的寶可夢：

```python
pikachu = Pokemon()
charmender = Pokemon()
squirtle = Pokemon()
```

在各自的狀態下:

```python
# 主角pikachu遇到敵人後
print('我要跟你戰鬥了！')
pikachu.switch_state(Battle())
print(pikachu.do_action())
```
输出:
```
我要跟你戰鬥了！
戰鬥
```

```python
# 版金德遇到敵人後
print('我要跟你戰鬥了！')
charmender.switch_state(Battle())
print(charmender.do_action())
```
输出:
```
我要跟你戰鬥了！
戰鬥
```

```python
# 小火龍遇到敵人後
print('我要跟你戰鬥了！')
squirtle.switch_state(Battle())
print(squirtle.do_action())
```
输出:
```
我要跟你戰鬥了！
戰鬥
```

這就是思路，类似的，你可以在其他場景中使用這種模式以達到更好的代碼可維護性和可擴展性。