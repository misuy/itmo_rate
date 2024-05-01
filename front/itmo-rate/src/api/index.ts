import { headers } from "@/utils";
import type { SubjectPreview } from "./subjects";
import type { TeacherPreview } from "./teachers";

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

interface SearchResult {
  subjects: SubjectPreview[],
  teachers: TeacherPreview[]
}

async function fetchSearch(text: string) 
  : Promise<ApiResult<SearchResult>> {
    const url = `http://localhost:8888/api/search?type=both&text=${text}`;
    const response = await fetch(url, {
      method: 'GET',
      mode: "cors",
      headers: headers,
    }) 
    if (response.ok) {
      const s: SearchResult = await response.json(); 
      console.log("ok: ", s);
      return ApiResult.ok(s, response.status);
    } else {
      console.error("error " + await response.text());
      return ApiResult.error("Cant get search result!", response.status);
    }
}


export {ApiResult, fetchSearch}
export type {Criterion}