module.exports = {
    lintOnSave: true,

    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:8000',
                changeOrigin: true,
                pathRewrite: {
                    // '^/api': '/'
                }
            }
        }
    }
}
