interface IClassroom {
  id: number;
  name: string;
  org_id: number;
  org_name: string;
}

interface IUserOrgsAndClassroomsResponse {
  orgs_and_classrooms: Map<IOrganization, IClassroom[]>;
}

interface IClassroomListResponse {
  classrooms: IClassroom[];
}