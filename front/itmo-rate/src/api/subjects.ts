import { headers, sleep } from "@/utils";
import { ApiResult } from ".";
import type { Criterion } from ".";

interface Subject {
  id: number;
  name: string;
	criteria: Criterion[];
	avg_rating: number;
	teachers: string[];
  lecturers: string[];
  faculties: string[];
  reviews_count: number;
}

interface SubjectPreview {
  id: number;
  name: string;
  tags: string[];
  score: number;
}

async function fetchSubject(id: number) : Promise<ApiResult<Subject>> {
  const url = `http://localhost:8888/api/subject/${id}`;
  const response = await fetch(url, {
    method: 'GET',
    mode: "cors",
    headers: headers,
  }) 
  if (response.ok) {
    const s: Subject = (await response.json()).subject;
    console.log(s);
    return ApiResult.ok(s, response.status);
  } else {
    console.error("error " + await response.text());
    return ApiResult.error("Cant get subjects list!", response.status);
  }
}

async function fetchSubjectsList(offset: number, amount: number) 
  : Promise<ApiResult<SubjectPreview[]>> {
    const url = `http://localhost:8888/api/subjects?amount=${amount}&offset=${offset}`;
    const response = await fetch(url, {
      method: 'GET',
      mode: "cors",
      headers: headers,
    }) 
    if (response.ok) {
      const s: SubjectPreview[] = (await response.json()).subjects; 
      console.log(s);
      return ApiResult.ok(s, response.status);
    } else {
      console.error("error " + await response.text());
      return ApiResult.error("Cant get subjects list!", response.status);
    }
}

export { fetchSubjectsList, fetchSubject }
export type { Subject, SubjectPreview }