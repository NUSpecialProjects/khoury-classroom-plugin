import React, { useEffect, useState } from "react";
import "./styles.css";
import { useNavigate } from "react-router-dom";
import OrganizationDropdown from "@/components/Dropdown/Organization";
import Panel from "@/components/Panel";
import Button from "@/components/Button";
import {
  getAppInstallations,
} from "@/api/organizations";


const OrganizationSelection: React.FC = () => {
  const [orgsWithApp, setOrgsWithApp] = useState<IOrganization[]>([]);
  const [orgsWithoutApp, setOrgsWithoutApp] = useState<IOrganization[]>([]);
  const [loadingOrganizations, setLoadingOrganizations] = useState(true);
  const [selectedOrg, setSelectedOrg] = useState<IOrganization | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchOrganizations = async () => {
      try {
        const data: IOrganizationsResponse = await getAppInstallations();
        console.log("data", data);

        //checking if data exists before populating and setting lists
        if (data.orgs_with_app) {
          setOrgsWithApp(data.orgs_with_app);
          console.log("orgsWithApp", orgsWithApp);
        }
        if (data.orgs_without_app) {
          setOrgsWithoutApp(data.orgs_without_app);
          console.log("orgsWithoutApp", orgsWithoutApp);
        }
      } catch (error) {
        console.error("Error fetching organizations:", error);
      } finally {
        setLoadingOrganizations(false);
      }
    };

    void fetchOrganizations();
  }, []);


  const handleOrganizationSelect = async (org: IOrganization) => {
    try {
      setSelectedOrg(org);
    } catch (error) {
      console.error("Error fetching organization details:", error);
    }
  };

  return (
    <Panel title="Your Organizations" logo={true}>
        <>
          <OrganizationDropdown
            orgsWithApp={orgsWithApp}
            orgsWithoutApp={orgsWithoutApp}
            selectedOrg={selectedOrg}
            loading={loadingOrganizations}
            onSelect={handleOrganizationSelect}
          />
            {selectedOrg && orgsWithApp.some((org) => org.login === selectedOrg.login) && (
                 <Button variant="primary" onClick={() => navigate(`/app/classroom/select?org_id=${selectedOrg.id}`)}>
              View Classrooms for {selectedOrg.login}
            </Button>
              )}
        </>
      <div className="Creation__buttonWrapper">
        {selectedOrg &&
          orgsWithoutApp.some((org) => org.login === selectedOrg.login) && (
            <Button variant="primary" href={selectedOrg.html_url}>
              Install GitGrader for {selectedOrg.login}
            </Button>
          )}
      </div>
    </Panel>
  );
};

export default OrganizationSelection;
