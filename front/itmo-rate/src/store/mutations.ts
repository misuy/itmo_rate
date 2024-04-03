import type { TeacherPreview } from "@/api/teachers";
import type { State } from ".";

export default {
  gotTeachers(state: State, teachers: TeacherPreview[]) {
    state.teachers = teachers;
  }
}