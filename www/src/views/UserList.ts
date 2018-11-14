import m, { ClassComponent } from "mithril";
import User from "../models/User";

export default class UserList implements ClassComponent {
  oninit = User.loadList;

  view() {
    return m(
      ".user-list",
      User.list.map((user: any) => {
        return m(".user-list-item", user.firstName + " " + user.lastName);
      })
    );
  }
}
