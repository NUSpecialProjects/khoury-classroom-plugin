import { useEffect, useContext, useState } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getOrganizationTemplates } from "@/api/organizations";
import { useNavigate } from "react-router-dom";

import MultiStepForm from '@/components/MultiStepForm';
import AssignmentDetails from '@/components/MultiStepForm/CreateAssignment/AssignmentDetails';
import StarterCodeDetails from '@/components/MultiStepForm/CreateAssignment/StarterCodeDetails';
import { createAssignment } from "@/api/assignments";

import './styles.css'

const CreateAssignment: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const orgName = selectedClassroom?.org_name;

  const [templateRepos, setTemplateRepos] = useState<ITemplateRepo[]>([]);
  const [loadingTemplates, setLoadingTemplates] = useState(true);

  const navigate = useNavigate();

  useEffect(() => {
    const fetchTemplates = async (orgName: string | undefined) => {
      if (orgName) {
        setLoadingTemplates(true);

        // TODO: Implement dynamic pagination in template dropdown
        getOrganizationTemplates(orgName, "100", "1")
          .then((response) => {
            setTemplateRepos(response.templates);
          })
          .catch((_: unknown) => {
            // do nothing
          })
          .finally(() => {
            setLoadingTemplates(false);
          });
      }
    };

    fetchTemplates(orgName);
  }, [orgName]);

  const steps: IStep<IAssignmentFormData>[] = [
    { title: 'Assignment Details', component: AssignmentDetails },
    {
      title: 'Starter Code Repository',
      component: (props: IStepComponentProps<IAssignmentFormData>) => (
        <StarterCodeDetails
          {...props}
          templateRepos={templateRepos}
          isLoading={loadingTemplates}
        />
      )
    },
  ];

  const initialData: IAssignmentFormData = {
    assignmentName: '',
    classroomId: selectedClassroom?.id || -1,
    groupAssignment: false,
    mainDueDate: null,
    templateRepo: null
  };

  const handleSubmit = async (data: IAssignmentFormData): Promise<boolean> => {
    if (!data?.templateRepo?.template_repo_id) return false;

    await createAssignment(data.templateRepo.template_repo_id, data);

    navigate('/app/dashboard');
    return true;
  }

  return (
    <div className="CreateAssignment">
      <div className="CreateAssignment__header">
        <h1>Create Assignment</h1>
      </div>
      <MultiStepForm
        steps={steps}
        submitFunc={handleSubmit}
        cancelLink="/app/dashboard"
        initialData={initialData}
      />
    </div>
  );
};

export default CreateAssignment;
