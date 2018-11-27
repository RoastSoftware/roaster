import m, {ClassComponent, CVnode} from 'mithril';
import nav from './nav';
import header from './header';

export default class Base implements ClassComponent {
  view(vnode: CVnode) {
    return [
      m(nav),
      m('.ui.main.inverted[style=height: calc(100% - 51.5px); margin: 0]', vnode.children),
    ];
  }
};
