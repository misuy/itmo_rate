import { ApiResult, sleep } from "@/utils";

interface Criterion {
  name: string;
  rating: number;
}

interface Teacher {
  id: number;
  avatar: string;
	creteria: Criterion[];
	avg_rating: number;
	subjects: string[];
}

interface TeacherPreview {
  id: number;
  name: string;
  tags: string[];
  score: number;
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

export {fetchTeachersList}
export type {Teacher, TeacherPreview}