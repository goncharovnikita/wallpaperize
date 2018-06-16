import * as webpack from 'webpack';
import config from './webpack.config';
// import rendererConf from './webpack.electron-front';

const mainConf: webpack.Configuration = {
  ...config,
  target: 'electron-main',
  entry: './src/main.ts',
  mode: 'development',
  output: {
    path: __dirname + '/dist',
    filename: '[name].js'
  }
};

export default mainConf;
