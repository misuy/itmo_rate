import type { State } from ".";
import { fetchTeachersList, fetchTeacher} from "@/api/teachers"
import { fetchSubject, fetchSubjectsList } from "@/api/subjects";
import { emptySubject, emptyTeacher } from "@/utils";
import { fetchTeacherReviews, fetchSubjectReviews } from "@/api/reviews";
import type { ApiResult } from "@/api";

interface Params {
  commit: any;
  state: State
}

function commitResult(commit: any, key: string, result: ApiResult<any>) {
  if (result.ok) {
    commit(key, result.payload!);
  } else {
    commit("setError", result.code);
  }
}

export default {
  async getTeachers ({ commit } : Params, { offset, amount } : any) {
    const result = await fetchTeachersList(offset, amount);
    commitResult(commit, "gotTeachers", result);
  },
  async getTeacher ({ commit } : Params, id: number) {
    commit("gotTeacher", emptyTeacher());
    const result = await fetchTeacher(id);
    commitResult(commit, "gotTeacher", result);
  },
  async getTeacherReviews ({ commit } : Params, {id, offset, amount} : any) {
    commit("gotTeacherReviews", []);
    const result = await fetchTeacherReviews(id, offset, amount);
    commitResult(commit, "gotTeacherReviews", result);
  },
  async getSubjects ({ commit } : Params, { offset, amount } : any) {
    const result = await fetchSubjectsList(offset, amount);
    commitResult(commit, "gotSubjects", result);
  },
  async getSubjectReviews ({ commit } : Params, {id, offset, amount} : any) {
    commit("getSubjectReviews", []);
    const result = await fetchSubjectReviews(id, offset, amount);
    commitResult(commit, "gotSubjectReviews", result);
  },
  async getSubject ({ commit } : Params, id: number) {
    commit("gotSubject", emptySubject());
    const result = await fetchSubject(id);
    commitResult(commit, "gotSubject", result);
  }
}