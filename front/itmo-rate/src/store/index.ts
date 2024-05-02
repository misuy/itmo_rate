import { createStore } from "vuex"
import mutations from "./mutations"
import actions from "./actions"
import type { Teacher, TeacherPreview } from "@/api/teachers"
import type { Review, SubjectReview, TeacherReview } from "@/api/reviews"
import type { Subject, SubjectPreview } from "@/api/subjects"
import { emptySubject, emptyTeacher } from "@/utils"

interface State {
  error: number;
  teachers: TeacherPreview[];
  subjects: SubjectPreview[];
  currentTeacher: Teacher;
  currentSubject: Subject;
  teacherReviews: TeacherReview[];
  subjectReviews: SubjectReview[];
  searchState: SearchState;
  openReview: Review | null;
  modalWindowState: ModalState;
}

enum SearchState {
  IDLE,
  NOTHING_FOUND,
  FETCHING,
  FOUND
}

enum ModalState {
  NOTHING,
  REVIEW_READ_MODE,
  REVIEW_WRITE_MODE,
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
      currentSubject: emptySubject(),
      searchState: SearchState.IDLE,
      modalWindowState: ModalState.NOTHING,
      openReview: null
    }
  },
  mutations: mutations,
  actions: actions
})

export type { State }
export { SearchState, ModalState }