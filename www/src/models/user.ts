import m from "mithril";

export interface User {
    email: string;
    userName: string;
    password: string;
    fullName: string;
}

const UserModel = {
    current: {} as User,
    load(userName: string) {
        return m.request<User>({
            method: "GET",
            url: "/user/" + userName,

        })
            .then(result => {
                UserModel.current = result;
            })
    },

    save() {
        return m.request ({
            method: "PATCH",
            url: "/user/" + userName,
            data: UserModel.current,
        })
    },
    // TODO: make shure request data is sent as JSON, applies for all commands.
    create() {
        return m.request ({
            method: "POST",
            url: "/user/",
            data: UserModel.current,
        })
    },

    authenticateSession() {
        return m.request ({
            method: "POST",
            url: "/session/",
            data: UserModel.current,
        })
    },

    removeSession() {
        return m.request ({
            method: "DELETE",
            url: "/session/",
            data: UserModel.current,
        })
    }
}

type UserModel = typeof UserModel;

export default UserModel;
