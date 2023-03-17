

JS Event Handling refers to the process of writing code that responds to certain actions or events that occur in a web page, such as clicks or mouse movements. 

To handle events properly, JavaScript provides a set of APIs to attach an event listener to an HTML element. These APIs are used to listen for events, such as button clicks, page loading, and form submissions.

Here is an example of JS event handling using the button click event:

HTML code:

```
<button id="myButton">Click me</button>
```

JS code:

```
const button = document.querySelector('#myButton');

button.addEventListener('click', function() {
   alert('You clicked the button!');
});
```

This code attaches an event listener to the `myButton` button element. When the button is clicked, the code inside the event listener function is executed, which in this case shows an alert message saying "You clicked the button!".