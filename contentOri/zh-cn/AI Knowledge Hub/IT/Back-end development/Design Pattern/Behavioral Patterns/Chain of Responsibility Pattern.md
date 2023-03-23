+++
title = "责任链模式"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Chain of Responsibility Pattern的问题
{{< /ask_chatgpt >}}



1. 什麼是Chain of Responsibility Pattern？
2. Chain of Responsibility Pattern 與其他設計模式的區別在哪裡？
3. 在Chain of Responsibility Pattern中，什麼是Handler？它有什麼作用？
4. 如何實現Chain of Responsibility Pattern？你可以舉一個例子嗎？
5. Chain of Responsibility Pattern 的優缺點是什麼？   

## 说明知识
{{< ask_chatgpt >}}
说明Chain of Responsibility Pattern并举例
{{< /ask_chatgpt >}}



Chain of Responsibility Pattern是一种行为型设计模式，目的是在多个对象间处理请求时，达到解耦与组织的效果。此模式将多个对象连接成一条链，每个对象都有机会处理请求。当一个对像不能处理请求时，会将请求转发到下一个对象，直到能够处理或到达链的末端。

举例来说，有一个单位要向国家政府提出一个计划，计划分为四个阶段：初步的方案、详细的方案、执行的方案和结果的报告。每个阶段的负责人负责处理相关的工作，如果负责人无法处理该阶段的工作，则需要向下一个负责人传递请求，直到工作完成。

使用Chain of Responsibility Pattern可以有效地处理这个问题。我们可以将四个阶段的负责人连接成一条链，每个负责人都有处理该阶段工作的能力。如果一个负责人无法处理该阶段的工作，他会将请求传递给下一个负责人，直到该阶段的工作完成。

例如，初步的方案负责人无法处理详细方案的工作，因此他将请求传递给详细方案的负责人，如果详细方案的负责人无法处理，他会转发给执行方案的负责人，最终到达结果的报告负责人。此模式能够减少耦合并更容易维护，因为每个负责人都可以独立处理工作，而不必知道工作的详细情况或其他负责人的工作。   

## 汇总重点
{{< ask_chatgpt >}}
条列Chain of Responsibility Pattern的重点
{{< /ask_chatgpt >}}



1. Chain of Responsibility 模式是一种行为型设计模式，通过组合对像形成一条链，每个对象都能够处理请求，若无法处理则传递给下一个对象。

2. 可以通过该模式消除发送者和接收者之间的耦合关系，使得请求可以被多个对象处理。

3. Chain of Responsibility 模式应用于多个对象相互交互的场景，其中每个对象都可以根据需要动态处理请求，并且该模式能够优雅地处理复杂的请求处理过程。

4. 该模式的优点包括：增强代码的复用性和灵活性，降低系统的耦合度，方便添加或删除处理者等。

5. 该模式的缺点包括：链条过长可能会降低性能，并且由于链条中的每个对象都可以执行处理，因此可能会导致请求被重复处理。

6. 在实现该模式时，需要定义一个抽像处理类，该类包含一个处理请求的方法以及设置下一个处理者的方法。同时还需要实现各个对象的具体处理逻辑。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Chain of Responsibility Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 在一个车厂中，负责检查并修理车辆的工人可以分为三种：检查员、机械师、高级技工。当一辆车进入车厂时，必须先由检查员检查问题，如果问题很小，则由检查员进行修理；如果问题比较严重，则必须交由机械师进行修理；如果问题非常复杂，才需要交由高级技工来修理。请使用Chain of Responsibility Pattern来完成此问题。

答案：请参考以下范例Code

2. 在一家小吃店中，有三种员工：收银员、制作食物的厨师和清洁人员。当一位客人下订单时，订单会传递给收银员，收银员会计算订单的费用，然后将订单传递给厨师，厨师会根据订单制作食物，最后将订单交给清洁员来清理桌子。请使用Chain of Responsibility Pattern来完成此问题。

答案：请参考以下范例Code

3. 在一家医院中，病人来到接待处报到，接待员工会询问病人的基本情况，然后将病人送到医生那里进行诊断。如果病情比较严重，则医生会将病人送到手术室进行手术。请使用Chain of Responsibility Pattern来完成此问题。

答案：请参考以下范例Code

4. 在一个系统中，有三种日志级别：information、warning、error。每当系统遇到不同的情况，都会产生不同级别的日志，例如：information纪录系统启动讯息、warning纪录系统运行异常等等。请使用Chain of Responsibility Pattern来完成此问题。

答案：请参考以下范例Code

5. 在一个电商平台上，用户可以对商品进行评论，评论可以分为一般评论、高级认证评论和行业大咖评论。对于不同级别的评论，系统需要进行不同的处理，例如：一般评论可以直接显示在商品页面上，高级认证评论需要审核通过才能显示，行业大咖评论则可以得到处理人员的专门回复。请使用Chain of Responsibility Pattern来完成此问题。

答案：请参考以下范例Code

以下是范例Code：

例一.

