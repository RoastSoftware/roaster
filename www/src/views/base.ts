import m, { ClassComponent, CVnode} from "mithril";
import nav from "./nav";
import header from "./header";

export default class Base implements ClassComponent {
    view(vnode: CVnode) {
        return [
            m(nav),
            m(".ui.container",
                m(header)
            ),
            m(".ui.main.container.segment.inverted", vnode.children)
        ];
    }
};
