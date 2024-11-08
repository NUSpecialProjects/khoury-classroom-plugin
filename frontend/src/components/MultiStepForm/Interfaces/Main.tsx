export interface StepComponentProps<T> {
  data: T;
  onChange: (newData: Partial<T>) => void;
}

export interface Step<T> {
  title: string;
  component: React.ComponentType<StepComponentProps<T>>;
}

export interface MultiStepFormProps<T> {
  steps: Step<T>[];
  submitFunc: (data: T) => void;
  initialData: T;
}