import { sleep } from "@/utils";
import { ApiResult } from ".";
import type { Criterion } from ".";

interface Subject {
  id: number;
	creteria: Criterion[];
	avg_rating: number;
	teachers: string[];
  lecturers: string[];
}

interface SubjectPreview {
  id: number;
  name: string;
  tags: string[];
  score: number;
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

export { fetchSubjectsList }
export type { Subject, SubjectPreview }