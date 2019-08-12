const path = require("path");
const merge = require("webpack-merge");
const { StatsWriterPlugin } = require("webpack-stats-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const TerserPlugin = require("terser-webpack-plugin");
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");
const common = require("./common.js");

function getPathnames(chunk) {
  let { vendor, app } = chunk

  return {
    vendor: `/${vendor}`,
    app: `/${app}`
  }
}

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
          name: "vendor"
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
    }),
    new StatsWriterPlugin({
      filename: "hashes.json",
      transform({ assetsByChunkName }) {
        return Promise.resolve().then(() =>
          JSON.stringify(getPathnames(assetsByChunkName), null, 2)
        );
      }
    })
  ]
});
