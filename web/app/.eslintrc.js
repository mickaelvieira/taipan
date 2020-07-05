module.exports = {
  env: {
    es6: true,
    browser: true,
  },
  parser: "@typescript-eslint/parser",
  extends: [
    "plugin:react/recommended",
    "plugin:@typescript-eslint/recommended",
    "prettier/@typescript-eslint",
    "plugin:prettier/recommended",
    "plugin:jsx-a11y/recommended",
  ],
  parserOptions: {
    ecmaVersion: 2018,
    sourceType: "module",
    ecmaFeatures: {
      jsx: true,
    },
  },
  rules: {
    "@typescript-eslint/explicit-function-return-type": [
      "error",
      {
        allowExpressions: true,
        allowHigherOrderFunctions: true,
      },
    ],
    "@typescript-eslint/explicit-member-accessibility": [
      "error",
      {
        accessibility: "no-public",
      },
    ],
    "@typescript-eslint/no-unused-vars": [
      "error",
      {
        argsIgnorePattern: "_",
      },
    ],
    "@typescript-eslint/ban-types": "off",
    "react/prop-types": "off",
    "react/display-name": "off",
    // "graphql/template-strings": ["error", {
    //   env: "apollo",
    //   schemaJson: require("./schema.json"),
    // }],
    "react-hooks/rules-of-hooks": "error",
    "react-hooks/exhaustive-deps": "warn",
    "jsx-a11y/no-autofocus": "off",
  },
  settings: {
    react: {
      version: "detect",
    },
  },
  plugins: ["graphql", "react", "react-hooks", "jsx-a11y"],
};
