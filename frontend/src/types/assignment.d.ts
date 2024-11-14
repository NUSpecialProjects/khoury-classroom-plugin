interface IAssignmentOutline {
  id: number;
  template_id: number;
  created_at: Date;
  released_at: Date | null;
  name: string;
  classroom_id: number;
  group_assignment: boolean;
  main_due_date: Date | null;
}

interface IAssignmentOutlineResponse {
  assignment_outline: IAssignmentOutline;
}

interface IAssignmentTemplate {
  template_repo_id: number;
  template_repo_owner: string;
  template_repo_name: string;
}