import ξ from 'mithril';
import home from './views/home';
import about from './views/about';
import register from './views/register';
import profile from './views/profile';
import statistics from './views/statistics';
import login from './views/login';
import user from './views/user';

import Auth from './services/auth';
import {UserModel} from './models/user';

function redirectMatcher(
    view: ξ.ClassComponent,
    policy: () => boolean,
    redirect: string): () => ξ.ClassComponent {
  return () => {
    if (!policy()) ξ.route.set(redirect);
    else return view;
  };
};

class Roaster {
  private body: HTMLBodyElement;

  constructor(private body: HTMLBodyElement) {
    Auth.resume<User>().then((user) => {
      if (user) {
        Object.assign(UserModel, user);
        UserModel.setLoggedIn(true);
      } else {
        UserModel.setLoggedIn(false);
      }
    });
  };

  start() {
    ξ.route(document.body, '/', {
      '/': home,
      '/about': about,
      '/statistics': statistics,
      '/profile': {onmatch: redirectMatcher(profile, (): boolean => {
        return UserModel.isLoggedIn();
      }, '/')},
      '/register': {onmatch: redirectMatcher(register, (): boolean => {
        return !UserModel.isLoggedIn();
      }, '/')},
      '/login': {onmatch: redirectMatcher(login, (): boolean => {
        return !UserModel.isLoggedIn();
      }, '/')},
        '/user/:username': user,
    });
  };
}

const roaster = new Roaster(document.body as HTMLBodyElement);
roaster.start();
