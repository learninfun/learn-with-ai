

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