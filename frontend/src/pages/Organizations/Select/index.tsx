import React, { useState } from "react";
import "./styles.css";
import OrganizationDropdown from "@/components/Dropdown/Organization";
import Panel from "@/components/Panel";
import Button from "@/components/Button";
import {
  getAppInstallations,
  getOrganizationDetails,
} from "@/api/organizations";
import { getCallbackURL } from "@/api/auth";
import { useQuery } from "@tanstack/react-query";

const OrganizationSelection: React.FC = () => {
  const [selectedOrg, setSelectedOrg] = useState<IOrganization | null>(null);

  const { data: installationsData, isLoading: loadingOrganizations, error: installationsError } = useQuery({
    queryKey: ['organizations'],
    queryFn: getAppInstallations,
  });

  const { data: callbackData } = useQuery({
    queryKey: ['callback-url'],
    queryFn: getCallbackURL,
  });

  const { error: orgDetailsError } = useQuery({
    queryKey: ['org-details', selectedOrg?.login],
    queryFn: async () => {
      if (!selectedOrg?.login) return null;
      return await getOrganizationDetails(selectedOrg.login);
    },
    enabled: !!selectedOrg?.login,
  });

  const orgsWithApp = installationsData?.orgs_with_app || [];
  const orgsWithoutApp = installationsData?.orgs_without_app || [];
  const consentUrl = callbackData?.consent_url || null;
  const error = installationsError || orgDetailsError;

  const handleOrganizationSelect = async (org: IOrganization) => {
    setSelectedOrg(org);
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
                href="/app/classroom/select"
                state={{ orgID: selectedOrg.id }}
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
                Install Marks on {selectedOrg.login}
              </Button>
            )}
        </div>
        {consentUrl && (
          <a className="Organization__link" href={consentUrl}>
            {"Don't see your organization?"}
          </a>
        )}
        {error && <div className="error">{error instanceof Error ? error.message : "An error occurred"}</div>}
      </div>
    </Panel>
  );
};

export default OrganizationSelection;
