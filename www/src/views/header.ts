import m, { ClassComponent, CVnode} from "mithril";

export default class Header implements ClassComponent {
    view(vnode: CVnode) {
        return m('.header', 
            m('h1', 
                m("a", {href: "/", oncreate: m.route.link}, "ROASTER")),
        );
    }
};
