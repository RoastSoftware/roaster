import ξ, {ClassComponent, CVnode} from 'mithril';
import nav from './nav';
import header from './header';

const fillMainAreaStyle = `\
display: flex;\
flex-flow: column;\
height: 100%;\
`;

const mainContentContainer = `\
margin: 0;\
flex: 2;\
`;

export default class Base implements ClassComponent {
  view(vnode: CVnode) {
    return [
      ξ('div', {style: fillMainAreaStyle},
          ξ(nav),
          ξ('.ui.main.inverted', {style: mainContentContainer}, vnode.children),
      ),
    ];
  }
};
