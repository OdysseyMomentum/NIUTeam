// eslint-disable-next-line @typescript-eslint/no-var-requires
const webpack = require('webpack')

module.exports = {
  lintOnSave: true,
  devServer: {
    port: 3000,
    proxy: {
      '^/api': {
        target: 'https://9ga0hfscwk.execute-api.eu-central-1.amazonaws.com/dev'
      }
    }
  },

  pluginOptions: {
    i18n: {
      locale: 'en',
      fallbackLocale: 'en',
      localeDir: 'locales',
      enableInSFC: true
    }
  }
}
