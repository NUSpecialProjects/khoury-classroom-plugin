export interface Classroom {
    id: number;
    name: string;
    url: string;
}

export interface ClassroomResponse {
    available_classrooms: Classroom[];
    unavailable_classrooms: Classroom[];
}

