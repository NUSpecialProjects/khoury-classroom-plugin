interface ISemester {
  id: number;
  classroom_id: number;
  org_id: number;
  name: string;
  active: boolean;
}

interface ISemestersResponse {
  active_semesters: Semester[];
  inactive_semesters: Semester[];
}
