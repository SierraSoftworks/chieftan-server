import {TestApplication} from "./support/app";
import * as request from "supertest-as-promised";
import * as chai from "chai";

describe("api", () => {
    describe("user", () => {
        const app = new TestApplication();
        before(() => app.setup());
        after(() => app.teardown());

        it("should allow you to get the users list", () => {
            return request(app.server)
                .get("/api/v1/users")
                .set("Authorization", "Token admin")
                .expect(200)
                .toPromise();
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get("/api/v1/users")
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't have an admin/users permission entry", () => {
            return request(app.server)
                .get("/api/v1/users")
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should allow you to get a specific user", () => {
            return request(app.server)
                .get(`/api/v1/user/${app.testUser._id}`)
                .set("Authorization", "Token admin")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.eql(app.testUser.toJSON());
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/user/${app.testUser._id}`)
                .expect(401)
                .toPromise();
        });

        it("should allow you to get your user details", () => {
            return request(app.server)
                .get(`/api/v1/user`)
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.eql(app.testUser.toJSON());
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/user`)
                .expect(401)
                .toPromise();
        });

        it("should return 401 if you don't have an admin/users permission entry", () => {
            return request(app.server)
                .get(`/api/v1/user/${app.testUser._id}`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should allow you to create a new user", () => {
            return request(app.server)
                .post(`/api/v1/users`)
                .send({
                    name: "New Test User",
                    email: "test123@test.com"
                })
                .set("Authorization", "Token admin")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("name", "New Test User");
                });
        });

        it("should create the user in the database", () => {
            return app.db.Users.get({ email: "test123@test.com" }).then(user => {
                chai.expect(user).to.exist;
            });
        });

        it("should create an auditlog entry", () => {
            return app.db.AuditLog.get({ type: "user.create" }).then(entry => {
                chai.expect(entry).to.exist;
                chai.expect(entry).to.have.property("context");

                chai.expect(entry.context).to.have.property("user").eql(app.adminUser.summary);
                chai.expect(entry.context).to.have.property("request").eql({
                    name: "New Test User",
                    email: "test123@test.com"
                });
            });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .post(`/api/v1/users`)
                .send({
                    name: "New Test User",
                    email: "test123@test.com"
                })
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't have the admin/users permission", () => {
            return request(app.server)
                .post(`/api/v1/users`)
                .send({
                    name: "New Test User",
                    email: "test123@test.com"
                })
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should allow you to get a user's tokens", () => {
            return request(app.server)
                .get(`/api/v1/user/${app.testUser._id}/tokens`)
                .set("Authorization", "Token admin")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.eql(app.testUser.tokens);
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/user/${app.testUser._id}/tokens`)
                .expect(401)
                .toPromise();
        });

        it("should return 401 if you don't have an admin/users permission entry", () => {
            return request(app.server)
                .get(`/api/v1/user/${app.testUser._id}/tokens`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should create an auditlog entry", () => {
            return app.db.AuditLog.get({ type: "user.tokens.view" }).then(entry => {
                chai.expect(entry).to.exist;
                chai.expect(entry).to.have.property("context");

                chai.expect(entry.context).to.have.property("user").eql(app.adminUser.summary);
                chai.expect(entry.context).to.have.property("request").eql({
                    user: app.testUser._id
                });
            });
        });

        it("should allow you to add a token to a user's account", () => {
            return request(app.server)
                .post(`/api/v1/user/${app.testUser._id}/tokens`)
                .send({ token: "test2" })
                .set("Authorization", "Token admin")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.eql([...app.testUser.tokens, "test2"]);
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .post(`/api/v1/user/${app.testUser._id}/tokens`)
                .send({ token: "test2" })
                .expect(401)
                .toPromise();
        });

        it("should return 401 if you don't have an admin/users permission entry", () => {
            return request(app.server)
                .post(`/api/v1/user/${app.testUser._id}/tokens`)
                .send({ token: "test2" })
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should create an auditlog entry", () => {
            return app.db.AuditLog.get({ type: "user.tokens.create" }).then(entry => {
                chai.expect(entry).to.exist;
                chai.expect(entry).to.have.property("context");

                chai.expect(entry.context).to.have.property("user").eql(app.adminUser.summary);
                chai.expect(entry.context).to.have.property("request").eql({
                    user: app.testUser._id
                });
            });
        });
    });
});