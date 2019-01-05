const path = require('path');
const HtmlWebpack = require('html-webpack-plugin');
const HtmlWebpackIncludeAssets = require('html-webpack-include-assets-plugin');
const WebappWebpack = require('webapp-webpack-plugin');
const CleanWebpack = require('clean-webpack-plugin');
const MonacoWebPack = require('monaco-editor-webpack-plugin');

module.exports = {
  entry: './src/index.ts',
  output: {
    filename: '[name].[contenthash].js',
    path: __dirname + '/dist',
    publicPath: './',
  },

  resolve: {
    // Add '.ts' as resolvable extensions.
    extensions: ['.ts', '.js', '.json'],
    // Alias ../../theme.config to semantic theme.config in src/styles/semantic
    alias: {
      '../../theme.config$': path.join(__dirname,
          'src/styles/semantic/theme.config'),
    },
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

      {
        test: /\.less$/,
        use: ['style-loader', 'css-loader', 'less-loader'],
      },

      // All files with a '.css' extension are handled style-loader.
      {test: /\.css$/, use: ['style-loader', 'css-loader']},

      // Media and fonts are handled by the url-loader.
      {
        test: /\.(png|woff|woff2|eot|ttf|svg)$/,
        loader: 'url-loader?limit=8000', // Limit inline file size to 8 kb.
      },
    ],
  },
  plugins: [
    new HtmlWebpack({title: 'Get Roasted! - Roaster Inc.'}),
    new HtmlWebpackIncludeAssets({
      assets: [{
        path: 'https://fonts.googleapis.com/css?family=Source+Sans+Pro',
        type: 'css',
      }],
      append: false,
      publicPath: '',
    }),
    new WebappWebpack('./src/assets/icons/roaster-icon-teal.svg'),
    new CleanWebpack(['dist']),
    new MonacoWebPack(),
  ],
};
