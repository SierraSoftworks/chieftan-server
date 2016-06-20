import * as restify from "restify";
import * as Iridium from "iridium";
import {Database} from "./models/Database";
import {AllRoutes} from "./routes/all";

export class Application {
    constructor(protected options: ApplicationOptions) {
        this.db = new Database(options.connectionString);
        this.server = restify.createServer({
            name: "Chief Server"
        });

        this.server
        .use(restify.queryParser())
        .use(restify.bodyParser());

        new AllRoutes(this.server, this.db).register();
    }

    public server: restify.Server;
    public db: Database;

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

    stop() {
        this.server.close();
        return this.db.close();
    }
}

export interface ApplicationOptions {
    port?: number;
    connectionString: string;
}