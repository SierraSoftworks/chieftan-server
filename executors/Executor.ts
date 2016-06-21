import {Database} from "../models/Database";
import {Action} from "../models/Action";
import {Task, TaskState} from "../models/Task";
import * as handlebars from "handlebars";
import {cloneDeep, assign} from "lodash";

export class ExecutorBase {
    constructor(protected db: Database, protected action: Action, protected task: Task, vars: { [key: string]: string; } = {}) {
        assign(this.vars, action.vars, task.vars, vars);
    }

    static config = null;

    vars: { [key: string]: string; } = {};

    start() {
        this.task.executed = new Date();
        this.task.output = `::[info] Running task::`;
        this.task.state = TaskState.Executing;
        return this.task.save().then(() => {
            return this.run();
        }).then(() => {
            this.task.output += `\n::[info] Task complete in ${new Date().valueOf() - this.task.executed.valueOf()}ms::`
            this.task.state = TaskState.Passed;
        }, (err) => {
            this.task.output += `\n::[error] Task failed in ${new Date().valueOf() - this.task.executed.valueOf()}ms::`
            this.task.state = TaskState.Failed;
            this.task.output = `${this.task.output || ""}\n${err.message}\n${err.stack}`.trim();
        }).then(() => {
            this.task.completed = new Date();
            return this.task.save()
        });
    }

    protected run(): Promise<void> {
        return Promise.reject(new Error("No executor implemented for this task type."));
    }

    protected interpolate<T>(item: T): T {
        if(typeof item === "string") {
            let template = handlebars.compile(item);
            return <T><any>template(this.vars);
        } else if(typeof item === "object") {
            let target = cloneDeep(item)
            for(let k in target) {
                target[k] = this.interpolate(target[k]);
            }

            return target;
        } else {
            return item;
        }
    }
}

interface InterpolatableObject {
    [prop: string]: string|InterpolatableObject;
}