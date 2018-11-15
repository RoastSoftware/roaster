import m, { ClassComponent } from "mithril";

export interface User {
  id: number;
  firstName: string;
  lastName: string;
}

export default class Users {
  static list: User[] = [];

  static loadList() {
    return m
      .request({
        method: "GET",
        url: "https://rem-rest-api.herokuapp.com/api/users",
        withCredentials: true
      })
      .then((result: m.RequestOptions<User[]>) => {
        Users.list = result.data;
      });
  }
}
