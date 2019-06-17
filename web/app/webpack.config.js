const path = require("path");
const { StatsWriterPlugin } = require("webpack-stats-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const TerserPlugin = require("terser-webpack-plugin");
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");
const Dotenv = require("dotenv-webpack");

const srcDir = path.resolve(__dirname, "js");
const tgtDir = path.resolve(__dirname, "../static");

const isProd = process.env.NODE_ENV === "production";
const isProduction = process.env.NODE_ENV === "production";
const patternJsFiles = isProduction ? "[name].[contenthash].js" : "[name].js";
const patternCssFiles = isProduction
  ? "[name].[contenthash].css"
  : "[name].css";

module.exports = {
  target: "web",
  mode: process.env.NODE_ENV,
  devtool: "source-map",
  entry: {
    app: srcDir + "/app.ts"
  },
  performance: {
    maxEntrypointSize: 614400, // 600 KiB
    maxAssetSize: 614400 // 600 KiB
  },
  output: {
    filename: `js/${patternJsFiles}`,
    chunkFilename: `js/${patternJsFiles}`,
    path: tgtDir
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js"]
  },
  optimization: {
    splitChunks: {
      name: "vendor",
      chunks: "all",
      cacheGroups: {
        default: {
          test: /[\\/]node_modules[\\/]/,
          name: "vendor"
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
    new Dotenv({
      path: isProd ? "../../.env" : "../../.env.local"
    }),
    new MiniCssExtractPlugin({
      filename: `css/${patternCssFiles}`,
      chunkFilename: `css/${patternCssFiles}`
    }),
    new StatsWriterPlugin({
      filename: "hashes.json",
      transform(data) {
        return Promise.resolve().then(() =>
          JSON.stringify(
            {
              app: `/static/${data.assetsByChunkName.app[0]}`,
              vendor: `/static/${data.assetsByChunkName.vendor[0]}`
            },
            null,
            2
          )
        );
      }
    })
  ],
  module: {
    rules: [
      {
        test: /\.graphql?$/,
        use: [
          {
            loader: "webpack-graphql-loader",
            options: {
              output: "document",
              validate: false,
              schema: "./schema.json",
              removeUnusedFragments: true
            }
          }
        ]
      },
      {
        test: /\.(js|jsx|ts|tsx)$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: "babel-loader"
        }
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
