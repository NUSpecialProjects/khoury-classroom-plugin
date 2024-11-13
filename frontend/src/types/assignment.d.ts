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