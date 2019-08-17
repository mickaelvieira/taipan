const development = require("./webpack/development");
const production = require("./webpack/production");

module.exports = process.env.NODE_ENV === "production" ? production : development;
