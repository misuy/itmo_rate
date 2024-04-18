import type { Subject } from "./api/subjects";
import type { Teacher } from "./api/teachers";

function sleep(time: number) {
  return new Promise((resolve) => setTimeout(resolve, time));
}

const emptyTeacher = () : Teacher => ({
    id: 0,
    name: "-",
    avatar: "",
    criteria: [],
    avgRating: 0,
    subjects: []
})

const emptySubject = () : Subject => ({
  id: 0,
  name: "-",
  criteria: [],
  faculties: [],
  lecturers: [],
  teachers:  [],
  avgRating: 0
})


export {sleep, emptyTeacher, emptySubject}