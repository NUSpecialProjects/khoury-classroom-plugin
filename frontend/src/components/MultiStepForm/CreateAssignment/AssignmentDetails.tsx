import React, { useCallback } from 'react';

const AssignmentDetails: React.FC<IStepComponentProps<IAssignmentFormData>> = ({ data, onChange }) => {
  const handleCheckboxChange = (target: HTMLInputElement) => {
    onChange({ [target.name]: target.checked });
  };
  const handleDateChange = (target: HTMLInputElement) => {
    onChange({ [target.name]: target.value ? new Date(target.value) : null });
  };
  const handleTextChange = (target: HTMLInputElement | HTMLTextAreaElement) => {
    onChange({ [target.name]: target.value });
  };

  const handleInputChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
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
