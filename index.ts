import {Application} from "./Application";

let application = new Application({
    connectionString: process.env.CHIEFTAN_MONGODB || "mongodb://localhost/chieftan",
    port: process.env.port || 80 
});

application.start().then(() => {
    console.log("Listening on %s", application.url);
}).catch(err => {
    console.log(err);
});