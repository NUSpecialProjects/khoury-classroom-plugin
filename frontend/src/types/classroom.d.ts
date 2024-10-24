interface IClassroom {
  id: number;
  name: string;
  url: string;
}

interface IClassroomResponse {
  available_classrooms: IClassroom[];
  unavailable_classrooms: IClassroom[];
}
