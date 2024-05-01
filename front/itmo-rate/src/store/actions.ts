import { SearchState, type State } from ".";
import { fetchTeachersList, fetchTeacher} from "@/api/teachers"
import { fetchSubject, fetchSubjectsList } from "@/api/subjects";
import { emptySubject, emptyTeacher } from "@/utils";
import { fetchTeacherReviews, fetchSubjectReviews } from "@/api/reviews";
import { fetchSearch, type ApiResult } from "@/api";

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
  },
  async search ({ commit } : Params, text: string) {
    commit("setSearchState", SearchState.FETCHING);
    const result = await fetchSearch(text);
    if (result.ok) {
      if (result.payload?.subjects.length == 0 && result.payload?.teachers.length == 0) {
        commit("setSearchState", SearchState.NOTHING_FOUND);
      } else {
        commit("setSearchState", SearchState.FOUND);
        commit("gotSubjects", result.payload?.subjects);
        commit("gotTeachers", result.payload?.teachers);
      }
    } else {
      commit("setSearchState", SearchState.NOTHING_FOUND);
    }
  }
}