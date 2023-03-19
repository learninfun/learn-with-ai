

The Interpreter Pattern is a behavioral pattern that defines the way to interpret a domain-specific language or grammar and on this basis create a tree structure for that language. It provides a way to define different language expressions while also defining how to evaluate them.

For instance, consider a simple mathematical expression like (3 + 4) * 5. Here, the expression can be interpreted in different ways. The Interpreter pattern provides a structured way to read and evaluate expressions like this. Using this pattern, we can create a parser that can read the expression and create an appropriate tree structure based on the expression.

Here is an example of implementing this pattern:

Suppose we have to create a mathematical expression evaluator that can evaluate the expressions involving addition and subtraction operation. We can use the Interpreter pattern to implement this. In this example, we will create an interface called Expression that consists of two methods, interpret and toString(). The interpret method will evaluate the expression and return the result.

public interface Expression {

    int interpret();

    String toString();

}

We will have two implementations of that Expression interface, namely PlusExpression and MinusExpression. The code for these is given below:

public class PlusExpression implements Expression {

    private Expression leftExpression;

    private Expression rightExpression;

    public PlusExpression(Expression leftExpression, Expression rightExpression) {
        this.leftExpression = leftExpression;
        this.rightExpression = rightExpression;
    }

    @Override
    public int interpret() {
        return leftExpression.interpret() + rightExpression.interpret();
    }

    @Override
    public String toString() {
        return "(" + leftExpression.toString() + " + " + rightExpression.toString() + ")";
    }

}


public class MinusExpression implements Expression {

    private Expression leftExpression;

    private Expression rightExpression;

    public MinusExpression(Expression leftExpression, Expression rightExpression) {
        this.leftExpression = leftExpression;
        this.rightExpression = rightExpression;
    }

    @Override
    public int interpret() {
        return leftExpression.interpret() - rightExpression.interpret();
    }

    @Override
    public String toString() {
        return "(" + leftExpression.toString() + " - " + rightExpression.toString() + ")";
    }
}

Using these classes, we can now create expressions, evaluate them and get the result. Here is some sample code to do that:

Expression exp = new PlusExpression(
    new PlusExpression(new NumberExpression(3), new NumberExpression(4)),
    new NumberExpression(5));

System.out.println(exp.toString() + " = " + exp.interpret());

In this example, we are creating an expression which adds 3 and 4 and then subtracts their sum from 5. The expected output is:

((3 + 4) + 5) = 12

Overall, the Interpreter Pattern provides a flexible way to define and evaluate expressions with different operations. It helps in creating a robust and maintainable code that can be extended easily.