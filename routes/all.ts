import {RouteBase} from "./base";
import {Actions} from "./actions";
import {Projects} from "./projects";
import {Tasks} from "./tasks";

export class AllRoutes extends RouteBase {
    static routeTypes: typeof RouteBase[] = [
        Actions,
        Projects,
        Tasks
    ];

    routes: RouteBase[] = AllRoutes.routeTypes.map(R => new R(this.server, this.db));

    register() {
        this.routes.forEach(r => r.register());
    }
}