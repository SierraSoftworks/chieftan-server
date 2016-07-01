import {Application, ApplicationOptions} from "../../Application";
import {Project} from "../../models/Project";
import {Action} from "../../models/Action";
import {Task} from "../../models/Task";
import {User} from "../../models/User";

export class TestApplication extends Application {
    constructor(options: ApplicationOptions = {
        connectionString: "mongodb://localhost/chieftan_test"
    }) {
        super(options);
    }
    
    testProject: Project;
    testProject2: Project;
    testAction: Action;
    testAction2: Action;
    testTask: Task;
    testTask2: Task;
    adminUser: User;
    testUser: User;

    setup() {
        return this.db.connect().then(() => this.cleanDB()).then(() => this.seedDB());
    }

    teardown() {
        return this.db.close();
    }

    reset() {
        return this.cleanDB().then(() => this.seedDB());
    }

    protected cleanDB() {
        return Promise.all([
            this.db.Actions.remove(),
            this.db.AuditLog.remove(),
            this.db.Projects.remove(),
            this.db.Tasks.remove(),
            this.db.Users.remove(),
        ]);
    }

    protected seedDB() {
        return this.db.Projects.insert({
            name: "Test Project",
            description: "This is a quick test project",
            url: "https://github.com/SierraSoftworks/Chieftan"
        })
        .then(project => this.testProject = project)
        .then(() => this.db.Actions.insert({
            name: "Test Action",
            description: "A test action",
            project: this.testProject.summary,
            vars: {
                query: "default_query_value",
                header: "default_header_value",
                data: "default_data_value"
            },
            configurations: [

            ],
            http: {
                method: "POST",
                url: "http://localhost:1234/test/path?query={{query}}",
                headers: {
                    "X-Header": "{{header}}"
                },
                data: {
                    value: "{{data}}"
                }
            }
        }))
        .then(action => this.testAction = action)
        .then(() => this.db.Tasks.insert({
            action: this.testAction.summary,
            metadata: {
                description: "A test task",
                url: "https://github.com/SierraSoftworks/Chieftan"
            },
            project: this.testProject.summary,
            vars: {
                query: "query_param",
                header: "header_value"
            }
        }))
        .then(task => this.testTask = task)

        /* Second set of project->action->task for permissions checking */
        .then(() => this.db.Projects.insert({
            name: "Test Project 2",
            description: "This is a quick test project",
            url: "https://github.com/SierraSoftworks/Chieftan"
        }))
        .then(project => this.testProject2 = project)
        .then(() => this.db.Actions.insert({
            name: "Test Action 2",
            description: "A test action",
            project: this.testProject2.summary,
            vars: {
                query: "default_query_value",
                header: "default_header_value",
                data: "default_data_value"
            },
            configurations: [

            ],
            http: {
                method: "POST",
                url: "http://localhost:1234/test/path?query={{query}}",
                headers: {
                    "X-Header": "{{header}}"
                },
                data: {
                    value: "{{data}}"
                }
            }
        }))
        .then(action => this.testAction2 = action)
        .then(() => this.db.Tasks.insert({
            action: this.testAction2.summary,
            metadata: {
                description: "A test task",
                url: "https://github.com/SierraSoftworks/Chieftan"
            },
            project: this.testProject2.summary,
            vars: {
                query: "query_param",
                header: "header_value"
            }
        }))
        .then(task => this.testTask2 = task)

        /* Users */
        .then(() => this.db.Users.insert({
            name: "Admin User",
            email: "admim@test.com",
            permissions: [
                "admin",
                "admin/projects",
                "admin/users"
            ],
            tokens: [
                "admin"
            ]
        })).then(user => this.adminUser = user)
        .then(() => this.db.Users.insert({
            name: "Test User",
            email: "test@test.com",
            permissions: [
                `project/000000000000000000000000`,
                `project/000000000000000000000000/admin`,
                `project/${this.testProject._id}`,
                `project/${this.testProject._id}/admin`
            ],
            tokens: [
                "test"
            ]
        })).then(user => this.testUser = user);
    }

    private logThrough<T>(item: T): T {
        console.log(require('util').inspect(item));
        return item;
    }
}