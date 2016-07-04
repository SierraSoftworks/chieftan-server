import * as Iridium from "iridium";
import {ProjectSummaryDoc, ProjectSummarySchema} from "./Project";
import {ActionSummaryDoc, ActionSummarySchema} from "./Action";
import {TaskSummaryDoc, TaskSummarySchema} from "./Task";
import {UserSummaryDoc, UserSummarySchema} from "./User";

export interface AuditLogDoc {
    _id?: string;
    
    type: string;
    user: UserSummaryDoc;
    token: string;
    timestamp?: Date;

    context: AuditLogContextDoc;
}

export interface AuditLogContextDoc {
    user?: UserSummaryDoc;
    project?: ProjectSummaryDoc;
    action?: ActionSummaryDoc;
    task?: TaskSummaryDoc;
    request?: {};
}

export const AuditLogContextSchema = {
    user: { $required: false, $type: UserSummarySchema },
    project: { $required: false, $type: ProjectSummarySchema },
    action: { $required: false, $type: ActionSummarySchema },
    task: { $required: false, $type: TaskSummarySchema },
    changes: { $required: false, $type: Object },
}

@Iridium.Collection("logs")
@Iridium.Index({ timestamp: -1 })
export class AuditLog extends Iridium.Instance<AuditLogDoc, AuditLog> implements AuditLogDoc {
    @Iridium.ObjectID
    _id: string;
    
    @Iridium.Property(String)
    type: string;

    @Iridium.Property(UserSummarySchema)
    user: UserSummaryDoc;
    
    @Iridium.Property(String)
    token: string;
    
    @Iridium.Property(Date)
    timestamp: Date;
    
    @Iridium.Property(AuditLogContextSchema)
    context: AuditLogContextDoc;

    static onCreating(doc: AuditLogDoc) {
        doc.timestamp = doc.timestamp || new Date();
    }

    toJSON() {
        return {
            id: this._id,
            type: this.type,
            user: this.user,
            token: this.token,
            timestamp: this.timestamp,
            context: this.context
        };
    }
}