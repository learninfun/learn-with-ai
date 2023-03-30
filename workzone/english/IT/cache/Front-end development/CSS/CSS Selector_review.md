1. What is the difference between the descendant selector and the direct descendant selector in CSS? 
Answer: The descendant selector selects all elements that are descendants of a specified parent element, while the direct descendant selector selects only the immediate child elements of a parent element. 

2. Can you use the class selector and the ID selector together in the same CSS rule? 
Answer: Yes, you can combine different selectors in a CSS rule. For example, you can target an element with a specific class and ID like this: #myID.myClass { /* CSS properties */ } 

3. How does the attribute selector work in CSS? 
Answer: The attribute selector allows you to target elements based on the presence or value of an attribute. For example, you can select all anchor elements with a specific href attribute value like this: a[href="https://example.com"] { /* CSS properties */ } 

4. What is the difference between the :nth-child() selector and the :nth-of-type() selector in CSS? 
Answer: The :nth-child() selector targets elements based on their position relative to their parent container, while the :nth-of-type() selector targets elements based on their position relative to their siblings of the same element type. 

5. How can you use the + and ~ combinators in CSS to target elements? 
Answer: The + combinator selects the next sibling element that comes immediately after the first element, while the ~ combinator selects all the following sibling elements that match the given selector. For example, you can select all paragraphs that come immediately after a heading like this: h2 + p { /* CSS properties */ }