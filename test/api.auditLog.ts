import {TestApplication} from "./support/app";
import * as request from "supertest-as-promised";
import * as chai from "chai";

describe("api", () => {
    describe("audit", () => {
        const app = new TestApplication();
        before(() => app.setup());
        after(() => app.teardown());

        it("should allow you to get the audit log", () => {
            return request(app.server)
                .get("/api/v1/audit")
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise();
        });

        it("should return a 401 if you aren't authorized", () => {
            return request(app.server)
                .get("/api/v1/audit")
                .expect(401)
                .toPromise();
        });
    });
});