import { useEffect, useContext, useState } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getOrganizationTemplates } from "@/api/organizations";
import RepositoryDropdown from "@/components/Dropdown/Repository";

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
          .catch((err: unknown) => {
            console.error("Error fetching templates:", err);
          })
          .finally(() => {
            setLoadingTemplates(false);
          });
      }
    };

    fetchTemplates(orgName);
  }, [orgName]);

  return (
    <div>
      <h1>Create Assignment</h1>
      <RepositoryDropdown
        repositories={templates}
        onChange={(selectedRepoId: number) => {
          console.log("Selected repo ID:", selectedRepoId);
          // PLACEHOLDER
        }}
        loading={loadingTemplates}
      />
    </div>
  );
};

export default CreateAssignment;
