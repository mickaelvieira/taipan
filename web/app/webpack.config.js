const path = require("path");
const { StatsWriterPlugin } = require("webpack-stats-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const TerserPlugin = require('terser-webpack-plugin');
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");

const srcDir = path.resolve(__dirname, "js");
const tgtDir = path.resolve(__dirname, "../static");

const isProduction = process.env.NODE_ENV === "production";
const patternJsFiles = isProduction ? "[name].[contenthash].js" : "[name].js"
const patternCssFiles = isProduction ? "[name].[contenthash].css" : "[name].css"

module.exports = {
  target: "web",
  mode: process.env.NODE_ENV,
  devtool: "source-map",
  entry: {
    app: srcDir + "/app.tsx"
    // login: srcDir + "/login.ts"
  },
  output: {
    filename: `js/${patternJsFiles}`,
    chunkFilename: `js/${patternJsFiles}`,
    path: tgtDir
  },
  resolve: {
    extensions: [".js", ".ts", ".tsx"],
    alias: {
      lib: path.resolve(srcDir, "lib/"),
      components: path.resolve(srcDir, "components/"),
      store: path.resolve(srcDir, "store/"),
      services: path.resolve(srcDir, "services/"),
      collection: path.resolve(srcDir, "collection/")
    }
  },
  optimization: {
    splitChunks: {
      cacheGroups: {
        vendor: {
          test: /[\\/]node_modules[\\/]/,
          name: "vendor",
          chunks: "all"
        }
      }
    },
    minimize: isProduction,
    minimizer: [
      new TerserPlugin({
        sourceMap: true
      }),
      new OptimizeCSSAssetsPlugin({})
    ]
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: `css/${patternCssFiles}`,
      chunkFilename: `css/${patternCssFiles}`
    }),
    new StatsWriterPlugin({
      filename: "hashes.json",
      transform(data) {
        return Promise.resolve().then(() => JSON.stringify({
          styles: `/static/${data.assetsByChunkName.app[0]}`,
          app: `/static/${data.assetsByChunkName.app[1]}`,
          vendor: `/static/${data.assetsByChunkName.vendor[0]}`
        }, null, 2));
      }
    })
  ],
  module: {
    rules: [
      {
        test: /\.(js|jsx|ts|tsx)$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: "babel-loader"
        }
      },
      {
        test: /\.scss$/,
        use: [
          MiniCssExtractPlugin.loader, {
            loader: "css-loader", options: {
              sourceMap: true
            }
          }, {
            loader: "sass-loader", options: {
              sourceMap: true
            }
          }]
      },
      {
        test: /\.(woff(2)?|ttf|svg|eot)(\?v=\d+\.\d+\.\d+)?$/,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "[name].[ext]",
              outputPath: "dist/fonts",
              publicPath: "/dist/fonts"
            }
          }
        ]
      }
    ]
  }
};
