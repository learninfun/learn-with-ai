

1. 在一個車廠中，負責檢查並修理車輛的工人可以分為三種：檢查員、機械師、高級技工。當一輛車進入車廠時，必須先由檢查員檢查問題，如果問題很小，則由檢查員進行修理；如果問題比較嚴重，則必須交由機械師進行修理；如果問題非常複雜，才需要交由高級技工來修理。請使用Chain of Responsibility Pattern來完成此問題。

答案：請參考以下範例Code

2. 在一家小吃店中，有三種員工：收銀員、製作食物的廚師和清潔人員。當一位客人下訂單時，訂單會傳遞給收銀員，收銀員會計算訂單的費用，然後將訂單傳遞給廚師，廚師會根據訂單製作食物，最後將訂單交給清潔員來清理桌子。請使用Chain of Responsibility Pattern來完成此問題。

答案：請參考以下範例Code

3. 在一家醫院中，病人來到接待處報到，接待員工會詢問病人的基本情況，然後將病人送到醫生那裡進行診斷。如果病情比較嚴重，則醫生會將病人送到手術室進行手術。請使用Chain of Responsibility Pattern來完成此問題。

答案：請參考以下範例Code

4. 在一個系統中，有三種日誌級別：information、warning、error。每當系統遇到不同的情況，都會產生不同級別的日誌，例如：information紀錄系統啟動訊息、warning紀錄系統運行異常等等。請使用Chain of Responsibility Pattern來完成此問題。

答案：請參考以下範例Code

5. 在一個電商平台上，用戶可以對商品進行評論，評論可以分為一般評論、高級認證評論和行業大咖評論。對於不同級別的評論，系統需要進行不同的處理，例如：一般評論可以直接顯示在商品頁面上，高級認證評論需要審核通過才能顯示，行業大咖評論則可以得到處理人員的專門回復。請使用Chain of Responsibility Pattern來完成此問題。

答案：請參考以下範例Code

以下是範例Code：

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
使用範例：
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
輸出結果：
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
使用範例：
<?php
$cashier = new Cashier();
$chef = new Chef();
$cleaner = new Cleaner();
$cashier->setNext($chef);
$chef->setNext($cleaner);
$cashier->work('hamburger');
?>
輸出結果：
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
使用範例：
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
輸出結果：
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
使用範例：
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
輸出結果：
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
使用範例：
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
輸出結果：
Basic review is displayed directly.
Certified review has to be approved before being displayed.
Pro review is sent to a specialist for reply.