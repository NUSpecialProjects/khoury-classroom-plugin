import { useState, useCallback, ReactElement } from 'react';
import { MultiStepFormProps } from './Interfaces/Main';

const MultiStepForm = <T,>({ steps, submitFunc, initialData }: MultiStepFormProps<T>): ReactElement => {
  const [currentStepIndex, setCurrentStepIndex] = useState<number>(0);
  const [formData, setFormData] = useState<T>(initialData);

  const totalSteps = steps.length;
  const isFirstStep = currentStepIndex === 0;
  const isLastStep = currentStepIndex === totalSteps - 1;

  const handleNext = useCallback(() => {
    setCurrentStepIndex((prev) => Math.min(prev + 1, totalSteps - 1));
  }, [totalSteps]);

  const handlePrevious = useCallback(() => {
    setCurrentStepIndex((prev) => Math.max(prev - 1, 0));
  }, []);

  const handleDataChange = useCallback(
    (newData: Partial<T>) => {
      setFormData((prevData) => ({
        ...prevData,
        ...newData,
      }));
    }, []
  );

  const handleSubmit = useCallback(() => {
    submitFunc(formData);
  }, [submitFunc, formData]);

  const CurrentStepComponent = steps[currentStepIndex].component;

  return (
    <div className="multi-step-form">
      <div className="form-step">
        <CurrentStepComponent data={formData} onChange={handleDataChange} />
      </div>

      <div className="form-navigation">
        {!isFirstStep && (
          <button type="button" onClick={handlePrevious} className="btn btn-secondary">
            Previous
          </button>
        )}
        {!isLastStep && (
          <button type="button" onClick={handleNext} className="btn btn-primary">
            Next
          </button>
        )}
        {isLastStep && (
          <button type="button" onClick={handleSubmit} className="btn btn-success">
            Submit
          </button>
        )}
      </div>
    </div>
  );
};

export default MultiStepForm;
