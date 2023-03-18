

Factory Method is a design pattern that allows the creation of objects without exposing the object instantiation logic to the client. Instead of calling the class constructor directly, the client calls a factory method that returns an instance of the class. This provides a layer of abstraction and allows for flexibility in creating objects.

For example, let's consider an abstract class called "Animal" with two concrete subclasses "Dog" and "Cat". The factory method could be implemented in a separate class called "AnimalFactory". By calling the factory method "createAnimal", the client can create an instance of either "Dog" or "Cat" without needing to know the specific details of how each class is constructed.

```
abstract class Animal {
  public abstract String makeSound();
}

class Dog extends Animal {
  public String makeSound() {
    return "Woof!";
  }
}

class Cat extends Animal {
  public String makeSound() {
    return "Meow!";
  }
}

class AnimalFactory {
  public Animal createAnimal(String animalType) {
    if (animalType.equalsIgnoreCase("Dog")) {
      return new Dog();
    } else {
      return new Cat();
    }
  }
}

// Client code
AnimalFactory factory = new AnimalFactory();
Animal dog = factory.createAnimal("Dog");
System.out.println(dog.makeSound()); // output: Woof!

Animal cat = factory.createAnimal("Cat");
System.out.println(cat.makeSound()); // output: Meow!
```