import m, { ClassComponent, CVnode} from "mithril";

export default class Header implements ClassComponent {
    view(vnode: CVnode) {
        return m('h1.ui.header', 
            m("a", {href: "/", oncreate: m.route.link},
                "ROASTER INC."
            )
        );
    }
};
