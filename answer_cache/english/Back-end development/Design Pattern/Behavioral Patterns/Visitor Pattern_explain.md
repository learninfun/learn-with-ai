

The Visitor Pattern is a design pattern in which a set of operations (functions or methods) can be applied to each element of a data structure regardless of its type. The Visitor pattern is a way to separate a data structure from the algorithms or operations that can be performed on it.

The pattern is commonly used in situations where a data structure contains multiple types of objects and the behavior or operations applied to each type of object is different. In such cases, it can be difficult to create a generic processing function that can handle all the types; the Visitor Pattern solves this by creating separate function or class for each type of object.

For example, let's consider a program that simulates a zoo. The zoo consists of different types of animals, such as lions, tigers, elephants, and monkeys. Each animal has different attributes such as height, weight, speed, and favorite food.

In this scenario, the Visitor Pattern can be used to calculate the total weight of all the animals in the zoo. Instead of having the main program iterate through each animal and calculate the weight, we can have a separate visitor function or class to do this.

The Animal class hierarchy would have an accept() method, that takes a Visitor object as a parameter. Each animal subclass would implement the accept() method and call the corresponding visit() method in the Visitor object.

The TotalWeightVisitor class would have a visit() method for each animal subclass, that would update the total weight, and a getWeight() method that would return the result.

class Animal {
  public void accept(Visitor v) { v.visit(this); }
}

class Lion extends Animal {}
class Tiger extends Animal {}
class Elephant extends Animal {}

interface Visitor {
  public void visit(Lion lion);
  public void visit(Tiger tiger);
  public void visit(Elephant elephant);
}

class TotalWeightVisitor implements Visitor {
  private int totalWeight = 0;

  public void visit(Lion lion) {
    totalWeight += lion.getWeight();
  }

  public void visit(Tiger tiger) {
    totalWeight += tiger.getWeight();
  }

  public void visit(Elephant elephant) {
    totalWeight += elephant.getWeight();
  }

  public int getWeight() {
    return totalWeight;
  }
}

Usage:

List<Animal> animals = new ArrayList<>();
animals.add(new Lion(200));
animals.add(new Tiger(150));
animals.add(new Elephant(500));

TotalWeightVisitor visitor = new TotalWeightVisitor();
for (Animal animal : animals) {
  animal.accept(visitor);
}

int totalWeight = visitor.getWeight(); // 850