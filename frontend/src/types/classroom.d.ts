interface IClassroom {
  id: number;
  name: string;
  url: string;
}

interface IClassroomResponse {
  available_classrooms: Classroom[];
  unavailable_classrooms: Classroom[];
}
