"use strict";
var chai = require("chai");
var chaiAsPromised = require("chai-as-promised");
var Bluebird = require("bluebird");
Bluebird.longStackTraces();
chai.config.includeStack = true;
chai.use(chaiAsPromised);
