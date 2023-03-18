+++
title = "享元模式"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Flyweight Pattern的问题
{{< /ask_chatgpt >}}



1. 何谓Flyweight Pattern？
2. 什么状况下适用Flyweight Pattern？
3. 与其他设计模式相比，Flyweight Pattern有哪些优点和缺点？
4. 如何实现Flyweight Pattern？
5. 在实际应用中，有哪些例子适合使用Flyweight Pattern？   

## 说明知识
{{< ask_chatgpt >}}
说明Flyweight Pattern并举例
{{< /ask_chatgpt >}}



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

## 汇总重点
{{< ask_chatgpt >}}
条列Flyweight Pattern的重点
{{< /ask_chatgpt >}}



1. Flyweight pattern 是一种设计模式，主要用于减少大量相似物件的记忆体使用量。
2. 该模式将物件分为两类，一类是共享的，一类是独立的。
3. 共享的物件可以在多个地方共用，节省记忆体空间。
4. 独立的物件则会被重复建立，但不影响其他物件的使用。
5. Flyweight pattern 的实现可以用工厂模式来简化程式码。
6. 使用 Flyweight pattern 可以提高应用程式效能，尤其在需要建立大量物件时。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Flyweight Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 假設有一個圖書館系統，你想實現一個 Book 類，每個 Book 對象都包含了一些固定的基本信息（例如書名、作者、ISBN 編號等等）。請使用 Flyweight Pattern 實現這個 Book 類。

答案：

```
class Book {
    private String title;
    private String author;
    private String isbn;
    // ...其他基本信息

    public Book(String title, String author, String isbn) {
        this.title = title;
        this.author = author;
        this.isbn = isbn;
    }

    // ...其他方法
}
```

2. 假設你正在開發一款 RPG 遊戲，其中有許多種類的武器和防具。你希望使用 Flyweight Pattern 來最小化物件的數量，同時讓玩家能輕鬆使用這些武器和防具。請寫出建議的類別設計。

答案：

```
interface Weapon {
    void use();
}

interface Armor {
    void use();
}

class Sword implements Weapon {
    private int damage;

    public Sword(int damage) {
        this.damage = damage;
    }

    @Override
    public void use() {
        // 使用劍進行攻擊
    }
}

class Shield implements Armor {
    private int defence;

    public Shield(int defence) {
        this.defence = defence;
    }

    @Override
    public void use() {
        // 使用盾牌進行防禦
    }
}

class WeaponFactory {
    // 保存已創建的武器實例
    private static final Map<Integer, Weapon> weapons = new HashMap<>();

    public static Weapon create(int damage) {
        Weapon weapon = weapons.get(damage);
        if (weapon == null) {
            weapon = new Sword(damage);
            weapons.put(damage, weapon);
        }
        return weapon;
    }
}

class ArmorFactory {
    // 保存已創建的防具實例
    private static final Map<Integer, Armor> armors = new HashMap<>();

    public static Armor create(int defence) {
        Armor armor = armors.get(defence);
        if (armor == null) {
            armor = new Shield(defence);
            armors.put(defence, armor);
        }
        return armor;
    }
}
```

3. 假設你正在開發一個基於 JavaFX 的圖像編輯器，當使用者選擇某種顏色時，它將應用到圖像中的所有像素。為了減少顏色對象的數量，你想使用 Flyweight Pattern。請描述如何實現這一概念，並給出適當的程式碼。

答案：

```
// 邏輯上代表不同的顏色，但實際上它們是同一個對象
class Color {
    private final javafx.scene.paint.Color value;

    public Color(javafx.scene.paint.Color value) {
        this.value = value;
    }

    public javafx.scene.paint.Color getValue() {
        return value;
    }
}

class ColorFactory {
    // 儲存已生成的顏色對象，避免重複創建
    private static final Map<String, Color> colors = new HashMap<>();

    public static Color create(javafx.scene.paint.Color value) {
        String key = value.toString();
        Color color = colors.get(key);
        if (color == null) {
            color = new Color(value);
            colors.put(key, color);
        }
        return color;
    }
}

class Pixel {
    // 儲存代表像素的(x, y)座標和其對應的顏色
    private final int x;
    private final int y;
    private final Color color;

    public Pixel(int x, int y, Color color) {
        this.x = x;
        this.y = y;
        this.color = color;
    }

    public void apply() {
        // 將顏色應用到圖片上
    }
}

class ImageEditor {
    private final Pixel[][] pixels;

    public ImageEditor(int width, int height) {
        pixels = new Pixel[width][height];
    }

    public void setColor(int x, int y, javafx.scene.paint.Color value) {
        Color color = ColorFactory.create(value);
        pixels[x][y] = new Pixel(x, y, color);
    }

    public void applyColor() {
        for (Pixel[] row : pixels) {
            for (Pixel pixel : row) {
                pixel.apply();
            }
        }
    }
}
```

