import { createStore } from "vuex"
import mutations from "./mutations"
import actions from "./actions"
import type { Teacher, TeacherPreview } from "@/api/teachers"
import type { SubjectReview, TeacherReview } from "@/api/reviews"
import type { Subject, SubjectPreview } from "@/api/subjects"
import { emptySubject, emptyTeacher } from "@/utils"

interface State {
  error: number;
  teachers: TeacherPreview[];
  subjects: SubjectPreview[];
  currentTeacher: Teacher;
  currentSubject: Subject;
  teacherReviews: TeacherReview[];
  subjectReviews: SubjectReview[]
}

export default createStore({
  state () : State {
    return {
      error: NaN,
      teachers: [],
      subjects: [],
      teacherReviews: [],
      subjectReviews: [],
      currentTeacher: emptyTeacher(),
      currentSubject: emptySubject()
    }
  },
  mutations: mutations,
  actions: actions
})

export type { State }