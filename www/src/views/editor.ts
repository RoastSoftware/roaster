import ξ from 'mithril';

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

/**
 * Editor component wraps a Monaco Editor.
 */
export default class Editor implements ClassComponent {
  ready: boolean = false;
  minimap: boolean = false;
  editor: monaco.editor.IStandaloneEditor;
  language: string = 'python';

  /**
   * Loads and adds a Monaco Editor to the empty div#editor created by view.
   * @param {CVnode} vnode - Virtual node.
   */
  oncreate(vnode: CVnodeDOM) {
    loadMonacoEditor('Solarized-dark').then((monaco) => {
      this.ready = true;
      ξ.redraw();

      this.editor = monaco.editor.create(vnode.dom as HTMLElement, {
        value: `\
"""
Roaster roasts your code with static code analysis, for free!
"""
def welcome(ξ):
    print('Please write your fabulous code here!')`,
        language: this.language,
        minimap: {
          enabled: this.minimap,
        },
      });

      window.addEventListener('resize', () => this.editor.layout());
    });
  }

  /**
   * Creates a empty div#editor that is used by oncreate.
   * @param {CVnode} vnode - Virtual node.
   * @return {CVnode}
   */
  view(vnode: CVnode) {
    return [
      (this.ready ?
          ξ('#editor[style=height: 100%;]')
        :
          ξ('.ui.active[style=height: 100%;]',
              ξ('.ui.large.loader[style=display: block;]')
          )
      ),
    ];
  }
}
