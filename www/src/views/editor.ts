import m, { ClassComponent, CVnode, CVnodeDOM} from "mithril";
import * as monaco from 'monaco-editor';
import * as SolarizedTheme from 'monaco-themes/themes/Solarized-dark.json';
// TODO: import languages for syntaxhighlighting.

declare global {
  interface Window { MonacoEnvironment: any }
}

export default class Editor implements ClassComponent {
    oncreate(vnode: CVnodeDOM) {
        monaco.editor.defineTheme('solarized', SolarizedTheme);
        monaco.editor.setTheme('solarized');

        let editor = monaco.editor.create(vnode.dom as HTMLElement);
        window.addEventListener("resize", () => editor.layout());
    };
    
    view(vnode: CVnode) {
        return m("#editor[style=min-height: 50rem;]");
    }
};

self.MonacoEnvironment = {
  getWorkerUrl(moduleId: number, label: string) {
    switch(label) {
      case 'json': return './json.worker.bundle.js'
      case 'css': return './css.worker.bundle.js'
      case 'html': return './html.worker.bundle.js'

      case 'javascript':
      case 'typescript': return './ts.worker.bundle.js'

      default: return './editor.worker.bundle.js'
    }
  }
}
