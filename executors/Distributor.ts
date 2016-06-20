import {Application} from "../Application";
import {ExecutorBase} from "./Executor";
import {Action} from "../models/Action";
import {Task} from "../models/Task";

export class Distributor {
    constructor(protected app: Application) {

    }

    getExecutors(action: Action, task: Task, vars: { [key: string]: string; } = {}): ExecutorBase[] {
        let executors: ExecutorBase[] = [];

        for(let configProperty in this.app.executors) {
            if(action[configProperty]) {
                const Executor = this.app.executors[configProperty];
                executors.push(new Executor(this.app.db, action, task, vars));
            }
        }

        return executors;
    }
}