import m from 'mithril';
import home from './views/home';
import about from './views/about';
import register from './views/register';
import profile from './views/profile';
import statistics from './views/statistics';
import login from './views/login';

m.route(document.body, '/', {
  '/': home,
  '/about': about,
  '/register': register,
  '/profile': profile,
  '/statistics': statistics,
  '/login': login,
});
