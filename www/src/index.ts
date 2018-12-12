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
    policy: (args: any[]) => boolean,
    redirect: string): () => ξ.ClassComponent {
  return (...args: any[]) => {
    if (!policy(args)) ξ.route.set(redirect);
    else return view;
  };
};

class Roaster {
  private body: HTMLBodyElement;

  constructor(private body: HTMLBodyElement) {};

  start() {
    Auth.resume<User>().then((user) => {
      if (user) {
        Object.assign(UserModel, user);
        UserModel.setLoggedIn(true);
      } else {
        UserModel.setLoggedIn(false);
      }
    }).then(() => {
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
        '/login': {
          onmatch: redirectMatcher(login, (): boolean => {
            return !UserModel.isLoggedIn();
          }, '/')},
        '/user/:username': {
          onmatch: redirectMatcher(user, (args: any[]): boolean => {
            return args[0].username != UserModel.getUsername();
          }, '/profile'),
        },
      });
    });
  };
}

const roaster = new Roaster(document.body as HTMLBodyElement);
roaster.start();
