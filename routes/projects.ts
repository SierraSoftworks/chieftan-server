import {RouteBase} from "./base";

export class Projects extends RouteBase {
    register() {
        this.server.get("/api/v1/projects", this.authorize(), (req, res) => {
            this.db.Projects.find().toArray().then(projects => {
                res.send(200, projects);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.get("/api/v1/project/:id", this.authorize(), (req, res) => {
            this.db.Projects.get(req.params.id).then(project => {
                if(!project) return this.notFound();
                res.send(200, project);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.post("/api/v1/projects", this.authorize(), (req, res) => {
            this.db.AuditLog.insert({
                type: "project.create",
                context: {
                    request: req.body
                }
            })
            .then(() => this.db.Projects.insert(req.body))
            .then(project => {
                if(this.isAuthorizedRequest(req)) {
                    req.user.permissions.push(`project/${project._id}`, `project/${project._id}/admin`);
                    return req.user.save().then(() => project);
                }

                return project;
            }).then(project => {
                res.send(200, project);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
    }
}