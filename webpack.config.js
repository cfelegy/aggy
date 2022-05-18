const path = require('path');

const config = {
    cache: false,
    mode: process.env.NODE_ENV || 'development',
    entry: './src/frontend/index.tsx',
    devtool: 'inline-source-map',
    devServer: {
        compress: true,
        static: {
            directory: path.join(__dirname, 'public')
        },
        port: 3000,
        watchFiles: ['src/frontend/**/*']
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                loader: 'ts-loader',
                exclude: /node-modules/,
                options: {
                    configFile: 'tsconfig.frontend.json'
                }
            },
            {
                test: /\.scss$/,
                use: ['style-loader', 'css-loader', 'sass-loader']
            }
        ]
    },
    resolve: {
        extensions: ['.tsx', '.ts', '.js', '.css', '.scss']
    },
    output: {
        filename: 'app.js',
        path: path.resolve(__dirname, 'public/dist'),
        publicPath: '/dist/'
    }
};

module.exports = [config];