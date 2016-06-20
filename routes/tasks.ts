import {RouteBase} from "./base";
import {TaskDoc} from "../models/Task";
import {Distributor} from "../executors/Distributor";
import {TaskState} from "../models/Task";

export class Tasks extends RouteBase {
    distributor = new Distributor(this.app);

    register() {
        this.server.get("/api/v1/tasks", (req, res) => {
            this.db.Tasks.find().toArray().then(tasks => {
                res.send(200, tasks);
            }).catch(err => this.catch(err).databaseError(res, err));
        });
        
        this.server.get("/api/v1/project/:project/tasks", (req, res) => {
            this.db.Tasks.find({
                "project.id": req.params.project
            }).toArray().then(tasks => {
                res.send(200, tasks);
            }).catch(err => this.catch(err).databaseError(res, err));
        });
        
        this.server.get("/api/v1/project/:project/action/:action/tasks", (req, res) => {
            this.db.Tasks.find({
                "project.id": req.params.project,
                'action.id': req.params.action
            }).toArray().then(tasks => {
                res.send(200, tasks);
            }).catch(err => this.catch(err).databaseError(res, err));
        });
        
        this.server.post("/api/v1/project/:project/action/:action/tasks", (req, res) => {
            this.db.Projects.get(req.params.project).then(project => {
                if (!project) return this.notFound(res);

                return this.db.Actions.get(req.params.action).then(action => {
                    if (!action) return this.notFound(res);
                    if (action.project.id !== project._id) return this.notFound(res);

                    let newTask: TaskDoc = {
                        action: {
                            id: action._id,
                            name: action.name,
                            description: action.description
                        },
                        project: {
                            id: project._id,
                            name: project.name,
                            url: project.url
                        },
                        metadata: {
                            description: req.body.metadata.description,
                            url: req.body.metadata.url
                        },
                        vars: req.body.vars
                    };

                    return this.db.Tasks.insert(newTask);
                }).then(task => {
                    res.send(200, task);
                });
            }).catch(err => this.catch(err).databaseError(res, err));
        });
        
        this.server.head("/api/v1/task/:id", (req, res) => {
            this.db.Tasks.get(req.params.id, {
                fields: { _id: 1 }
             }).then(task => {
                if (!task) res.status(404);
                else res.status(200);
                
                return res.end();
            }).catch(err => {
                this.catch(err);
                res.status(500);
                return res.end();
            });
        });
        
        this.server.get("/api/v1/task/:id", (req, res) => {
            this.db.Tasks.get(req.params.id).then(task => {
                if (!task) return this.notFound(res);
                
                res.send(200, task);
            }).catch(err => this.catch(err).databaseError(res, err));
        });
        
        this.server.post("/api/v1/task/:id/run", (req, res) => {
            this.db.Tasks.get(req.params.id).then(task => {
                if (!task) return this.notFound(res);

                return this.db.Actions.get(task.action.id).then(action => {
                    if (!action) return this.notFound(res);

                    const executors = this.distributor.getExecutors(action, task, req.body);
                    executors.forEach(executor => {
                        console.log(`START ${task.project.name}:${task.action.name}:${task._id} - ${executor.toString()}`);
                        executor.start().then(() => {
                            console.log(`STOP ${task.project.name}:${task.action.name}:${task._id} - ${executor.toString()} (${TaskState[task.state]})`);
                        });
                    });
                });
            }).catch(err => this.catch(err).databaseError(res, err));
        });
        
        this.server.del("/api/v1/task/:id", (req, res) => {
            this.db.Tasks.get(req.params.id).then(task => {
                if (!task) return this.notFound(res);
                
                return task.remove().then(() => {
                    res.status(200);
                    res.end();
                });
            }).catch(err => this.catch(err).databaseError(res, err));
        });
    }
}