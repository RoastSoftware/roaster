import m, { ClassComponent, CVnode} from "mithril";
import base from "./base";

function onsubmit() {
    // TODO
}

function input(fieldName: string, type: string) {
    return m(".field",
        m("label", fieldName + ":"),
        m("input", {
            type: type,
            placeholder: fieldName,
            // TODO: inplement on input
        })
    );
}
                        
export default class Register implements ClassComponent {
    view(vnode: CVnode) {
        return m(base, 
            m(".ui.grid",
                m(".ui.container.six.wide.column.centered",
                    m(".ui.segments",
                        m(".ui.segment", m("h2", "REGISTER / LOGIN")),
                        m(".ui.segment", m("form.ui.form",  { onsubmit }, [
                            input("Full name", "text"),
                            input("Username", "text"),
                            input("E-mail", "text"),
                            input("Password", "password"),
                            m("button.ui.teal.basic.button", "GET ROASTED!")
                        ]))
                    )
                )
            )
        );
    }
};
