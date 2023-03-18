

策略模式（Strategy Pattern）是一种行为型模式，用以定义不同的算法和演算法，并在需要时 dynamically interchange during runtime for different scenarios。

使用策略模式的目的是将相关的算法和演算法封装成不同的策略，让程式能够更加灵活和易于维护。此外，策略模式遵循开放封闭原则（Open-Closed Principle, OCP），允许新增新的策略而不会修改原有的程式码。

以下为举例：

例如，假设我们有一个以动物为主题的游戏，里面有不同类型的动物，如狗、猫、鸟等，每种动物都有自己的攻击技能。当我们在游戏中使用不同的动物时，会需要对应不同的攻击技能。

为此，我们可以使用策略模式。我们可以定义一个名为 AttackStrategy 的介面，并在其中定义攻击行为。接着，对每个动物实现一个不同的 AttackStrategy 策略，以封装不同的攻击技能。当使用不同的动物时，我们只需要动态选择对应的策略即可。

AttackStrategy.java

```
public interface AttackStrategy {
   public void attack();
}
```

DogAttackStrategy.java 做了咬的动作

```
public class DogAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Dog attacks by biting!");
   }
}
```

CatAttackStrategy.java 做了抓的动作

```
public class CatAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Cat attacks by scratching!");
   }
}
```

BirdAttackStrategy.java 做了咬和飞的动作

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