import {RouteBase} from "./base";

export class Users extends RouteBase {
    register() {
        this.server.get("/api/v1/users", this.authorize(), this.permission("admin/users"), (req, res) => {
            this.db.Users.find().toArray().then(users => {
                res.send(200, users);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.post("/api/v1/users", this.authorize(), this.permission("admin/users"), (req, res) => {
            return this.db.AuditLog.insert({
                type: "user.create",
                context: {
                    user: this.isAuthorizedRequest(req) ? req.user.summary : null,
                    request: req.body
                }
            })
            .then(() => this.db.Users.insert(req.body))
            .then(user => {
                res.send(200, user);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.get("/api/v1/user", this.authorize(), (req, res) => {
            Promise.resolve().then(() => {
                if (this.isAuthorizedRequest(req)) {
                    return res.send(200, req.user);
                }

                return this.unauthorized();
            }).then(err => this.catch(res, err));
        });

        this.server.get("/api/v1/user/:user", this.authorize(), this.permission("admin/users"), (req, res) => {
            this.db.Users.get(req.params.user).then(user => {
                if (!user) return this.notFound();

                return this.db.AuditLog.insert({
                    type: "user.tokens.view",
                    context: {
                        user: this.isAuthorizedRequest(req) ? req.user.summary : null,
                        request: {
                            user: req.params.user
                        }
                    }
                }).then(() => {
                    res.send(200, user);
                    return user;
                });
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.get("/api/v1/user/:user/tokens", this.authorize(), this.permission("admin/users"), (req, res) => {
            this.db.Users.get(req.params.user).then(user => {
                if (!user) return this.notFound();

                res.send(200, user.tokens);
            }, err => this.databaseError(err)).catch(err => this.catch(res, err));
        });

        this.server.post("/api/v1/user/:user/tokens", this.authorize(), this.permission("admin/users"), (req, res) => {
            this.db.Users.get(req.params.user).then(user => {
                if (!user) return this.notFound();

                const { token } = <{ token: string; }>req.body;

                user.tokens.push(token);

                return this.db.AuditLog.insert({
                    type: "user.tokens.create",
                    context: {
                        user: this.isAuthorizedRequest(req) ? req.user.summary : null,
                        request: {
                            user: req.params.user
                        }
                    }
                }).then(() => user.save()).then(user => {
                    res.send(200, user.tokens);
                    return user;
                }, err => this.databaseError(err));
            }).catch(err => this.catch(res, err));
        });
    }
}