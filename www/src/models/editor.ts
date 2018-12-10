export class EditorModel {
  static code: string = '';
  static ready: boolean = false;
  static language: string = 'python3';
  static editorLanguage: string = 'python';
  static monacoModel: monaco.editor.ITextModel;

  static setCode(code: string) {
    EditorModel.code = code;
  };

  static getCode(): string {
    return EditorModel.code;
  }

  static setEditorLanguage(language: string) {
    monaco.editor.setModelLanguage(
        monaco.editor.getModel(EditorModel.monacoModel.uri),
        language);
  }

  static createModel() {
    EditorModel.monacoModel = monaco.editor.createModel(
        EditorModel.getCode(),
        EditorModel.getEditorLanguage());
  }

  static getModel(): any {
    return EditorModel.monacoModel;
  }

  static setLanguage(language: string) {
    EditorModel.language = language;

    switch (language) {
      case 'python3', 'python2':
        EditorModel.setEditorLanguage('python');
        break;
      default:
        EditorModel.setEditorLanguage('plaintext');
    }
  }

  static getEditorLanguage(): string {
    return EditorModel.editorLanguage;
  }

  static getLanguage(): string {
    return EditorModel.language;
  }

  static setReady() {
    EditorModel.ready = true;
  }
}
