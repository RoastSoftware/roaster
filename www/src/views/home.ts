import m, { ClassComponent, CVnode} from "mithril";
import base from "./base";
import editor from "./editor";

export default class Home implements ClassComponent {
    view(vnode: CVnode) {
        return m(base,
            m("p", "Please write your fabulous code below:"),
            m(".ui.two.column.grid", [
                m(".ui.column", 
                    m(editor)),
                m(".ui.column", 
                    m("p", "Really roasting error msg"))
            ])
        );
    }
};