import React, { useEffect, useState } from "react";
import "./styles.css";
import OrganizationDropdown from "@/components/Dropdown/Organization";
import Panel from "@/components/Panel";
import Button from "@/components/Button";
import {
  getAppInstallations,
  getOrganizationDetails,
} from "@/api/organizations";

const OrganizationSelection: React.FC = () => {
  const [orgsWithApp, setOrgsWithApp] = useState<IOrganization[]>([]);
  const [orgsWithoutApp, setOrgsWithoutApp] = useState<IOrganization[]>([]);
  const [loadingOrganizations, setLoadingOrganizations] = useState(true);
  const [selectedOrg, setSelectedOrg] = useState<IOrganization | null>(null);
  

  useEffect(() => {
    const fetchOrganizations = async () => {
      try {
        const data: IOrganizationsResponse = await getAppInstallations();
        //checking if data exists before populating and setting lists
        if (data.orgs_with_app) {
          setOrgsWithApp(data.orgs_with_app);
        }
        if (data.orgs_without_app) {
          setOrgsWithoutApp(data.orgs_without_app);
        }
      } catch (_) {
        // do nothing
      } finally {
        setLoadingOrganizations(false);
      }
    };

    void fetchOrganizations();
  }, []);

  const handleOrganizationSelect = async (org: IOrganization) => {
    setSelectedOrg(org);
    await getOrganizationDetails(org.login)
      .then((orgDetails) => {
        setSelectedOrg(orgDetails);
      })
      .catch((_) => {
        // do nothing
      });
  };

  //TODO: visually disable the button while it's loading the org details
  return (
    <Panel title="Your Organizations" logo={true}>
      <div className="Organization">
        <OrganizationDropdown
          orgsWithApp={orgsWithApp}
          orgsWithoutApp={orgsWithoutApp}
          selectedOrg={selectedOrg}
          loading={loadingOrganizations}
          onSelect={handleOrganizationSelect}
        />

        <div className="Organization__buttonWrapper">
          {selectedOrg &&
            orgsWithApp.some((org) => org.login === selectedOrg.login) && (
              <Button
                variant="primary"
                href={`/app/classroom/select?org_id=${selectedOrg.id}`}
              >
                View Classrooms for {selectedOrg.login}
              </Button>
            )}
          {selectedOrg &&
            orgsWithoutApp.some((org) => org.login === selectedOrg.login) && (
              <Button href={selectedOrg.html_url}>
                Install GitGrader for {selectedOrg.login}
              </Button>
            )}
        </div>
      </div>
    </Panel>
  );
};

export default OrganizationSelection;
