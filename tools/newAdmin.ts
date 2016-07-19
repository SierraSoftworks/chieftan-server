import {Application} from "../Application";

let application = new Application({
    connectionString: process.env.CHIEFTAN_MONGODB || "mongodb://localhost/chieftan",
    port: process.env.port || 80 
});

function showHelp() {
    console.error("npm run create:admin FULL_NAME EMAIL");
    process.exit(1);
}

if (process.argv.length < 4) {
    showHelp();
}

const name = process.argv[2],
      email = process.argv[3];

if (!name || !email) {
    showHelp();
}

application.db.connect().then(() => application.db.Users.create({
    name: name,
    email: email,
    permissions: [
        "admin",
        "admin/users",
        "project/:project",
        "project/:project/admin"
    ]
})).then(user => {
    console.error("User created!");
    console.log(`${user._id}`);
    console.error(`Run 'create:token ${user._id}' to create an access token.'`);

    return application.db.close();
}).catch(err => {
    console.error("User not created!");
    console.error(err);

    return application.db.close().then(() => process.exit(1));
});