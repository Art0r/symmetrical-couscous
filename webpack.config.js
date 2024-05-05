const path = require('path');

module.exports = {
    entry: { bulma: 'bulma/css/bulma.css' },
    mode: 'production',
    output: {
        filename: '[name].bundle.js',
        path: path.resolve(__dirname, 'static/'),
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                },
            },
            {
                test: /\.(scss|sass)$/,
                use: [
                    'style-loader',  // Inject styles into DOM
                    'css-loader',    // Translates CSS into CommonJS
                    'sass-loader',   // Compiles Sass to CSS
                ],
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif|ico)$/i,
                use: {
                    loader: 'file-loader',
                    options: {
                        name: '[name].[ext]',
                        outputPath: 'images/',  // Specify the output path for images
                    },
                },
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader'],
            },
        ],
    },
};
