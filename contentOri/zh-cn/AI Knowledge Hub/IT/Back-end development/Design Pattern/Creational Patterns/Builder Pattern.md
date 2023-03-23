+++
title = "建造者模式"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Builder Pattern的问题
{{< /ask_chatgpt >}}



1. Builder Pattern是什麼設計模式？它的作用是什麼？

2. Builder Pattern和Factory Pattern有什麼區別？它們何時應該使用？

3. 一個完整的Builder Pattern通常包括哪些元素？請詳細描述這些元素的作用。

4. Builder Pattern可以幫助我們解決哪些設計問題？舉例說明。

5. Builder Pattern的優勢是什麼？它的缺陷又是什麼？如何在開發中選擇是否使用Builder Pattern？   

## 说明知识
{{< ask_chatgpt >}}
说明Builder Pattern并举例
{{< /ask_chatgpt >}}



Builder Pattern是一種軟體設計模式，用於在創建複雜的物件時簡化其構建過程。通過分解物件創建過程，它使得客戶端程式能夠更直觀地構建添加參數。

這個模式包含建造者，指導者和產品。建造者是負責創建複雜物件的介面，指導者負責使用建造者來創建物件，產品是建造完成的複雜物件。

以下是Builder Pattern的範例：

假設我們需要建立一個汽車物件，這個物件有各種屬性，如顏色，引擎類型，座位數量等。使用Builder Pattern，我們可以定義一個建造者介面`CarBuilder`，並實現不同的建造者類別，以構建不同的汽車屬性。

```
public interface CarBuilder {
    void setColor(String color);
    void setEngine(String engine);
    void setSeats(int seats);
    Car getResult();
}

public class SportsCarBuilder implements CarBuilder {
    private Car car;
    public SportsCarBuilder() {
        car = new Car();
    }
    public void setColor(String color) {
        car.setColor(color);
    }
    public void setEngine(String engine) {
        car.setEngine(engine);
    }
    public void setSeats(int seats) {
        car.setSeats(seats);
    }
    public Car getResult() {
        return car;
    }
}

public class SUVBuilder implements CarBuilder {
    private Car car;
    public SUVBuilder() {
        car = new Car();
    }
    public void setColor(String color) {
        car.setColor(color);
    }
    public void setEngine(String engine) {
        car.setEngine(engine);
    }
    public void setSeats(int seats) {
        car.setSeats(seats);
    }
    public Car getResult() {
        return car;
    }
}
```

建造者類別實現`CarBuilder`介面，該介面定義了用於設置汽車屬性的方法。每個建造者都有一個屬於自己的汽車物件實例，它最終將被返回為結果。 

最後，創建指導者類`Director`，它負責將汽車物件創建出來。

```
public class Director {
    private CarBuilder builder;
    public void setBuilder(CarBuilder builder) {
        this.builder = builder;
    }
    public Car getCar() {
        return builder.getResult();
    }
    public void constructCar(String color, String engine, int seats) {
        builder.setColor(color);
        builder.setEngine(engine);
        builder.setSeats(seats);
    }
}

public class Client {
    public static void main(String[] args) {
        Director director = new Director();
        CarBuilder sportsCarBuilder = new SportsCarBuilder();
        CarBuilder suvBuilder = new SUVBuilder();

        director.setBuilder(sportsCarBuilder);
        director.constructCar("Red", "V8", 2);
        Car sportsCar = director.getCar();

        director.setBuilder(suvBuilder);
        director.constructCar("Blue", "V6", 7);
        Car suv = director.getCar();
    }
}
```

客戶端使用指導者來創建汽車物件，使用構造方法傳遞汽車屬性，指導者則使用建造者創建完成的汽車物件。這樣，客戶端不需要知道汽車物件是如何創建的，而是通過使用建造者和指導者來隔離物件創建過程，使得程式碼更加清晰、簡潔。   

## 汇总重点
{{< ask_chatgpt >}}
条列Builder Pattern的重点
{{< /ask_chatgpt >}}



以下是Builder Pattern的重点：

1. Builder Pattern是一种创建型设计模式，它允许你创建复杂对象的不同部分，并将它们组装起来以创建完整的对象。

2. 它通常用于创建复杂的对象，这些对象可能由许多不同的部分组成，这些部分之间可能有很多关联。

3. Builder Pattern将对象的建立过程分为多个步骤，每个步骤都由一个不同的Builder类负责。

