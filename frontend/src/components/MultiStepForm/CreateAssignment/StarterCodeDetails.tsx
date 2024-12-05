import React from 'react';
import GenericDropdown from '@/components/Dropdown';

interface StarterCodeDetailsProps extends IStepComponentProps<IAssignmentFormData> {
  templateRepos: ITemplateRepo[];
  isLoading: boolean;
}

const StarterCodeDetails: React.FC<StarterCodeDetailsProps> = ({ data, onChange, templateRepos, isLoading }) => {
  const formattedOptions = templateRepos.map(repo => repo.template_repo_name);

  const selectedOption = data.templateRepo ? data.templateRepo.template_repo_name : null;

  return (
    <div className="CreateAssignmentForms">
      <h2 className="CreateAssignmentForms__header">Starter Code Repository</h2>
      <div>
        <GenericDropdown
          options={formattedOptions.map(option => ({ value: option, label: option }))}
          onChange={(selected) => {
            const selectedRepo = templateRepos.find(repo => repo.template_repo_name === selected);
            if (selectedRepo) {
              onChange({ templateRepo: selectedRepo });
            }
          }}
          selectedOption={selectedOption}
          loading={isLoading}
          labelText="Pick a template repository to use as the starter code"
          captionText="The template repository must be owned by the organization you are in"
        />
      </div>
    </div>
  );
};

export default StarterCodeDetails;
