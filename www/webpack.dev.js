const merge = require("webpack-merge");
const common = require("./webpack.common.js");

module.exports = merge(common, {
    // Run webpack in development mode.
    mode: "development",

    // Enable sourcemaps for debugging webpack's output.
    devtool: "source-map",

    // Set ./dist folder as content base.
    devServer: {
        contentBase: "./dist"
    }
});
