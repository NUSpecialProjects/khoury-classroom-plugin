import { StepComponentProps } from "./Main";

export interface AssignmentFormData {
    assignmentName: string
    classroomId: number
    groupAssignment: boolean
    mainDueDate: Date | null
    templateRepo: IRepository | null
}

export interface StarterCodeDetailsProps extends StepComponentProps<AssignmentFormData> {
    repositories: IRepository[];
    isLoading: boolean;
  }