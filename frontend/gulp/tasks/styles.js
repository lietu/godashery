var gulp = require('gulp');
var sass = require('gulp-sass');

gulp.task('styles', function () {
    return gulp.src(STYLES + "/main.scss")
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest(PUBLIC));
});
