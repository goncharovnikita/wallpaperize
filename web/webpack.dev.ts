import * as webpack from 'webpack';
import * as path from 'path';
import * as baseConfig from './webpack.config';

const devConfig: webpack.Configuration = {
    ...baseConfig.default,
    mode: 'development',
    resolve: {
        extensions: [ '.tsx', '.ts', '.js' ],
        modules: [
          path.resolve(__dirname, "src"),
          "node_modules"
        ]
      },
};

export default devConfig;
