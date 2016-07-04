import {TestApplication} from "./support/app";
import * as request from "supertest-as-promised";
import * as chai from "chai";

describe("api", () => {
    describe("audit", () => {
        const app = new TestApplication();
        before(() => app.setup());
        after(() => app.teardown());

        let auditLogID: string;
        before(() => app.db.AuditLog.insert({
            type: "test",
            user: app.testUser.summary,
            token: "test",
            context: {

            }
        }).then(log => auditLogID = log._id));

        it("should allow you to get the audit log", () => {
            return request(app.server)
                .get("/api/v1/audit")
                .set("Authorization", "Token admin")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.length(1);
                });
        });

        it("should return a 401 if you aren't authorized", () => {
            return request(app.server)
                .get("/api/v1/audit")
                .expect(401)
                .toPromise();
        });

        it("should return a 403 if you aren't an admin", () => {
            return request(app.server)
                .get("/api/v1/audit")
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should allow you to get a specific audit log entry", () => {
            return request(app.server)
                .get(`/api/v1/audit/${auditLogID}`)
                .set("Authorization", "Token admin")
                .expect(200)
                .toPromise();
        });

        it("should return a 401 if you aren't authorized", () => {
            return request(app.server)
                .get(`/api/v1/audit/${auditLogID}`)
                .expect(401)
                .toPromise();
        });

        it("should return a 403 if you aren't an admin", () => {
            return request(app.server)
                .get(`/api/v1/audit/${auditLogID}`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });
    });
});