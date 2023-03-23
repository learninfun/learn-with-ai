+++
title = "享元模式"
weight = "6"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Flyweight Pattern的问题
{{< /ask_chatgpt >}}



1. 何謂Flyweight Pattern？
2. 什麼狀況下適用Flyweight Pattern？
3. 與其他設計模式相比，Flyweight Pattern有哪些優點和缺點？
4. 如何實現Flyweight Pattern？
5. 在實際應用中，有哪些例子適合使用Flyweight Pattern？   

## 说明知识
{{< ask_chatgpt >}}
说明Flyweight Pattern并举例
{{< /ask_chatgpt >}}



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

## 汇总重点
{{< ask_chatgpt >}}
条列Flyweight Pattern的重点
{{< /ask_chatgpt >}}



1. Flyweight pattern 是一種設計模式，主要用於減少大量相似物件的記憶體使用量。
2. 該模式將物件分為兩類，一類是共享的，一類是獨立的。
3. 共享的物件可以在多個地方共用，節省記憶體空間。
4. 獨立的物件則會被重複建立，但不影響其他物件的使用。
5. Flyweight pattern 的實現可以用工廠模式來簡化程式碼。
6. 使用 Flyweight pattern 可以提高應用程式效能，尤其在需要建立大量物件時。   

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

