class ApiResult<T> {
  ok: boolean;
  payload: T | null;
  errorMessage: string;

  constructor(success: boolean = true, payload?: T, errorMessage = "") {
    this.errorMessage = errorMessage;
    this.ok = success;
    this.payload = payload ? payload : null;
  }

  public static error = (message: string) : ApiResult<any> => new ApiResult(false, null, message);
  public static ok = <T>(payload: T) => new ApiResult<T>(true, payload);
}

function sleep(time: number) {
  return new Promise((resolve) => setTimeout(resolve, time));
}

export {sleep, ApiResult}