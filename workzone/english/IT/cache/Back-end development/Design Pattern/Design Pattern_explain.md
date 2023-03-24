

Design patterns are reusable solutions to common programming problems which provide solutions in a standardize way to common software development problems. It means Design patterns help you solve common design problems and increase code reusability by providing tested and tried approaches.

A well-known example of the design pattern is the Singleton pattern. The Singleton pattern is used when you need only a single instance of a class, especially when the class takes up too many resources, such as database connections. It ensures that there is only one instance of a class throughout the application and provides a global point of access to that instance. In PHP, the implementation of Singleton design pattern looks like the below example:

class Singleton
{
   // Hold the class instance.
   private static $instance = null;

   private function __construct()
   {
      // The expensive process (e.g.,db connection) goes here
   }

   public static function getInstance()
   {
      if (self::$instance == null)
      {
         self::$instance = new Singleton();
      }

      return self::$instance;
   }

   public function avoidClone()
   {
        trigger_error('Clone is not allowed ', E_USER_ERROR);
   }

   public function avoidWakeUpSerializing()
   {
        throw new \Exception('Wake up and Serializing is not allowed ');
   }
}

//how to use singleton class
$instance1 = Singleton::getInstance();
$instance2 = Singleton::getInstance();
$instance3 = Singleton::getInstance();

var_dump($instance1 === $instance2); // true

var_dump($instance1 === $instance3); // true

$instance4 = clone $instance1; // Cloning forbidden
$instance5 = unserialize(serialize($instance1)); //Unserialization forbidden
$instance6 = new Singleton(); //Constructor is private, can't be instantiated from outside.


In this example, the Singleton design pattern is used to limit the number of instances that can be created from a class to one. This ensures that there is only one instance of the class in the application, which can be accessed by other classes that need it.