class ApiResult<T> {
  ok: boolean;
  payload: T | null;
  errorMessage: string;
  code: number;

  constructor(code: number, payload?: T, errorMessage = "") {
    this.errorMessage = errorMessage;
    this.ok = code >= 200 && code < 300;
    this.payload = payload ? payload : null;
    this.code = code;
  }

  public static error = (message: string, code: number = 400) : ApiResult<any> => new ApiResult(code, null, message);
  public static ok = <T>(payload: T, code: number = 200) => new ApiResult<T>(code, payload);
}

interface Criterion {
  name: string;
  rating: number;
}


export {ApiResult}
export type {Criterion}