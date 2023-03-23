+++
title = "工厂模式"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Factory Pattern的问题
{{< /ask_chatgpt >}}



1. 什麼是工廠模式（Factory Pattern）？請舉例說明。
2. 工廠模式和單例模式有什麼區別？什麼情況下應該使用工廠模式？
3. 工廠模式中的三種角色（Creator、Product、ConcreteProduct）分別代表什麼？他們之間是如何協作的？
4. 工廠模式有哪些不同的變形？各有什麼特點和應用場景？
5. 從性能、可維護性和擴展性等角度分析，工廠模式有什麼優缺點？在設計中應該如何抉擇是否要使用工廠模式？   

## 说明知识
{{< ask_chatgpt >}}
说明Factory Pattern并举例
{{< /ask_chatgpt >}}



工廠模式（Factory Pattern）是一種常見的軟體設計模式，用於優化對象的創建過程。事實上，當需要動態創建複雜對象時，尤其是當這些對象具有共同的特徵時，Factory Pattern可能是最佳選擇。

Factory Pattern的基本思想是將對象的創建過程（與使用者隔離的異步），這樣能夠使得向應用程序添加新類的過程更加簡單。透過使用工廠模式，用戶端端不需要關注如何創建對象，而只需要專注於使用對象。

舉個例子，假設我們正在開發一個遊戲，該遊戲中有許多不同的敵方角色，包括僵屍、魔鬼和巨魔等。每個敵人都有自己獨特的特徵，例如血量、攻擊力和速度等。我們可以使用工廠模式來創建這些角色，從而避免大量的重複代碼。

首先，我們可以定義一個名為Enemy的抽象基類，該基類包含敵方角色具有的共同屬性和方法。然後，我們定義敵方角色的具體子類（如Zombie，Devil和Troll），並重寫其父類的屬性和方法以滿足其獨特的特徵。最後，我們實現一個EnemyFactory類，該類用於創建特定類型的敵人，例如：

```python
class Enemy:
    """
    敵方角色基類
    """
    def __init__(self, name):
        self.name = name
        self.health = None
        self.attack_power = None
        self.speed = None
        
    def attack(self):
        pass

class Zombie(Enemy):
    """
    僵屍類型
    """
    def __init__(self):
        super().__init__("Zombie")
        self.health = 100
        self.attack_power = 10
        self.speed = 5
        
    def attack(self):
        print("The Zombie attacks!")

class Devil(Enemy):
    """
    魔鬼類型
    """
    def __init__(self):
        super().__init__("Devil")
        self.health = 150
        self.attack_power = 15
        self.speed = 8
        
    def attack(self):
        print("The Devil attacks!")
        
class Troll(Enemy):
    """
    巨魔類型
    """
    def __init__(self):
        super().__init__("Troll")
        self.health = 250
        self.attack_power = 25
        self.speed = 3
        
    def attack(self):
        print("The Troll attacks!")
        
class EnemyFactory:
    """
    敵方角色工廠類
    """
    def create_enemy(enemy_type):
        if enemy_type == 'Zombie':
            return Zombie()
        elif enemy_type == 'Devil':
            return Devil()
        elif enemy_type == 'Troll':
            return Troll()
        else:
            raise ValueError("Invalid enemy type")
```

通過上述設計，我們可以創建出不同的敵方角色，具體如下：

```python
zombie = EnemyFactory.create_enemy('Zombie')
devil = EnemyFactory.create_enemy('Devil')
troll = EnemyFactory.create_enemy('Troll')

zombie.attack()  # output: The Zombie attacks!
devil.attack()  # output: The Devil attacks!
troll.attack()  # output: The Troll attacks!
```

當需要添加一個新敵方角色時，我們只需要創建一個新的敵方角色子類並實現其相關屬性和方法，然後在EnemyFactory中添加創建新角色的代碼即可。這樣做可以大大簡化對象的創建過程，並且可以使代碼更加簡潔易讀。   

## 汇总重点
{{< ask_chatgpt >}}
条列Factory Pattern的重点
{{< /ask_chatgpt >}}



1. Factory Design Pattern是一种创建型模式，用于创建具有相同属性的物件，并隔离客户端与具体产品的创建过程。

2. 工厂模式将创建产品的过程封装在一个工厂类中，客户端只需要知道工厂方法，而不需要知道产品的具体实现细节。

3. 工厂模式有三种形式：简单工厂模式、工厂方法模式和抽像工厂模式。简单工厂模式是最基本的工厂模式，而工厂方法模式和抽像工厂模式则是进一步的扩展。

