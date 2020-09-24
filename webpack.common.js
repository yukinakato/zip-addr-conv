const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');

module.exports = {
  entry: {
    main: './web/js/main.js',
    stat: './web/js/stat.js',
    style: './web/js/style.js',
  },
  output: {
    path: path.resolve(__dirname, 'web', 'dist'),
    filename: 'js/[name].[contenthash].js'
  },
  module: {
    rules: [
      {
        enforce: 'pre',
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'eslint-loader',
        options: {
          fix: true,
        },
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
        }
      },
      {
        test: /\.html$/,
        use: ['html-loader'],
      },
      {
        test: /\.(jpe?g|gif|png|svg)$/i,
        use: [
          {
            loader: 'file-loader',
            options: {
              name: '[name].[contenthash].[ext]',
              outputPath: 'img',
            },
          }
        ],
      },
      {
        test: /\.scss$/,
        use: [MiniCssExtractPlugin.loader, 'css-loader', 'postcss-loader', 'sass-loader'],
      },
    ],
  },
  plugins: [
    new CleanWebpackPlugin({ cleanStaleWebpackAssets: false }),
    new HtmlWebpackPlugin({
      filename: 'index.html',
      template: './web/index.html',
      favicon: './web/favicon.ico',
      chunks: ['main', 'style'],
    }),
    new HtmlWebpackPlugin({
      filename: 'stat.html',
      template: './web/stat.html',
      favicon: './web/favicon.ico',
      chunks: ['stat', 'style'],
    }),
    new HtmlWebpackPlugin({
      filename: 'terms.html',
      template: './web/terms.html',
      favicon: './web/favicon.ico',
      chunks: ['style'],
    }),
    new HtmlWebpackPlugin({
      filename: '40x.html',
      template: './web/40x.html',
      favicon: './web/favicon.ico',
      chunks: ['style'],
    }),
    new HtmlWebpackPlugin({
      filename: '50x.html',
      template: './web/50x.html',
      favicon: './web/favicon.ico',
      chunks: ['style'],
    }),
    new MiniCssExtractPlugin({
      filename: 'css/[name].[contenthash].css',
    }),
  ],
  // optimization: {
  //   splitChunks: {
  //     chunks: 'initial',
  //     name: false,
  //     minSize: 0,
  //   },
  // },
};
