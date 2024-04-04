import type { TeacherPreview } from "@/api/teachers";
import type { State } from ".";
import type { SubjectPreview } from "@/api/subjects";

export default {
  gotTeachers(state: State, teachers: TeacherPreview[]) {
    state.teachers = teachers;
  },
  gotSubjects(state: State, subjects: SubjectPreview[]) {
    state.subjects = subjects;
  }
}