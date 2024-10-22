export interface Semester {
  org_id: number;
  classroom_id: number;
  org_name: string;
  classroom_name: string;
  active: boolean;
}

//TODO: refactor this to come in as a map of {orgid: Semester[]}
export interface UserSemestersResponse {
  active_semesters: Semester[];
  inactive_semesters: Semester[];
}

export interface OrgSemestersResponse {
  semesters: Semester[];
}
