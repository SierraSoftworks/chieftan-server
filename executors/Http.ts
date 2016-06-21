import {ExecutorBase} from "./Executor";
import {RequestDoc} from "../models/Request";

import * as http from "http";
import * as https from "https";
import * as url from "url";

export class HttpExecutor extends ExecutorBase {
    static config = "http";

    run() {
        const action = this.action.http;

        return new Promise<void>((resolve, reject) => {
            
            const request = this.buildRequest(action, (res => {
                let responseData = [
                    `::[info] Response Received::`,
                    `HTTP/${res.httpVersion} ${res.statusCode} ${res.statusMessage}`,
                    this.renderHeaders(res.rawHeaders),
                    ""
                ];

                res.setEncoding("utf8");
                res.on("data", (chunk) => {
                    responseData += chunk;
                });

                res.on("end", () => {
                    this.task.output += responseData;

                    if(res.statusCode >= 400)
                        return reject(new Error(`::[error] Request failed with status ${res.statusCode} ${res.statusMessage}::`));
                    return resolve();
                });
            }));

            if(action.data) {
                const data = this.interpolate(action.data);
                if (typeof action.data === "string") request.write(data);
                else request.write(JSON.stringify(data));
            }

            request.on("error", (err) => {
                this.task.output += `\n::[error] ${err.message}::`;
                reject(err);
            });

            request.end();
        });
    }

    protected renderHeaders(headers: string[]) {
        let headerData: string[] = [];
        for (let i = 0; i < headers.length; i += 2) {
            headerData.push(`${headers[i]}: ${headers[i + 1]}`)
        }

        return headerData.join("\n");
    }

    protected buildRequest(request: RequestDoc, callback: (res: http.IncomingMessage) => void) {
        const options = this.buildRequestOptions(request);

        switch(options.protocol) {
            case "http:":
                return http.request(options, callback);
            case "https:":
                return https.request(options, callback);
            default:
                throw new Error(`The "${options.protocol}" protocol is not supported.`);
        }
    }

    protected buildRequestOptions(request: RequestDoc) {
        const parsedUrl = url.parse(this.interpolate(request.url));
        const headers = this.interpolate(request.headers || {});

        return {
            path: parsedUrl.path,
            hostname: parsedUrl.hostname,
            port: parseInt(parsedUrl.port) || 80,
            protocol: parsedUrl.protocol,
            method: request.method,
            headers: headers
        };
    }

    toString() {
        return "HTTP";
    }
}