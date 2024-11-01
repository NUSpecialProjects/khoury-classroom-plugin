export function getClassroomsInOrg(orgId: number): Promise<IClassroomListResponse> {
  console.log("Using mocked API call for org: ", orgId);
  return Promise.resolve({
    classrooms: [
      {
        id: 1,
        name: "Spring 2024",
        org_id: orgId,
        org_name: "Organization One",
      },
      {
        id: 2,
        name: "Fall 2023",
        org_id: orgId,
        org_name: "Organization One",
      },
    ],
  });
}


  export function postClassroom(classroom: Omit<IClassroom, "id">): Promise<IClassroom> {
    console.log("Using mocked API call for creating classroom: ", classroom);
    return Promise.resolve({
      id: 5,
      name: classroom.name,
      org_id: classroom.org_id,
      org_name: classroom.org_name,
    });
  }
