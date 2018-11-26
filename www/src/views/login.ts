
import m from 'mithril';
import base from './base';
import user from '../models/user';

/**
 *  Onsubmit function sends form data
 */
function onsubmit() {
  console.log('On submit function called, but not implemented');
  user.authenticateSession();
}
/**
 * Validate checks field input data
 */
function validate() {
  console.log('Un-implemented validator');
}
/**
 * input generates input form fields
 * @param {string} fieldName name of input field
 * @param {string} type HTML type of input field
 *
 * @return {undefined}
 */
function input(fieldName: string, type: string) {
  return m('.field',
      m('label', fieldName + ':'),
      m('input', {
        type: type,
        placeholder: fieldName,
        oninput: m.withAttr('value', validate),
        // TODO: inplement on input
      })
  );
}

export default {
  view(vnode: CVnode) {
    return m(base,
        m('.ui.grid',
            m('.ui.container.six.wide.column.centered',
                m('.ui.segments',
                    m('.ui.segment', m('h2', 'LOGIN')),
                    m('.ui.segment', m('form.ui.form', {onsubmit}, [
                      input('Username', 'text'),
                      input('Password', 'password'),
                      m('button.ui.teal.basic.button', 'LOGIN!'),
                    ]))
                )
            )
        )
    );
  },
} as m.Component;
