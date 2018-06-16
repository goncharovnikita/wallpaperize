import * as webpack from 'webpack';
import * as HtmlPlugin from 'html-webpack-plugin';
import config from './webpack.config';
import * as path from 'path';

const rendererConf: webpack.Configuration = {
  ...config,
  target: 'electron-renderer',
  entry: {
    renderer: __dirname + '/src/index.tsx',
    vendor: __dirname + '/src/vendor/vendor.ts'
  },
  mode: 'development',
  plugins: [new HtmlPlugin({ template: 'src/index.html' })],
  output: {
    path: __dirname + '/dist/renderer',
    filename: '[name].js'
  },
  devServer: {
    port: 4200,
    contentBase: path.join(__dirname, 'dist'),
    compress: true
  }
};

export default rendererConf;
