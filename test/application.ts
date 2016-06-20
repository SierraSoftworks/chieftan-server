import {TestApplication} from "./support/app";
import {ExecutorBase} from "../executors/Executor";
import * as request from "supertest-as-promised";
import * as chai from "chai";

describe("application", () => {
    const app = new TestApplication();
    before(() => app.setup());
    after(() => app.teardown());

    it("should reject executors which don't have a config field set", () => {
        chai.expect(() => app.use(TestExecutor)).to.throw();
    });

    it("should allow you to register an executor", () => {
        TestExecutor.config = "test";
        chai.expect(() => app.use(TestExecutor)).to.not.throw();
    });

    it("should allow chaining of executor registration", () => {
        TestExecutor.config = "test";
        chai.expect(() => app.use(TestExecutor)).to.equal(app);
    });
});

class TestExecutor extends ExecutorBase {
    run() {
        return Promise.resolve();
    }
}