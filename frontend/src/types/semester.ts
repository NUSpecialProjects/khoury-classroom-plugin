export interface Semester {
    id: number;
    classroom_id: number;
    org_id: number;
    name: string;
    active: boolean;
}

export interface SemestersResponse {
    active_semesters: Semester[];
    inactive_semesters: Semester[];
}