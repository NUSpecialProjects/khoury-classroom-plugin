import { useEffect, useContext, useState } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getOrganizationTemplates } from "@/api/organizations";

import MultiStepForm from '@/components/MultiStepForm';
import { Step, StepComponentProps } from '@/components/MultiStepForm/Interfaces/main';
import { AssignmentFormData } from "@/components/MultiStepForm/Interfaces/CreateAssignment";
import AssignmentDetails from '@/components/MultiStepForm/CreateAssignment/AssignmentDetails';
import StarterCodeDetails from '@/components/MultiStepForm/CreateAssignment/StarterCodeDetails';

const CreateAssignment: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const orgName = selectedClassroom?.org_name;

  const [templates, setTemplates] = useState<IRepository[]>([]);
  const [loadingTemplates, setLoadingTemplates] = useState(true);

  useEffect(() => {
    const fetchTemplates = async (orgName: string | undefined) => {
      if (orgName) {
        setLoadingTemplates(true);

        // TODO: Implement dynamic pagination in template dropdown
        getOrganizationTemplates(orgName, "100", "1")
          .then((response) => {
            setTemplates(response.template_repos);
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

    const steps: Step<AssignmentFormData>[] = [
        { title: 'Assignment Details', component: AssignmentDetails },
        {
            title: 'Starter Code Repository',
            component: (props: StepComponentProps<AssignmentFormData>) => (
                <StarterCodeDetails
                    {...props}
                    repositories={templates}
                    isLoading={loadingTemplates}
                />
            )
        },
    ];

    const initialData: AssignmentFormData = {
        name: '',
        description: '',
        dueDate: null,
        isGroupAssignment: false,
        selectedRepoId: 0,
    };

    return (
        <div>
            <h1>Create Assignment</h1>
            <MultiStepForm
                steps={steps}
                submitFunc={(data) => {
                    console.log(data);
                }}
                initialData={initialData}
            />
        </div>
    );
};

export default CreateAssignment;
