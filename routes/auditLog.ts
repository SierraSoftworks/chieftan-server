import {RouteBase} from "./base";
import {ActionDoc} from "../models/Action";
import {assign, pick} from "lodash";

export class AuditLog extends RouteBase {
    register() {
        this.server.get("/api/v1/audit", this.authorize(), this.permission("admin"), (req, res) => {
            this.db.AuditLog.find().sort({ timestamp: -1 }).toArray().then(logs => {
                res.send(logs);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.get("/api/v1/audit/:id", this.authorize(), this.permission("admin"), (req, res) => {
            this.db.AuditLog.get(req.params.id).then(log => {
                if (!log) return this.notFound();

                res.send(log);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
    }
}