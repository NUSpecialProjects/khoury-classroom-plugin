interface IFullRubric {
    rubric: IRubric;
    rubric_items: [IRubricItem];
}

interface IRubric {
    id: number;
    name: string;
    org_id: number;
    classroom_id: number;
    reusable: boolean;
    created_at: Date;
}

interface IRubricItem {
    id: number;
    rubric_id: number | null;
    point_value: number;
    explanation: string;
    created_at: Date;
}

interface IFullRubricResponse {
    full_rubric: IFullRubric;
}