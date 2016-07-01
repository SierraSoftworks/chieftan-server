import {TestApplication} from "./support/app";
import * as request from "supertest-as-promised";
import * as chai from "chai";

describe("api", () => {
    describe("actions", () => {
        const app = new TestApplication();
        before(() => app.setup());
        after(() => app.teardown());

        it("should allow you to get a list of all the actions for a project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/actions`)
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.be.a("array");
                    chai.expect(res.body).to.have.length(1);

                    res.body.forEach(item => {
                        chai.expect(item).to.have.property("id", app.testAction._id);
                        chai.expect(item).to.have.property("name", app.testAction.name);
                        chai.expect(item).to.have.property("description", app.testAction.description);
                        chai.expect(item).to.have.property("project").eql(app.testAction.project);
                        chai.expect(item).to.have.property("configurations").eql(app.testAction.configurations);
                        chai.expect(item).to.have.property("http").eql(app.testAction.http);
                    });
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/actions`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't have a permission to access the project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject2._id}/actions`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should allow you to get a specific action by its ID", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction._id}`)
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("id", app.testAction._id);
                    chai.expect(res.body).to.have.property("name", app.testAction.name);
                    chai.expect(res.body).to.have.property("description", app.testAction.description);
                    chai.expect(res.body).to.have.property("project").eql(app.testAction.project);
                    chai.expect(res.body).to.have.property("configurations").eql(app.testAction.configurations);
                    chai.expect(res.body).to.have.property("http").eql(app.testAction.http);
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction._id}`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't have a permission to access the project", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction2._id}`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should return 404 if the action doesn't exist", () => {
            return request(app.server)
                .get(`/api/v1/action/000000000000000000000000`)
                .set("Authorization", "Token test")
                .expect(404)
                .toPromise();
        });

        it("should allow you to create a new action for a project", () => {
            return request(app.server)
                .post(`/api/v1/project/${app.testProject._id}/actions`)
                .send({
                    name: "test2",
                    description: "Second test action",
                    configurations: [
                        {
                            name: "Default",
                            vars: {

                            }
                        }
                    ]
                })
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("name", "test2");
                    chai.expect(res.body).to.have.property("description", "Second test action");
                    chai.expect(res.body).to.have.property("project").eql(app.testAction.project);
                    chai.expect(res.body).to.have.property("configurations").eql([{
                        name: "Default",
                        vars: {}
                    }]);
                });
        });

        it("should add the new action to the database", () => {
            return app.db.Actions.get({ name: "test2" }).then(action => {
                chai.expect(action).to.exist;
            });
        });

        it("should create an auditlog entry", () => {
            return app.db.AuditLog.get({ type: "action.create" }).then(entry => {
                chai.expect(entry).to.exist;
                chai.expect(entry).to.have.property("context");

                chai.expect(entry.context).to.have.property("project").eql(app.testProject.summary);
                chai.expect(entry.context).to.have.property("request");
            });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .post(`/api/v1/project/${app.testProject._id}/actions`)
                .send({
                    name: "test2",
                    description: "Second test action",
                    configurations: [
                        {
                            name: "Default",
                            vars: {

                            }
                        }
                    ]
                })
                .expect(401)
                .toPromise();
        });

        it("should return 404 if the project doesn't exist", () => {
            return request(app.server)
                .post(`/api/v1/project/000000000000000000000000/actions`)
                .send({
                    name: "test2",
                    description: "Second test action"
                })
                .set("Authorization", "Token test")
                .expect(404)
                .toPromise();
        });

        it("should allow you to update an action by id", () => {
            return request(app.server)
                .put(`/api/v1/action/${app.testAction._id}`)
                .send({
                    name: "Tested"
                })
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("name", "Tested");
                });
        });

        it("should update the database entry", () => {
            return app.db.Actions.get(app.testAction._id).then(action => {
                chai.expect(action).to.have.property("name", "Tested");
            });
        });

        it("should create an auditlog entry", () => {
            return app.db.AuditLog.get({ type: "action.update" }).then(entry => {
                chai.expect(entry).to.exist;
                chai.expect(entry).to.have.property("context");

                chai.expect(entry.context).to.have.property("project").eql(app.testProject.summary);
                chai.expect(entry.context).to.have.property("action").eql(app.testAction.summary);
                chai.expect(entry.context).to.have.property("request").eql({ name: "Tested" });
            });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .put(`/api/v1/action/${app.testAction._id}`)
                .send({
                    name: "Tested"
                })
                .expect(401)
                .toPromise();
        });
    });
});