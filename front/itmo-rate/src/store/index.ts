import { createStore } from "vuex"
import mutations from "./mutations"
import actions from "./actions"
import type { TeacherPreview } from "@/api/teachers"
import type { SubjectPreview } from "@/api/subjects"

interface State {
  teachers: TeacherPreview[];
  subjects: SubjectPreview[];
}

export default createStore({
  state () : State {
    return {
      teachers: [],
      subjects: []
    }
  },
  mutations: mutations,
  actions: actions
})

export type { State }