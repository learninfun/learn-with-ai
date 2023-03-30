1. What is the difference between margin and padding in CSS?
Answer: Margin is the space outside of an element's border, while padding is the space inside of an element's border.

2. What is the purpose of the CSS display property and give an example of its usage?
Answer: The display property is used to specify the type of box an element is displayed as. For example, setting display: inline-block allows an element to have block-level properties (e.g. height, width, padding) while flowing as an inline element on the page.

3. What is the box-sizing property in CSS and what does it do?
Answer: The box-sizing property specifies how an element's total width and height are calculated. When set to "border-box", the total width and height includes the padding and border, but not the margin.

4. What is the difference between absolute and relative positioning in CSS?
Answer: Absolute positioning positions an element relative to its closest positioned ancestor, while relative positioning positions an element relative to its normal position in the document flow.

5. How can you center an element horizontally and vertically in CSS?
Answer: One way to center an element horizontally and vertically is by using the following CSS:

.element {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
} 

This sets the element's top and left positions to 50% of the parent element's height and width, respectively, and uses the transform property to center it both horizontally and vertically.