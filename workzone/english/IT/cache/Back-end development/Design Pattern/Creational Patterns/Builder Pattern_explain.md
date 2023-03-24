

The Builder Pattern is a creational pattern that separates the creation of an object from its representation. In other words, it allows you to create an object step by step and set different values for its properties without having to expose its construction logic to the client.

An example of the builder pattern would be building a custom burger at a restaurant. The client can choose the type of bun, patty, toppings and sauces, and the builder will assemble the burger in the correct order. The builder can also validate the order to ensure that all required components are present.

Here's a simplified code example:

```
public class Burger {
    private String bun;
    private String patty;
    private List<String> toppings;
    private List<String> sauces;

    // Builder
    public static class Builder {
        private String bun;
        private String patty;
        private List<String> toppings = new ArrayList<>();
        private List<String> sauces = new ArrayList<>();

        public Builder(String bun, String patty) {
            this.bun = bun;
            this.patty = patty;
        }

        public Builder addTopping(String topping) {
            toppings.add(topping);
            return this;
        }

        public Builder addSauce(String sauce) {
            sauces.add(sauce);
            return this;
        }

        public Burger build() {
            if (toppings.isEmpty() || sauces.isEmpty()) {
                throw new IllegalStateException("Burger is incomplete");
            }
            Burger burger = new Burger();
            burger.bun = this.bun;
            burger.patty = this.patty;
            burger.toppings = this.toppings;
            burger.sauces = this.sauces;
            return burger;
        }
    }

    // Getters
    public String getBun() {
        return bun;
    }

    public String getPatty() {
        return patty;
    }

    public List<String> getToppings() {
        return toppings;
    }

    public List<String> getSauces() {
        return sauces;
    }
}
```

The client can now create a burger like this:

```
Burger burger = new Burger.Builder("sesame seed bun", "beef patty")
                    .addTopping("lettuce")
                    .addTopping("tomato")
                    .addSauce("ketchup")
                    .addSauce("mayonnaise")
                    .build();
```

If the client misses any of the required components, like toppings or sauces, the builder will throw an exception. This ensures that the burger is always built according to the restaurant's standards.