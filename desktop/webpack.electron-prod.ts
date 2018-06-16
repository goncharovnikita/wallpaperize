import * as webpack from 'webpack';
import rendererConf from './webpack.electron-front';
import mainConf from './webpack.electron';

const prodConfElectron: webpack.Configuration = {
  ...mainConf,
  mode: 'production'
};

const prodConfRenderer: webpack.Configuration = {
  ...rendererConf,
  mode: 'production'
};

export default [prodConfElectron, prodConfRenderer];
