

1. How do you configure a Spring Boot application to use a custom, non-default port?
Answer: 
To set a custom port, add the following line to your application.properties or application.yml file: server.port=8081 (replace 8081 with your preferred port number)

2. What is the difference between @ComponentScan and @SpringBootApplication annotations?
Answer: 
@ComponentScan is used to specify which package(s) should be scanned for Spring components (such as classes annotated with @Controller, @Repository, or @Service). @SpringBootApplication includes @ComponentScan and other auto-configuration annotations to specify the Spring Boot application as the root of all application context beans.

3. How can you access Spring Boot Actuator endpoints?
Answer: 
Spring Boot Actuator endpoints are available at the /actuator URL path of your application. For example, http://localhost:8080/actuator/health or http://localhost:8080/actuator/info.

4. Can you create a custom error page for your Spring Boot application?
Answer: 
Yes, you can create a custom error page by creating an HTML file named error.html (or error.whateverExtensionYouPrefer) in your src/main/resources/templates/ directory. This file will be used as the error page for all HTTP status codes.

5. What is the purpose of @EnableAutoConfiguration in Spring Boot?
Answer: 
@EnableAutoConfiguration tells Spring Boot to automatically configure the Spring application context based on the classpath dependencies and other settings in the application.properties or application.yml files. This allows developers to focus on developing their business logic without having to worry about configuring the lower-level framework components.