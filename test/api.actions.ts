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
                        chai.expect(item).to.have.property("http").eql(app.testAction.http);
                    });
                });
        });

        it("should allow you to get a specific action for a project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/action/${app.testAction._id}`)
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("id", app.testAction._id);
                    chai.expect(res.body).to.have.property("name", app.testAction.name);
                    chai.expect(res.body).to.have.property("description", app.testAction.description);
                    chai.expect(res.body).to.have.property("project").eql(app.testAction.project);
                    chai.expect(res.body).to.have.property("http").eql(app.testAction.http);
                });
        });

        it("should return 404 if the action doesn't exist", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/action/000000000000000000000000`)
                .expect(404)
                .toPromise();
        });

        it("should allow you to create a new action for a project", () => {
            return request(app.server)
                .post(`/api/v1/project/${app.testProject._id}/actions`)
                .send({
                    name: "test2",
                    description: "Second test action"
                })
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("name", "test2");
                    chai.expect(res.body).to.have.property("description", "Second test action");
                    chai.expect(res.body).to.have.property("project").eql(app.testAction.project);
                });
        });

        it("should add the new action to the database", () => {
            return app.db.Actions.get({ name: "test2" }).then(action => {
                chai.expect(action).to.exist;
            });
        });

        it("should return 404 if the project doesn't exist", () => {
            return request(app.server)
                .post(`/api/v1/project/000000000000000000000000/actions`)
                .send({
                    name: "test2",
                    description: "Second test action"
                })
                .expect(404)
                .toPromise();
        });
    });
});