import { fetchTeachersList } from "@/api/teachers"
import type { State } from ".";

interface Params {
  commit: any;
  state: State
}

export default {
  async gotTeachers ({ commit, state } : Params, { offset, amount } : any) {
    const result = await fetchTeachersList(offset, amount)
    if (result.ok) {
      commit("gotTeachers", result.payload!)
    } else {
      // error!
    }
  }
}