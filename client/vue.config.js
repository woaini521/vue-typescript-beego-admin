/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 18:34:32
 * @LastEdit: Sy.
 * @FilePath: \client\vue.config.js
 * @Description: 项目配置
 */
const path = require('path')
const name = 'Vue Typescript Admin'

module.exports = {
  publicPath: process.env.NODE_ENV === 'production' ? '/vue-typescript-admin-template/' : '/', // TODO: Remember to change this to fit your need
  // lintOnSave: process.env.NODE_ENV === 'development',
  lintOnSave: true,
  pwa: {
    name: name
  },
  pluginOptions: {
    'style-resources-loader': {
      preProcessor: 'scss',
      patterns: [
        path.resolve(__dirname, 'src/styles/_variables.scss'),
        path.resolve(__dirname, 'src/styles/_mixins.scss')
      ]
    }
  },
  productionSourceMap: false,
  configureWebpack: {
    plugins: []
  },
  chainWebpack(config) {
    // Provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    config.set('name', name)
  },
  css: {
    sourceMap: false
  },
  devServer: {
    open: false,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8000/api/admin/',
        changeOrigin: true,
        ws: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
}
