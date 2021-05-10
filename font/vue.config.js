module.exports = {
  css: {
    loaderOptions: {
      less: {
        javascriptEnabled: true
      }
    }
  },
  devServer: {
    open: false,
    port: 8081,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8024',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
}
