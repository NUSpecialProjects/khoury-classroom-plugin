import React, { useCallback, useState } from 'react';

const AssignmentDetails: React.FC<IStepComponentProps<IAssignmentFormData>> = ({ data, onChange }) => {
  // const [isValid, setIsValid] = React.useState(false);
  // const [errors, setErrors] = useState<{ assignmentName?: string; mainDueDate?: string }>({});
  // const assignmentNameRegex = /^[a-z0-9-]+$/;
  
  // const validateAssignmentTitle = (title: string) => {
  //   if (!title) {
  //     setErrors((prevErrors) => ({ ...prevErrors, assignmentName: 'Assignment name is required' }));
  //   } else if (!assignmentNameRegex.test(title)) {
  //     setErrors((prevErrors) => ({
  //       ...prevErrors,
  //       assignmentName: 'Assignment name can only contain lowercase letters, numbers, and hyphens',
  //     }));
  //   } else {
  //     setErrors((prevErrors) => ({ ...prevErrors, assignmentName: undefined }));
  //   }
  // };
  // const validateDueDate = (date: Date | null) => {
  //   if (!date) {
  //     setErrors((prevErrors) => ({ ...prevErrors, mainDueDate: 'Due date is required' }));
  //   } else if (date < new Date()) {
  //     setErrors((prevErrors) => ({ ...prevErrors, mainDueDate: 'Due date must be in the future' }));
  //   } else {
  //     setErrors((prevErrors) => ({ ...prevErrors, mainDueDate: undefined }));
  //   }
  // };

  const handleCheckboxChange = (target: HTMLInputElement) => {
    onChange({ [target.name]: target.checked });
  };
  const handleDateChange = (target: HTMLInputElement) => {
    const newDate = target.value ? new Date(target.value) : null;
    onChange({ [target.name]: newDate });
    // validateDueDate(newDate);
  };
  const handleTextChange = (target: HTMLInputElement) => {
    const newName = target.value;
    onChange({ [target.name]: newName });
    // validateAssignmentTitle(newName);
  };

  const handleInputChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const { type } = e.target;
      const target = e.target as HTMLInputElement;

      switch (type) {
        case 'checkbox':
          handleCheckboxChange(target);
          break;
        case 'date':
          handleDateChange(target);
          break;
        default:
          handleTextChange(target);
      }
    },
    [onChange]
  );

  return (
    <form className="assignment-details">
      <h2>Assignment Details</h2>

      <div className="form-group">
        <label htmlFor="assignmentName">Assignment Name:</label>
        <input
          id="assignmentName"
          type="text"
          name="assignmentName"
          value={data.assignmentName}
          onChange={handleInputChange}
          required
        />
      </div>

      <div className="form-group">
        <label htmlFor="mainDueDate">Due Date:</label>
        <input
          id="mainDueDate"
          type="date"
          name="mainDueDate"
          value={data.mainDueDate ? data.mainDueDate.toISOString().split('T')[0] : ''}
          onChange={handleInputChange}
          required
        />
      </div>

      <div className="form-group checkbox-group">
        <input
          id="groupAssignment"
          type="checkbox"
          name="groupAssignment"
          checked={data.groupAssignment}
          onChange={handleInputChange}
        />
        <label htmlFor="groupAssignment">Group Assignment</label>
      </div>
    </form>
  );
};

export default AssignmentDetails;