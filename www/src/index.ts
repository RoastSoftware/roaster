import ξ from 'mithril';
import home from './views/home';
import about from './views/about';
import register from './views/register';
import profile from './views/profile';
import statistics from './views/statistics';
import login from './views/login';

import Auth from './services/auth';
import {User} from './models/user';

function redirectMatcher(
    view: ξ.ClassComponent,
    policy: () => boolean,
    redirect: string): () => ξ.ClassComponent {
  return () => {
    if (!policy) ξ.route.set(redirect);
    else return view;
  };
};

class Roaster {
  private body: HTMLBodyElement;

  constructor(private body: HTMLBodyElement) {
    Auth.resume<string>().then((user) => {
      console.log(user);
      // TODO
      User.setUsername(user.username);
      User.setFullname(user.fullname);
      User.setEmail(user.email);
    });
  };

  start() {
    ξ.route(document.body, '/', {
      '/': home,
      '/about': about,
      '/statistics': statistics,
      '/profile': {onmatch: redirectMatcher(profile, () => {
        return User.isLoggedIn();
      }, '/')},
      '/register': {onmatch: redirectMatcher(register, () => {
        return !User.isLoggedIn();
      }, '/')},
      '/login': {onmatch: redirectMatcher(login, () => {
        return !User.isLoggedIn();
      }, '/')},
    });
  };
}

const roaster = new Roaster(document.body as HTMLBodyElement);
roaster.start();
