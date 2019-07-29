'use strict';d
// Template version: 1.2.7
// see http://vuejs-templates.github.io/webpack for documentation.

const path = require('path');

module.exports = {
    dev: {
        env: require('./dev.env'),
        host: 'localhost',
        port: 8099,
        // Paths
        assetsSubDirectory: 'http://localhost:8099/static',
        assetsPublicPath: 'http://localhost:8099/',
        proxyTable: {
            '/perm': {
                // 本机
                target: 'http://127.0.0.1:8098',
                changeOrigin: true,
                pathRewrite: {
                    '/perm': '/perm'
                }
            }
        },
        cssSourceMap: false,
    },

    build: {
        // Template for index.html
        index: path.resolve(__dirname, '../perm-vue-manage/index.html'),

        // Paths
        assetsRoot: path.resolve(__dirname, '../perm-vue-manage'),
        assetsSubDirectory: 'static',
        assetsPublicPath: './',

        /**
         * Source Maps
         */

        productionSourceMap: false,
        // https://webpack.js.org/configuration/devtool/#production
        devtool: '#source-map',

        // Gzip off by default as many popular static hosts such as
        // Surge or Netlify already gzip all static assets for you.
        // Before setting to `true`, make sure to:
        // npm install --save-dev compression-webpack-plugin
        productionGzip: false,
        productionGzipExtensions: ['js', 'css'],

        // Run the build command with an extra argument to
        // View the bundle analyzer report after build finishes:
        // `npm run build --report`
        // Set to `true` or `false` to always turn it on or off
        bundleAnalyzerReport: process.env.npm_config_report
    }
};
