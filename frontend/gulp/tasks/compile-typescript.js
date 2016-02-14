var gulp = require('gulp');
var fs = require("fs");
var Handlebars = require("handlebars");

var source = require('vinyl-source-stream');
var browserify = require('browserify');
var tsify = require('tsify');
var glob = require('glob');

gulp.task('compile-typescript', function () {
    var bundler = browserify({
        basedir: TYPESCRIPT,
        debug: true,
        paths: [
            "vendor/blocks/dist"
        ]
    });

    // Load all the definitions
    var definitions = glob.sync(TYPESCRIPT + "/definitions/**/*.d.ts");
    definitions.forEach(function (file) {
        file = file.replace(new RegExp("^" + TYPESCRIPT + "/"), '');
        bundler.add(file);
    });

    // Find widgets
    var widgets = [];
    var widgetFiles = glob.sync(TYPESCRIPT + "/widgets/**/*.ts");
    widgetFiles.forEach(function (file, index) {
        file = file.replace(new RegExp("^" + TYPESCRIPT + "/"), '');

        widgets.push({
            path: "./" + file.replace(new RegExp("\.ts$"), ""),
            name: "widget" + index
        });

        bundler.add(file);
    });

    // Generate widget loader
    var loaderTemplate = fs.readFileSync(__dirname + "/widgetloader.hbs");
    var tmpl = Handlebars.compile(String(loaderTemplate));
    var loader = tmpl({widgets: widgets});
    fs.writeFileSync(TYPESCRIPT + "/widgetloader.ts", loader);

    // Load the main script
    bundler.add('main.ts');

    bundler.plugin(tsify, {noImplicitAny: true});

    return bundler.bundle()
        .on('error', function (err) {
            console.log(err.toString());
            this.emit("end");
        })
        .pipe(source("script.js"))
        .pipe(gulp.dest(PUBLIC));
});