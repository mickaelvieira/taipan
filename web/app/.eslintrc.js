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
      env: "literal",

      // Import your schema JSON here
      schemaJson: require("./schema.json"),

      // OR provide absolute path to your schema JSON (but not if using `eslint --cache`!)
      // schemaJsonFilepath: path.resolve(__dirname, "./schema.json"),

      // OR provide the schema in the Schema Language format
      // schemaString: printSchema(schema),

      // tagName is set automatically
    }]
  },
  settings: {
    react: {
      version: "detect"
    },
  },
  plugins: [
    "graphql"
  ]
};
