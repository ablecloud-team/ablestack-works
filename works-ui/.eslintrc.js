module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: ["eslint:recommended", "@vue/prettier"],
  parserOptions: {
    ecmaVersion: 2020,
    parser: "babel-eslint",
    sourceType: "module",
  },
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
  },
  globals: {
    localStorage: true,
  },
};