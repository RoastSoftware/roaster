import m, { ClassComponent, CVnode} from "mithril";
import base from "./base";

export default class Statistics implements ClassComponent {
    view(vnode: CVnode) {
        return m(base,
            m("p", "Let there be GRAPHS! "),
            m("p", "later...")
        );
    }
};
