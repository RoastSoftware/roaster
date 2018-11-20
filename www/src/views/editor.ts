import m, { ClassComponent, CVnode, CVnodeDOM} from "mithril";
import * as monaco from 'monaco-editor';

declare global {
  interface Window { MonacoEnvironment: any }
}

export default class Editor implements ClassComponent {
    editor: monaco.editor.IStandaloneCodeEditor;

    oncreate(vnode: CVnodeDOM) {
        this.editor = monaco.editor.create(vnode.dom as HTMLElement);
    };
    
    onupdate(vnode: CVnodeDOM){
        // TODO: call editor.layout() here to get the new size of the editor when
        // the container changes size.
        this.editor.layout();
        console.log("LOLOLO");
    };

    view(vnode: CVnode) {
        return m("#editor");
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
