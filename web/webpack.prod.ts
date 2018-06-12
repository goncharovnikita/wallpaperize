import * as webpack from 'webpack';
import * as path from 'path';
import devConfig from './webpack.dev';

const prodConfig: webpack.Configuration[] = [{
    ...devConfig,
    mode: 'production',
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: '[name].[hash].js'
    },
}];

export default prodConfig;
