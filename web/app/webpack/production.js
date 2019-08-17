const path = require("path");
const merge = require("webpack-merge");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const TerserPlugin = require("terser-webpack-plugin");
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");
const common = require("./common.js");

module.exports = merge(common, {
  mode: "production",
  output: {
    filename: "js/[name].[contenthash].js",
    chunkFilename: "js/[name].[contenthash].js",
    path: path.resolve(__dirname, "../../static"),
    publicPath: "/"
  },
  optimization: {
    splitChunks: {
      name: true,
      chunks: "all",
      cacheGroups: {
        default: {
          test: /[\\/]node_modules[\\/]/,
          name(module) {
            let name = "vendor";
            const packageName = module.context.match(/[\\/]node_modules[\\/](.*?)([\\/]|$)/)[1];
            if (packageName.indexOf("material-ui") !== -1) {
              name = "materialui";
            } else if (packageName.indexOf("react") !== -1) {
              name = "react";
            }
            return name;
          }
        }
      }
    },
    minimize: true,
    minimizer: [
      new TerserPlugin({
        sourceMap: false
      }),
      new OptimizeCSSAssetsPlugin({})
    ]
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: "css/[name].[contenthash].css",
      chunkFilename: "css/[name].[contenthash].css"
    })
  ]
});
