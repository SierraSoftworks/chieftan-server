import {RouteBase} from "./base";

export class Releases extends RouteBase {
    register() {
        this.server.get("/releases", (req, res) => {
            this.db.Releases.find().toArray().then(releases => {
                res.json(200, releases);
            }).catch(err => this.serverError(res));
        });
        
        this.server.get("/releases/:project", (req, res) => {
            this.db.Releases.find({
                "project.id": req.params.project
            }).toArray().then(releases => {
                res.json(200, releases);
            }).catch(err => this.serverError(res));
        });
    }
}