module.exports = {
  "plugins": [
    ["lodash"],
    ["@babel/plugin-syntax-class-properties"],
    ["@babel/plugin-transform-regenerator", {
      "asyncGenerators": false,
      "generators": false,
      "async": false
    }],
    ["transform-class-properties"],
    ["styled-components"]
  ],
  "env": {
    "production": {
      "presets": [
        ["@babel/env", {
          "modules": false,
          "targets": ["last 2 versions"]
        }],
        ["@babel/typescript"],
        ["@babel/react"]
      ]
    },
    "development": {
      "presets": [
        ["@babel/env", {
          "debug": true,
          "modules": false,
          "targets": ["last 2 versions"]
        }],
        ["@babel/typescript"],
        ["@babel/react", {
          "development": true
        }]
      ]
    },
    "test": {
      "presets": [
        ["@babel/env", {
          "modules": false,
          "targets": ["last 2 versions"]
        }],
        ["@babel/typescript"],
        ["@babel/react"]
      ]
    }
  }
}