4. 工厂模式具有以下优点：封装了创建产品的过程，提高了系统的可维护性和可扩展性；隔离了客户端和具体产品的依赖关系，降低系统的耦合性；提高了系统的灵活性和可重用性。

5. 工厂模式的缺点是：增加了系统的抽像性和理解难度；需要设计抽像工厂和产品族，增加了系统的代码量和设计难度。

6. 工厂模式适用于以下情况：对象的创建过程比较复杂，需要隔离客户端和具体产品之间的相互依赖关系；需要在不同地方创建同类型的对象时，可以通过工厂类对像创建；需要通过配置文件等方式动态生成对像时。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Factory Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 请实现一个工厂方法，其为指定的产品创建对象，但每个产品都有不同的数量和价值属性。
答案：

```python
from typing import List

class Product:
    def __init__(self, quantity: int, price: float):
        self.quantity = quantity
        self.price = price

class ProductFactory:
    @staticmethod
    def create_product(name: str, quantity: int, price: float) -> Product:
        if name == 'A':
            return Product(quantity * 2, price * 2)
        elif name == 'B':
            return Product(quantity * 4, price / 2)
        elif name == 'C':
            return Product(quantity - 1, price * 3)
        else:
            raise ValueError('Invalid product name')
```
这个工厂方法可以创建诸如 A、B 和 C 的产品对象，其具有不同的数量和价值属性。

2. 实现一个工厂方法，可以创建不同的动物对象，例如狗、猫和猪。
答案：

```python
from typing import Union

class Animal:
    def speak(self):
        pass

class Dog(Animal):
    def speak(self):
        return "Woof"

class Cat(Animal):
    def speak(self):
        return "Meow"

class Pig(Animal):
    def speak(self):
        return "Oink"

class AnimalFactory:
    @staticmethod
    def create_animal(name: str) -> Union[Dog, Cat, Pig]:
        if name == 'dog':
            return Dog()
        elif name == 'cat':
            return Cat()
        elif name == 'pig':
            return Pig()
        else:
            raise ValueError('Invalid animal name')
```
可以使用这个工厂方法来创建 Dog、Cat 或 Pig 的对象。

3. 请实现一个工厂方法，用于创建不同类型的 UI 元素，例如按钮、文本框和标签。
答案：

```python
from typing import Union

class UIElement:
    def render(self):
        pass

class Button(UIElement):
    def render(self):
        return "Rendering button..."

class TextBox(UIElement):
    def render(self):
        return "Rendering text box..."

class Label(UIElement):
    def render(self):
        return "Rendering label..."

class UIElementFactory:
    @staticmethod
    def create_element(name: str) -> Union[Button, TextBox, Label]:
        if name == 'button':
            return Button()
        elif name == 'textbox':
            return TextBox()
        elif name == 'label':
            return Label()
        else:
            raise ValueError('Invalid UI element name')
```
这个工厂方法可以创建 Button、TextBox 或 Label 的对象。

4. 实现一个工厂方法，用于创建不同类型的图形对象，例如圆形、矩形和三角形。
答案：

```python
from typing import Union

class Shape:
    def draw(self):
        pass

class Circle(Shape):
    def draw(self):
        return "Drawing circle..."

class Rectangle(Shape):
    def draw(self):
        return "Drawing rectangle..."

class Triangle(Shape):
    def draw(self):
        return "Drawing triangle..."

class ShapeFactory:
    @staticmethod
    def create_shape(name: str) -> Union[Circle, Rectangle, Triangle]:
        if name == 'circle':
            return Circle()
        elif name == 'rectangle':
            return Rectangle()
        elif name == 'triangle':
            return Triangle()
        else:
            raise ValueError('Invalid shape name')
```
这个工厂方法可以创建 Circle、Rectangle 或 Triangle 的对象。

5. 实现一个工厂方法，用于创建不同类型的文件对象，例如文本文件、PDF 文件和图像文件。
答案：

```python
from typing import Union

class File:
    def open(self):
        pass

class TextFile(File):
    def open(self):
        return "Opening text file..."

class PDFFile(File):
    def open(self):
        return "Opening PDF file..."

class ImageFile(File):
    def open(self):
        return "Opening image file..."

class FileFactory:
    @staticmethod
    def create_file(name: str) -> Union[TextFile, PDFFile, ImageFile]:
        if name == 'text':
            return TextFile()
        elif name == 'pdf':
            return PDFFile()
        elif name == 'image':
            return ImageFile()
        else:
            raise ValueError('Invalid file name')
```
这个工厂方法可以创建 TextFile、PDFFile 或 ImageFile 的对象。   

