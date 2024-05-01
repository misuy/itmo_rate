import { headers, sleep } from "@/utils";
import { ApiResult } from ".";
import type { Criterion } from ".";

interface Teacher {
  id: number;
  name: string;
  avatar: string;
	criteria: Criterion[];
	avg_rating: number;
	subjects: string[];
}

interface TeacherPreview {
  id: number;
  name: string;
  tags: string[];
  score: number;
}

async function fetchTeacher(id: number) : Promise<ApiResult<Teacher>>  {
  const url = `http://localhost:8888/api/teacher/${id}`;
  const response = await fetch(url, {
    method: 'GET',
    mode: "cors",
    headers: headers,
  }) 
  if (response.ok) {
    const s: Teacher = (await response.json()).teacher;
    s.avatar = "https://photo.itmo.su/avatar/8864a7e8b3d285daa5e12eec5ab6a82c782b2804/cover/320/320/"
    console.log(s);
    return ApiResult.ok(s, response.status);
  } else {
    console.error("error " + await response.text());
    return ApiResult.error("Cant get subjects list!", response.status);
  }
}

async function fetchTeachersList(offset: number, amount: number) 
  : Promise<ApiResult<TeacherPreview[]>> {
    const url = `http://localhost:8888/api/teachers?amount=${amount}&offset=${offset}`;
    const response = await fetch(url, {
      method: 'GET',
      mode: "cors",
      headers: headers,
    }) 
    if (response.ok) {
      const s: TeacherPreview[] = (await response.json()).teachers; 
      console.log(s);
      return ApiResult.ok(s, response.status);
    } else {
      console.error("error " + await response.text());
      return ApiResult.error("Cant get subjects list!", response.status);
    }
}

export {fetchTeachersList, fetchTeacher}
export type {Teacher, TeacherPreview}