const DotenvFlow = require("dotenv-flow-webpack");
const path = require("path");
const { StatsWriterPlugin } = require("webpack-stats-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const TerserPlugin = require("terser-webpack-plugin");
const OptimizeCSSAssetsPlugin = require("optimize-css-assets-webpack-plugin");

const srcDir = path.resolve(__dirname, "js");
const tgtDir = path.resolve(__dirname, "../static");

const isTesting = process.env.NODE_ENV === "test";
const isProduction = process.env.NODE_ENV === "production";
const patternJsFiles = isProduction ? "[name].[contenthash].js" : "[name].js";
const patternCssFiles = isProduction
  ? "[name].[contenthash].css"
  : "[name].css";

function getPathnames(chunk) {
  const prefix = "static"
  let { vendor, app } = chunk

  if (!isProduction) {
    vendor = vendor[0];
    app = app[0];
  }

  return {
    vendor: `/${prefix}/${vendor}`,
    app: `/${prefix}/${app}`,
  }
}

module.exports = {
  target: "web",
  mode: isProduction || isTesting ? "production" : "development",
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
        sourceMap: false
      }),
      new OptimizeCSSAssetsPlugin({})
    ]
  },
  plugins: [
    new DotenvFlow({
      path: "../.."
    }),
    new MiniCssExtractPlugin({
      filename: `css/${patternCssFiles}`,
      chunkFilename: `css/${patternCssFiles}`
    }),
    new StatsWriterPlugin({
      filename: "hashes.json",
      transform({ assetsByChunkName }) {
        return Promise.resolve().then(() =>
          JSON.stringify(getPathnames(assetsByChunkName), null, 2)
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
