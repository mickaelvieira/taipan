module.exports = {
  env: {
    es6: true,
    browser: true
  },
  parser: "@typescript-eslint/parser",
  extends: [
    "plugin:react/recommended",
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
    "@typescript-eslint/explicit-function-return-type": ["error", {
      allowExpressions: true,
      allowHigherOrderFunctions: true
    }],
    "@typescript-eslint/explicit-member-accessibility": ["error", {
      accessibility: "no-public"
    }],
    "@typescript-eslint/no-unused-vars": ["error", {
      argsIgnorePattern: "_"
    }],
    "react/display-name": "off",
    // "graphql/template-strings": ["error", {
    //   env: "apollo",
    //   schemaJson: require("./schema.json"),
    // }],
    "react-hooks/rules-of-hooks": "error",
    "react-hooks/exhaustive-deps": "warn"
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
