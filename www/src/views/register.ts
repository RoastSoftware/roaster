import m from 'mithril';
import base from './base';
import user from '../models/user';

/**
 *  Onsubmit function sends form data
 */
function onsubmit() {
  console.log('On submit function called, but not implemented');
  user.create();
}

/**
 * Validate checks field input data
 */
function validate() {
    console.log('Un-implemented validator');
    // TODO: implement local model in this view, make databind and on submit only works if all fields has passed validator
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

interface inputField {
    value: string;
    valid(): boolean;
    error: string;
    inputType: string;
}

class userInputView implements m.ClassComponent {
    inputField: inputField;
    fieldName: string;

    constructor(fieldName: string, inputField: inputField) {
        this.inputField = inputField;
        this.fieldName = fieldName;
    }

    view() {
        return [
            m('input', {
                type: this.inputField.inputType,
                className: this.inputField.valid() ? 'error' : '',
                oninput: m.withAttr('value', this.inputField.valid)
            }),
            m('p.errorMessage', this.inputField.error)
        ];
    };
}

class passwordInputField implements inputField {
    inputType = 'password';

    value = '';

    valid(password: string): boolean {
        if (this.value.length > 500) {
            this.error = `do you really need more than 1.3*10^(1204)\
                possible entropy? We're blocking this length of \
                passwords due to DoS attacks.`;
            return false;
        }

        if (this.value.length < 8) {
            this.error ='expected a password of atleast 8 characters';
            return false;
        }

        return true;
    }

    error: string;
}

    /*
class validatedUserInput implements m.ClassComponent<userInput> {
    inputType: string;
    fieldName: string;

    constructor(fieldName: string, inputType: string, field: Field) {
        this.inputType = inputType;
    }

    view() {
        return [
            m('input', {
                type: this.inputType,
                className: this.error ? 'error' : '',
                value: this.value(),
                oninput: m.withAttr('value', this.value)
            }),
            m('p.errorMessage', this.error)
        ];
    };
}
     */


export interface State {
    password: string;
    username: string;
    email: string;
}

export default {
    password: "",
    username: "",
    email: "",

  view(vnode: m.CVnode) {
    return m(base,
        m('.ui.grid',
            m('.ui.container.six.wide.column.centered',
                m('.ui.segments',
                    m('.ui.segment', m('h2', 'REGISTER')),
                    m('.ui.segment', m('form.ui.form', {onsubmit}, [
                        m('input', {
                            type: "password",
                            onkeypress: m.withAttr('value', e => this.password = e),
                        }),
                        /*
                      input('Full name', 'text'),
                      input('Username', 'text'),
                      input('E-mail', 'text'),
                      input('Password', 'password'),
                         */
                      m('button.ui.teal.basic.button', 'GET ROASTED!'),
                    ]))
                )
            )
        )
    );
  },
} as m.Component<State>;