4. 假設你正在開發一個網路遊戲，當玩家加入遊戲時，你需要為每個玩家建立一個對象。為了減少佔用的內存，你想使用 Flyweight Pattern。請寫出建議的類別設計。

答案：

```
class Player {
    private final String name;
    private final int level;
    // ...其他玩家屬性

    public Player(String name, int level) {
        this.name = name;
        this.level = level;
    }

    // ...其他方法
}

class PlayerFactory {
    // 儲存已建立的玩家對象，避免重複建立
    private static final Map<String, Player> players = new HashMap<>();

    public static Player create(String name, int level) {
        String key = name + ":" + level;
        Player player = players.get(key);
        if (player == null) {
            player = new Player(name, level);
            players.put(key, player);
        }
        return player;
    }
}
```

5. 假設你正在開發一個購物車系統，每個產品都有一個唯一的數字和一個價格。當使用者將一個產品添加到購物車中時，你希望使用 Flyweight Pattern 從 Cache 中取回該產品的相關資料。請寫出建議的類別設計。

答案：

```
class Product {
    private final int id;
    private final double price;

    public Product(int id, double price) {
        this.id = id;
        this.price = price;
    }

    public double getPrice() {
        return price;
    }

    // ...其他方法
}

class ProductFactory {
    // 儲存已創建的產品對象，避免重複創建
    private static final Map<Integer, Product> products = new HashMap<>();

    public static Product create(int id, double price) {
        Product product = products.get(id);
        if (product == null) {
            product = new Product(id, price);
            products.put(id, product);
        }
        return product;
    }
}

class ShoppingCart {
    // 儲存使用者已添加的所有產品
    private final List<Product> products = new ArrayList<>();

    public void addProduct(int id) {
        // 透過產品工廠取得產品對象
        Product product = ProductFactory.create(id, getPriceFromCache(id));
        products.add(product);
    }

    private double getPriceFromCache(int id) {
        // 從 Cache 中取回產品的價格（如果適用）
        return 0.0;
    }

    public double getTotalPrice() {
        double total = 0.0;
        for (Product product : products) {
            total += product.getPrice();
        }
        return total;
    }
}
```   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Flyweight Pattern的网络数据
{{< /ask_chatgpt >}}



1. "Flyweight Design Pattern" by JavaTpoint

https://www.javatpoint.com/flyweight-design-pattern

This website provides a detailed explanation of the Flyweight pattern, including the concept, structure, and advantages of using this pattern. It also includes examples using Java programming language.

2. "Flyweight Design Pattern" by TutorialsPoint

https://www.tutorialspoint.com/design_pattern/flyweight_pattern.htm

TutorialsPoint provides an in-depth introduction and explanation of the Flyweight pattern, including how it works, its features, and when to use it. The website also includes code examples in Java programming language.

3. "Flyweight Design Pattern" by GeeksforGeeks

https://www.geeksforgeeks.org/flyweight-design-pattern/

GeeksforGeeks explains the Flyweight pattern in detail, including its definition, use case, and implementation using code examples in Java programming language. The website also includes a comparison of this pattern with other patterns.

4. "Flyweight Pattern" by Refactoring.Guru

https://refactoring.guru/design-patterns/flyweight

Refactoring.Guru provides an article on the Flyweight pattern, including its structure, concept, and benefits. The website also includes a comparison of this pattern with other patterns and code examples in Java programming language.

5. "Flyweight Design Pattern" by DZone

https://dzone.com/articles/flyweight-pattern

DZone explains the Flyweight pattern, including its definition, use case, and implementation using code examples in Java programming language. The website also includes the pros and cons of using this pattern in software development.   

