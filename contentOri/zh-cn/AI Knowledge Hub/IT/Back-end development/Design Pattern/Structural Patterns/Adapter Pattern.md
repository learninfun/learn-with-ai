+++
title = "適配器模式"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Adapter Pattern的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Adapter Pattern？
2. 合成與繼承在Adapter Pattern中的不同使用方式是什麼？
3. 如何實現類型與方法的適配器？
4. Adapter Pattern的使用場景是什麼？
5. 如何解決不同類型之間的接口兼容性問題？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Adapter Pattern並舉例
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

## 條列重點
{{< ask_chatgpt >}}
以中文條列Adapter Pattern的重點
{{< /ask_chatgpt >}}



1.目標介面(Target Interface)
- 定義客戶端希望使用的介面。
- 通常是一個抽像類別或介面。

2.適配器(Adapter)
- 實現目標介面。
- 將不相容的類別(即要轉換的類別)轉換成目標介面。
- 可以是類別型態或物件型態。

3.不相容的類別(Adaptee)
- 不能直接使用的類別。
- 通常是一個現有的類別。
- 介面和目標介面不相同。

4.客戶端(Client)
- 想要使用目標介面的類別。

5.轉換過程
- 客戶端呼叫適配器的方法。
- 適配器接收到客戶端的呼叫，轉換成對不相容的類別的方法呼叫。
- 不相容的類別回傳結果給適配器。
- 適配器將不相容的類別的結果轉換成目標介面的結果，回傳給客戶端。

6.類別適配器和物件適配器
- 類別適配器: 繼承適配器和不相容的類別。
- 物件適配器: 組合適配器和不相容的類別。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Adapter Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 現有一個遊戲中的角色類，其中攻擊方法名為attack()，現在需要將其轉換為一個可應用於另一個遊戲中的角色類。使用適當的設計模式，實現需要的類適配器類。

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

2. 假設現有一個印表機接口，但系統需要將接口改為支持打印機及掃瞄功能。使用可擴展的對象適配器，滿足系統的需求。

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

3. 現有一個報表產生器，其支援的數據庫為MySQL，現需要擴展以支援PostgreSQL數據庫。使用對像適配器，滿足報表產生系統的需求。

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

4. 現有一個商店購物車系統，現在需要將購物車系統擴展以支援開卷模式與包裝模式。使用類適配器，滿足系統需求。

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

5. 假設現有一個軟體系統需要擴展以支援三個不同的加密算法，現有兩個具有不同加密算法的類別，使用介面適配器，為原始系統增加三種加密算法的支援。

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

