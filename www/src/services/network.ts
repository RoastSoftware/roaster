import ξ from 'mithril';

const xCsrfToken: string = 'X-Csrf-Token';

export default class Network {
  private static nextCSRFToken: string = '';

  private static extractCSRFToken(xhr, xhrOptions): string {
    const token: string = xhr.getResponseHeader(xCsrfToken);

    if (token == '') {
      throw new Error('empty CSRF token received');
    }

    Network.nextCSRFToken = token;
  };

  private static async initCSRFToken(): Promise {
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
      extract: Network.extractCSRFToken, // TODO: Apparently JSON data isn't returned.
    }).then((result: T) => {
      console.log(result);
      return result;
    });
  };
}
