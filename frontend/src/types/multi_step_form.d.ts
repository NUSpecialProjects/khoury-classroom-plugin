// ==============================
// MultiStepForm types
// ==============================

interface IStepComponentProps<T> {
    data: T;
    onChange: (newData: Partial<T>) => void;
}
  
interface IStep<T> {
    title: string;
    component: React.ComponentType<IStepComponentProps<T>>;
}

interface IMultiStepFormProps<T> {
    steps: IStep<T>[];
    submitFunc: (data: T) => void;
    initialData: T;
}

// ==============================
// Specific form data types
// ==============================

interface IAssignmentFormData {
    assignmentName: string
    classroomId: number
    groupAssignment: boolean
    mainDueDate: Date | null
    templateRepo: IAssignmentTemplate | null
}