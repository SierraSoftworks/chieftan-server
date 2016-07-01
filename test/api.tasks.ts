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
                .set("Authorization", "Token admin")
                .expect(200)
                .toPromise()
                .then(res => {
                    chai.expect(res.body).to.exist;
                    chai.expect(res.body).to.be.a("array");
                    chai.expect(res.body).to.have.length(2);

                    res.body.forEach(item => {
                        const { id } = item;
                        let task = app.testTask;
                        if (id === app.testTask2._id) task = app.testTask2;

                        chai.expect(item).to.have.property("id", task._id);
                        chai.expect(item).to.have.property("metadata").eql(task.metadata);
                        chai.expect(item).to.have.property("action").eql(task.action);
                        chai.expect(item).to.have.property("project").eql(task.project);
                        chai.expect(item).to.have.property("vars").eql(task.vars);
                        chai.expect(item).to.have.property("state").eql("NotExecuted");
                        chai.expect(item).to.have.property("output", "");
                        chai.expect(item).to.have.property("created").which.exist;
                    });
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/tasks`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't have the admin permission", () => {
            return request(app.server)
                .get(`/api/v1/tasks`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });
        
        it("should return the list of tasks for a project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/tasks`)
                .set("Authorization", "Token test")
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
                        chai.expect(item).to.have.property("created").which.exist;
                    });
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/tasks`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't permission to access the project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject2._id}/tasks`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });
        
        it("should return the recent list of tasks for a project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/tasks/recent`)
                .set("Authorization", "Token test")
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
                        chai.expect(item).to.have.property("created").which.exist;
                    });
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject._id}/tasks/recent`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't permission to access the project", () => {
            return request(app.server)
                .get(`/api/v1/project/${app.testProject2._id}/tasks/recent`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });
        
        it("should return the list of tasks for a project and action", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction._id}/tasks`)
                .set("Authorization", "Token test")
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
                        chai.expect(item).to.have.property("created").which.exist;
                    });
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction._id}/tasks`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't permission to access the project", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction2._id}/tasks`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });
        
        it("should return the recent list of tasks for a project and action", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction._id}/tasks/recent`)
                .set("Authorization", "Token test")
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
                        chai.expect(item).to.have.property("created").which.exist;
                    });
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction._id}/tasks/recent`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't permission to access the project", () => {
            return request(app.server)
                .get(`/api/v1/action/${app.testAction2._id}/tasks/recent`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should allow you to check whether a task exists by ID", () => {
            return request(app.server)
                .head(`/api/v1/task/${app.testTask._id}`)
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise();
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .head(`/api/v1/task/${app.testTask._id}`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't permission to access the project", () => {
            return request(app.server)
                .head(`/api/v1/task/${app.testTask2._id}`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should allow you to check whether a task doesn't exist by ID", () => {
            return request(app.server)
                .head(`/api/v1/task/000000000000000000000000`)
                .set("Authorization", "Token test")
                .expect(404)
                .toPromise();
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .head(`/api/v1/task/000000000000000000000000`)
                .expect(401)
                .toPromise();
        });

        it("should let you request a specific task by ID", () => {
            return request(app.server)
                .get(`/api/v1/task/${app.testTask._id}`)
                .set("Authorization", "Token test")
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
                    chai.expect(res.body).to.have.property("created").which.exist;
                });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .get(`/api/v1/task/${app.testTask._id}`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't permission to access the project", () => {
            return request(app.server)
                .get(`/api/v1/task/${app.testTask2._id}`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should return 404 if the task doesn't exist", () => {
            return request(app.server)
                .get(`/api/v1/task/000000000000000000000000`)
                .set("Authorization", "Token test")
                .expect(404)
                .toPromise();
        });

        it("should allow you to create a new task for an action", () => {
            return request(app.server)
                .post(`/api/v1/action/${app.testAction._id}/tasks`)
                .set("Authorization", "Token test")
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
                    chai.expect(res.body).to.have.property("created").which.exist;
                });
        });

        it("should add the task to the database", () => {
            return chai.expect(app.db.Tasks.count()).to.eventually.eql(3);
        });

        it("should create an auditlog entry", () => {
            return app.db.AuditLog.get({ type: "task.create" }).then(entry => {
                chai.expect(entry).to.exist;
                chai.expect(entry).to.have.property("context");

                chai.expect(entry).to.have.property("user").eql(app.testUser.summary);
                chai.expect(entry.context).to.have.property("project").eql(app.testProject.summary);
                chai.expect(entry.context).to.have.property("action").eql(app.testAction.summary);
                chai.expect(entry.context).to.have.property("request").eql({
                    metadata: {
                        description: "This is a quick test task",
                        url: "http://localhost:9000/"
                    },
                    vars: {
                        x: "1"
                    }
                });
            });
        });

        it("should return 401 if you aren't authenticated", () => {
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
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't permission to access the project", () => {
            return request(app.server)
                .post(`/api/v1/action/${app.testAction2._id}/tasks`)
                .send({
                    metadata: {
                        description: "This is a quick test task",
                        url: "http://localhost:9000/"
                    },
                    vars: {
                        x: "1"
                    }
                })
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });

        it("should allow you to remove a task", () => {
            return request(app.server)
                .del(`/api/v1/task/${app.testTask._id}`)
                .set("Authorization", "Token test")
                .expect(200)
                .toPromise();
        });

        it("should remove the task from the DB", () => {
            return chai.expect(app.db.Tasks.get(app.testTask._id)).to.eventually.not.exist;
        });

        it("should create an auditlog entry", () => {
            return app.db.AuditLog.get({ type: "task.remove" }).then(entry => {
                chai.expect(entry).to.exist;
                chai.expect(entry).to.have.property("context");

                chai.expect(entry).to.have.property("user").eql(app.testUser.summary);
                chai.expect(entry.context).to.have.property("project").eql(app.testProject.summary);
                chai.expect(entry.context).to.have.property("action").eql(app.testAction.summary);
                chai.expect(entry.context).to.have.property("task").eql(app.testTask.summary);
            });
        });

        it("should return 401 if you aren't authenticated", () => {
            return request(app.server)
                .del(`/api/v1/task/${app.testTask._id}`)
                .expect(401)
                .toPromise();
        });

        it("should return 403 if you don't permission to access the project", () => {
            return request(app.server)
                .del(`/api/v1/task/${app.testTask2._id}`)
                .set("Authorization", "Token test")
                .expect(403)
                .toPromise();
        });
    });
});