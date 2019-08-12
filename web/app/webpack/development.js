const merge = require("webpack-merge");
const path = require("path");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const common = require("./common.js");

module.exports = merge(common, {
  mode: "development",
  devtool: "source-map",
  output: {
    filename: "js/[name].js",
    chunkFilename: "js/[name].js",
    path: path.resolve(__dirname, "../../static"),
    publicPath: "/static/"
  },
  optimization: {
    splitChunks: {
      name: true,
      chunks: "all",
      cacheGroups: {
        default: {
          test: /[\\/]node_modules[\\/]/,
          name: "vendor"
        }
      }
    }
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: "css/[name].css",
      chunkFilename: "css/[name].css"
    })
  ]
});
