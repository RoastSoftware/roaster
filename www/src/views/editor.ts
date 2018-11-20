import m, { ClassComponent, CVnode, CVnodeDOM} from "mithril";
import * as monaco from 'monaco-editor';
import 'monaco-editor/esm/vs/basic-languages/python/python.contribution';
// @ts-ignore
import solarizedMonacoTheme from 'monaco-themes/themes/Solarized-dark.json';

declare global {
  interface Window { MonacoEnvironment: any }
}

export default class Editor implements ClassComponent {
    oncreate(vnode: CVnodeDOM) {
        monaco.editor.defineTheme('solarized', solarizedMonacoTheme as monaco.editor.IStandaloneThemeData);
        monaco.editor.setTheme('solarized');

        let editor = monaco.editor.create(vnode.dom as HTMLElement, {language: 'python'});
        window.addEventListener("resize", () => editor.layout());
    };
    
    view(vnode: CVnode) {
        return m("#editor[style=min-height: 50rem;]");
    }
};

self.MonacoEnvironment = {
    getWorkerUrl(moduleId: number, label: string) {
        switch(label) {
            case 'json': return './dist/json.worker.js'
            case 'css': return './dist/css.worker.js'
            case 'html': return './dist/html.worker.js'
            case 'javascript':
            case 'typescript': return './dist/ts.worker.js'

            default: return './dist/editor.worker.js'
        }
    }
}
