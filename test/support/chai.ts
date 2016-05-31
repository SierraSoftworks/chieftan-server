import * as chai from "chai";
import * as chaiAsPromised from "chai-as-promised";
import * as Bluebird from "bluebird";

Bluebird.longStackTraces();

chai.config.includeStack = true;
chai.use(chaiAsPromised);
