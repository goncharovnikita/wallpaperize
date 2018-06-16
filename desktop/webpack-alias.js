module.exports = {
  entry: {
    renderer: __dirname + '/src/renderer/index.tsx',
    vendor: __dirname + '/src/vendor/vendor.ts'
  },
  node: {
    __dirname: false
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js', '.jsx'],
    alias: {
      '@app': __dirname + '/src/renderer/',
      '@approot': __dirname + '/src/'
    }
  },
  module: {
    rules: [
      {
        test: /\.sass/,
        use: ['style-loader', 'css-loader', 'sass-loader']
      }
    ]
  }
};
