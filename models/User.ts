import * as Iridium from "iridium";
import * as crypto from "crypto";

export interface UserSummaryDoc {
    id: string;
    name: string;
    email: string;
}

export const UserSummarySchema = {
    id: String,
    name: String,
    email: String
};

export interface UserDoc {
    _id?: string;

    name: string;
    email: string;

    permissions?: string[];

    tokens?: string[];
}

@Iridium.Collection("users")
@Iridium.Index({ tokens: 1 }, { unique: true })
export class User extends Iridium.Instance<UserDoc, User> implements UserDoc {
    @Iridium.Property(/^[0-9a-f]{32}$/)
    _id: string;

    @Iridium.Property(String)
    name: string;

    @Iridium.Property(/^.+@.*\.[a-z]+$/)
    email: string;

    @Iridium.Property([String])
    permissions: string[];

    @Iridium.Property([String])
    tokens: string[];

    get summary(): UserSummaryDoc {
        return {
            id: this._id,
            name: this.name,
            email: this.email
        };
    }

    can(permission: string, context: { [name: string]: string; } = {}) {
        if (~this.permissions.indexOf(permission)) return true;
        if (~this.permissions.indexOf(permission.replace(/:(\w+)/g, (match, name) => {
            return context[name] || match;
        }))) return true;
        return false;
    }

    static onCreating(doc: UserDoc) {
        if (!doc._id) {
            const hash = crypto.createHash("md5");
            hash.update((doc.email || "").trim().toLowerCase())
            doc._id = hash.digest("hex");
        }

        doc.permissions = doc.permissions || [];
        doc.tokens = doc.tokens || [];
    }

    toJSON() {
        return {
            id: this._id,
            name: this.name,
            email: this.email,
            permissions: this.permissions
        };
    }
}