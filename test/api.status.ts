import {TestApplication} from "./support/app";
import * as request from "supertest-as-promised";
import * as chai from "chai";

describe("api", () => {
    describe("status", () => {
        const app = new TestApplication();
        before(() => app.setup());
        after(() => app.teardown());

        it("should return 200", () => {
            return request(app.server)
                .get(`/api/v1/status`)
                .expect(200)
                .toPromise();
        });
    });
});