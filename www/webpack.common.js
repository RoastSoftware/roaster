const webpack = require('webpack');
const CopyWebpack = require('copy-webpack-plugin');
const HtmlWebpack = require('html-webpack-plugin');
const HtmlWebpackIncludeAssets = require('html-webpack-include-assets-plugin');
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

      // All files with a '.css' extension are handled style-loader.
      {test: /\.css$/, use: ['style-loader', 'css-loader']},

      // Media and fonts are handled by the url-loader.
      {
        test: /\.(png|woff|woff2|eot|ttf|svg)$/,
        loader: 'url-loader?limit=100000',
      },
    ],
  },
  plugins: [
    new CopyWebpack([
      {from: 'node_modules/semantic-ui-site', to: 'semantic/'},
      {
        from: 'node_modules/semantic-ui-forest-themes',
        to: 'semantic-themes/',
      },
    ]),
    new HtmlWebpack({title: 'Get Roasted!'}),
    new HtmlWebpackIncludeAssets({
      assets: [
        'semantic/site.min.css',
        'semantic-themes/semantic.solar.min.css',
      ],
      append: false,
    }),
    new CleanWebpack(['dist']),
    new MonacoWebPack(),
  ],
};
