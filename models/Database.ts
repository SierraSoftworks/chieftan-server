import * as Iridium from "iridium";
import {Release, ReleaseDoc} from "./Release";
import {Target, TargetDoc} from "./Target";

export class Database extends Iridium.Core {
    Releases = new Iridium.Model<ReleaseDoc, Release>(this, Release);
    Target = new Iridium.Model<TargetDoc, Target>(this, Target);
}