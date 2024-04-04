import { createStore } from "vuex"
import mutations from "./mutations"
import actions from "./actions"
import type { Teacher, TeacherPreview } from "@/api/teachers"
import type { TeacherReview } from "@/api/reviews"
import type { SubjectPreview } from "@/api/subjects"
import { emptyTeacher } from "@/utils"

interface State {
  error: number;
  teachers: TeacherPreview[];
  subjects: SubjectPreview[];
  currentTeacher: Teacher;
  // currentSubject: Teacher;
  teacherReviews: TeacherReview[]
}

export default createStore({
  state () : State {
    return {
      error: NaN,
      teachers: [],
      subjects: [],
      teacherReviews: [],
      currentTeacher: emptyTeacher(),
      // currentSubject: null;
    }
  },
  mutations: mutations,
  actions: actions
})

export type { State }