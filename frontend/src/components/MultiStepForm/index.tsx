import { useState, useCallback, ReactElement } from "react";
import Button from "../Button";
import { useNavigate } from "react-router-dom";
import './styles.css';

const MultiStepForm = <T,>({
  steps,
  cancelLink,
  initialData,
}: IMultiStepFormProps<T>): ReactElement => {
  // Default form state
  const [formData, setFormData] = useState<T>(initialData);
  const [error, setError] = useState<string | null>(null);

  // Step navigation
  const [currentStepIndex, setCurrentStepIndex] = useState<number>(0);
  const totalSteps = steps.length;
  const isFirstStep = currentStepIndex === 0;
  const isLastStep = currentStepIndex === totalSteps - 1;

  // Site navigation
  const navigate = useNavigate();
  const handleCancel = () => navigate(cancelLink);

  // Navigate backwards when the page exists
  const handlePrevious = useCallback(() => {
    setCurrentStepIndex((prev) => Math.max(prev - 1, 0));
  }, []);

  // Update form data, clearing errors on form change
  const handleDataChange = useCallback(
    (newData: Partial<T>) => {
      setFormData((prevData) => ({
        ...prevData,
        ...newData,
      }));

      setError(null);
    },
    []
  );

  // Handle next button click, preventing progression on error
  const handleNext = useCallback(async () => {
    const currentStep = steps[currentStepIndex];

    try {
      await currentStep.onNext(formData);
      setError(null);

      if (!isLastStep) {
        setCurrentStepIndex((prev) => Math.min(prev + 1, totalSteps - 1));
      }
    } catch (e) {
      const err: string = e instanceof Error
        ? e.message
        : "An error occurred. Please try again.";
      setError(err);
    }
  }, [steps, currentStepIndex, formData, isLastStep, totalSteps]);

  const CurrentStepComponent = steps[currentStepIndex].component;
  return (
    <div className="MultiStepForm">
      {/* Display current step */}
      <div className="MultiStepForm__formStep">
        <CurrentStepComponent data={formData} onChange={handleDataChange} />
      </div>

      {/* Display error message */}
      {error && <p style={{ color: 'red' }}>{error}</p>}

      {/* Render navigation buttons */}
      <div className="MultiStepForm__formNavigationButtonsWrapper">
        {isFirstStep && (
          <Button onClick={handleCancel} variant="secondary">
            Cancel
          </Button>
        )}
        {!isFirstStep && (
          <Button onClick={handlePrevious} variant="secondary">
            Previous
          </Button>
        )}
        <Button onClick={handleNext} variant="primary">
          {isLastStep ? 'Finish' : 'Continue'}
        </Button>
      </div>
    </div>
  );
};

export default MultiStepForm;