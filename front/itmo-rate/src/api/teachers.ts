import { sleep } from "@/utils";
import { ApiResult } from ".";
import type { Criterion } from ".";

interface Teacher {
  id: number;
  name: string;
  avatar: string;
	criteria: Criterion[];
	avgRating: number;
	subjects: string[];
}

interface TeacherPreview {
  id: number;
  name: string;
  tags: string[];
  score: number;
}

async function fetchTeacher(id: number) : Promise<ApiResult<Teacher>>  {
  await sleep(500);
  return ApiResult.ok({
    id: id,
    name: "Клименков Сергей Викторович",
    avatar: "https://photo.itmo.su/avatar/8864a7e8b3d285daa5e12eec5ab6a82c782b2804/cover/320/320/",
    criteria: [
      {
        name: "Критерий 1",
        rating: 5.7
      },
      {
        name: "Критерий 2",
        rating: 5.7
      },
      {
        name: "Критерий 3",
        rating: 8.2
      },
      {
        name: "Критерий 4",
        rating: 1.4
      },
      {
        name: "Критерий 5",
        rating: 4.4
      },
    ],
    avgRating: 5.7,
    subjects: ["Основые проектной деятельности"]
  })
}

async function fetchTeachersList(offset: number, amount: number) 
  : Promise<ApiResult<TeacherPreview[]>> {
  await sleep(1000);
  return ApiResult.ok([
    {
      id: 1,
      name: "Соснов Николай Федорович",
      tags: ["Методы криптографии", "ТПО", "Компьютерные сети"],
      score: 4.4
    },
    {
      id: 1,
      name: "Соснов Семен Федорович",
      tags: ["Методы криптографии", "ТПО", "Компьютерные сети"],
      score: 7.2
    }
  ])
}

export {fetchTeachersList, fetchTeacher}
export type {Teacher, TeacherPreview}