import { StepComponentProps } from "./main";

export interface AssignmentFormData {
    name: string;
    description: string;
    dueDate: Date | null;
    isGroupAssignment: boolean;
    selectedRepoId: number;
}

export interface StarterCodeDetailsProps extends StepComponentProps<AssignmentFormData> {
    repositories: IRepository[];
    isLoading: boolean;
  }