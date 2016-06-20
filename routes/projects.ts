import {RouteBase} from "./base";

export class Projects extends RouteBase {
    register() {
        this.server.get("/api/v1/projects", (req, res) => {
            this.db.Projects.find().toArray().then(projects => {
                res.send(200, projects);
            }).catch(err => this.catch(err).databaseError(res, err));
        });

        this.server.get("/api/v1/project/:id", (req, res) => {
            this.db.Projects.get(req.params.id).then(project => {
                if(!project) return this.notFound(res);
                res.send(200, project);
            }).catch(err => this.catch(err).databaseError(res, err));
        });

        this.server.post("/api/v1/projects", (req, res) => {
            this.db.Projects.insert(req.body).then(project => {
                res.send(200, project);
            }).catch(err => this.catch(err).databaseError(res, err));
        });
    }
}