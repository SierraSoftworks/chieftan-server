"use strict";
var RouteBase = (function () {
    function RouteBase(server, db) {
        this.server = server;
        this.db = db;
    }
    RouteBase.prototype.register = function () {
    };
    RouteBase.prototype.notFound = function (res) {
        return this.error(res, 404, "Not Found", "The entity you requested could not be found.");
    };
    RouteBase.prototype.serverError = function (res) {
        return this.error(res, 500, "Server Error", "The server has encountered an error preventing it from serving your request.");
    };
    RouteBase.prototype.error = function (res, code, error, message) {
        return res.json(code, {
            code: code,
            error: error,
            message: message
        });
    };
    return RouteBase;
}());
exports.RouteBase = RouteBase;
