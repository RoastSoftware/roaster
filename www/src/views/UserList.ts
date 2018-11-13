import * as m from "mithril";
import { User } from "../model/User";

module.exports = {
	oninit: User.loadList,
	view: function() {
		return m(".user-list", User.list.map(function(user) {
			return m(".user-list-item", user.firstName + " " + user.lastName);
		}));
	}
};
