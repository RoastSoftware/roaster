import ξ from 'mithril';
import base from './base';

export default {
  view(vnode: CVnode) {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 2em;]',
            ξ('p', 'Let there be GRAPHS! '),
            ξ('p', 'later...')
        )
    );
  },
} as ξ.Component;
