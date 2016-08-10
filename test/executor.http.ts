import {TestApplication} from "./support/app";
import {HttpExecutor} from "../executors/Http";
import {RequestDoc} from "../models/Request";
import * as chai from "chai";
import * as http from "http";

class TestHttpExecutor extends HttpExecutor {
    static config = "http";

    testBuildRequestOptions(request: RequestDoc) {
        return this.buildRequestOptions(request);
    }
}

class TestServer {
    constructor() {
        this.server = http.createServer((req, res) => {
            this.handle(req, res);
        });
    }

    server: http.Server;

    statusCode: number = 200;
    sendData: any;

    reqMethod: string;
    reqUrl: string;
    reqHeaders: { [header: string]: string; };
    reqData: Promise<string> = Promise.resolve(null);

    handle(req: http.IncomingMessage, res: http.ServerResponse) {
        this.reqMethod = req.method;
        this.reqUrl = req.url;
        this.reqHeaders = req.headers;
        req.setEncoding("utf8");

        this.reqData = new Promise<string>((resolve, reject) => {
            let data = "";
            req.on("data", d => {
                data += d;
            });

            req.on("end", () => {
                resolve(data || null);
            });
        });

        res.statusCode = this.statusCode;

        if (this.sendData !== undefined)
            res.write(JSON.stringify(this.sendData));

        res.end();
    }

    listen(port: number) {
        this.server.listen(port);
    }

    reset() {
        this.reqData = Promise.resolve(null);
        this.reqMethod = undefined;
        this.reqHeaders = undefined;
        this.reqUrl = undefined;
    }

    close() {
        this.server.close();
    }
}

describe("executor", () => {

    describe("http", () => {
        const app = new TestApplication();

        before(() => app.setup());
        beforeEach(() => app.reset());
        after(() => app.teardown());

        it("should be marked as the handler for the http config", () => {
            chai.expect(HttpExecutor.config).to.eql("http");
        });

        describe("run", () => {
            let port: number = Math.floor(Math.random() * 40000 + 10000);

            let server: TestServer;
            before(() => server = new TestServer());
            before(() => server.listen(port));
            afterEach(() => server.reset());
            after(() => server.close());

            let executor: HttpExecutor;
            beforeEach(() => app.testAction.http.url = `http://localhost:${port}/test/path?query={{query}}`);
            beforeEach(() => executor = new HttpExecutor(app.db, app.testAction, app.testTask));
            after(() => app.reset());

            it("should make a request to the correct url", () => {
                return executor.run().then(() => {
                    chai.expect(server.reqUrl).to.eql("/test/path?query=query_param");
                });
            });
            
            it("should make a request using the correct verb", () => {
                return executor.run().then(() => {
                    chai.expect(server.reqMethod).to.eql("POST");
                });
            });
            
            it("should make a request using the correct data", () => {
                return executor.run().then(() => {
                    return server.reqData.then(dataString => { 
                        chai.expect(dataString).to.not.eql("");
                        const data = JSON.parse(dataString);

                        chai.expect(data).to.eql({
                            value: "default_data_value"
                        });
                    });
                });
            });

            it("should use the correct headers", () => {
                return executor.run().then(() => {
                    chai.expect(server.reqHeaders).to.have.property("x-header", "header_value");
                });
            });
        });

        describe("buildRequestOptions", () => {
            let executor: TestHttpExecutor;
            const request: RequestDoc = {
                method: "POST",
                url: "http://localhost:1234/test/path?query={{query}}",
                headers: {
                    "X-Header": "{{header}}"
                },
                data: {
                    value: "{{data}}"
                }
            };
            
            beforeEach(() => executor = new TestHttpExecutor(app.db, app.testAction, app.testTask));

            it("should pass the correct protocol to the request", () => {
                let options = executor.testBuildRequestOptions(request);
                chai.expect(options).to.have.property("protocol", "http:");
            });

            it("should pass the correct hostname to the request", () => {
                let options = executor.testBuildRequestOptions(request);
                chai.expect(options).to.have.property("hostname", "localhost");
            });

            it("should pass the correct port to the request", () => {
                let options = executor.testBuildRequestOptions(request);
                chai.expect(options).to.have.property("port", 1234);
            });

            it("should pass the correct path to the request", () => {
                let options = executor.testBuildRequestOptions(request);
                chai.expect(options).to.have.property("path", "/test/path?query=query_param");
            });

            it("should pass the correct headers to the request", () => {
                let options = executor.testBuildRequestOptions(request);
                chai.expect(options).to.have.property("headers").eql({
                    "X-Header": "header_value" 
                });
            });
        });
    });
});