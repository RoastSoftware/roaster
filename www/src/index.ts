import ξ from 'mithril';
import home from './views/home';
import about from './views/about';
import register from './views/register';
import profile from './views/profile';
import statistics from './views/statistics';
import login from './views/login';

const xCsrfToken: string = 'X-Csrf-Token';

class Network {
  private static nextCSRFToken: string = '';

  private static extractCSRFToken(xhr, xhrOptions): string {
    const token: string = xhr.getResponseHeader(xCsrfToken);

    if (token == '') {
      throw new Error('empty CSRF token received');
    }

    Network.nextCSRFToken = token;
  };

  private static async initCSRFToken(): Promise {
    return ξ.request({
      method: 'HEAD',
      url: '/',
      extract: Network.extractCSRFToken,
    });
  };

  public static async request<T>(method: string, url: string): Promise<T> {
    if (Network.nextCSRFToken == '') {
      await Network.initCSRFToken();
    }

    return ξ.request<T>({
      method: method,
      url: url,
      headers: {[xCsrfToken]: Network.nextCSRFToken},
      extract: Network.extractCSRFToken,
    });
  };
}

class Auth {
  private user: User;

  async authenticate(): Promise<boolean> {
    // TODO
    return Network.request<boolean>('POST', '/user');
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
