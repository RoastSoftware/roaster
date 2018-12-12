import ξ from 'mithril';

const xCsrfToken: string = 'X-Csrf-Token';
const contentType: string = 'Content-Type';

export default class Network {
  private static nextCSRFToken: string = '';

  private static extractCSRFToken(xhr, xhrOptions): string {
    const token: string = xhr.getResponseHeader(xCsrfToken);

    if (token == '') {
      throw new Error('empty CSRF token received');
    }

    Network.nextCSRFToken = token;

    if (xhr.responseText.length > 0
      && xhr.getResponseHeader(contentType) == 'application/json') {
      return JSON.parse(xhr.responseText);
    }

    return '';
  };

  private static async initCSRFToken() {
    return ξ.request({
      method: 'HEAD',
      url: '/',
      extract: Network.extractCSRFToken,
    });
  };

  public static async request<T>(
      method: string,
      url: string,
      data?: any): Promise<T> {
    if (Network.nextCSRFToken == '') {
      await Network.initCSRFToken();
    }

    return ξ.request<T>({
      method: method,
      url: url,
      data: data,
      headers: {[xCsrfToken]: Network.nextCSRFToken},
      extract: Network.extractCSRFToken,
    }).then((result: T) => {
      return result;
    });
  };
}
