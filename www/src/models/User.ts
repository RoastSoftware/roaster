import m, { ClassComponent } from "mithril";

export default class User {
  static list: string[] = [];

  static loadList() {
    return m
      .request({
        method: "GET",
        url: "https://rem-rest-api.herokuapp.com/api/users",
        withCredentials: true
      })
      .then((result: m.RequestOptions<any>) => {
        User.list = result.data;
      });
  }
}
