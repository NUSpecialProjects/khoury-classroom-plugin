import React, { useCallback } from 'react';
import { StepComponentProps } from '../Interfaces/main';
import { AssignmentFormData } from '../Interfaces/CreateAssignment';

const AssignmentDetails: React.FC<StepComponentProps<AssignmentFormData>> = ({ data, onChange }) => {
  const handleInputChange = useCallback(
    (
      e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
    ) => {
      const { name, value, type } = e.target;

      if (type === 'checkbox') {
        const target = e.target as HTMLInputElement;
        onChange({ [name]: target.checked });
      } else if (type === 'date') {
        onChange({ [name]: value ? new Date(value) : null });
      } else {
        onChange({ [name]: value });
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
