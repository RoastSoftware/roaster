import m, { ClassComponent, CVnode} from "mithril";
import base from "./base";

export default class About implements ClassComponent {
    view(vnode: CVnode) {
        return [
            m(base, 
                m("p", "this is information about this wierd roaster thingie. Here you can analyze all of your code, much wow.")
            )
        ];
    }
};
