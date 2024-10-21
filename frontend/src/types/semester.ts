export interface Semester {
    id: number;
    classroom_id: number;
    org_id: number;
    name: string;
    active: boolean;
}

//TODO: refactor this to come in as a map of {orgid: Semester[]}
export interface UserSemestersResponse {
    active_semesters: Semester[];
    inactive_semesters: Semester[];
}

export interface OrgSemestersResponse {
    semesters: Semester[];
}