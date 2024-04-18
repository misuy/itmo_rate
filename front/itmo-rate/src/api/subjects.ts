import { sleep } from "@/utils";
import { ApiResult } from ".";
import type { Criterion } from ".";

interface Subject {
  id: number;
  name: string;
	criteria: Criterion[];
	avgRating: number;
	teachers: string[];
  lecturers: string[];
  faculties: string[]
}

interface SubjectPreview {
  id: number;
  name: string;
  tags: string[];
  score: number;
}

async function fetchSubject(id: number) : Promise<ApiResult<Subject>> {
  await sleep(1000);
  return ApiResult.ok({
    id: id,
    name: "Основы проектной деятельности",
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
    faculties: ["Вт", "ПИиКТ"],
    lecturers: ["Клименков Сергей Викторович", "Соснов Николай Федорович"],
    teachers:  ["Тимофеев Тихон Александрович", "Колпакова Екатерина Александровна", "Рябов Лука Макарович"],
    avgRating: 10
  });
}

async function fetchSubjectsList(offset: number, amount: number) 
  : Promise<ApiResult<SubjectPreview[]>> {
  await sleep(1000);
  return ApiResult.ok([
    {
      id: 1,
      name: "Основы проектной деятельности",
      tags: ["ПИиКТ", "Вт"],
      score: 4.4
    },
    {
      id: 2,
      name: "ТПО",
      tags: ["Вт", "ПИиКТ"],
      score: 7.2
    }
  ])
}

export { fetchSubjectsList, fetchSubject }
export type { Subject, SubjectPreview }