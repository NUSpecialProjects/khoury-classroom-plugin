import { useState, useCallback, ReactElement } from "react";
import Button from "../Button";
import { useNavigate } from "react-router-dom";
import './styles.css';

const MultiStepForm = <T,>({
  steps,
  submitFunc,
  cancelLink,
  initialData,
}: IMultiStepFormProps<T>): ReactElement => {
  const [currentStepIndex, setCurrentStepIndex] = useState<number>(0);
  const [formData, setFormData] = useState<T>(initialData);
  const [error, setError] = useState<string | null>(null);

  const totalSteps = steps.length;
  const isFirstStep = currentStepIndex === 0;
  const isLastStep = currentStepIndex === totalSteps - 1;

  const handleNext = useCallback(() => {
    setCurrentStepIndex((prev) => Math.min(prev + 1, totalSteps - 1));
  }, [totalSteps]);

  const navigate = useNavigate();
  const handleCancel = () => navigate(cancelLink);

  const handlePrevious = useCallback(() => {
    setCurrentStepIndex((prev) => Math.max(prev - 1, 0));
  }, []);

  const handleDataChange = useCallback(
    // Update form data
    (newData: Partial<T>) => {
      setFormData((prevData) => ({
        ...prevData,
        ...newData,
      }));

      // Clear errors on form change
      setError(null);
    },
    []
  );

  const handleSubmit = useCallback(async () => {
    const success = await submitFunc(formData);
    if (success) {
      setError(null);
    } else {
      setError("Submission failed. Please try again.");
    }
  }, [submitFunc, formData]);

  const CurrentStepComponent = steps[currentStepIndex].component;

  return (
    <div className="MultiStepForm">
      <div className="MultiStepForm__formStep">
        <CurrentStepComponent data={formData} onChange={handleDataChange} />
      </div>

      {error && <p>{error}</p>}

      <div className="MultiStepForm__formNavigationButtonsWrapper">
        {
          isFirstStep && (
            <Button onClick={handleCancel} variant="secondary">
              Cancel
            </Button>
          )
        }
        {
          !isFirstStep && (
            <Button onClick={handlePrevious} variant="secondary">
              Previous
            </Button>
          )
        }
        {!isLastStep && (
          <Button onClick={handleNext} variant="primary">
            Continue
          </Button>
        )}
        {isLastStep && (
          <Button onClick={handleSubmit} variant="primary">
            Finish
          </Button>
        )}
      </div>
    </div>
  );
};

export default MultiStepForm;
