import * as webpack from 'webpack';
import * as baseConfig from './webpack.config';

const prodConfig: webpack.Configuration = {
    ...baseConfig.default,
    mode: 'production'
};

export default prodConfig;
