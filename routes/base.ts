import {Database} from "../models/Database";
import {Server, Response} from "restify";

export class RouteBase {
    constructor(protected server: Server, protected db: Database) {
        
    }
    
    register() {
        
    }
    
    notFound(res: Response) {
        return this.error(res, 404, "Not Found", "The entity you requested could not be found.");
    }
    
    serverError(res: Response) {
        return this.error(res, 500, "Server Error", "The server has encountered an error preventing it from serving your request.");
    }
    
    error(res: Response, code: number, error: string, message: string) {
        return res.json(code, {
            code,
            error,
            message
        });
    }
}