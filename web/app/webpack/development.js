const merge = require("webpack-merge");
const path = require("path");
const { StatsWriterPlugin } = require("webpack-stats-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const common = require("./common.js");

function getPathnames(chunk) {
  let { vendor, app } = chunk

  return {
    vendor: `/${vendor[0]}`,
    app: `/${app[0]}`
  }
}

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
