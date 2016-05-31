export interface Project {
    id: string;
    name: string;
    url: string;
}

export const ProjectSchema = {
    id: String,
    name: String,
    url: String
};