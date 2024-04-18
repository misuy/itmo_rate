import { headers, sleep } from "@/utils";
import { ApiResult, type Criterion } from ".";

interface TeacherReview {
  id: number;
  rating: number;
  subject: string;
  created: string;
  text: string;
  author: string;
}

interface SubjectReview {
  id: number;
  rating: number[];
  created: string;
  text: string;
  author: string;
}

interface FullReview {
  id: number;
  lecturer: string;
  teacher: string;
  rating: number[];
  created: string;
  text: string;
  criteria: Criterion[];
  author: string;
}

async function fetchTeacherReviews(id: number, offset: number, amount: number): Promise<ApiResult<TeacherReview[]>> {
  const url = `http://localhost:8888/api/teacher/${id}/reviews?amount=${amount}&offset=${offset}`;
  const response = await fetch(url, {
    method: 'GET',
    mode: "cors",
    headers: headers,
  }) 
  if (response.ok) {
    const s: TeacherReview[] = (await response.json()).reviews;
    console.log(s);
    return ApiResult.ok(s, response.status);
  } else {
    console.error("error " + await response.text());
    return ApiResult.error("Cant get subjects list!", response.status);
  }
}

async function fetchSubjectReviews(id: number, offset: number, amount: number): Promise<ApiResult<SubjectReview[]>> {
  const url = `http://localhost:8888/api/subject/${id}/reviews?amount=${amount}&offset=${offset}`;
  const response = await fetch(url, {
    method: 'GET',
    mode: "cors",
    headers: headers,
  }) 
  if (response.ok) {
    const s: SubjectReview[] = (await response.json()).reviews;
    s.push(s[0])
    console.log(s);
    return ApiResult.ok(s, response.status);
  } else {
    console.error("error " + await response.text());
    return ApiResult.error("Cant get subjects list!", response.status);
  }
}

export type {TeacherReview, SubjectReview, FullReview};
export {fetchTeacherReviews, fetchSubjectReviews}