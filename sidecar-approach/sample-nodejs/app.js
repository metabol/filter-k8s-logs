var express = require('express');
var fs = require('fs');

var app = express();

// redirect stdout and stderr to file
var access = fs.createWriteStream('logs.txt');
process.stdout.write = process.stderr.write = access.write.bind(access);

// handle uncaught exceptions
process.on('uncaughtException', function(err) {
    console.error((err && err.stack) ? err.stack : err);
});

app.get('/', function (req, res) {
    console.log(req);
    res.send('Hello World!');
});

app.listen(8080, function () {
    console.log('Example app listening on port 8080!');
});