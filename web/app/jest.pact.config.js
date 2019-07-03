const { defaults: tsjPreset } = require('ts-jest/presets');

module.exports = {
  verbose: true,
  collectCoverage: false,
  coverageDirectory: "./coverage",
  preset: "ts-jest",
  transform: {
    ...tsjPreset.transform,
    "\\.(gql|graphql)$": "jest-transform-graphql",
  },
  setupFiles: [
    "./pact.setup.js"
  ],
  setupFilesAfterEnv: [
    "./pact.env.js"
  ],
  moduleFileExtensions: ["js", "jsx", "ts", "tsx", "graphql"],
  testRegex: "/*(.pact)\\.[jt]sx?$"
};
