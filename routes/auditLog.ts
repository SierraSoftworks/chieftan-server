import {RouteBase} from "./base";
import {ActionDoc} from "../models/Action";
import {assign, pick} from "lodash";

export class AuditLog extends RouteBase {
    register() {
        this.server.get("/api/v1/audit", this.authorize(), (req, res) => {
            this.db.AuditLog.find().sort({ timestamp: -1 }).toArray().then(logs => {
                res.send(200, logs);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });
    }
}