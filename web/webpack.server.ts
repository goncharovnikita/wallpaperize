import * as webpack from 'webpack';
import * as baseConfig from './webpack.config';
import * as path from 'path';
import * as NodeExternals from 'webpack-node-externals';
import devConfig from './webpack.dev';

const prodConfig: webpack.Configuration[] = [{
    ...baseConfig.default,
    entry: {
        server: path.resolve(__dirname, 'src', 'Server.tsx')
    },
    target: 'node',
    mode: 'production',
    externals: [
        NodeExternals(),
    ],
    output: {
        filename: '[name].js'
    },
    devtool: 'source-map',
    resolve: {
        extensions: [ '.tsx', '.ts', '.js' ],
        modules: [
          path.resolve(__dirname, "src"),
          "node_modules"
        ]
      },
    // plugins: []
}, {
    ...devConfig,
}];

export default prodConfig;
