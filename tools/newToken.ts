import {Application} from "../Application";
import * as crypto from "crypto";

let application = new Application({
    connectionString: process.env.CHIEFTAN_MONGODB || "mongodb://localhost/chieftan",
    port: process.env.port || 80 
});

function showHelp() {
    console.error("npm run create:token USER_ID");
    process.exit(1);
}

if (process.argv.length < 3) {
    showHelp();
}

const id = process.argv[2];

if (!id) {
    showHelp();
}

application.db.connect().then(() => application.db.Users.get(id)).then(user => {
    if (!user) {
        console.error("User not found!");
        process.exit(2);
    }

    const newToken = crypto.randomBytes(16).toString("hex");
    user.tokens.push(newToken);

    return user.save().then(() => {
        console.error(`Token created for ${user.name}.`);
        console.log(newToken);
        return application.db.close();
    });
}).catch(err => {
    console.error("Token not added!");
    console.error(err);

    return application.db.close().then(() => process.exit(1));
});