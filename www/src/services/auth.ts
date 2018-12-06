import Network from './network';

export default class Auth {
  public static async login<T>(username: string, password: string): Promise<T> {
    return Network.request<T>('POST', '/session', {
      username: username,
      password: password,
    });
  };

  public static async resume<T>(): Promise<T> {
    return Network.request<T>('GET', '/session');
  };

  public static async logout() {
    return Network.request('DELETE', '/session');
  };

  public static async register<T>(user: T, password: string): Promise<T> {
    return Network.request<T>('POST', '/user', user); // TODO: Handle password.
  };
}
