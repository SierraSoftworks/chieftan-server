import {RouteBase} from "./base";
import {Actions} from "./actions";
import {Projects} from "./projects";
import {Tasks} from "./tasks";
import {Status} from "./status";

export class AllRoutes extends RouteBase {
    static routeTypes: typeof RouteBase[] = [
        Actions,
        Projects,
        Tasks,
        Status
    ];

    routes: RouteBase[] = AllRoutes.routeTypes.map(R => new R(this.app, this.server, this.db));

    register() {
        this.routes.forEach(r => r.register());
    }
}