

1. What is the difference between event bubbling and event capturing in JavaScript? 
Answer: Event bubbling is when an event is first captured by the innermost element and then propagated to its parent elements. Event capturing, on the other hand, is when the outermost element captures the event and then passes it to the innermost element. 

2. How do you prevent an event from bubbling up the DOM tree? 
Answer: You can prevent an event from bubbling up the DOM tree using the event.stopPropagation() method. 

3. What is event delegation in JavaScript? 
Answer: Event delegation is a technique where you attach an event listener to a parent element, instead of attaching an event listener to individual child elements. This is useful when you have a large number of child elements or dynamically created elements. 

4. What is the difference between passive and non-passive event listeners? 
Answer: Passive event listeners are a way to improve scroll performance by notifying the browser that the listener will not prevent scrolling. Non-passive event listeners, on the other hand, block scrolling until the listener has finished executing. 

5. How do you attach an event listener to multiple elements in JavaScript? 
Answer: You can attach an event listener to multiple elements using the forEach method to loop through and attach the listener to each element in the array. Alternatively, you can use event delegation by attaching the listener to a parent element and filtering for the specific child elements using event.target.