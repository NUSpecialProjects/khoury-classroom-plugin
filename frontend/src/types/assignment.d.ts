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

interface IStudentAssignment {
    id: number;
    local_id: number;
    rubric_id: number;
    assignment_id: number;
    repo_name: string;
    student_gh_username: string;
    ta_gh_username: string;
    complted: boolean;
    started: boolean;

}