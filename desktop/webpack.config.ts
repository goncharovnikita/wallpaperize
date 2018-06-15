import * as webpack from 'webpack';

const config: webpack.Configuration = {
    entry: './src/back/main.ts',
    node: {
        __dirname: false
    },
    resolve: {
        extensions: [ '.tsx', '.ts', '.js', '.jsx' ],
        alias: {
            '@app': __dirname + '/src/'
        }
    },
    module: {
        rules: [
            {
                test: /\.tsx?/,
                use: [
                    'babel-loader',
                    'ts-loader'
                ]
            },
            {
              test: /\.jsx?$/,
              loader: 'babel-loader'
            },
            {
                test: /\.css/,
                use: 'css-loader'
            },
            {
                test: /\.sass/,
                use: [
                    'style-loader',
                    'css-loader',
                    'sass-loader'
                ]
            },
            {
                test: /\.(woff(2)?|ttf|eot|svg)(\?v=\d+\.\d+\.\d+)?$/,
                use: [{
                    loader: 'file-loader',
                    options: {
                        name: '[name].[ext]',
                        outputPath: 'fonts/'
                    }
                }]
            }
        ]
    },
    output: {
        filename: '[name].js',
        path: __dirname + '/dist'
    }
};

export default config;