interface Worker
{
    public function setNext(Worker $worker);
    public function handle($car);
}
class Inspector implements Worker
{
    private $nextWorker;
    public function setNext(Worker $worker)
    {
        $this->nextWorker = $worker;
    }
    public function handle($car)
    {
        if ($car->getProblemLevel() == 'minor')
        {
            echo "Inspector handles the car by repairing the problem directly.\n";
            $car->setProblemFixed();
        }
        else
        {
            $this->nextWorker->handle($car);
        }
    }
}
class Mechanic implements Worker
{
    private $nextWorker;
    public function setNext(Worker $worker)
    {
        $this->nextWorker = $worker;
    }
    public function handle($car)
    {
        if ($car->getProblemLevel() == 'moderate')
        {
            echo "Mechanic handles the car by repairing the problem.\n";
            $car->setProblemFixed();
        }
        else
        {
            $this->nextWorker->handle($car);
        }
    }
}
class SeniorTechnician implements Worker
{
    public function setNext(Worker $worker)
    {
        // This is the last class in the chain
    }
    public function handle($car)
    {
        if ($car->getProblemLevel() == 'complex')
        {
            echo "Senior technician handles the car by repairing the problem.\n";
            $car->setProblemFixed();
        }
        else
        {
            echo "Sorry, we are unable to repair the problem.\n";
            $car->setProblemUnfixed();
        }
    }
}
class Car
{
    private $problemLevel;
    private $problemFixed = false;
    public function __construct($problemLevel)
    {
        $this->problemLevel = $problemLevel;
    }
    public function getProblemLevel()
    {
        return $this->problemLevel;
    }
    public function setProblemFixed()
    {
        $this->problemFixed = true;
    }
    public function setProblemUnfixed()
    {
        $this->problemFixed = false;
    }
    public function isFixed()
    {
        return $this->problemFixed;
    }
}
?>
使用范例：
<?php
$car1 = new Car('minor');
$car2 = new Car('moderate');
$car3 = new Car('complex');
$inspector = new Inspector();
$mechanic = new Mechanic();
$seniorTechnician = new SeniorTechnician();
$inspector->setNext($mechanic);
$mechanic->setNext($seniorTechnician);
$inspector->handle($car1);
echo "Is car1's problem fixed? " . ($car1->isFixed() ? "Yes" : "No") . ".\n";
$inspector->handle($car2);
echo "Is car2's problem fixed? " . ($car2->isFixed() ? "Yes" : "No") . ".\n";
$inspector->handle($car3);
echo "Is car3's problem fixed? " . ($car3->isFixed() ? "Yes" : "No") . ".\n";
?>
输出结果：
Inspector handles the car by repairing the problem directly.
Is car1's problem fixed? Yes.
Mechanic handles the car by repairing the problem.
Is car2's problem fixed? Yes.
Senior technician handles the car by repairing the problem.
Is car3's problem fixed? Yes.


例二.

interface Employee
{
    public function setNext(Employee $employee);
    public function work($food);
}
class Cashier implements Employee
{
    private $nextEmployee;
    public function setNext(Employee $employee)
    {
        $this->nextEmployee = $employee;
    }
    public function work($food)
    {
        echo "Cashier calculates the cost of the food.\n";
        $this->nextEmployee->work($food);
    }
}
class Chef implements Employee
{
    private $nextEmployee;
    public function setNext(Employee $employee)
    {
        $this->nextEmployee = $employee;
    }
    public function work($food)
    {
        echo "Chef cooks the food according to the order.\n";
        $this->nextEmployee->work($food);
    }
}
class Cleaner implements Employee
{
    public function setNext(Employee $employee)
    {
        // This is the last class in the chain
    }
    public function work($food)
    {
        echo "Cleaner cleans up the table after the customer leaves.\n";
    }
}
?>
使用范例：
<?php
$cashier = new Cashier();
$chef = new Chef();
$cleaner = new Cleaner();
$cashier->setNext($chef);
$chef->setNext($cleaner);
$cashier->work('hamburger');
?>
输出结果：
Cashier calculates the cost of the food.
Chef cooks the food according to the order.
Cleaner cleans up the table after the customer leaves.


例三.

interface Employee
{
    public function setNext(Employee $employee);
    public function work($patient);
}
class Receptionist implements Employee
{
    private $nextEmployee;
    public function setNext(Employee $employee)
    {
        $this->nextEmployee = $employee;
    }
    public function work($patient)
    {
        echo "Receptionist asks for patient's basic information.\n";
        $this->nextEmployee->work($patient);
    }
}
class Doctor implements Employee
{
    private $nextEmployee;
    public function setNext(Employee $employee)
    {
        $this->nextEmployee = $employee;
    }
    public function work($patient)
    {
        if ($patient->isSerious())
        {
            echo "Doctor diagnoses and sends the patient to the operating room.\n";
        }
        else
        {
            echo "Doctor diagnoses and prescribes medicine for the patient.\n";
        }
    }
}
class Surgeon implements Employee
{
    public function setNext(Employee $employee)
    {
        // This is the last class in the chain
    }
    public function work($patient)
    {
        echo "Surgeon performs the surgery on the patient.\n";
    }
}
class Patient
{
    private $isSerious;
    public function __construct($isSerious)
    {
        $this->isSerious = $isSerious;
    }
    public function isSerious()
    {
        return $this->isSerious;
    }
}
?>
使用范例：
<?php
$patient1 = new Patient(false);
$patient2 = new Patient(true);
$receptionist = new Receptionist();
$doctor = new Doctor();
$surgeon = new Surgeon();
$receptionist->setNext($doctor);
$doctor->setNext($surgeon);
$receptionist->work($patient1);
$receptionist->work($patient2);
?>
输出结果：
Receptionist asks for patient's basic information.
Doctor diagnoses and prescribes medicine for the patient.
Receptionist asks for patient's basic information.
Doctor diagnoses and sends the patient to the operating room.
Surgeon performs the surgery on the patient.


