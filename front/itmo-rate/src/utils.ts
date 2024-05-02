import type { Review } from "./api/reviews";
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
    avg_rating: 0,
    subjects: []
})

const emptySubject = () : Subject => ({
  id: 0,
  name: "-",
  criteria: [],
  faculties: [],
  lecturers: [],
  teachers:  [],
  avg_rating: 0,
  reviews_count: 0
})

const emptyReview = () : Review => ({
  "id": -1,
  "lecturers": [
    ""
  ],
  "teachers": [
    ""
  ],
  "subject": "",
  "text": "",
  "author": "",
  "created": "",
  "criteria": [
    {
      "name": "Критерий 1",
      "rating": 0
    },
    {
      "name": "Критерий 2",
      "rating": 0
    },
    {
      "name": "Критерий 3",
      "rating": 0
    },
    {
      "name": "Критерий 4",
      "rating": 0
    },
    {
      "name": "Критерий 5",
      "rating": 0
    }
  ]
})

const headers = {
  "Content-Type": "application/json",
  'Access-Control-Allow-Credentials' : "true",
  'Access-Control-Allow-Origin':'*',
  'Access-Control-Allow-Methods':'POST, OPTIONS, GET, PUT',
  'Access-Control-Allow-Headers':'application/json',
};

export {sleep, emptyTeacher, emptySubject, emptyReview, headers}