import React from 'react';
import TemplateRepoDropdown from "@/components/Dropdown/Repository";

interface StarterCodeDetailsProps extends IStepComponentProps<IAssignmentFormData> {
  repositories: IAssignmentTemplate[];
  isLoading: boolean;
}

const StarterCodeDetails: React.FC<StarterCodeDetailsProps> = ({ data, onChange, repositories, isLoading }) => {
  return (
    <div>
      <h2>Starter Code Repository</h2>
      <div>
        <TemplateRepoDropdown
          repositories={repositories}
          onChange={(selectedTemplate: IAssignmentTemplate) => {
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
