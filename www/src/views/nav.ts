import m, { ClassComponent, CVnode} from "mithril";

export default class Nav implements ClassComponent {
    view(vnode: CVnode) {
        return m('nav.ui.massive.borderless.menu', [
            m(".ui.container",
                m("a.header.item", {href: "/", oncreate: m.route.link},
                    m("i.coffee.icon.logo"),
                    "Roaster"
                ),
                m("a.item", {href: "/about", oncreate: m.route.link}, "About"),
                m("a.item", {href: "/register", oncreate: m.route.link}, "REGISTER"),
                m("a.item", {href: "/profile", oncreate: m.route.link}, "Profile"),
                m("a.item", {href: "/statistics", oncreate: m.route.link}, "Statistics")
            )
        ]);
    }
};
