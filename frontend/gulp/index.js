var glob = require('glob');
var tasks = glob.sync('./gulp/tasks/*.js');

require('./config');

tasks.forEach(function (task) {
  require(task.replace(/^\.\/gulp/, "."));
});
