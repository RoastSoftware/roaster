import ξ from 'mithril';
import nav from './nav';

const fillMainAreaStyle = `\
display: flex;\
flex-flow: column;\
height: 100%;\
`;

const mainContentContainer = `\
margin: 0;\
flex: 2;\
`;

// Instead of only using a { display: none; }, the element is positioned
// off-screen, this allows screen readers to still see the content.
// For more information, see:
// https://developer.paciellogroup.com/blog/2012/05/html5-accessibility-chops-hidden-and-aria-hidden/
const hideMonacoAriaContainer = `\
clip-path: inset(100%);\
clip: rect(1px 1px 1px 1px); /* IE 6/7 */\
clip: rect(1px, 1px, 1px, 1px);\
height: 1px;\
overflow: hidden;\
position: absolute;\
white-space: nowrap; /* Added line */\
width: 1px;\
`;

export default class Base implements ξ.ClassComponent {
  view(vnode: ξ.CVnode) {
    return [
      ξ('div', {style: fillMainAreaStyle},
          ξ(nav),
          ξ('.ui.main.inverted', {style: mainContentContainer}, vnode.children),
      ),
      ξ('style', '.monaco-aria-container {' + hideMonacoAriaContainer + '}'),
    ];
  }
};
