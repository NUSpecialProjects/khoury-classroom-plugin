interface IStudentWork {
  student_work_id: number;
  classroom_id: number;
  assignment_name?: string;
  assignment_outline_id: number;
  repo_name: string;
  org_name: string;
  unique_due_date?: Date;
  submitted_pr_number?: number;
  manual_feedback_score?: number;
  auto_grader_score?: number;
  submission_timestamp?: Date;
  grades_published_timestamp?: Date;
  work_state: string;
  created_at: Date;
  contributors: string[];
}

interface IPaginatedStudentWork extends IStudentWork {
  row_num: number;
  total_student_works: number;
  previous_student_work_id: number | null;
  next_student_work_id: number | null;
}

interface IStudentWorkResponses {
  student_works: IStudentWork[];
}

interface IPaginatedStudentWorkResponse {
  student_work: IPaginatedStudentWork;
  feedback: IGraderFeedback[];
}


interface ICommitsPerDayResponse {
  dated_commits: Map<Date, number>;
}
