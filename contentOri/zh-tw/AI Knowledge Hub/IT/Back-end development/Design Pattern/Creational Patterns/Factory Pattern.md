+++
title = "工廠模式"
weight = "3"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Factory Pattern的中文問題
{{< /ask_chatgpt >}}



1. 什麼是工廠模式（Factory Pattern）？請舉例說明。
2. 工廠模式和單例模式有什麼區別？什麼情況下應該使用工廠模式？
3. 工廠模式中的三種角色（Creator、Product、ConcreteProduct）分別代表什麼？他們之間是如何協作的？
4. 工廠模式有哪些不同的變形？各有什麼特點和應用場景？
5. 從性能、可維護性和擴展性等角度分析，工廠模式有什麼優缺點？在設計中應該如何抉擇是否要使用工廠模式？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Factory Pattern並舉例
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

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Factory Pattern的重點
{{< /ask_chatgpt >}}



1. Factory Design Pattern是一種創建型模式，用於創建具有相同屬性的物件，並隔離客戶端與具體產品的創建過程。

2. 工廠模式將創建產品的過程封裝在一個工廠類中，客戶端只需要知道工廠方法，而不需要知道產品的具體實現細節。

3. 工廠模式有三種形式：簡單工廠模式、工廠方法模式和抽像工廠模式。簡單工廠模式是最基本的工廠模式，而工廠方法模式和抽像工廠模式則是進一步的擴展。

4. 工廠模式具有以下優點：封裝了創建產品的過程，提高了系統的可維護性和可擴展性；隔離了客戶端和具體產品的依賴關係，降低系統的耦合性；提高了系統的靈活性和可重用性。

5. 工廠模式的缺點是：增加了系統的抽像性和理解難度；需要設計抽像工廠和產品族，增加了系統的代碼量和設計難度。

6. 工廠模式適用於以下情況：對象的創建過程比較複雜，需要隔離客戶端和具體產品之間的相互依賴關係；需要在不同地方創建同類型的對象時，可以通過工廠類對像創建；需要通過配置文件等方式動態生成對像時。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Factory Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 請實現一個工廠方法，其為指定的產品創建對象，但每個產品都有不同的數量和價值屬性。
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
這個工廠方法可以創建諸如 A、B 和 C 的產品對象，其具有不同的數量和價值屬性。

2. 實現一個工廠方法，可以創建不同的動物對象，例如狗、貓和豬。
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
可以使用這個工廠方法來創建 Dog、Cat 或 Pig 的對象。

3. 請實現一個工廠方法，用於創建不同類型的 UI 元素，例如按鈕、文本框和標籤。
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
這個工廠方法可以創建 Button、TextBox 或 Label 的對象。

4. 實現一個工廠方法，用於創建不同類型的圖形對象，例如圓形、矩形和三角形。
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
這個工廠方法可以創建 Circle、Rectangle 或 Triangle 的對象。

5. 實現一個工廠方法，用於創建不同類型的文件對象，例如文本文件、PDF 文件和圖像文件。
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
這個工廠方法可以創建 TextFile、PDFFile 或 ImageFile 的對象。   

