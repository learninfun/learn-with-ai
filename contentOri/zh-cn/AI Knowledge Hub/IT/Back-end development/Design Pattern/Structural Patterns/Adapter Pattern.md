+++
title = "适配器模式"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Adapter Pattern的问题
{{< /ask_chatgpt >}}



1. 什麼是Adapter Pattern？
2. 合成與繼承在Adapter Pattern中的不同使用方式是什麼？
3. 如何實現類型與方法的適配器？
4. Adapter Pattern的使用場景是什麼？
5. 如何解決不同類型之間的接口兼容性問題？   

## 说明知识
{{< ask_chatgpt >}}
说明Adapter Pattern并举例
{{< /ask_chatgpt >}}



適配器模式是一種常見的設計模式，它通常是用來把一個類的接口轉換成客戶端所期待的另一種接口的模式。這意味著你可以使用原本不相容的類來一起工作，並且無需對現有的程式碼進行修改或重構。

舉例來說，在某個程式裡面有一個現成的組件可以呈現圓形，但你需要的是呈現正方形。這時，你可以使用適配器模式去轉換這兩種不相容的接口，並且達到你需要的結果。

以下是一個簡單的範例：

假設有一個工具箱，裡面有兩個工具：一把圓鉗和一把方鉗。現在，我們需要從這些鉗子當中選擇一把，並用它來剪開一個方形形狀的材料。

但問題是：圓鉗只適合剪開圓形形狀的材料，而方鉗只適合剪開方形形狀的材料。那該怎麼辦呢？我們可以使用適配器模式來轉換這兩種不相容的接口，並且讓它們一起工作。

我們可以實現一個適配器，它可以接收圓鉗的輸入，並將其轉換為正方形，最後再使用方鉗來剪開它。這樣，我們就可以用適配器來完成這個任務，而不必去重寫或重構原本的程式碼。

以下是示例程式碼：

```
// 鉗子介面
interface ITool {
    void cut();
}

// 圓形鉗子
class CircleTool implements ITool {
    @Override
    public void cut() {
        System.out.println("剪開圓形的材料");
    }
}

// 方形鉗子
class SquareTool implements ITool {
    @Override
    public void cut() {
        System.out.println("剪開方形的材料");
    }
}

// 適配器
class Adapter implements ITool {
    private CircleTool circleTool;

    public Adapter(CircleTool circleTool) {
        this.circleTool = circleTool;
    }

    @Override
    public void cut() {
        System.out.println("使用適配器進行轉換");
        System.out.println("將圓形材料轉換為方形材料");
        new SquareTool().cut();
    }
}

// 客户端
public class Client {
    public static void main(String[] args) {
        CircleTool circleTool = new CircleTool();
        ITool adapter = new Adapter(circleTool);
        adapter.cut();
    }
}
```

上述程式碼中，我們設計了一個 ITool 介面，裡面包含一個剪開的方法。並且定義了兩種鉗子：圓形鉗子和方形鉗子，分別實現了 ITool 介面的方法。

接著，我們又新增了一個 Adapter 類，它實現了 ITool 介面，但卻是通過圓形鉗子去轉換成方形鉗子。最後，在客戶端中，我們只需實例化 Adapter 類，就可以通過圓形鉗子去剪開方形形狀的材料了。

總結：適配器模式讓原本不相容的類可以共同協作，提高了程式的靈活性和可擴展性，使得系統的修改和維護變得更加方便。   

## 汇总重点
{{< ask_chatgpt >}}
条列Adapter Pattern的重点
{{< /ask_chatgpt >}}



1.目标介面(Target Interface)
- 定义客户端希望使用的介面。
- 通常是一个抽像类别或介面。

2.适配器(Adapter)
- 实现目标介面。
- 将不相容的类别(即要转换的类别)转换成目标介面。
- 可以是类别型态或物件型态。

3.不相容的类别(Adaptee)
- 不能直接使用的类别。
- 通常是一个现有的类别。
- 介面和目标介面不相同。

4.客户端(Client)
- 想要使用目标介面的类别。

5.转换过程
- 客户端呼叫适配器的方法。
- 适配器接收到客户端的呼叫，转换成对不相容的类别的方法呼叫。
- 不相容的类别回传结果给适配器。
- 适配器将不相容的类别的结果转换成目标介面的结果，回传给客户端。

6.类别适配器和物件适配器
- 类别适配器: 继承适配器和不相容的类别。
- 物件适配器: 组合适配器和不相容的类别。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Adapter Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 现有一个游戏中的角色类，其中攻击方法名为attack()，现在需要将其转换为一个可应用于另一个游戏中的角色类。使用适当的设计模式，实现需要的类适配器类。

答案：
```
public interface Character {
    void attack();
}

public class ExistingCharacter {
    public void attack() {
        System.out.println("Attacking with existing character!");
    }
}

public class TargetCharacterAdapter extends ExistingCharacter implements Character {
    public void attack() {
        System.out.println("Adapting existing character attack to target character attack...");
        super.attack();
    }
}

```

2. 假设现有一个印表机接口，但系统需要将接口改为支持打印机及扫瞄功能。使用可扩展的对象适配器，满足系统的需求。

答案：
```
public interface Printer {
    void print();
}

public class BasicPrinter implements Printer {
    public void print() {
        System.out.println("Basic printing...");
    }
}

public interface Scanner {
    void scan();
}

public class BasicScanner implements Scanner {
    public void scan() {
        System.out.println("Basic scanning...");
    }
}


public class PrinterScannerAdapter implements printer {
    Scanner scanner;

    public PrinterScannerAdapter(Scanner scanner) {
        this.scanner = scanner;
    }

    public void print() {
        System.out.println("Adapter is printing...");
    }

    public void scanAndPrint() {
        scanner.scan();
        print();
    }
}
```

