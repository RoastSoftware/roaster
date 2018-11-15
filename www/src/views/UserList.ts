import m, { ClassComponent } from "mithril";
import Users, { User } from "../models/Users";

export default class UserList implements ClassComponent {
  oninit = Users.loadList;

  view() {
    return m(
      ".user-list",
      Users.list.map((user: User) => {
        return m(".user-list-item", user.firstName + " " + user.lastName);
      })
    );
  }
}
