

Java annotations are a form of metadata that can be added to Java code to provide additional information about the code to the Java compiler or other tools. Annotations are defined using the "@" symbol followed by the annotation name and any additional parameters, and are typically placed just before the declaration of a class, method, or variable.

Here is an example of a Java annotation:

```
@Deprecated
public void oldMethod() {
    // code for the old method here
}
```

In this example, the annotation "@Deprecated" is added to the declaration of a method named "oldMethod". This annotation indicates that the method is no longer recommended for use and may be removed in a future version of the code. By adding this annotation, developers who use the method will see a warning message or error in their Java IDE, reminding them to update their code to use a newer alternative method.