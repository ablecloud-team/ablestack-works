module.exports = {
  pluginOptions: {
    i18n: {
      localeDir: "locales",
      enableLegacy: true,
      runtimeOnly: false,
      compositionOnly: true,
      fullInstall: true,
    },
  },
  css: {
    extract: true,
  },
  devServer: {
    historyApiFallback: true,
  },
};
