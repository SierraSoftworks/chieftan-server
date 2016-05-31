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
var Request_1 = require("./Request");
var Project_1 = require("./Project");
var Target = (function (_super) {
    __extends(Target, _super);
    function Target() {
        _super.apply(this, arguments);
    }
    __decorate([
        Iridium.ObjectID
    ], Target.prototype, "_id", void 0);
    __decorate([
        Iridium.Property(String)
    ], Target.prototype, "name", void 0);
    __decorate([
        Iridium.Property(String)
    ], Target.prototype, "description", void 0);
    __decorate([
        Iridium.Property(Project_1.ProjectSchema)
    ], Target.prototype, "project", void 0);
    __decorate([
        Iridium.Property(Request_1.RequestSchema)
    ], Target.prototype, "deploy", void 0);
    Target = __decorate([
        Iridium.Collection("targets")
    ], Target);
    return Target;
}(Iridium.Instance));
exports.Target = Target;
