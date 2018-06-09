import * as webpack from 'webpack';
import * as baseConfig from './webpack.config';

const devConfig: webpack.Configuration = {
    ...baseConfig.default,
    mode: 'development'
};

export default devConfig;
