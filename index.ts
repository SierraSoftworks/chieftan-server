import {Application} from "./Application";

let application = new Application({
    connectionString: process.env.CHIEFTAN_MONGODB || "mongodb://localhost/chieftan",
    port: process.env.PORT || 3000 
});

application.start().then(() => {
    console.log("Listening on %s", application.url);
}).catch(err => {
    console.error(err.stack);
});