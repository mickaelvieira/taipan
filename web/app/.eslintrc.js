module.exports = {
  env: {
    es6: true,
    browser: true
  },
  parser: "@typescript-eslint/parser",
  extends: [
    "plugin:@typescript-eslint/recommended",
    "prettier/@typescript-eslint",
    "plugin:prettier/recommended"
  ],
  parserOptions: {
    ecmaVersion: 2018,
    sourceType: "module",
    ecmaFeatures: {
      jsx: true
    },
  },
  rules: {
    // both rules don"t seem to be working as expected
    "@typescript-eslint/explicit-function-return-type": "off",
    "@typescript-eslint/no-unused-vars": "off",
    "@typescript-eslint/explicit-member-accessibility": ["off", {
      accessibility: "no-public"
    }],
    "graphql/template-strings": ["error", {
      // Import default settings for your GraphQL client. Supported values:
      // "apollo", "relay", "lokka", "fraql", "literal"
      env: "apollo",

      // Import your schema JSON here
      schemaJson: require("./schema.json"),
    }],
    "react-hooks/rules-of-hooks": "error", // Checks rules of Hooks
    "react-hooks/exhaustive-deps": "warn" // Checks effect dependencies
  },
  settings: {
    react: {
      version: "detect"
    },
  },
  plugins: [
    "graphql",
    "react-hooks"
  ]
};
