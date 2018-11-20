const MonacoWebPackPlugin = require('monaco-editor-webpack-plugin');

module.exports = {
  entry: './src/index.ts',
  output: {
    filename: 'bundle.js',
    path: __dirname + '/dist',
    publicPath: './dist/',
  },

  resolve: {
    // Add '.ts' as resolvable extensions.
    extensions: ['.ts', '.js', '.json'],
  },

  module: {
    rules: [
      // All files with a '.ts' extension will be handled by "babel-loader"
      // and 'awesome-typescript-loader'.
      {
        test: /\.ts?$/,
        loaders: ['babel-loader'],
      },

      // All output '.js' files will have any sourcemaps re-processed by
      // 'source-map-loader'.
      {
        enforce: 'pre',
        test: /\.js$/,
        loader: 'source-map-loader',
        exclude: [/.*monaco-editor.*/],
      },

      // All files with a '.css' extension will be handled by 'style-loader'.
      {test: /\.css$/, use: ['style-loader', 'css-loader']},
    ],
  },
  plugins: [new MonacoWebPackPlugin()],
};
