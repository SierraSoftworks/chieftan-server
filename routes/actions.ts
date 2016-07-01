import {RouteBase} from "./base";
import {ActionDoc} from "../models/Action";
import {assign, pick} from "lodash";

export class Actions extends RouteBase {
    register() {
        this.server.get("/api/v1/project/:project/actions", this.authorize(), this.permission("project/:project"), (req, res) => {
            this.db.Actions.find({
                "project.id": req.params.project
            }).toArray().then(actions => {
                res.send(200, actions);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.get("/api/v1/action/:id", this.authorize(), (req, res) => {
            this.db.Actions.get(req.params.id).then(action => {
                if (!action) return this.notFound();
                if (!this.hasPermission(req, "project/:project", { project: action.project.id })) return this.forbidden();
                res.send(200, action);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.put("/api/v1/action/:id", this.authorize(), (req, res) => {
            this.db.Actions.get(req.params.id).then(action => {
                if (!action) return this.notFound();
                if (!this.hasPermission(req, "project/:project/admin", { project: action.project.id })) return this.forbidden();

                const changes = pick(req.body, "name", "description", "vars", "configurations", "http");

                return this.db.AuditLog.insert({
                    type: "action.update",
                    user: this.isAuthorizedRequest(req) ? req.user.summary : null,
                    context: {
                        project: action.project,
                        action: action.summary,
                        request: req.body
                    }
                }).then(() => {
                    assign(action, changes);
                    return action.save();
                });
            }, err => this.databaseError(err)).then(action => {
                res.send(200, action);
            }).catch(err => this.catch(res, err));
        });

        this.server.post("/api/v1/project/:project/actions", this.authorize(), this.permission("project/:project/admin"), (req, res) => {
            this.db.Projects.get(req.params.project).then(project => {
                if (!project) return this.notFound();

                let newAction: ActionDoc = {
                    name: req.body.name,
                    description: req.body.description,
                    project: {
                        id: project._id,
                        name: project.name,
                        url: project.url
                    },
                    vars: req.body.vars || {},
                    configurations: req.body.configurations || {},
                    http: req.body.http
                };

                return this.db.AuditLog.insert({
                    type: "action.create",
                    user: this.isAuthorizedRequest(req) ? req.user.summary : null,
                    context: {
                        project: project.summary,
                        request: req.body
                    }
                })
                .then(() => this.db.Actions.insert(newAction))
                .then(action => {
                    res.send(200, action);
                    return action;
                }, err => this.databaseError(err));
            }).catch(err => this.catch(res, err));
        });
    }
}