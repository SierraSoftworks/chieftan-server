import {TestApplication} from "./support/app";
import * as request from "supertest-as-promised";
import * as chai from "chai";

describe("api", () => {
    describe("projects", () => {
        const app = new TestApplication();
        before(() => app.setup());
        after(() => app.teardown());

        it("should allow you to get a list of all the projects", () => {
            return request(app.server)
                .get(`/api/v1/projects`)
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.be.a("array");
                    chai.expect(res.body).to.have.length(2);

                    res.body.forEach(item => {
                        const { id } = item;
                        let project = app.testProject;
                        if (id === app.testProject2._id) project = app.testProject2;

                        chai.expect(item).to.have.property("id", project._id);
                        chai.expect(item).to.have.property("name", project.name);
                        chai.expect(item).to.have.property("description", project.description);
                        chai.expect(item).to.have.property("url", project.url);
                    });
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/projects`)
                .expect(401)
                .toPromise();
        });

        it("should allow you to get a specific project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}`)
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("id", app.testProject._id);
                    chai.expect(res.body).to.have.property("name", app.testProject.name);
                    chai.expect(res.body).to.have.property("url", app.testProject.url);
                });
        });

        it("should return 404 if the project doesn't exist", () => {
            return request(app.server)
                .get(`/api/v1/project/000000000000000000000000`)
                .set("Authorization", "Token test")
                .expect(404)
                .toPromise();
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/project/000000000000000000000000`)
                .expect(401)
                .toPromise();
        });

        it("should allow you to create a new project", () => {
            return request(app.server)
                .post(`/api/v1/projects`)
                .send({
                    name: "Test Project",
                    description: "This is a test project",
                    url: "http://localhost:9000/"
                })
                .set("Authorization", "Token test")
                .expect(200)
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("id");
                    chai.expect(res.body).to.have.property("name", "Test Project");
                    chai.expect(res.body).to.have.property("description", "This is a test project");
                    chai.expect(res.body).to.have.property("url", "http://localhost:9000/");
                });
        });

        it("should add the project to the database", () => {
            return chai.expect(app.db.Projects.count()).to.eventually.eql(3);
        });

        it("should grant you permissions for the new project", () => {
            return app.db.Users.get(app.testUser._id).then(user => {
                chai.expect(user.permissions).to.have.length(app.testUser.permissions.length + 2);
            });
        });

        it("should create an auditlog entry", () => {
            return app.db.AuditLog.get({ type: "project.create" }).then(entry => {
                chai.expect(entry).to.exist;
                chai.expect(entry).to.have.property("context");

                chai.expect(entry).to.have.property("user").eql(app.testUser.summary);
                chai.expect(entry.context).to.have.property("request").eql({
                    name: "Test Project",
                    description: "This is a test project",
                    url: "http://localhost:9000/"
                });
            });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .post(`/api/v1/projects`)
                .send({
                    name: "Test Project",
                    description: "This is a test project",
                    url: "http://localhost:9000/"
                })
                .expect(401)
                .toPromise();
        });
    });
});