import m, { ClassComponent, CVnode} from "mithril";
import base from "./base";
import editor from "./editor";

export default class Home implements ClassComponent {
    view(vnode: CVnode) {
        return m(base,
            m("p", "this is a fabulous placeholder"),
            m(editor)
        );
    }
};
