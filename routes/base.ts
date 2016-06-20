import {Application} from "../Application";
import {Database} from "../models/Database";
import {Server, Response} from "restify";

export class RouteBase {
    constructor(protected app: Application, protected server: Server, protected db: Database) {
        
    }
    
    register() {
        
    }

    catch(err: Error) {
        console.error(`\nRoute Error Handler\n${err.stack}\n`);
        return this;
    }

    databaseError(res: Response, error: Error | SkmatcError | MongoDBError) {
        if(this.isSkmatcError(error)) {
            return this.badRequest(res, error.message);
        } else if(this.isMongoDBError(error)) {
            if(error.code === 11000) {
                this.conflict(res);
            } else {
                this.serverError(res);
            }
        }
    }

    conflict(res: Response) {
        return this.error(res, 409, "Conflict", "The entry you attempted to create already exists.");
    }

    badRequest(res: Response, message: string = "You did not provide a valid set of values for this request. Please check them and try again.") {
        return this.error(res, 400, "Bad Request", message);
    }
    
    notFound(res: Response) {
        return this.error(res, 404, "Not Found", "The entity you requested could not be found.");
    }
    
    serverError(res: Response) {
        return this.error(res, 500, "Server Error", "The server has encountered an error preventing it from serving your request.");
    }
    
    error(res: Response, code: number, error: string, message: string) {
        return res.send(code, {
            code,
            error,
            message
        });
    }

    private isSkmatcError(error: any): error is SkmatcError {
        return error && error.isValidationError;
    }

    private isMongoDBError(error: any): error is MongoDBError {
        return error && error.code;
    }
}

interface SkmatcError extends Error {
    isValidationError: boolean;
}

interface MongoDBError extends Error {
    code: number;
}