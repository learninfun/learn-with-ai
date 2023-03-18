

Flyweight模式是一種設計模式，用於減少系統中重複的物體（Object）數量，從而提高系統的效率。Flyweight可在大量具有相似或相同的內容的對象中有效地減少內容共享，以減少系統內部對象的數量。

該模式通常由工廠（Factory）對像創建和管理，它維護一個Flyweight Cache，以確保只有一個flyweight實例的各個實例在該Cache中被共享。當需要訪問單個Flyweight對像時，工廠對像先檢查Flyweight Cache，是否已經存在相同的對象，如果是則返回已存在的對象，否則創建一個新的對象並添加到Flyweight Cache中。

下面是一個Flyweight模式的例子，假設我們需要創建一個文字編輯工具，其中有各種不同的字體和顏色可供選擇，但我們希望盡可能減少類的數量以提高效率：

```python
class Font:
    def __init__(self, name, size):
        self.name = name
        self.size = size

class Color:
    def __init__(self, red, green, blue):
        self.red = red
        self.green = green
        self.blue = blue

class Character:
    def __init__(self, char, font, color):
        self.char = char
        self.font = font
        self.color = color

class CharacterFactory:
    def __init__(self):
        self.cache = {}

    def get_char(self, char, font, color):
        key = (char, font.name, font.size, color.red, color.green, color.blue)
        if key not in self.cache:
            self.cache[key] = Character(char, font, color)
        return self.cache[key]
```

在上述案例中，字體和顏色都是Flyweight物件，CharacterFactory負責管理這些物件，並且只創建必要的物件。當客戶端需要一個新的Character時，CharacterFactory會創建一個具有相同參數的另一個Character，或者返回現有的Character，以便減少系統中的物件數量。