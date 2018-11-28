import ξ from 'mithril';
import home from './views/home';
import about from './views/about';
import register from './views/register';
import profile from './views/profile';
import statistics from './views/statistics';
import login from './views/login';

const xCSRFTokenHeader: string = 'X-Csrf-Token';

class Network {
  private static csrfToken: string = '';

  private static extractCSRFToken(xhr, xhrOptions): string {
    return xhr.getResponseHeader(xCSRFTokenHeader);
  };

  private static async initCSRFToken(): Promise<string> {
    return ξ.request<string>({
      method: 'HEAD',
      url: '/',
      withCredentials: true,
      extract: Network.extractCSRFToken,
    }).then((result: string) => {
      if (result == '') {
        throw new Error('unable to fetch CSRF token');
      }

      Network.csrfToken = result;
    });
  };

  public static async request<T>(method: string, url: string): Promise<T> {
    if (Network.csrfToken == '') {
      await Network.initCSRFToken();
    }

    return ξ.request<T>({
      method: method,
      url: url,
      withCredentials: true,
      headers: {xCSRFTokenHeader: Network.csrfToken},
    });
  };
}

class Auth {
  private user: User;

  async authenticate(): Promise<boolean> {
    // TODO
    return Network.request<boolean>('GET', '/');
  }

  getUser(): User {
    return this.user;
  }
}

class Roaster {
  private body: HTMLBodyElement;
  private auth: Auth;

  constructor(private body: HTMLBodyElement) {
    this.auth = new Auth();
    this.auth.authenticate().then((loggedIn) => {
      if (loggedIn) {
        ξ.redraw();
      }
    });
  };

  start() {
    ξ.route(document.body, '/', {
      '/': home,
      '/about': about,
      '/register': register,
      '/profile': profile,
      '/statistics': statistics,
      '/login': login,
    });
  };
}

const roaster = new Roaster(document.body as HTMLBodyElement);
roaster.start();
