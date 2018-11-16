import m, { ClassComponent, CVnode} from "mithril";
import nav from "./nav";
import header from "./header";

export default class Home implements ClassComponent {
    view(vnode: CVnode) {
        return [
            m(nav),
            m(header),
            m('.page', [
                m("p", "this is a fabulous placeholder")
            ])
        ];
    }
};
