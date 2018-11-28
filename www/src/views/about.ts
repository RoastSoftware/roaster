import ξ from 'mithril';
import base from './base';

export default {
  view(vnode: CVnode) {
    return [
      ξ(base,
          ξ('.ui.main.text.container[style=margin-top: 2em;]',
              ξ('p', `his is information about this wierd roaster thingie. \
Here you can analyze all of your code, much wow.`)
          ),
      ),
    ];
  },
} as ξ.Component;
