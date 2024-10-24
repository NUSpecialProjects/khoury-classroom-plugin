interface IAssignment {
  id: number;
  rubric_id: number | null;
  assignment_classroom_id: number;
  semester_id: number;
  name: string;
  inserted_date: Date | null;
  main_due_date: Date | null;
}
