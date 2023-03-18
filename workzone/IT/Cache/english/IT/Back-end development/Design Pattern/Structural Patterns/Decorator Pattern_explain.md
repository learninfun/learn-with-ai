

The Decorator Pattern is a structural design pattern that allows adding behavior or responsibilities to an object dynamically, at runtime, without changing its core structure. This pattern involves creating a wrapper or decorator object that wraps the original object and provides additional capabilities.

For example, imagine a coffee shop that offers different types of coffee drinks. Each type of coffee can have different add-ons, such as milk, sugar, whipped cream, etc. Instead of creating a separate class for each combination, the Decorator Pattern can be used. The base interface could be `Coffee`, and each coffee type could be represented by a concrete implementation of the interface, for example, `Espresso` and `Latte`. Then, decorators can be created for each add-on, such as `MilkDecorator`, `SugarDecorator`, etc. These decorators would implement the same interface as the base `Coffee` class and would have a reference to the decorated coffee object. When a customer orders a coffee with add-ons, the decorator objects can be stacked on top of the base coffee object, creating a chain of decorated objects, like in the following example:

```
Coffee latte = new Latte();
latte = new MilkDecorator(latte);
latte = new SugarDecorator(latte);
latte = new WhippedCreamDecorator(latte);
```

This would create a `Latte` object with added milk, sugar, and whipped cream.

By using the Decorator Pattern, we can add new features to our objects easily, without touching their original code, and keep our code flexible and easy to extend in the future.