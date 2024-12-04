import React, { useEffect, useState } from "react";
import "./styles.css";
import OrganizationDropdown from "@/components/Dropdown/Organization";
import Panel from "@/components/Panel";
import Button from "@/components/Button";
import {
  getAppInstallations,
  getOrganizationDetails,
} from "@/api/organizations";
import { getCallbackURL } from "@/api/auth";

const OrganizationSelection: React.FC = () => {
  const [orgsWithApp, setOrgsWithApp] = useState<IOrganization[]>([]);
  const [orgsWithoutApp, setOrgsWithoutApp] = useState<IOrganization[]>([]);
  const [loadingOrganizations, setLoadingOrganizations] = useState(true);
  const [selectedOrg, setSelectedOrg] = useState<IOrganization | null>(null);
  const [consentUrl, setConsentUrl] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

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

        const callbackData = await getCallbackURL();
        setConsentUrl(callbackData.consent_url);
        setError(null);
      } catch (e) {
        setError("Error fetching organizations");
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
        setError("Error fetching organization details");
      });
  };

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
              <Button
                href={`https://github.com/apps/khoury-classroom/installations/new/permissions?target_id=${selectedOrg.id}&target_type=Organization`}
                newTab={true}
              >
                Install GitGrader for {selectedOrg.login}
              </Button>
            )}
        </div>
        {consentUrl && (
          <a className="Organization__link" href={consentUrl}>
            {"Don't see your organization?"}
          </a>
        )}
        {error && <div className="error">{error}</div>}
      </div>
    </Panel>
  );
};

export default OrganizationSelection;
