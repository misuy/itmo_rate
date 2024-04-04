import type { State } from ".";
import { fetchTeachersList, fetchTeacher} from "@/api/teachers"
import { fetchSubjectsList } from "@/api/subjects";
import { emptyTeacher } from "@/utils";
import { fetchTeacherReviews } from "@/api/reviews";

interface Params {
  commit: any;
  state: State
}

export default {
  async getTeachers ({ commit, state } : Params, { offset, amount } : any) {
    const result = await fetchTeachersList(offset, amount)
    if (result.ok) {
      commit("gotTeachers", result.payload!)
    } else {
      // error!
    }
  },
  async getTeacher ({ commit, state } : Params, id: number) {
    commit("gotTeacher", emptyTeacher());
    const result = await fetchTeacher(id);
    if (result.ok) {
      commit("gotTeacher", result.payload!)
    } else {
      // error!
    }
  },
  async getTeacherReviews ({ commit, state } : Params, {id, offset, amount} : any) {
    commit("gotTeacherReviews", []);
    const result = await fetchTeacherReviews(id, offset, amount);
    if (result.ok) {
      commit("gotTeacherReviews", result.payload!)
    } else {
      // error!
    }
  },
  async getSubjects ({ commit, state } : Params, { offset, amount } : any) {
    const result = await fetchSubjectsList(offset, amount)
    if (result.ok) {
      commit("gotSubjects", result.payload!)
    } else {
      // error!
    }
  }
}