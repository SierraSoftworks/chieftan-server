"use strict";
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
var Iridium = require("iridium");
var Release_1 = require("./Release");
var Target_1 = require("./Target");
var Database = (function (_super) {
    __extends(Database, _super);
    function Database() {
        _super.apply(this, arguments);
        this.Releases = new Iridium.Model(this, Release_1.Release);
        this.Target = new Iridium.Model(this, Target_1.Target);
    }
    return Database;
}(Iridium.Core));
exports.Database = Database;
