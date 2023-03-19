

Semantic HTML refers to the practice of using HTML tags that reflect the meaning or purpose of the content they contain, rather than just using generic container tags. This approach helps improve the accessibility and usability of web pages, as well as making the content more meaningful and understandable to search engines.

For example, instead of using a div tag to group a list of items, you could use the semantic tag ul (unordered list) to indicate that the items are part of a list. Similarly, instead of using a div tag to contain a navigation menu, you could use the semantic tag nav to indicate that it is a navigation block. Other examples of semantic tags include header, section, article, footer, and figure – each of which conveys a specific meaning about the content they contain. 

Here’s an example of semantic HTML:

```
<article>
  <header>
    <h2>The Benefits of Eating Healthy</h2>
    <p>Published on <time datetime="2022-06-01">June 1, 2022</time></p>
  </header>
  <section>
    <p>Eating a balanced diet that includes plenty of fruits, vegetables, whole grains, and lean protein can have many benefits for your health.</p>
    <ul>
      <li>Improves digestion and gut health</li>
      <li>Reduces the risk of chronic diseases</li>
      <li>Boosts energy levels and brain function</li>
      <li>Helps maintain a healthy weight</li>
    </ul>
  </section>
  <footer>
    <p>Written by Jane Doe</p>
  </footer>
</article>
```

In this example, the article tag is used to contain a complete article on the topic of healthy eating. The header and footer tags are used to add metadata information such as the title of the article, publication date, and author name. The section tag is used to group together content that relates to the benefits of eating healthy, while the ul tag is used to create a bulleted list of these benefits. Overall, this code demonstrates the use of semantic HTML to convey meaning and maximize the accessibility and usability of the content.