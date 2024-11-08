import { useEffect, useContext, useState } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getOrganizationTemplates } from "@/api/organizations";
import { useNavigate } from "react-router-dom";

import MultiStepForm from '@/components/MultiStepForm';
import { Step, StepComponentProps } from '@/components/MultiStepForm/Interfaces/main';
import AssignmentDetails from '@/components/MultiStepForm/CreateAssignment/AssignmentDetails';
import StarterCodeDetails from '@/components/MultiStepForm/CreateAssignment/StarterCodeDetails';
import { createAssignment, createAssignmentTemplate } from "@/api/assignments";
import { AssignmentFormData } from "@/components/MultiStepForm/Interfaces/CreateAssignment";

const CreateAssignment: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const orgName = selectedClassroom?.org_name;

  const [templates, setTemplates] = useState<IRepository[]>([]);
  const [loadingTemplates, setLoadingTemplates] = useState(true);

  const navigate = useNavigate();

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
        assignmentName: '',
        classroomId: selectedClassroom?.id || -1,
        groupAssignment: false,
        mainDueDate: null,
        templateRepo: null
    };

    const handleSubmit = (data: AssignmentFormData) => {
        createAssignmentTemplate(data.classroomId, data.templateRepo!)
        createAssignment(data.templateRepo!.id, data)
        
        navigate('/app/dashboard');
    }

    return (
        <div>
            <h1>Create Assignment</h1>
            <MultiStepForm
                steps={steps}
                submitFunc={handleSubmit}
                initialData={initialData}
            />
        </div>
    );
};

export default CreateAssignment;
