interface ISemester {
  org_id: number;
  classroom_id: number;
  org_name: string;
  classroom_name: string;
  active: boolean;
}

interface IUserSemestersResponse {
  active_semesters: ISemester[];
  inactive_semesters: ISemester[];
}

interface IOrgSemestersResponse {
  semesters: ISemester[];
}