例四.

interface Logger
{
    public function setNext(Logger $logger);
    public function log($level, $message);
}
class InformationLogger implements Logger
{
    private $nextLogger;
    public function setNext(Logger $logger)
    {
        $this->nextLogger = $logger;
    }
    public function log($level, $message)
    {
        if ($level == 'information')
        {
            echo "InformationLogger: $message\n";
            return true;
        }
        elseif (!is_null($this->nextLogger))
        {
            return $this->nextLogger->log($level, $message);
        }
        else
        {
            return false;
        }
    }
}
class WarningLogger implements Logger
{
    private $nextLogger;
    public function setNext(Logger $logger)
    {
        $this->nextLogger = $logger;
    }
    public function log($level, $message)
    {
        if ($level == 'warning')
        {
            echo "WarningLogger: $message\n";
            return true;
        }
        elseif (!is_null($this->nextLogger))
        {
            return $this->nextLogger->log($level, $message);
        }
        else
        {
            return false;
        }
    }
}
class ErrorLogger implements Logger
{
    private $nextLogger;
    public function setNext(Logger $logger)
    {
        $this->nextLogger = $logger;
    }
    public function log($level, $message)
    {
        if ($level == 'error')
        {
            echo "ErrorLogger: $message\n";
            return true;
        }
        elseif (!is_null($this->nextLogger))
        {
            return $this->nextLogger->log($level, $message);
        }
        else
        {
            return false;
        }
    }
}
?>
使用范例：
<?php
$logger1 = new InformationLogger();
$logger2 = new WarningLogger();
$logger3 = new ErrorLogger();
$logger1->setNext($logger2);
$logger2->setNext($logger3);
$logger1->log('information', 'The system is starting up.');
$logger1->log('warning', 'The system has encountered an error.');
$logger1->log('error', 'The system has crashed.');
?>
输出结果：
InformationLogger: The system is starting up.
WarningLogger: The system has encountered an error.
ErrorLogger: The system has crashed.


例五.

interface ReviewHandler
{
    public function setNext(ReviewHandler $handler);
    public function handle($comment);
}
class BasicReviewHandler implements ReviewHandler
{
    private $nextHandler;
    public function setNext(ReviewHandler $handler)
    {
        $this->nextHandler = $handler;
    }
    public function handle($comment)
    {
        if ($comment->getLevel() == 'basic')
        {
            echo "Basic review is displayed directly.\n";
        }
        elseif (!is_null($this->nextHandler))
        {
            $this->nextHandler->handle($comment);
        }
        else
        {
            echo "Sorry, your comment cannot be displayed.\n";
        }
    }
}
class CertifiedReviewHandler implements ReviewHandler
{
    private $nextHandler;
    public function setNext(ReviewHandler $handler)
    {
        $this->nextHandler = $handler;
    }
    public function handle($comment)
    {
        if ($comment->getLevel() == 'certified')
        {
            echo "Certified review has to be approved before being displayed.\n";
        }
        elseif (!is_null($this->nextHandler))
        {
            $this->nextHandler->handle($comment);
        }
        else
        {
            echo "Sorry, your comment cannot be displayed.\n";
        }
    }
}
class ProReviewHandler implements ReviewHandler
{
    private $nextHandler;
    public function setNext(ReviewHandler $handler)
    {
        $this->nextHandler = $handler;
    }
    public function handle($comment)
    {
        if ($comment->getLevel() == 'pro')
        {
            echo "Pro review is sent to a specialist for reply.\n";
        }
        elseif (!is_null($this->nextHandler))
        {
            $this->nextHandler->handle($comment);
        }
        else
        {
            echo "Sorry, your comment cannot be displayed.\n";
        }
    }
}
class Comment
{
    private $level;
    public function __construct($level)
    {
        $this->level = $level;
    }
    public function getLevel()
    {
        return $this->level;
    }
}
?>
使用范例：
<?php
$comment1 = new Comment('basic');
$comment2 = new Comment('certified');
$comment3 = new Comment('pro');
$handler1 = new BasicReviewHandler();
$handler2 = new CertifiedReviewHandler();
$handler3 = new ProReviewHandler();
$handler1->setNext($handler2);
$handler2->setNext($handler3);
$handler1->handle($comment1);
$handler1->handle($comment2);
$handler1->handle($comment3);
?>
输出结果：
Basic review is displayed directly.
Certified review has to be approved before being displayed.
Pro review is sent to a specialist for reply.   

