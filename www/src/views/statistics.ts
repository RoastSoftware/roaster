import m, { ClassComponent, CVnode} from "mithril";
import nav from "./nav";
import header from "./header";

export default class Statistics implements ClassComponent {
    view(vnode: CVnode) {
        return [
            m(nav),
            m(header),
            m('.statistics', [
                m("p", "Let there be GRAPHS! "),
                m("p", "later...")
            ])
        ];
    }
};
