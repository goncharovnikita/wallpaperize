import * as path from 'path';
import * as HtmlPlugin from 'html-webpack-plugin';
import * as webpack from 'webpack';
import * as CopyPlugin from 'copy-webpack-plugin';

const config: webpack.Configuration = {
    mode: 'development',
    entry: {
        bundle: path.resolve(__dirname, 'src', 'Index.tsx'),
        vendor: path.resolve(__dirname, 'src', 'vendor', 'vendor.ts')
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/
            },
            {
                test: /\.sass$/,
                use: [
                    'style-loader',
                    'css-loader',
                    'sass-loader'
                ]
            },
            {
                test: /\.css$/,
                use: [
                    'style-loader',
                    'css-loader',
                ]
            }
        ]
    },
    plugins: [
        new HtmlPlugin(
            {
                template: path.resolve(__dirname, 'src', 'index.html'),
                excludeChunks: ['server']
            }
        ),
        new CopyPlugin([
            {
                from: path.resolve(__dirname, 'src', 'assets'),
                to: path.resolve(__dirname, 'dist', 'assets')
            }
        ])
    ],
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: '[name].js'
    },
    devServer: {
        contentBase: path.resolve(__dirname, 'dist'),
        port: 4200,
        compress: true
    }
};

export default config;
