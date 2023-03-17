

策略模式（Strategy Pattern）是一種行為型模式，用以定義不同的算法和演算法，並在需要時 dynamically interchange during runtime for different scenarios。

使用策略模式的目的是將相關的算法和演算法封裝成不同的策略，讓程式能夠更加靈活和易於維護。此外，策略模式遵循開放封閉原則（Open-Closed Principle, OCP），允許新增新的策略而不會修改原有的程式碼。

以下為舉例：

例如，假設我們有一個以動物為主題的遊戲，裡面有不同類型的動物，如狗、貓、鳥等，每種動物都有自己的攻擊技能。當我們在遊戲中使用不同的動物時，會需要對應不同的攻擊技能。

為此，我們可以使用策略模式。我們可以定義一個名為 AttackStrategy 的介面，並在其中定義攻擊行為。接著，對每個動物實現一個不同的 AttackStrategy 策略，以封裝不同的攻擊技能。當使用不同的動物時，我們只需要動態選擇對應的策略即可。

AttackStrategy.java

```
public interface AttackStrategy {
   public void attack();
}
```

DogAttackStrategy.java 做了咬的動作

```
public class DogAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Dog attacks by biting!");
   }
}
```

CatAttackStrategy.java 做了抓的動作

```
public class CatAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Cat attacks by scratching!");
   }
}
```

BirdAttackStrategy.java 做了咬和飛的動作

```
public class BirdAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Bird attacks by biting and flying!");
   }
}
```

Animal.java

```
public class Animal {
   private String name;
   private AttackStrategy attackStrategy;

   public Animal(String name, AttackStrategy attackStrategy) {
      this.name = name;
      this.attackStrategy = attackStrategy;
   }

   public void attack() {
      System.out.println(name + " attacks.");
      attackStrategy.attack();
   }
}
```

Client.java

```
public class Client {
   public static void main(String[] args) {
      AttackStrategy dogStrategy = new DogAttackStrategy();
      AttackStrategy catStrategy = new CatAttackStrategy();
      AttackStrategy birdStrategy = new BirdAttackStrategy();

      Animal dog = new Animal("Dog", dogStrategy);
      Animal cat = new Animal("Cat", catStrategy);
      Animal bird = new Animal("Bird", birdStrategy);

      dog.attack();    // Dog attacks by biting!
      cat.attack();    // Cat attacks by scratching!
      bird.attack();   // Bird attacks by biting and flying!
   }
}
```