import React from 'react';
import TemplateRepoDropdown from "@/components/Dropdown/Repository";

interface StarterCodeDetailsProps extends IStepComponentProps<IAssignmentFormData> {
  templateRepos: ITemplateRepo[];
  isLoading: boolean;
}

const StarterCodeDetails: React.FC<StarterCodeDetailsProps> = ({ data, onChange, templateRepos, isLoading }) => {
  return (
    <div className="CreateAssignmentForms">
      <h2 className="CreateAssignmentForms__header">Starter Code Repository</h2>
      <div>
        <TemplateRepoDropdown
          repositories={templateRepos}
          onChange={(selectedTemplate: ITemplateRepo) => {
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
