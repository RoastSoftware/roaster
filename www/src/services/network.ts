import ξ from 'mithril';

const xCsrfToken: string = 'X-Csrf-Token';
const contentType: string = 'Content-Type';

// APIError is used as an helper to create a Error with a code property.
//
// This class should _not_ be used outside of this service. TypeScript has no
// standard way of narrowing errors with the `cause` clause, i.e. no type-safety
// inside the `cause` clause.
//
// Solution: Just check if the `code` property exists using standard JavaScript,
// like:
//    if ('code' in error && error.code == 404) {
//      doSomething(error.code);
//    }
class APIError extends Error {
  code: number

  constructor(message: string, code?: number) {
    super(message);
    this.code = code || 0;
  }
};

export function encodeURL(...url: string): string {
  let u = '';

  if (url.length > 1) {
    u = '/' + url.join('/');
  } else {
    u = url[0];
  }

  return u;
};

export default class Network {
  private static nextCSRFToken: string = '';

  private static extract(xhr, xhrOptions): string {
    const token: string = xhr.getResponseHeader(xCsrfToken);

    if (token) {
      Network.nextCSRFToken = token;
    }

    let response: any;
    if (xhr.responseText.length > 0
      && xhr.getResponseHeader(contentType).includes('application/json')) {
      response = JSON.parse(xhr.responseText);
    } else {
      response = xhr.responseText; // Unparsable data, use raw data instead.
    }


    if (xhr.status < 300 || xhr.status == 304) {
      return response;
    }

    throw new APIError(response.message || response, xhr.status);
  };

  private static async initCSRFToken() {
    return ξ.request({
      method: 'HEAD',
      url: '/',
      extract: Network.extract,
    });
  };

  public static async request<T>(
      method: string,
      url: string | Array<string>,
      data?: any): Promise<T> {
    if (Network.nextCSRFToken == '') {
      await Network.initCSRFToken();
    }

    return ξ.request<T>({
      method: method,
      url: Array.isArray(url) ? encodeURL(...url) : encodeURL(url),
      data: data,
      headers: {[xCsrfToken]: Network.nextCSRFToken},
      extract: Network.extract,
    });
  };
}
