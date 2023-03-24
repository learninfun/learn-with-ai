

1. 请问什么是 Tree Shaking？

答：Tree Shaking 是一个在模组打包时，移除没有被使用的程式码的技术。

2. 请问什么是 Code Splitting？

答：Code Splitting 是将一个大的 JavaScript 模组分成多个小的模组，并且只加载需要的部分，从而减少整体页面载入时间。

3. 请问 Webpack 的 entry、output、loader 与 plugin 分别有什么作用？

答：entry 为定义进入点，output 为定义输出文件的位置与文件名称，loader 为处理非 JavaScript 文件，将其转换成 JavaScript 可以执行的形式，plugin 则是扩展 Webpack 的功能，并进行更进阶的设定。

4. 请问 Webpack 中常用的优化方式有哪些？

答：常用的优化方式有：使用 Tree Shaking、Code Splitting、使用缓存、使用懒加载（Lazy Loading）等。

5. 请问 Rollup 与 Webpack 有哪些不同点？

答：Rollup 的优点是产生的 bundle 更小、更快，并且更适合用于打包开源库等程式库，而 Webpack 则更适合用于打包应用程式，并且支援更多的优化方式，如 Tree Shaking、Code Splitting 等。