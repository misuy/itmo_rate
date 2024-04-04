import type { Teacher } from "./api/teachers";

function sleep(time: number) {
  return new Promise((resolve) => setTimeout(resolve, time));
}

function emptyTeacher() : Teacher {
  return {
    id: 0,
    name: "-",
    avatar: "",
    criteria: [
      {
        name: "-",
        rating: 0
      },
      {
        name: "-",
        rating: 0
      },
      {
        name: "-",
        rating: 0
      },
      {
        name: "-",
        rating: 0
      },
      {
        name: "-",
        rating: 0
      },
    ],
    avgRating: 0,
    subjects: []
  };
}

export {sleep, emptyTeacher}