

1. What is a JS module and how does it differ from traditional scripts?
- A JS module is a set of code that can be reused in multiple locations without affecting other code in the same file. It is different from traditional scripts in that each module has its own scope, and it can be imported and exported to other modules as needed.

2. How do you import a module in JavaScript?
- To import a module in JavaScript, you can use the 'import' keyword followed by the name of the module you want to import. For example: 'import { myFunction } from './myModule';'.

3. What are the benefits of using JS modules in a project?
- JS modules can help make code more modular and easier to maintain by breaking it up into smaller, reusable pieces. They can also improve performance by only loading the modules that are needed for a specific piece of code.

4. Can you have multiple exports in a single JS module?
- Yes, a single JS module can have multiple exports by using the 'export' keyword before each export. For example: 'export const myConstant = 'Hello world'; export function myFunction() { // code here }'.

5. How do you create a default export in a JS module?
- To create a default export in a JS module, you can use the 'export default' keyword followed by the value you want to export. For example: 'const myVariable = 'Hello world'; export default myVariable;'.