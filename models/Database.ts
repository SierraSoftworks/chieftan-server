import * as Iridium from "iridium";
import {Action, ActionDoc} from "./Action";
import {AuditLog, AuditLogDoc} from "./AuditLog";
import {Project, ProjectDoc} from "./Project";
import {Task, TaskDoc} from "./Task";
import {User, UserDoc} from "./User";

export class Database extends Iridium.Core {
    Actions: Iridium.Model<ActionDoc, Action> = new Iridium.Model<ActionDoc, Action>(this, Action);
    AuditLog: Iridium.Model<AuditLogDoc, AuditLog> = new Iridium.Model<AuditLogDoc, AuditLog>(this, AuditLog);
    Projects: Iridium.Model<ProjectDoc, Project> = new Iridium.Model<ProjectDoc, Project>(this, Project);
    Tasks: Iridium.Model<TaskDoc, Task> = new Iridium.Model<TaskDoc, Task>(this, Task);
    Users: Iridium.Model<UserDoc, User> = new Iridium.Model<UserDoc, User>(this, User);

    onConnected() {
        return super.onConnected().then(() => Promise.all([
            this.Actions.ensureIndexes(),
            this.AuditLog.ensureIndexes(),
            this.Projects.ensureIndexes(),
            this.Tasks.ensureIndexes(),
            this.Users.ensureIndexes()
        ])).then(() => {
            
        });
    }
}