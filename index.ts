import * as restify from "restify";
import * as Iridium from "iridium";
import {Database} from "./models/Database";

let server = restify.createServer({
    name: "Chief Server"
});

let db = new Database("mongodb://localhost/chief");

db.connect().then(() => {
    server.listen(process.env.port || 80, () => {
        console.log("%s listening at %s", server.name, server.url);
    });
}).catch(err => {
    console.error(err);
});