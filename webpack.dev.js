const path = require('path');
const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
  mode: 'development',
  devtool: 'eval-cheap-module-source-map',
  watch: true,
  devServer: {
    port: 8080,
    host: '0.0.0.0',
    contentBase: false,
  },
});
