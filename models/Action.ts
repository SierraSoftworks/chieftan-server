import * as Iridium from "iridium";
import {RequestDoc, RequestSchema} from "./Request";
import {ProjectSummaryDoc, ProjectSummarySchema} from "./Project";

export interface ActionDoc {
    _id?: string;
    name: string;
    description: string;
    
    project: ProjectSummaryDoc;
    
    vars: {
        [name: string]: string;
    };
    
    http?: RequestDoc;
}

export interface ActionSummaryDoc {
    id: string;
    name: string;
    description: string;
}

export const ActionSummarySchema = {
    id: String,
    name: String,
    description: String
};

@Iridium.Collection("actions")
export class Action extends Iridium.Instance<ActionDoc, Action> implements ActionDoc {
    @Iridium.ObjectID
    _id: string;
    
    @Iridium.Property(String)
    name: string;
    
    @Iridium.Property(String)
    description: string;
    
    @Iridium.Property(ProjectSummarySchema)
    project: ProjectSummaryDoc;
    
    @Iridium.Property({
        $propertyType: String
    })
    vars: {
        [name: string]: string;
    };
    
    @Iridium.Property(RequestSchema, false)
    http: RequestDoc;

    get summary(): ActionSummaryDoc {
        return {
            id: this._id,
            name: this.name,
            description: this.description
        };
    }

    toJSON() {
        return {
            id: this._id,
            name: this.name,
            description: this.description,
            project: this.project,
            vars: this.vars,
            http: this.http
        };
    }
}