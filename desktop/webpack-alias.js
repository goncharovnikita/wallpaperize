const HtmlPlugin = require('html-webpack-plugin');

module.exports = {
  entry: {
    renderer: __dirname + '/src/renderer/index.tsx',
    vendor: __dirname + '/src/vendor/vendor.ts'
  },
  node: {
    __dirname: false
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js', '.jsx'],
    alias: {
      '@app': __dirname + '/src/renderer/',
      '@approot': __dirname + '/src/'
    }
  },
  module: {
    rules: [
      {
        test: /\.sass/,
        use: ['style-loader', 'css-loader', 'sass-loader']
      },
      {
        test: /\.scss/,
        use: ['css-loader', 'sass-loader']
      },
      {
        test: /.(ttf|otf|eot|svg|woff(2)?)(\?[a-z0-9]+)?$/,
        use: [
          {
            loader: 'file-loader',
            options: {
              name: '[name].[ext]',
              outputPath: 'fonts/', // where the fonts will go
              publicPath: '../' // override the default path
            }
          }
        ]
      }
    ]
  },
  plugins: [new HtmlPlugin({ template: 'src/index.html' })]
};
