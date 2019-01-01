import ξ from 'mithril';
import base from './base';
import editor from './editor';
import {EditorModel} from '../models/editor';
import Network from '../services/network';
import roasterLogo from '../assets/icons/roaster-icon-white.svg';

const containerStyle = `\
min-height: 100%;\
display: flex;\
flex-direction: row;\
`;

const columnStyle = `\
display: flex;\
`;

const roastColumnStyle = `\
border-left: 1px solid #094959;\
flex: 1 0 35%;\
flex-flow: column;\
` + columnStyle;

const editorColumnStyle = `\
flex: 1 0 65%;\
overflow: hidden;\
`;

const messagesRowStyle = `\
flex: 2;\
padding: 8px;\
`;

const messagesListStyle = `\
display: flex;
flex-direction: column;
`;

const controlsRowStyle = `\
min-height: 4em;\
padding: 8px;\
justify-content: flex-end;\
`;

const dropdownStyle = `\
display: block;\
float: right;\
border: none;\
-webkit-appearance: none;\
-moz-appearance: none;\
height: 14px;\
appearance: none;\
`;

const roasterIconStyle = `\
height: 1.5em;\
margin: -0.6em 0.5em -0.3em 0em;\
`;

interface Snippet {
  code: string
  language: string
};

interface RoastMessage {
  hash: string
  row: number
  column: number
  engine: string
  name: string
  description: string
};

interface RoastError {
  RoastMessage
};

interface RoastWarning {
  RoastMessage
}

interface RoastResult {
  username: string
  code: string
  score: number
  language: string
  errors: Array<RoastError>
  warnings: Array<RoastWarning>
  createTime: string
};

class RoastMessageList implements ξ.ClassComponent {
  view(vnode: ξ.CVnode) {
    const errors = vnode.attrs.errors;
    const warnings = vnode.attrs.warnings;

    return [
      (errors && errors.length > 0) ?
      errors.map((msg) => {
        return ξ('.item',
            ξ('i.big.frown.icon[style=color: #dc322f;]'),
            ξ('.content',
                ξ('.header[style=color: #dc322f;]',
                    msg.name),
                ξ('.description',
                    `${msg.description} [${msg.row}:${msg.column}]`),
            ),
        );
      }) : '',
      (warnings && warnings.length > 0) ?
      warnings.map((msg) => {
        return ξ('.item',
            ξ('i.big.meh.icon[style=color: #b58900;]'),
            ξ('.content',
                ξ('.header[style=color: #b58900;]',
                    msg.name),
                ξ('.description',
                    `${msg.description} [${msg.row}:${msg.column}]`),
            ),
        );
      }) : '',
    ];
  }
}


export default class Home implements ξ.ClassComponent {
  roast: RoastResult = {} as RoastResult;
  analyzing: boolean = false;
  error: Error;

  roastMe() {
    if (!this.analyzing) {
      this.reset(); // Clean up errors.
      this.analyzing = true;

      Network.request<RoastResult>('POST', '/roast', {
        'code': EditorModel.getCode(),
        'language': EditorModel.getLanguage(),
      }).then((roast: RoastResult) => {
        this.roast = roast;
        this.analyzing = false;
        ξ.redraw();
      }).catch((error: Error) => {
        this.error = error;
        this.analyzing = false;
        ξ.redraw();
      });
    }
  };

  reset() {
    this.roast = {} as RoastResult;
    this.error = undefined;
    this.analyzing = false;
  };

  view(vnode: ξ.CVnode): ξ.Children {
    return ξ(base,
        ξ('section#roast-container', {style: containerStyle},

            ξ('#editor-column', {style: editorColumnStyle},
                ξ(editor),
            ),

            ξ('#roast-column', {style: roastColumnStyle},
                ξ('#messages-row', {style: messagesRowStyle},
                    ξ('.ui.vertical.basic.segment.clearing', {
                      style: 'padding: 0; margin: 0;',
                    },
                    ξ('h3.ui.header.teal.left.floated[style=margin: 0;]',
                        ξ('i.bug.icon'),
                        ξ('.content',
                            'Code Result',
                            ξ('.sub.header',
                                'Warnings and errors found in your code.'
                            )
                        )
                    ),
                    this.roast.score ?
                    ξ('.ui.small.right.floated.statistic', {
                      style: 'margin-bottom: -2px;',
                    },
                    ξ('.value',
                        ξ('i.trophy.icon[style=color: gold;]'),
                        ` ${this.roast.score}`,
                    ),
                    ): ''),
                    ξ('.ui.divider'),
                    ξ('.ui.relaxed.list.divided', {style: messagesListStyle}, [
                      ξ(RoastMessageList, {
                        errors: this.roast.errors,
                        warnings: this.roast.warnings,
                      }),
                    ]),
                ),

                ξ('#controls-row', {style: controlsRowStyle},
                    this.error ?
                    ξ('.ui.error.message',
                        ξ('.header',
                            'We had some problems with your submission!',
                        ),
                        ξ('p', `${this.error.message}.`),
                    ): '',
                    ξ('.ui.buttons.large',
                        ξ('button.ui.primary.button', {
                          class: this.analyzing ? 'loading disabled' : '',
                          onclick: () => {
                            this.roastMe();
                          },
                        },
                        ξ('img.ui.image.spaced', {
                          src: roasterLogo,
                          style: roasterIconStyle,
                        }),
                        'ROAST ME!'),
                        ξ('.or'),
                        ξ('button.ui.button', {
                          onclick: () => {
                            this.reset();
                          },
                        },
                        'Reset'),
                    ),
                    ξ('select.ui.compact.selection.dropdown', {
                      style: dropdownStyle,
                    }, [
                      ξ('option', {value: 'python3'}, 'Python 3'),
                      // Python 2.7 is not implemented.
                      // ξ('option', {value: 'python2'}, 'Python 2.7'),
                    ]),
                ),

            ),
        ),
    );
  }
};
