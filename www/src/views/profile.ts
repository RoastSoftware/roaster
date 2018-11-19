import m, { ClassComponent, CVnode} from "mithril";
import base from "./base";

export default class Profile implements ClassComponent {
    view(vnode: CVnode) {
        return m(base,
            m("img.ui.image.rounded.medium", {src: "http://c0419384.cdn2.cloudfiles.rackspacecloud.com/adjnbivz-27337_l-avatar-main.jpg"}, "User profile picture."),
            m("h2",
                "Mr. Bean-A-Tar"),
            m("h3",
                m("i.user.icon"),
                "mr-bean-a-tar"),
            m("p", 
                m("i.mail.icon"),
                "mr@bean-a-tar.example.org")
        );
    }
};
