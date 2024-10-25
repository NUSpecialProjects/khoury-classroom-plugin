interface IStudentAssignment {
  id: number;
  assignment_id: number;
  repo_name: string;
  student_gh_username: [string];
  ta_gh_username: string | null;
  completed: boolean;
  started: boolean;
}
