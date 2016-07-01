import * as Iridium from "iridium";
import {Action, ActionDoc} from "./Action";
import {Project, ProjectDoc} from "./Project";
import {Task, TaskDoc} from "./Task";

export class Database extends Iridium.Core {
    Actions: Iridium.Model<ActionDoc, Action> = new Iridium.Model<ActionDoc, Action>(this, Action);
    Projects: Iridium.Model<ProjectDoc, Project> = new Iridium.Model<ProjectDoc, Project>(this, Project);
    Tasks: Iridium.Model<TaskDoc, Task> = new Iridium.Model<TaskDoc, Task>(this, Task);

    onConnected() {
        return super.onConnected().then(() => Promise.all([
            this.Actions.ensureIndexes(),
            this.Projects.ensureIndexes(),
            this.Tasks.ensureIndexes(),
        ])).then(() => {
            
        });
    }
}