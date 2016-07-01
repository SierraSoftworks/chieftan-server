import {Application} from "../Application";
import {Database} from "../models/Database";
import {User} from "../models/User";
import {Server, Request, Response, RequestHandler} from "restify";
import {assign} from "lodash";

export interface AuthorizedRequest extends Request {
    user: User;
}

export class RouteBase {
    constructor(protected app: Application, protected server: Server, protected db: Database) {
        
    }
    
    register() {
        
    }

    authorize(): RequestHandler {
        return (req: Request, res: Response, next: Function) => {
            if (req.authorization.scheme !== "Token") return this.unauthorized().catch(err => this.catch(res, err));

            this.db.Users.get({
                tokens: req.authorization.credentials
            }).then(user => {
                if (!user) return this.unauthorized();

                req.username = user._id;
                (<AuthorizedRequest>req).user = user;

                return next();
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        }
    }

    isAuthorizedRequest(req: Request): req is AuthorizedRequest {
        return !!(req && (<AuthorizedRequest>req).user);
    }

    permission(permission: string, context: { [name: string]: string; } = {}): RequestHandler {
        return (req: Request, res: Response, next: Function) => {
            if (!this.hasPermission(req, permission, context)) return this.forbidden().catch(err => this.catch(res, err));
            return next();
        }
    }

    hasPermission(req: Request, permission: string, context: { [name: string]: string; } = {}): boolean {
        if (!this.isAuthorizedRequest(req)) return false;
        else {
            if (!req.user.can(permission, assign<{}, { [name: string]: string; }>({}, context, req.params))) return false;
            return true;
        }
    }

    catch(res: Response, err: Error): Promise<Error> {
        if (res.headersSent) return Promise.resolve(err);

        if (err instanceof ServerError) {
            res.send(err.code, err);
            return Promise.resolve(err);
        } else {
            console.error(err);
            return this.serverError().then(err => res.send(500, err));
        }
    }

    databaseError(error: Error | SkmatcError | MongoDBError | ServerError) {
        return Promise.reject(error).catch(err => {
            if (err instanceof ServerError) return Promise.reject(err);
            if(this.isSkmatcError(error)) {
                return this.badRequest(error.message);
            } 
            if(this.isMongoDBError(error)) {
                if(error.code === 11000) {
                    return this.conflict();
                } else {
                    return this.serverError();
                }
            }

            return Promise.reject(err);
        });
    }

    unauthorized<T>() {
        return this.error<T>(401, "Unauthorized", "You haven't provided a required authentication header.");
    }

    forbidden<T>() {
        return this.error<T>(403, "Forbidden", "You do not have permission to access this method.");
    }

    conflict<T>() {
        return this.error<T>(409, "Conflict", "The entry you attempted to create already exists.");
    }

    badRequest<T>(message: string = "You did not provide a valid set of values for this request. Please check them and try again.") {
        return this.error<T>(400, "Bad Request", message);
    }
    
    notFound<T>() {
        return this.error<T>(404, "Not Found", "The entity you requested could not be found.");
    }
    
    serverError<T>() {
        return this.error<T>(500, "Server Error", "The server has encountered an error preventing it from serving your request.");
    }
    
    error<T>(code: number, error: string, message: string) {
        return Promise.reject<T>(new ServerError(code, error, message));
    }

    private isSkmatcError(error: any): error is SkmatcError {
        return error && error.isValidationError;
    }

    private isMongoDBError(error: any): error is MongoDBError {
        return error && error.code;
    }
}

export class ServerError extends Error {
    constructor(public code: number, name: string, message: string) {
        super(message);
        this.name = name;
    }

    toJSON() {
        return {
            code: this.code,
            error: this.name,
            message: this.message
        };
    }
}

interface SkmatcError extends Error {
    isValidationError: boolean;
}

interface MongoDBError extends Error {
    code: number;
}