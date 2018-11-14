import * as m from "mithril";
import { User } from "../models/User";

export default {
    oninit() {
        return User.loadList();
    },

    view(vnode) {
        return m(".user-list", User.list.map((user: any) => {
			return m(".user-list-item", user.firstName + " " + user.lastName);
		}));
    }
} as m.Component<{}, {}>;
