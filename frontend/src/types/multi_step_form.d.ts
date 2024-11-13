// ==============================
// MultiStepForm types
// ==============================

interface StepComponentProps<T> {
    data: T;
    onChange: (newData: Partial<T>) => void;
}
  
interface Step<T> {
    title: string;
    component: React.ComponentType<StepComponentProps<T>>;
}

interface MultiStepFormProps<T> {
    steps: Step<T>[];
    submitFunc: (data: T) => void;
    initialData: T;
}

// ==============================
// Specific form data types
// ==============================

interface AssignmentFormData {
    assignmentName: string
    classroomId: number
    groupAssignment: boolean
    mainDueDate: Date | null
    templateRepo: IRepository | null
}