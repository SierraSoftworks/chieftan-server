import * as restify from "restify";
import * as Iridium from "iridium";
import {Database} from "./models/Database";
import {AllRoutes} from "./routes/all";
import {ExecutorBase} from "./executors/Executor";
import {RegisterExecutors} from "./executors/all";

export class Application {
    constructor(protected options: ApplicationOptions) {
        this.db = new Database(options.connectionString);
        this.server = restify.createServer({
            name: "Chief Server"
        });

        this.server
            .use(restify.queryParser())
            .use(restify.bodyParser())
            .use(restify.CORS());

        new AllRoutes(this, this.server, this.db).register();
        RegisterExecutors(this);
    }

    public server: restify.Server;
    public db: Database;
    public executors: {
        [configProperty: string]: typeof ExecutorBase
    } = {};

    get url(): string {
        return this.server.url;
    }

    start() {
        return this.db.connect().then(() => {
            return new Promise((resolve, reject) => {
                this.server.listen(this.options.port, err => {
                    if (err) return reject(err);
                    return resolve();
                });
            });
        });
    }

    use(executor: typeof ExecutorBase) {
        if (!executor.config)
            throw new Error("Your executor should specify a config field name.");

        this.executors[executor.config] = executor;
        return this;
    }

    stop() {
        this.server.close();
        return this.db.close();
    }
}

export interface ApplicationOptions {
    port?: number;
    connectionString: string;
}