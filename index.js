"use strict";
var restify = require("restify");
var Database_1 = require("./models/Database");
var server = restify.createServer({
    name: "Chief Server"
});
var db = new Database_1.Database("mongodb://localhost/chief");
db.connect().then(function () {
    server.listen(process.env.port || 80, function () {
        console.log("%s listening at %s", server.name, server.url);
    });
}).catch(function (err) {
    console.error(err);
});
