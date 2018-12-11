import ξ from 'mithril';

import {EditorModel} from '../models/editor';

declare global {
  interface Window {
    MonacoEnvironment: any;
  }
}

self.MonacoEnvironment = {
  getWorkerUrl(moduleId: number, label: string) {
    switch (label) {
      case 'json':
        return './dist/json.worker.js';
      case 'css':
        return './dist/css.worker.js';
      case 'html':
        return './dist/html.worker.js';
      case 'javascript':
      case 'typescript':
        return './dist/ts.worker.js';

      default:
        return './dist/editor.worker.js';
    }
  },
};

/**
 * Lazy-load the Monaco Editor assets.
 * @param {string} themeName - The theme name to use (see monaco-themes/themes/)
 */
async function loadMonacoEditor(
    themeName: string
): Promise<monaco.editor.IStandaloneEditor> {
  const [monaco, themeData] = await Promise.all([
    import('monaco-editor'), // @ts-ignore
    import('monaco-themes/themes/' + themeName + '.json'),
  ]);

  monaco.editor.defineTheme(
      themeName,
    themeData as monaco.editor.IStandaloneThemeData
  );

  monaco.editor.setTheme(themeName);

  return monaco;
}

const fillAreaStyle = `\
height: 100%;\
width: 100%;\
padding: 0;\
margin: 0;\
border: 0;\
box-shadown: none;\
border-radius: 0;\
`;

/**
 * Editor component wraps a Monaco Editor.
 */
export default class Editor implements ClassComponent {
  editor: monaco.editor.IStandaloneEditor;

  /**
   * Loads and adds a Monaco Editor to the empty div#editor created by view.
   * @param {CVnode} vnode - Virtual node.
   */
  oncreate(vnode: ξ.CVnodeDOM) {
    loadMonacoEditor('Solarized-dark').then((monaco) => {
      EditorModel.setCode(`\
"""
Roaster roasts your code with static code analysis, for free!
"""
def welcome(ξ):
    print('Please write your fabulous code here!')


def too_complex():
    """
    Example of a too complex function!
    """
    def b(z, x, c, v, b, n, p, o, i, u, y, t):
        if z == x:
            pass
        if x == c:
            pass
        if c == v:
            pass
        if b == n:
            pass
        if n == p:
            pass
        if p == o:
            pass
        if o == i:
            pass
        if i == u:
            pass
        if u == y:
            pass
        if y == t:
            pass
        if z == t:
            pass

    b(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
`);

      EditorModel.createModel();
      EditorModel.setLanguage('python3');

      this.editor = monaco.editor.create(vnode.dom as HTMLElement, {
        value: EditorModel.getCode(),
        minimap: {
          enabled: false,
        },
        language: 'python',
        scrollBeyondLastLine: false, // Display scrollbar only on overflow.
        // TODO: For some reason the language isn't set correctly.
        //       In the mean time the language is hardcoded above.
        //        - Probably because the local `monaco` parameter should be used
        //          instead of the global `monaco` that is used in the model (?)
        // model: EditorModel.getModel(),
      });

      // Update the EditorModel with the new code when the user writes in the
      // editor.
      this.editor.onDidChangeModelContent((_: Event) => {
        EditorModel.setCode(this.editor.getValue());
      });

      // Update the size of the editor on resize events.
      window.addEventListener('resize', () => this.editor.layout());

      EditorModel.setReady();
      ξ.redraw();
    });
  };

  /**
   * Creates a empty div#editor that is used by oncreate.
   * @param {CVnode} vnode - Virtual node.
   * @return {CVnode}
   */
  view(vnode: ξ.CVnode) {
    return EditorModel.ready ?
          ξ('#editor', {style: fillAreaStyle})
          :
          ξ('.ui.segment', {style: fillAreaStyle},
              ξ('.ui.large.loader', {style: 'display: block;'})
          );
  };
};
