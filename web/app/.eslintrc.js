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
    // This rule makes sense with helpers function but it is silly with React components
    "@typescript-eslint/explicit-function-return-type": ["off", {
      allowExpressions: true
    }],
    "@typescript-eslint/explicit-member-accessibility": ["error", {
      accessibility: "no-public"
    }],
    "@typescript-eslint/no-unused-vars": ["error", {
      argsIgnorePattern: "_"
    }],
    "react/display-name": "off",
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
