import * as webpack from 'webpack';
import config from './webpack.config';
// import rendererConf from './webpack.electron-front';

const mainConf: webpack.Configuration = {
  ...config,
  entry: {
    renderer: __dirname + '/src/renderer/index.tsx',
    vendor: __dirname + '/src/vendor/vendor.ts'
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js', '.jsx'],
    alias: {
      '@app': __dirname + '/src/renderer/',
      '@approot': __dirname + '/src/'
    }
  },
  mode: 'development',
  output: {
    path: __dirname + '/dist/main',
    filename: 'main.js'
  }
};

export default mainConf;
