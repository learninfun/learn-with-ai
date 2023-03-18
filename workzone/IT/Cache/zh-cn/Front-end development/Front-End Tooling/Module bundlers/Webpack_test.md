

1. 如何在Webpack中使用CSS模组化？
答案： 在Webpack中，可以使用CSS模组化（CSS Modules）来避免CSS样式冲突的问题。在Webpack的配置文件中，需要添加以下内容：
```javascript
{
  test: /\.css$/,
  use: [
    'style-loader',
    {
      loader: 'css-loader',
      options: {
        modules: true
      }
    }
  ]
}
```

2. 如何在Webpack中实现代码分割？
答案： 在Webpack中可以使用`import()`函数来实现代码分割。使用`import()`函数动态加载模块时，Webpack会自动将引用的模块进行分割，生成不同的chunk。（注意：需要使用babel插件，如@babel/plugin-syntax-dynamic-import）

3. 如何在Webpack中实现热加载（Hot Module Replacement）？
答案： 可以使用Webpack提供的`HotModuleReplacementPlugin`插件实现热加载。在Webpack的配置文件中，需要添加以下内容：
```javascript
const webpack = require('webpack');
module.exports = {
  // ...
  devServer: {
    hot: true
  },
  plugins: [
    new webpack.HotModuleReplacementPlugin()
  ]
};
```

4. 如何在Webpack中实现对图片等静态资源的处理？
答案： 可以使用`file-loader`或`url-loader`来处理图片等静态资源。`file-loader`会将文件复制到输出目录中，而`url-loader`则会将小文件转换成Base64编码嵌入到JS中。（注意：需要使用对应的loader配置相应的模块）

5. 如何在Webpack中实现代码压缩（minification）？
答案： 在Webpack中可以使用`uglifyjs-webpack-plugin`插件来实现代码压缩。在Webpack的配置文件中，需要添加以下内容：
```javascript
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
module.exports = {
  // ...
  optimization: {
    minimizer: [
      new UglifyJsPlugin({
        cache: true,
        parallel: true,
        sourceMap: true
      })
    ]
  }
};
```