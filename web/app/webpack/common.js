const path = require("path");
const DotenvFlow = require("dotenv-flow-webpack");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
  target: "web",
  entry: {
    app: path.resolve(__dirname, "../js") + "/app.ts",
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js"],
  },
  plugins: [
    new DotenvFlow({
      path: "../..",
    }),
  ],
  module: {
    rules: [
      {
        test: /\.(css)$/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
          },
          "css-loader",
        ],
      },
      {
        test: /\.graphql?$/,
        use: [
          {
            loader: "webpack-graphql-loader",
            options: {
              output: "document",
              validate: false,
              schema: "./schema.json",
              removeUnusedFragments: true,
            },
          },
        ],
      },
      {
        test: /\.(js|jsx|ts|tsx)$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: "babel-loader",
        },
      },
      {
        test: /\.(woff(2)?|ttf|svg|eot)(\?v=\d+\.\d+\.\d+)?$/,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "[name].[ext]",
              outputPath: "dist/fonts",
              publicPath: "/dist/fonts",
            },
          },
        ],
      },
    ],
  },
};
