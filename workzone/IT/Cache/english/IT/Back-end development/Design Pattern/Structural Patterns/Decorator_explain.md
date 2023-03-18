

Decorator is a design pattern that allows for the modification or enhancement of an object's behavior dynamically without altering its original source code. It is used to add new functionality to an existing object by creating a wrapper object that encapsulates the original object and adds new methods, properties, or behaviors to it.

For example, let's consider a coffee shop that offers different types of coffee drinks. Each drink has a price, and customers can add extra ingredients such as whipped cream, caramel, or chocolate syrup to their coffee. Using the decorator pattern, we can create a Coffee class as the base object, and add decorators for each extra ingredient that customers can choose from:

```
// Coffee base class
class Coffee {
  constructor() {
    this.description = "Unknown coffee";
  }

  cost() {
    return 0;
  }

  getDescription() {
    return this.description;
  }
}

// Decorator for whipped cream
class WhippedCream extends Coffee {
  constructor(coffee) {
    super();
    this.coffee = coffee;
    this.description = this.coffee.getDescription() + ", whipped cream";
  }

  cost() {
    return this.coffee.cost() + 0.50;
  }
}

// Decorator for caramel
class Caramel extends Coffee {
  constructor(coffee) {
    super();
    this.coffee = coffee;
    this.description = this.coffee.getDescription() + ", caramel";
  }

  cost() {
    return this.coffee.cost() + 0.75;
  }
}

// Client code
const myCoffee = new Coffee();
console.log(myCoffee.getDescription(), "costs", myCoffee.cost());

myCoffee = new WhippedCream(myCoffee);
console.log(myCoffee.getDescription(), "costs", myCoffee.cost());

myCoffee = new Caramel(myCoffee);
console.log(myCoffee.getDescription(), "costs", myCoffee.cost());
```

In this example, the Coffee class is the base object, while the WhippedCream and Caramel classes are the decorators. Each decorator adds a new functionality (whipped cream and caramel) to the base object (coffee), and the final cost and description of the coffee are calculated based on the selected decorators. The advantage of using the decorator pattern in this example is that we can easily add new decorators for different types of ingredients without modifying the Coffee class.