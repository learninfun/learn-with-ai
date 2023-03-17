

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