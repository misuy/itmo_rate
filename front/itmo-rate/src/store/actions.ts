import type { State } from ".";
import { fetchTeachersList } from "@/api/teachers"
import { fetchSubjectsList } from "@/api/subjects";

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
  async getSubjects ({ commit, state } : Params, { offset, amount } : any) {
    const result = await fetchSubjectsList(offset, amount)
    if (result.ok) {
      commit("gotSubjects", result.payload!)
    } else {
      // error!
    }
  }
}