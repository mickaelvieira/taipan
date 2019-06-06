module.exports = {
  plugins: [
    ["lodash"],
    ["@babel/plugin-syntax-class-properties"],
    [
      "@babel/plugin-transform-regenerator",
      {
        asyncGenerators: false,
        generators: false,
        async: false
      }
    ],
    ["transform-class-properties"]
  ],
  env: {
    production: {
      presets: [
        [
          "@babel/env",
          {
            modules: false
          }
        ],
        ["@babel/typescript"],
        ["@babel/react"]
      ]
    },
    development: {
      presets: [
        [
          "@babel/env",
          {
            debug: false,
            modules: false
          }
        ],
        ["@babel/typescript"],
        [
          "@babel/react",
          {
            development: true
          }
        ]
      ]
    },
    test: {
      presets: [
        [
          "@babel/env",
          {
            modules: false
          }
        ],
        ["@babel/typescript"],
        ["@babel/react"]
      ]
    }
  }
};
