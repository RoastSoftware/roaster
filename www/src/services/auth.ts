import Network from './network';

export default class Auth {
  public static async login<T>(
      username: string,
      password: string): Promise<T> {
    return Network.request<T>('POST', '/session', {
      username: username,
      password: password,
    });
  };

  public static async resume<T>(): Promise<T> {
    return Network.request('GET', '/session'); // TODO
  };

  public static async logout(): Promise {
    return Network.request('DELETE', '/session');
  };

  public static async register<T>(user: T, password: string): Promise<T> {
    return Network.request<T>('POST', '/user', user);
  };
}
