"use strict";
var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var Iridium = require("iridium");
var Project_1 = require("./Project");
var Release = (function (_super) {
    __extends(Release, _super);
    function Release() {
        _super.apply(this, arguments);
    }
    Release.onCreating = function (doc) {
        doc.vars = doc.vars || {};
    };
    __decorate([
        Iridium.ObjectID
    ], Release.prototype, "_id", void 0);
    __decorate([
        Iridium.Property({
            id: String,
            name: String,
            url: String
        })
    ], Release.prototype, "version", void 0);
    __decorate([
        Iridium.Property(Project_1.ProjectSchema)
    ], Release.prototype, "project", void 0);
    __decorate([
        Iridium.Property({
            $propertyType: String
        })
    ], Release.prototype, "vars", void 0);
    Release = __decorate([
        Iridium.Collection("releases")
    ], Release);
    return Release;
}(Iridium.Instance));
exports.Release = Release;
