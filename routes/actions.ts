import {RouteBase} from "./base";
import {ActionDoc} from "../models/Action";

export class Actions extends RouteBase {
    register() {
        this.server.get("/api/v1/project/:project/actions", (req, res) => {
            this.db.Actions.find({
                "project.id": req.params.project
            }).toArray().then(actions => {
                res.send(200, actions);
            }).catch(err => this.catch(err).databaseError(res, err));
        });

        this.server.get("/api/v1/project/:project/action/:id", (req, res) => {
            this.db.Actions.get(req.params.id).then(action => {
                if(!action) return this.notFound(res);
                if (action.project.id !== req.params.project) return this.notFound(res);
                res.send(200, action);
            }).catch(err => this.catch(err).databaseError(res, err));
        });

        this.server.post("/api/v1/project/:project/actions", (req, res) => {
            this.db.Projects.get(req.params.project).then(project => {
                if (!project) return this.notFound(res);

                let newAction: ActionDoc = {
                    name: req.body.name,
                    description: req.body.description,
                    project: {
                        id: project._id,
                        name: project.name,
                        url: project.url
                    },
                    vars: req.body.vars || {},
                    http: req.body.http
                };

                return this.db.Actions.insert(newAction).then(action => {
                    res.send(200, action);
                });
            }).catch(err => this.catch(err).databaseError(res, err));
        });
    }
}