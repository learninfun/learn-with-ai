

1. 假设有一个图书馆系统，你想实现一个 Book 类，每个 Book 对象都包含了一些固定的基本信息（例如书名、作者、ISBN 编号等等）。请使用 Flyweight Pattern 实现这个 Book 类。

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

2. 假设你正在开发一款 RPG 游戏，其中有许多种类的武器和防具。你希望使用 Flyweight Pattern 来最小化物件的数量，同时让玩家能轻松使用这些武器和防具。请写出建议的类别设计。

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
        // 使用剑进行攻击
    }
}

class Shield implements Armor {
    private int defence;

    public Shield(int defence) {
        this.defence = defence;
    }

    @Override
    public void use() {
        // 使用盾牌进行防御
    }
}

class WeaponFactory {
    // 保存已创建的武器实例
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
    // 保存已创建的防具实例
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

3. 假设你正在开发一个基于 JavaFX 的图像编辑器，当使用者选择某种颜色时，它将应用到图像中的所有像素。为了减少颜色对象的数量，你想使用 Flyweight Pattern。请描述如何实现这一概念，并给出适当的程式码。

答案：

```
// 逻辑上代表不同的颜色，但实际上它们是同一个对象
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
    // 储存已生成的颜色对象，避免重复创建
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
    // 储存代表像素的(x, y)座标和其对应的颜色
    private final int x;
    private final int y;
    private final Color color;

    public Pixel(int x, int y, Color color) {
        this.x = x;
        this.y = y;
        this.color = color;
    }

    public void apply() {
        // 将颜色应用到图片上
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

4. 假设你正在开发一个网路游戏，当玩家加入游戏时，你需要为每个玩家建立一个对象。为了减少占用的内存，你想使用 Flyweight Pattern。请写出建议的类别设计。

答案：

```
class Player {
    private final String name;
    private final int level;
    // ...其他玩家属性

    public Player(String name, int level) {
        this.name = name;
        this.level = level;
    }

    // ...其他方法
}

class PlayerFactory {
    // 储存已建立的玩家对象，避免重复建立
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

5. 假设你正在开发一个购物车系统，每个产品都有一个唯一的数字和一个价格。当使用者将一个产品添加到购物车中时，你希望使用 Flyweight Pattern 从 Cache 中取回该产品的相关资料。请写出建议的类别设计。

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
    // 储存已创建的产品对象，避免重复创建
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
    // 储存使用者已添加的所有产品
    private final List<Product> products = new ArrayList<>();

    public void addProduct(int id) {
        // 透过产品工厂取得产品对象
        Product product = ProductFactory.create(id, getPriceFromCache(id));
        products.add(product);
    }

    private double getPriceFromCache(int id) {
        // 从 Cache 中取回产品的价格（如果适用）
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