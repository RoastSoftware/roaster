import m from 'mithril';
import Home from './views/home';
import About from './views/about';
import Register from './views/register';
import Profile from './views/profile';
import Statistics from './views/statistics';

m.route(document.body, '/', {
  '/': Home,
  '/about': About,
  '/register': Register,
  '/profile': Profile,
  '/statistics': Statistics,
});
