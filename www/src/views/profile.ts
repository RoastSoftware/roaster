import m, { ClassComponent, CVnode} from "mithril";
import nav from "./nav";
import header from "./header";

export default class Profile implements ClassComponent {
    view(vnode: CVnode) {
        return [
            m(nav),
            m(header),
            m("img", {src: "http://c0419384.cdn2.cloudfiles.rackspacecloud.com/adjnbivz-27337_l-avatar-main.jpg"}, "Avatar"),
            m("p", "Full Name"),
            m("p", "E-mail"),
            m("p", "User Name")
        ];
    }
};
