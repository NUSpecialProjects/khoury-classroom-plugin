interface IAssignment {
    id: number; 
    rubric_id: number | null; 
    active: boolean;
    assignment_classroom_id: number;
    semester_id: number;
    name: string;
    local_id: number;
    main_due_date: Date | null;
  }