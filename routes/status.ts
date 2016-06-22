import {RouteBase} from "./base";

export class Status extends RouteBase {
    register() {
        const started = new Date();

        this.server.get("/api/v1/status", (req, res) => {
            res.send(200, {
                started
            });
        });
    }
}