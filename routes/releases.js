"use strict";
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
var base_1 = require("./base");
var Releases = (function (_super) {
    __extends(Releases, _super);
    function Releases() {
        _super.apply(this, arguments);
    }
    Releases.prototype.register = function () {
        var _this = this;
        this.server.get("/releases", function (req, res) {
            _this.db.Releases.find().toArray().then(function (releases) {
                res.json(200, releases);
            }).catch(function (err) { return _this.serverError(res); });
        });
        this.server.get("/releases/:project", function (req, res) {
            _this.db.Releases.find({
                "project.id": req.params.project
            }).toArray().then(function (releases) {
                res.json(200, releases);
            }).catch(function (err) { return _this.serverError(res); });
        });
    };
    return Releases;
}(base_1.RouteBase));
exports.Releases = Releases;
