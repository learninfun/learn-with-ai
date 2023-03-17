

Factory pattern is a creational design pattern that provides an interface for creating objects in a superclass but allows subclasses to alter the type of objects that will be created. It separates the responsibility of creating an object from the client that requests the object. 

Here is an example of the factory pattern:

Let's say we have a program that sells pizzas. We have different types of pizzas, such as cheese pizza, pepperoni pizza, and veggie pizza. Rather than creating individual classes for each type of pizza, we can use a factory pattern to create the pizzas.

1. First, we create an abstract Pizza class that will serve as the superclass for all pizzas. 

```
public abstract class Pizza{
    String name;
    String dough;
    String sauce;
    ArrayList<String> toppings = new ArrayList<String>();
    
    public void prepare(){
        System.out.println("Preparing " + name);
        System.out.println("Adding " + dough);
        System.out.println("Adding " + sauce);
        System.out.println("Adding toppings: ");
        
        for(String topping : toppings){
            System.out.println(" " + topping);
        }
    }
    
    public void bake(){
        System.out.println("Baking " + name);
    }
    
    public void cut(){
        System.out.println("Cutting " + name);
    }
    
    public void box(){
        System.out.println("Boxing " + name);
    }
}
```

2. Next, we create a PizzaFactory class that will handle the creation of pizzas. 

```
public class PizzaFactory{
    public Pizza createPizza(String type){
        Pizza pizza = null;
        
        if(type.equals("cheese")){
            pizza = new CheesePizza();
        }
        else if(type.equals("pepperoni")){
            pizza = new PepperoniPizza();
        }
        else if(type.equals("veggie")){
            pizza = new VeggiePizza();
        }
        
        return pizza;
    }
}
```

3. Finally, we create subclasses for each type of pizza that we want to create. 

```
public class CheesePizza extends Pizza{
    public CheesePizza(){
        name = "Cheese Pizza";
        dough = "Regular Crust";
        sauce = "Marinara Pizza Sauce";
        
        toppings.add("Fresh Mozzarella");
        toppings.add("Parmesan");
    }
}
public class PepperoniPizza extends Pizza{
    public PepperoniPizza(){
        name = "Pepperoni Pizza";
        dough = "Thick Crust";
        sauce = "Tomato Sauce";
        
        toppings.add("Pepperoni");
        toppings.add("Sausage");
        toppings.add("Mushrooms");
    }
}
public class VeggiePizza extends Pizza{
    public VeggiePizza(){
        name = "Veggie Pizza";
        dough = "Thin Crust";
        sauce = "Pesto Sauce";
        
        toppings.add("Red Peppers");
        toppings.add("Green Peppers");
        toppings.add("Mushrooms");
    }
}
```

Now, whenever we want to create a pizza, we simply call the createPizza() method of the PizzaFactory class and pass in the type of pizza we want to create. For example:

```
PizzaFactory pizzaFactory = new PizzaFactory();

Pizza cheesePizza = pizzaFactory.createPizza("cheese");
cheesePizza.prepare();
cheesePizza.bake();
cheesePizza.cut();
cheesePizza.box();

Pizza veggiePizza = pizzaFactory.createPizza("veggie");
veggiePizza.prepare();
veggiePizza.bake();
veggiePizza.cut();
veggiePizza.box();
```

This will output:

```
Preparing Cheese Pizza
Adding Regular Crust
Adding Marinara Pizza Sauce
Adding toppings:
 Fresh Mozzarella
 Parmesan
Baking Cheese Pizza
Cutting Cheese Pizza
Boxing Cheese Pizza
Preparing Veggie Pizza
Adding Thin Crust
Adding Pesto Sauce
Adding toppings:
 Red Peppers
 Green Peppers
 Mushrooms
Baking Veggie Pizza
Cutting Veggie Pizza
Boxing Veggie Pizza
```