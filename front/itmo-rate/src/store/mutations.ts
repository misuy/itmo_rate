import type { TeacherPreview, Teacher } from "@/api/teachers";
import type { State } from ".";
import type { SubjectPreview } from "@/api/subjects";
import type { TeacherReview } from "@/api/reviews";

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
  }
}