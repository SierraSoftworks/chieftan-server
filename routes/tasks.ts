import {RouteBase, ServerError} from "./base";
import {TaskDoc} from "../models/Task";
import {Distributor} from "../executors/Distributor";
import {TaskState} from "../models/Task";
import {assign} from "lodash";

export class Tasks extends RouteBase {
    distributor = new Distributor(this.app);

    register() {
        this.server.get("/api/v1/tasks", this.authorize(), this.permission("admin"), (req, res) => {
            this.db.Tasks.find().sort({
                created: -1
            }).toArray().then(tasks => {
                res.send(tasks);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
        
        this.server.get("/api/v1/project/:project/tasks", this.authorize(), this.permission("project/:project"), (req, res) => {
            this.db.Tasks.find({
                "project.id": req.params.project
            }).sort({
                created: -1
            }).toArray().then(tasks => {
                res.send(tasks);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
        
        this.server.get("/api/v1/project/:project/tasks/recent", this.authorize(), this.permission("project/:project"), (req, res) => {
            this.db.Tasks.find({
                "project.id": req.params.project
            }).sort({
                created: -1
            }).limit(50).toArray().then(tasks => {
                res.send(tasks);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
        
        this.server.get("/api/v1/action/:action/tasks", this.authorize(), (req, res) => {
            this.db.Actions.get(req.params.action).then(action => {
                if (!action) return this.notFound();
                if (!this.hasPermission(req, "project/:project", { project: action.project.id })) return this.forbidden();
            }).then(() => this.db.Tasks.find({
                    'action.id': req.params.action
                }).sort({
                    created: -1
                }).toArray())
            .then(tasks => {
                res.send(tasks);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
        
        this.server.get("/api/v1/action/:action/tasks/recent", this.authorize(), (req, res) => {
            this.db.Actions.get(req.params.action).then(action => {
                if (!action) return this.notFound();
                if (!this.hasPermission(req, "project/:project", { project: action.project.id })) return this.forbidden();
            }).then(() => this.db.Tasks.find({
                    'action.id': req.params.action
                }).sort({
                    created: -1
                }).limit(50).toArray())
            .then(tasks => {
                res.send(200, tasks);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
        
        this.server.post("/api/v1/action/:action/tasks", this.authorize(), (req, res) => {
            this.db.Actions.get(req.params.action).then(action => {
                if (!action) return this.notFound();
                if (!this.hasPermission(req, "project/:project/admin", { project: action.project.id })) return this.forbidden();

                let newTask: TaskDoc = {
                    action: action.summary,
                    project: action.project,
                    metadata: {
                        description: req.body.metadata.description,
                        url: req.body.metadata.url
                    },
                    vars: req.body.vars
                };

                return this.db.AuditLog.insert({
                    type: "task.create",
                    context: {
                        project: action.project,
                        action: action.summary,
                        request: req.body
                    }
                }).then(() => this.db.Tasks.insert(newTask));
            }).then(task => {
                res.send(task);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
        
        this.server.head("/api/v1/task/:id", this.authorize(), (req, res) => {
            this.db.Tasks.get(req.params.id, {
                fields: { _id: 1, project: 1 }
             }).then(task => {
                if (!task) return this.notFound();
                else if (!this.hasPermission(req, "project/:project", { project: task.project.id })) return this.forbidden();
                
                res.end();
            }).catch(err => {
                if (err instanceof ServerError) {
                    res.status(err.code);
                }
                else res.status(500);
                res.end();
            });
        });
        
        this.server.get("/api/v1/task/:id", this.authorize(), (req, res) => {
            this.db.Tasks.get(req.params.id).then(task => {
                if (!task) return this.notFound();
                if (!this.hasPermission(req, "project/:project", { project: task.project.id })) return this.forbidden();
                
                res.send(task);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
        
        this.server.post("/api/v1/task/:id/run", this.authorize(), (req, res) => {
            this.db.Tasks.get(req.params.id).then(task => {
                if (!task) return this.notFound();
                if (!this.hasPermission(req, "project/:project", { project: task.project.id })) return this.forbidden();

                const { vars, configuration } = <{
                    vars: { [name: string]: string; };
                    configuration?: string;
                }>req.body;

                return this.db.Actions.get(task.action.id).then(action => {
                    if (!action) return this.notFound();

                    let resolvedVars: { [name: string]: string; } = {};
                    assign(resolvedVars, action.vars, task.vars, vars);
                    if (configuration) {
                        const resolvedConfiguration = action.configurations.find(config => config.name === configuration);
                        if (resolvedConfiguration)
                            assign(resolvedVars, resolvedConfiguration.vars);
                        else
                            return this.notFound();
                    }

                    return this.db.AuditLog.insert({
                        type: "task.run",
                        context: {
                            project: action.project,
                            action: action.summary,
                            task: task.summary,
                            request: req.body
                        }
                    }).then(() => {
                        const executors = this.distributor.getExecutors(action, task, resolvedVars);
                        executors.forEach(executor => {
                            console.log(`START ${task.project.name}:${task.action.name}:${task._id} - ${executor.toString()}`);
                            executor.start().then(() => {
                                console.log(`STOP ${task.project.name}:${task.action.name}:${task._id} - ${executor.toString()} (${TaskState[task.state]})`);
                            });
                        });

                        res.send(task);
                        return task;
                    });
                });
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
        
        this.server.del("/api/v1/task/:id", this.authorize(), (req, res) => {
            this.db.Tasks.get(req.params.id).then(task => {
                if (!task) return this.notFound();
                if (!this.hasPermission(req, "project/:project/admin", { project: task.project.id })) return this.forbidden();

                
                return this.db.AuditLog.insert({
                    type: "task.remove",
                    context: {
                        project: task.project,
                        action: task.action,
                        task: task.summary
                    }
                })
                .then(() => task.remove())
            }, err => this.databaseError(err))
            .then(() => {
                res.status(200);
                res.end();
            }).catch(err => this.catch(res, err));
        });
    }
}