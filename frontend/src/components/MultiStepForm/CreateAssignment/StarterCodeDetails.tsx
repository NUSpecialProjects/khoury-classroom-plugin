import React from 'react';
import RepositoryDropdown from "@/components/Dropdown/Repository";

interface StarterCodeDetailsProps extends StepComponentProps<AssignmentFormData> {
  repositories: IRepository[];
  isLoading: boolean;
}

const StarterCodeDetails: React.FC<StarterCodeDetailsProps> = ({ data, onChange, repositories, isLoading }) => {
  return (
    <div>
      <h2>Starter Code Repository</h2>
      <div>
        <RepositoryDropdown
          repositories={repositories}
          onChange={(selectedTemplate: IRepository) => {
            onChange({ templateRepo: selectedTemplate });
          }}
          loading={isLoading}
          selectedRepo={data.templateRepo}
        />
      </div>
    </div>
  );
};

export default StarterCodeDetails;
