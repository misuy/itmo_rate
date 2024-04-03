import { createStore } from "vuex"
import mutations from "./mutations"
import actions from "./actions"
import type { TeacherPreview } from "@/api/teachers"

interface State {
  teachers: TeacherPreview[]
}

export default createStore({
  state () : State {
    return {
      teachers: []
    }
  },
  mutations: mutations,
  actions: actions
})

export type { State }