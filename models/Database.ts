import * as Iridium from "iridium";
import {Action, ActionDoc} from "./Action";
import {Project, ProjectDoc} from "./Project";
import {Task, TaskDoc} from "./Task";

export class Database extends Iridium.Core {
    Actions = new Iridium.Model<ActionDoc, Action>(this, Action);
    Projects = new Iridium.Model<ProjectDoc, Project>(this, Project);
    Tasks = new Iridium.Model<TaskDoc, Task>(this, Task);
}