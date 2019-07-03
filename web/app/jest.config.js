module.exports = {
  verbose: true,
  collectCoverage: false,
  coverageDirectory: "./coverage",
  preset: "ts-jest",
  transform: {
    "\\.(gql|graphql)$": "jest-transform-graphql",
  }
};
