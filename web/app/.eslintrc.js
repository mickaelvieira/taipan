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
    // both rules don't seem to be working as expected
    "@typescript-eslint/explicit-function-return-type": "off",
    "@typescript-eslint/no-unused-vars": "off"
  },
  settings: {
    react: {
      version: "detect"
    },
  },
};
