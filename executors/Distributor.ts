import {Database} from "../models/Database";
import {ExecutorBase} from "./Executor";
import {HttpExecutor} from "./Http";
import {Action} from "../models/Action";
import {Task, TaskState} from "../models/Task";

const executors: typeof ExecutorBase[] = [
    HttpExecutor
];

export class Distributor {
    constructor(protected db: Database) {
        for(const executor of executors) {
            this.executors[executor.config] = executor;
        }
    }

    executors : {
        [configProperty: string]: typeof ExecutorBase
    } = {};

    getExecutors(action: Action, task: Task, vars: { [key: string]: string; } = {}): ExecutorBase[] {
        let executors: ExecutorBase[] = [];

        for(let configProperty in this.executors) {
            if(action[configProperty]) {
                const Executor = this.executors[configProperty];
                executors.push(new Executor(this.db, action, task, vars));
            }
        }

        return executors;
    }
}