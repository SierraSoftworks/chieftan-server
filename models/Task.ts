import * as Iridium from "iridium";
import {ProjectSummaryDoc, ProjectSummarySchema} from "./Project";
import {ActionSummaryDoc, ActionSummarySchema} from "./Action";

export enum TaskState {
    NotExecuted = 0,
    Executing,
    Failed,
    Passed
}

export interface TaskDoc {
    _id?: string;

    created?: Date;
    executed?: Date;
    completed?: Date;
    
    action: ActionSummaryDoc;
    project: ProjectSummaryDoc;
    
    metadata?: {
        description?: string;
        url?: string;
    };
    
    vars: {
        [name: string]: string;
    };

    state?: TaskState;

    output?: string;
}

@Iridium.Collection("tasks")
export class Task extends Iridium.Instance<TaskDoc, Task> implements TaskDoc {
    @Iridium.ObjectID
    _id: string;

    @Iridium.Property(Date)
    created: Date;
    @Iridium.Property(Date, false)
    executed: Date;
    @Iridium.Property(Date, false)
    completed: Date;
    
    @Iridium.Property({
        description: { $required: false, $type: String },
        url: { $required: false, $type: String }
    })
    metadata: {
        description?: string;
        url?: string;
    };
    
    @Iridium.Property(ActionSummarySchema)
    action: ActionSummaryDoc;

    @Iridium.Property(ProjectSummarySchema)
    project: ProjectSummaryDoc;
    
    @Iridium.Property({
        $propertyType: String
    })
    vars: {
        [name: string]: string;
    };

    @Iridium.Property(Number)
    state: TaskState;

    @Iridium.Property(String)
    output: string;
    
    static onCreating(doc: TaskDoc) {
        doc.created = doc.created || new Date();
        doc.metadata = doc.metadata || {};
        doc.vars = doc.vars || {};
        doc.state = doc.state || TaskState.NotExecuted;
        doc.output = doc.output || "";
    }

    toJSON() {
        return {
            id: this._id,
            created: this.created,
            executed: this.executed,
            completed: this.completed,
            metadata: this.metadata,
            action: this.action,
            project: this.project,
            vars: this.vars,
            state: TaskState[this.state],
            output: this.output
        }
    }
}