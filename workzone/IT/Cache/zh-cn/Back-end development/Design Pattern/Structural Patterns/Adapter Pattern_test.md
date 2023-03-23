

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