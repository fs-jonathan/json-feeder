// vue.config.js
module.exports = {
    pluginOptions: {
        externals: [
            {
                module: 'liff',
                entry: 'https://d.line-scdn.net/liff/1.0/sdk.js',
                global: 'liff',
            },
        ]
    }
}
