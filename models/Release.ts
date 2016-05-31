import * as Iridium from "iridium";
import {Project, ProjectSchema} from "./Project";

export interface ReleaseDoc {
    _id?: string;
    
    project: Project;
    
    version: {
        id: string;
        name: string;
        url: string;
    };
    
    vars: {
        [name: string]: string;
    };
}

@Iridium.Collection("releases")
export class Release extends Iridium.Instance<ReleaseDoc, Release> implements ReleaseDoc {
    @Iridium.ObjectID
    _id: string;
    
    @Iridium.Property({
        id: String,
        name: String,
        url: String
    })
    version: {
        id: string;
        name: string;
        url: string;
    };
    
    @Iridium.Property(ProjectSchema)
    project: Project;
    
    @Iridium.Property({
        $propertyType: String
    })
    vars: {
        [name: string]: string;
    };
    
    static onCreating(doc: ReleaseDoc) {
        doc.vars = doc.vars || {};
    }
}