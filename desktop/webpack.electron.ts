import * as webpack from 'webpack';
import config from './webpack.config';
// import rendererConf from './webpack.electron-front';

const mainConf: webpack.Configuration = {
  ...config,
  target: 'electron-main',
  entry: './src/index.ts',
  mode: 'development',
  output: {
    path: __dirname + '/dist/main',
    filename: 'main.js'
  }
};

export default mainConf;
