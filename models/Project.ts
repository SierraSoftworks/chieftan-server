import * as Iridium from "iridium";

export interface ProjectSummaryDoc {
    id: string;
    name: string;
    url: string;
}

export const ProjectSummaryDocSchema = {
    id: String,
    name: String,
    url: String
};

export interface ProjectDoc {
    _id?: string;
    name: string;
    description: string;
    url: string;
}

@Iridium.Collection("projects")
export class Project extends Iridium.Instance<ProjectDoc, Project> {
    @Iridium.ObjectID
    _id: string;

    @Iridium.Property(String)
    name: string;

    @Iridium.Property(String)
    description: string;

    @Iridium.Property(String)
    url: string;

    get summary(): ProjectSummaryDoc {
        return {
            id: this._id,
            name: this.name,
            url: this.url
        };
    }

    toJSON() {
        return {
            id: this._id,
            name: this.name,
            description: this.description,
            url: this.url
        };
    }
}