4. Builder类将对象的属性保存到自己的成员变量中，最终通过一种方法将这些属性转移到最终对像中。

5. Builder Pattern使得建立一个复杂的对象变得更加简单和可控，同时也能够将对象的表示和建立过程分离开来。

6. Builder Pattern还可以通过方法链（Method Chaining）模式来简化建立复杂对象的代码，使代码更加简洁易读。

7. Builder Pattern的缺点是它需要创建多个Builder类，这可能会导致代码变得复杂和冗长。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Builder Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 请实现一个Builder Pattern，建立一个可定制的饮料店菜单。应当包含饮料种类、甜度、冰块量等选项。

答案：

```python
# 建立一个可定制的饮料店菜单

class Drink:
    def __init__(self, type, sweetness, ice):
        self.type = type
        self.sweetness = sweetness
        self.ice = ice

class DrinkBuilder:
    
    def __init__(self):
        self.type = None
        self.sweetness = None
        self.ice = None
    
    def with_type(self, type):
        self.type = type
        return self
    
    def with_sweetness(self, sweetness):
        self.sweetness = sweetness
        return self
    
    def with_ice(self, ice):
        self.ice = ice
        return self
    
    def build(self):
        return Drink(self.type, self.sweetness, self.ice)

# 创建一个带有定制选项的饮料
drink = DrinkBuilder().with_type('奶茶').with_sweetness('正常').with_ice('少冰').build()

print(f'您的饮料种类为：{drink.type}')
print(f'甜度选择为：{drink.sweetness}')
print(f'冰块量为：{drink.ice}')
```

2. 请实现一个Builder Pattern，建立一个可定制的简历。应该包含姓名、电话、邮箱、教育、工作经验等选项。

答案：

```python
# 创建一个可定制的简历

class Resume:
    
    def __init__(self, name, phone, email):
        self.name = name
        self.phone = phone
        self.email = email
        self.educations = []
        self.work_experiences = []
    
    def add_education(self, education):
        self.educations.append(education)
    
    def add_work_experience(self, work_experience):
        self.work_experiences.append(work_experience)

class Education:
    
    def __init__(self, school, degree, major, start_date, end_date):
        self.school = school
        self.degree = degree
        self.major = major
        self.start_date = start_date
        self.end_date = end_date

class WorkExperience:
    
    def __init__(self, company, title, description, start_date, end_date):
        self.company = company
        self.title = title
        self.description = description
        self.start_date = start_date
        self.end_date = end_date

class ResumeBuilder:
    
    def __init__(self):
        self.name = None
        self.phone = None
        self.email = None
        self.educations = []
        self.work_experiences = []
    
    def with_name(self, name):
        self.name = name
        return self
    
    def with_phone(self, phone):
        self.phone = phone
        return self
    
    def with_email(self, email):
        self.email = email
        return self
    
    def add_education(self, education):
        self.educations.append(education)
        return self
    
    def add_work_experience(self, work_experience):
        self.work_experiences.append(work_experience)
        return self
    
    def build(self):
        resume = Resume(self.name, self.phone, self.email)
        
        for education in self.educations:
            resume.add_education(education)
        
        for work_experience in self.work_experiences:
            resume.add_work_experience(work_experience)
        
        return resume

# 创建一个定制简历

resume = (
    ResumeBuilder()
    .with_name('Mary')
    .with_phone('123-456-7890')
    .with_email('mary@example.com')
    .add_education(
        Education(
            'University of California, Berkeley',
            'Bachelor of Science',
            'Computer Science',
            'September 2018',
            'May 2022'
        )
    )
    .add_work_experience(
        WorkExperience(
            'Google',
            'Software Engineer Intern',
            'Worked on Google Maps project',
            'June 2021',
            'August 2021'
        )
    )
    .build()
)

print(f'姓名：{resume.name}')
print(f'电话：{resume.phone}')
print(f'邮箱：{resume.email}')
print('教育')
for education in resume.educations:
    print(f'{education.degree} in {education.major} from {education.school} ({education.start_date} - {education.end_date})')
    
print('工作经验')
for work_experience in resume.work_experiences:
    print(f'{work_experience.title} at {work_experience.company} ({work_experience.start_date} - {work_experience.end_date}): {work_experience.description}')
```

3. 请实现一个Builder Pattern，建立一个可定制的汽车。应该包含品牌、型号、引擎、车身等选项。

答案：

