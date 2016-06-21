import {TestApplication} from "./support/app";
import * as request from "supertest-as-promised";
import * as chai from "chai";

describe("api", () => {
    describe("tasks", () => {
        const app = new TestApplication();
        before(() => app.setup());
        after(() => app.teardown());

        it("should return a list of tasks", () => {
            return request(app.server)
                .get(`/api/v1/tasks`)
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.be.a("array");
                    chai.expect(res.body).to.have.length(1);

                    res.body.forEach(item => {
                        chai.expect(item).to.have.property("id", app.testTask._id);
                        chai.expect(item).to.have.property("metadata").eql(app.testTask.metadata);
                        chai.expect(item).to.have.property("action").eql(app.testTask.action);
                        chai.expect(item).to.have.property("project").eql(app.testTask.project);
                        chai.expect(item).to.have.property("vars").eql(app.testTask.vars);
                        chai.expect(item).to.have.property("state").eql("NotExecuted");
                        chai.expect(item).to.have.property("output", "");
                    });
                });
        });
        
        it("should return the list of tasks for a project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/tasks`)
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.be.a("array");
                    chai.expect(res.body).to.have.length(1);

                    res.body.forEach(item => {
                        chai.expect(item).to.have.property("id", app.testTask._id);
                        chai.expect(item).to.have.property("metadata").eql(app.testTask.metadata);
                        chai.expect(item).to.have.property("action").eql(app.testTask.action);
                        chai.expect(item).to.have.property("project").eql(app.testTask.project);
                        chai.expect(item).to.have.property("vars").eql(app.testTask.vars);
                        chai.expect(item).to.have.property("state").eql("NotExecuted");
                        chai.expect(item).to.have.property("output", "");
                    });
                });
        });
        
        it("should return the list of tasks for a project and action", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction._id}/tasks`)
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.be.a("array");
                    chai.expect(res.body).to.have.length(1);

                    res.body.forEach(item => {
                        chai.expect(item).to.have.property("id", app.testTask._id);
                        chai.expect(item).to.have.property("metadata").eql(app.testTask.metadata);
                        chai.expect(item).to.have.property("action").eql(app.testTask.action);
                        chai.expect(item).to.have.property("project").eql(app.testTask.project);
                        chai.expect(item).to.have.property("vars").eql(app.testTask.vars);
                        chai.expect(item).to.have.property("state").eql("NotExecuted");
                        chai.expect(item).to.have.property("output", "");
                    });
                });
        });

        it("should allow you to check whether a task exists by ID", () => {
            return request(app.server)
                .head(`/api/v1/task/${app.testTask._id}`)
                .expect(200)
                .toPromise();
        });

        it("should allow you to check whether a task doesn't exit by ID", () => {
            return request(app.server)
                .head(`/api/v1/task/000000000000000000000000`)
                .expect(404)
                .toPromise();
        });

        it("should let you request a specific task by ID", () => {
            return request(app.server)
                .get(`/api/v1/task/${app.testTask._id}`)
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("id", app.testTask._id);
                    chai.expect(res.body).to.have.property("metadata").eql(app.testTask.metadata);
                    chai.expect(res.body).to.have.property("action").eql(app.testTask.action);
                    chai.expect(res.body).to.have.property("project").eql(app.testTask.project);
                    chai.expect(res.body).to.have.property("vars").eql(app.testTask.vars);
                    chai.expect(res.body).to.have.property("state").eql("NotExecuted");
                    chai.expect(res.body).to.have.property("output", "");
                });
        });

        it("should return 404 if the task doesn't exist", () => {
            return request(app.server)
                .get(`/api/v1/task/000000000000000000000000`)
                .expect(404)
                .toPromise();
        });

        it("should allow you to create a new task for an action", () => {
            return request(app.server)
                .post(`/api/v1/action/${app.testAction._id}/tasks`)
                .send({
                    metadata: {
                        description: "This is a quick test task",
                        url: "http://localhost:9000/"
                    },
                    vars: {
                        x: "1"
                    }
                })
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.have.property("id");
                    chai.expect(res.body).to.have.property("metadata").eql({
                        description: "This is a quick test task",
                        url: "http://localhost:9000/"
                    });
                    chai.expect(res.body).to.have.property("action").eql(app.testTask.action);
                    chai.expect(res.body).to.have.property("project").eql(app.testTask.project);
                    chai.expect(res.body).to.have.property("vars").eql({
                        x: "1"
                    });
                    chai.expect(res.body).to.have.property("state").eql("NotExecuted");
                    chai.expect(res.body).to.have.property("output", "");
                });
        });

        it("should add the task to the database", () => {
            return chai.expect(app.db.Tasks.count()).to.eventually.eql(2);
        });

        it("should allow you to remove a task", () => {
            return request(app.server)
                .del(`/api/v1/task/${app.testTask._id}`)
                .expect(200)
                .toPromise();
        });

        it("should remove the task from the DB", () => {
            return chai.expect(app.db.Tasks.get(app.testTask._id)).to.eventually.not.exist;
        });
    });
});