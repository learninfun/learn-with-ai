

DOM manipulation refers to the process of changing or modifying the document object model (DOM) of a web page using JavaScript. It involves accessing and modifying the elements and attributes of a web page in real-time, without the need for a page refresh. There are numerous methods and tools available to manipulate the DOM, including jQuery.

For example, imagine you have an HTML document that contains a button with the id "myButton". Using DOM manipulation, you could write a JavaScript function that is triggered when the button is clicked, and then use jQuery to add a red background color to all paragraphs on the page:

```
$(document).ready(function() {
  $("#myButton").click(function() {
    $("p").css("background-color", "red");
  });
});
```

This code binds a click event to the button with the ID "myButton", and when clicked, it selects all paragraph elements on the page and adds a red background color to them. This is just one example of the many ways in which DOM manipulation can be used to enhance the functionality and appearance of web pages.