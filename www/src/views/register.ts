import m, { ClassComponent, CVnode} from "mithril";
import nav from "./nav";
import header from "./header";

function onsubmit() {
    // TODO
}

function input(fieldName: string) {
    return m("label", 
        m("span", fieldName + ":"),
        m("br"),
        m("input", {
            placeholder: fieldName,
            // TODO: inplement on input
        }),
        m("br")
    );
}
                        
export default class Register implements ClassComponent {
    view(vnode: CVnode) {
        return [
            m(nav),
            m(header),
            m("h2", "REGISTER / LOGIN"),
            m("form",  { onsubmit }, [
                input("Full name"),
                input("Username"),
                input("E-mail"),
                input("Password"),
                m("button", "GET ROASTED!")
            ])
        ];
    }
};
