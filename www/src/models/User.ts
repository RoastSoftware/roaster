import * as m from "mithril";

export class User {
    static list: String[];
    
    static loadList() {
        return m.request({
			method: "GET",
			url: "https://rem-rest-api.herokuapp.com/api/users",
			withCredentials: true,
        }).then((result: m.RequestOptions<any>) => {
                User.list = result.data;
		});
    };
};
