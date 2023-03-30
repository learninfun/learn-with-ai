1. What is Java serialization? 
Answer: Java serialization is the process of converting an object's state into a byte stream, which can then be saved to a file, sent over the network, or stored in memory.

2. What is Java deserialization? 
Answer: Java deserialization is the process of converting a byte stream into an object instance, restoring the object's state to its original form.

3. What is a serialVersionUID in Java serialization? 
Answer: serialVersionUID is a unique identifier that ensures compatibility of serialized objects between different versions of an application. It is used to check whether the serialized object is compatible with the class definition at deserialization time.

4. What is the difference between Serializable and Externalizable interfaces in Java? 
Answer: Serializable interface is used for basic serialization, which writes the entire object to a stream. Externalizable interface is used for advanced serialization, which provides more control over serialization process and writes only selected objects to the stream.

5. What are the potential security risks associated with deserialization in Java? 
Answer: A malicious attacker can inject harmful code by sending a malicious byte stream at deserialization time, leading to code execution or denial-of-service attacks. Therefore, it is recommended to validate the incoming serialized data and use secure deserialization techniques to prevent these attacks.