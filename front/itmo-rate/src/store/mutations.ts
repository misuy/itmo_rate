import type { TeacherPreview, Teacher } from "@/api/teachers";
import type { ModalState, SearchState, State } from ".";
import type { Subject, SubjectPreview } from "@/api/subjects";
import type { Review, SubjectReview, TeacherReview } from "@/api/reviews";

export default {
  gotTeachers(state: State, teachers: TeacherPreview[]) {
    state.teachers = teachers;
  },
  gotTeacher(state: State, teacher: Teacher) {
    state.currentTeacher = teacher;
  },
  gotTeacherReviews(state: State, reviews: TeacherReview[]) {
    state.teacherReviews = reviews;
  },
  gotSubjects(state: State, subjects: SubjectPreview[]) {
    state.subjects = subjects;
  },
  gotSubjectReviews(state: State, reviews: SubjectReview[]) {
    state.subjectReviews = reviews;
  },
  gotSubject(state: State, subject: Subject) {
    state.currentSubject = subject;
  },
  gotReview(state: State, review: Review | null) {
    state.openReview = review;
  },
  setError(state: State, code: number) {
    state.error = code;
  },
  setSearchState(state: State, searchState: SearchState) {
    state.searchState = searchState;
  },
  setModalState(state: State, modalState: ModalState) {
    state.modalWindowState = modalState;
  }
}