import m, { ClassComponent, CVnode} from "mithril";
import nav from "./nav";
import header from "./header";

export default class About implements ClassComponent {
    view(vnode: CVnode) {
        return [
            m(nav),
            m(header),
            m("p", "this is information about this wierd roaster thingie. Here you can analyze all of your code, much wow.")
        ];
    }
};
