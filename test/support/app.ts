import {Application, ApplicationOptions} from "../../Application";
import {Project} from "../../models/Project";
import {Action} from "../../models/Action";
import {Task} from "../../models/Task";

export class TestApplication extends Application {
    constructor(options: ApplicationOptions = {
        connectionString: "mongodb://localhost/chieftan_test"
    }) {
        super(options);
    }
    
    testProject: Project;
    testAction: Action;
    testTask: Task;

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
            this.db.Projects.remove(),
            this.db.Actions.remove(),
            this.db.Tasks.remove()
        ]);
    }

    protected seedDB() {
        return this.db.Projects.insert({
            name: "Test Project",
            description: "This is a quick test project",
            url: "https://github.com/SierraSoftworks/Chief"
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
                url: "https://github.com/SierraSoftworks/Chef"
            },
            project: this.testProject.summary,
            vars: {
                query: "query_param",
                header: "header_value"
            }
        }))
        .then(task => this.testTask = task);
    }

    private logThrough<T>(item: T): T {
        console.log(require('util').inspect(item));
        return item;
    }
}