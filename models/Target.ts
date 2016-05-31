import * as Iridium from "iridium";
import {Request, RequestSchema} from "./Request";
import {Project, ProjectSchema} from "./Project";

export interface TargetDoc {
    _id?: string;
    name: string;
    description: string;
    
    project: Project;
    
    deploy: Request;
}

@Iridium.Collection("targets")
export class Target extends Iridium.Instance<TargetDoc, Target> implements TargetDoc {
    @Iridium.ObjectID
    _id: string;
    
    @Iridium.Property(String)
    name: string;
    
    @Iridium.Property(String)
    description: string;
    
    @Iridium.Property(ProjectSchema)
    project: Project;
    
    @Iridium.Property(RequestSchema)
    deploy: Request;
}