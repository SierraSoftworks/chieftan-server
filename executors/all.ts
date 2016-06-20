import {HttpExecutor} from "./Http";
import {Application} from "../Application";

export function RegisterExecutors(app: Application) {
    app.use(HttpExecutor);
}