import * as webpack from 'webpack';
import * as baseConfig from './webpack.config';
import * as path from 'path';
import * as NodeExternals from 'webpack-node-externals';
import devConfig from './webpack.dev';

const prodConfig: webpack.Configuration[] = [{
    ...devConfig,
    mode: 'production',
}];

export default prodConfig;
