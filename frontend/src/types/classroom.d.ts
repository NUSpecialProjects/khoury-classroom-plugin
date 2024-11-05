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

interface IClassroomToken {
  classroom_id: number;
  classroom_role: string;
  token: string;
  expires_at: string | null;
  created_at: string;
}
