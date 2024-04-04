import { sleep } from "@/utils";
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
  await sleep(1000);
  const reviews: TeacherReview[] = [...Array(amount).keys()].map(x => ({
    id: x,
    rating: 1.2,
    created: "10.03.2023",
    text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat laboris, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris...",
    author: "Анонимно",
    subject: "Предмет " + x
  }));
  return ApiResult.ok(reviews)
}

export type {TeacherReview, SubjectReview, FullReview};
export {fetchTeacherReviews}