```python
# 创建一个可定制的汽车

class Car:
    
    def __init__(self, brand, model, engine, body):
        self.brand = brand
        self.model = model
        self.engine = engine
        self.body = body

class Engine:
    
    def __init__(self, name, cylinders, displacement, horsepower):
        self.name = name
        self.cylinders = cylinders
        self.displacement = displacement
        self.horsepower = horsepower

class Body:
    
    def __init__(self, color, transmission, seats):
        self.color = color
        self.transmission = transmission
        self.seats = seats

class CarBuilder:
    
    def __init__(self):
        self.brand = None
        self.model = None
        self.engine = None
        self.body = None
    
    def with_brand(self, brand):
        self.brand = brand
        return self
    
    def with_model(self, model):
        self.model = model
        return self
    
    def with_engine(self, engine):
        self.engine = engine
        return self
    
    def with_body(self, body):
        self.body = body
        return self
    
    def build(self):
        return Car(self.brand, self.model, self.engine, self.body)

# 建立一辆定制汽车

car = (
    CarBuilder()
    .with_brand('BMW')
    .with_model('i8')
    .with_engine(
        Engine(
            'B38',
            3,
            '1.5L',
            228
        )
    )
    .with_body(
        Body(
            'Sophisto Grey',
            '6-speed',
            2
        )
    )
    .build()
)

print(f'品牌：{car.brand}')
print(f'型号：{car.model}')
print(f'引擎：{car.engine.name} ({car.engine.cylinders} cylinders, {car.engine.displacement}, {car.engine.horsepower} horsepower)')
print(f'车身：{car.body.color} {car.body.seats} seats, {car.body.transmission} transmission')
```

4. 请实现一个Builder Pattern，建立一个可定制的订单。应当包括商品、数量、付款方式、送货地址等选项。

答案：

```python
# 建立一个定制订单

class Order:
    
    def __init__(self, items, quantity, payment_method, delivery_address):
        self.items = items
        self.quantity = quantity
        self.payment_method = payment_method
        self.delivery_address = delivery_address

class Item:
    
    def __init__(self, name, price):
        self.name = name
        self.price = price

class OrderBuilder:
    
    def __init__(self):
        self.items = []
        self.quantity = []
        self.payment_method = None
        self.delivery_address = None
    
    def add_item(self, item, quantity):
        self.items.append(item)
        self.quantity.append(quantity)
        return self
    
    def with_payment_method(self, payment_method):
        self.payment_method = payment_method
        return self
    
    def with_delivery_address(self, delivery_address):
        self.delivery_address = delivery_address
        return self
    
    def build(self):
        order_items = [(item, quantity) for item, quantity in zip(self.items, self.quantity)]
        return Order(order_items, self.quantity, self.payment_method, self.delivery_address)

# 创建一个定制的订单

item1 = Item('book1', 10)
item2 = Item('book2', 20)

order = (
    OrderBuilder()
    .add_item(item1, 1)
    .add_item(item2, 2)
    .with_payment_method('Credit Card')
    .with_delivery_address('123 Main St, San Francisco, CA 94109')
    .build()
)

print('订单项目')
total_price = 0
for item, quantity in order.items:
    total_price += item.price * quantity
    print(f'{item.name} x{quantity}: ${item.price * quantity}')

print(f'总价格：${total_price}')
print(f'付款方式：{order.payment_method}')
print(f'送货地址：{order.delivery_address}')
```

5. 请实现一个Builder Pattern，建立一个可定制的室内装修。应当包括材料、光照、风格等选项。

答案：

```python
# 建立一个可定制的室内装修

class Interior:
    
    def __init__(self, material, lighting, style):
        self.material = material
        self.lighting = lighting
        self.style = style

class InteriorBuilder:
    
    def __init__(self):
        self.material = None
        self.lighting = None
        self.style = None
    
    def with_material(self, material):
        self.material = material
        return self
    
    def with_lighting(self, lighting):
        self.lighting = lighting
        return self
    
    def with_style(self, style):
        self.style = style
        return self
    
    def build(self):
        return Interior(self.material, self.lighting, self.style)

# 建立一个定制的室内装修

interior = (
    InteriorBuilder()
    .with_material('Wooden floors')
    .with_lighting('Dim lighting')
    .with_style('Modern')
    .build()
)

print(f'材料：{interior.material}')
print(f'光照：{interior.lighting}')
print(f'风格：{interior.style}')
```

以上是我为您提供的5个Builder Pattern的中等难度问题和答案，希望能够对您有所帮助。   

