interface IAssignmentOutline {
  id: number;
  template_id: number;
  created_at: Date;
  released_at: Date | null;
  name: string;
  classroom_id: number;
  rubric_id: number | null;
  group_assignment: boolean;
  main_due_data: Date | null;
}

interface IAssignmentOutlineResponse {
  assignment_outline: IAssignmentOutline
}