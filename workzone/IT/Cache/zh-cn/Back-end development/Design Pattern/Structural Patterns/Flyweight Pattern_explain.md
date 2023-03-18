

Flyweight模式是一种设计模式，用于减少系统中重复的物体（Object）数量，从而提高系统的效率。Flyweight可在大量具有相似或相同的内容的对象中有效地减少内容共享，以减少系统内部对象的数量。

该模式通常由工厂（Factory）对象创建和管理，它维护一个Flyweight Cache，以确保只有一个flyweight实例的各个实例在该Cache中被共享。当需要访问单个Flyweight对象时，工厂对象先检查Flyweight Cache，是否已经存在相同的对象，如果是则返回已存在的对象，否则创建一个新的对象并添加到Flyweight Cache中。

下面是一个Flyweight模式的例子，假设我们需要创建一个文字编辑工具，其中有各种不同的字体和颜色可供选择，但我们希望尽可能减少类的数量以提高效率：

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

在上述案例中，字体和颜色都是Flyweight物件，CharacterFactory负责管理这些物件，并且只创建必要的物件。当客户端需要一个新的Character时，CharacterFactory会创建一个具有相同参数的另一个Character，或者返回现有的Character，以便减少系统中的物件数量。