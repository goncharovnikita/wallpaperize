import * as path from 'path';
import * as HtmlPlugin from 'html-webpack-plugin';
import * as webpack from 'webpack';

const config: webpack.Configuration = {
    mode: 'development',
    entry: {
        bundle: path.resolve(__dirname, 'src', 'App.tsx')
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader'
            },
            {
                test: /\.sass$/,
                use: [
                    'style-loader',
                    'css-loader',
                    'sass-loader'
                ]
            }
        ]
    },
    plugins: [
        new HtmlPlugin({template: path.resolve(__dirname, 'src', 'index.html')})
    ],
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: 'bundle.js'
    },
    devServer: {
        contentBase: path.resolve(__dirname, 'dist'),
        port: 4200,
        compress: true
    }
};

export default config;
