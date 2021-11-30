module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: ["eslint:recommended", "@vue/prettier"],
  parser: "vue-eslint-parser",
  parserOptions: {
    parser: "@typescript-eslint/parser",
    ecmaVersion: "latest",
    sourceType: "module",
  },
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
  },
  globals: {
    localStorage: true,
    sessionStorage: true,
  },
};
