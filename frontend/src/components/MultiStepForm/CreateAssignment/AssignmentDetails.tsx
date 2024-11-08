import React, { useCallback } from 'react';
import { StepComponentProps } from '../Interfaces/main';
import { AssignmentFormData } from '../Interfaces/CreateAssignment';

const AssignmentDetails: React.FC<StepComponentProps<AssignmentFormData>> = ({ data, onChange }) => {
  const handleInputChange = useCallback(
    (
      e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
    ) => {
      const { name, value, type } = e.target;
      let updatedValue: string | boolean | Date | null = value;

      if (type === 'checkbox') {
        updatedValue = (e.target as HTMLInputElement).checked;
      } else if (type === 'date') {
        updatedValue = value ? new Date(value) : null;
      }

      onChange({ [name]: updatedValue });
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
          name="name"
          value={data.name}
          onChange={handleInputChange}
          required
        />
      </div>

      <div className="form-group">
        <label htmlFor="description">Description:</label>
        <textarea
          id="description"
          name="description"
          value={data.description}
          onChange={handleInputChange}
        />
      </div>

      <div className="form-group">
        <label htmlFor="dueDate">Due Date:</label>
        <input
          id="dueDate"
          type="date"
          name="dueDate"
          value={data.dueDate ? data.dueDate.toISOString().split('T')[0] : ''}
          onChange={handleInputChange}
          required
        />
      </div>

      <div className="form-group checkbox-group">
        <input
          id="isGroupAssignment"
          type="checkbox"
          name="isGroupAssignment"
          checked={data.isGroupAssignment}
          onChange={handleInputChange}
          required
        />
        <label htmlFor="isGroupAssignment">Group Assignment</label>
      </div>
    </form>
  );
};

export default AssignmentDetails;