3. 现有一个报表产生器，其支援的数据库为MySQL，现需要扩展以支援PostgreSQL数据库。使用对像适配器，满足报表产生系统的需求。

答案：
```
//MySQL connector
public interface MySQLConnection {
    void connectMySQL();
}

public class MySQLConnector implements MySQLConnection {
    public void connectMySQL() {
        System.out.println("Connecting to MySQL...");
    }
}

public interface ReportGenerator {
    void generateReport();
}

public class MySQLReportGenerator implements ReportGenerator {
    MySQLConnection mysql;

    public MySQLReportGenerator(MySQLConnection mysql) {
        this.mysql = mysql;
    }

    public void generateReport() {
        mysql.connectMySQL();
        System.out.println("Generating MySQL Report...");
    }
}

//PostgreSQL connector
public interface PostgreSQLConnection {
    void connectPostgreSQL();
}

public class PostgreSQLConnector implements PostgreSQLConnection {
    public void connectPostgreSQL() {
        System.out.println("Connecting to PostgreSQL...");
    }
}

public class PostgreSQLConnectorAdapter implements MySQLConnection {
    PostgreSQLConnection postgres;

    public PostgreSQLConnectorAdapter(PostgreSQLConnection postgres) {
        this.postgres = postgres;
    }

    public void connectMySQL() {
        postgres.connectPostgreSQL();
    }
}

//Test client
public class TestClient {
    public static void main(String[] args) {
        MySQLConnection mysql = new MySQLConnector();
        ReportGenerator mysqlGenerator = new MySQLReportGenerator(mysql);
        mysqlGenerator.generateReport();

        PostgreSQLConnection postgres = new PostgreSQLConnector();
        MySQLConnection postgresAdapter = new PostgreSQLConnectorAdapter(postgres);
        ReportGenerator postgresGenerator = new MySQLReportGenerator(postgresAdapter);
        postgresGenerator.generateReport();
    }
}
```

4. 现有一个商店购物车系统，现在需要将购物车系统扩展以支援开卷模式与包装模式。使用类适配器，满足系统需求。

答案：
```
public interface ShoppingCart {
    void addItem(String item);
    void removeItem(String item);
    void checkout();
}

public class BasicShoppingCart implements ShoppingCart {
    List<String> items = new ArrayList<>();

    public void addItem(String item) {
        items.add(item);
        System.out.println("Item \"" + item + "\" added to cart.");
    }

    public void removeItem(String item) {
        items.remove(item);
        System.out.println("Item \"" + item + "\" removed from cart.");
    }

    public void checkout() {
        System.out.println("Items in cart: " + items.toString());
        System.out.println("Basic checkout...");
    }
}

public interface OpenCart {
    void addOpenItem(String item);
    void removeOpenItem(String item);
    void openBulkCheckout();
}

public class OpenCartImpl implements OpenCart {
    List<String> items = new ArrayList<>();

    public void addOpenItem(String item) {
        items.add(item);
        System.out.println("Open item \"" + item + "\" added to cart.");
    }

    public void removeOpenItem(String item) {
        items.remove(item);
        System.out.println("Open item \"" + item + "\" removed from cart.");
    }

    public void openBulkCheckout() {
        System.out.println("Open items in cart: " + items.toString());
        System.out.println("Bulk checkout...");
    }
}

public class OpenCartAdapter extends OpenCartImpl implements ShoppingCart {
    public void addItem(String item) {
        addOpenItem(item);
    }

    public void removeItem(String item) {
        removeOpenItem(item);
    }

    public void checkout() {
        openBulkCheckout();
    }
}

public class WrappedCartAdapter extends BasicShoppingCart implements ShoppingCart {
    public void checkout() {
        System.out.println("Items in cart: " + items.toString());
        System.out.println("Wrapped checkout...");
    }
}
```

5. 假设现有一个软体系统需要扩展以支援三个不同的加密算法，现有两个具有不同加密算法的类别，使用介面适配器，为原始系统增加三种加密算法的支援。

答案：
```
public interface EncryptionAlgorithm {
    void encrypt();
}

public class XOREncryption {
    public void encrypt() {
        System.out.println("Encrypting using XOR...");
    }
}

public class ShiftEncryption {
    public void encrypt() {
        System.out.println("Encrypting using Shifting Cipher...");
    }
}

public class XOREncryptionAdapter implements EncryptionAlgorithm {
    XOREncryption xor;

    public XOREncryptionAdapter(XOREncryption xor) {
        this.xor = xor;
    }

    public void encrypt() {
        System.out.println("Adapter is encrypting using XOR...");
        xor.encrypt();
    }
}

public class ShiftEncryptionAdapter implements EncryptionAlgorithm {
    ShiftEncryption shift;

    public ShiftEncryptionAdapter(ShiftEncryption shift) {
        this.shift = shift;
    }

    public void encrypt() {
        System.out.println("Adapter is encrypting using Shifting Cipher...");
        shift.encrypt();
    }
}

public class SystemWithEncryption {
    EncryptionAlgorithm algorithm;

    public void setEncryptionAlgorithm(EncryptionAlgorithm algorithm) {
        this.algorithm = algorithm;
    }

    public void doSomethingWithEncryptedData() {
        algorithm.encrypt();
    }
}

```   

