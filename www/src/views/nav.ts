import m, { ClassComponent, CVnode} from "mithril";

export default class Nav implements ClassComponent {
    view(vnode: CVnode) {
        return m('.nav', [
            m("a", {href: "/", oncreate: m.route.link}, "Roaster"),
            m("span", " | "),
            m("a", {href: "/about", oncreate: m.route.link}, "About"),
            m("span", " | " ),
            m("a", {href: "/register", oncreate: m.route.link}, "REGISTER"),
            m("span", " | "),
            m("a", {href: "/profile", oncreate: m.route.link}, "Profile"),
            m("span", " | "),
            m("a", {href: "/statistics", oncreate: m.route.link}, "Statistics")
        ]);
    }
};
