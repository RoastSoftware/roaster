import m from 'mithril';
import base from './base';

export default {
  view(vnode: CVnode) {
    return [
      m(base,
          m('p', `his is information about this wierd roaster thingie. \
Here you can analyze all of your code, much wow.`)
      ),
    ];
  },
} as m.Component;
