import React from "react";

interface Props {
  orgsWithApp: IOrganization[];
  orgsWithoutApp: IOrganization[];
  selectedOrg: IOrganization | null;
  loading: boolean;
  onSelect: (org: IOrganization) => Promise<void>;
}

const OrganizationDropdown: React.FC<Props> = ({
  orgsWithApp,
  orgsWithoutApp,
  selectedOrg,
  loading,
  onSelect,
}) => {
  return (
    <div className="form-group">
      <label htmlFor="organization">Select Organization:</label>
      <select
        id="organization"
        value={selectedOrg ? selectedOrg.login : ""}
        onChange={async (e) => {
          const selectedLogin = e.target.value;
          if (selectedLogin === "create_new_org") {
            window.open("https://github.com/organizations/plan", "_blank");
          } else {
            const selected = [...orgsWithApp, ...orgsWithoutApp].find(
              (org) => org.login === selectedLogin
            );
            if (selected) {
              await onSelect(selected);
            }
          }
        }}
      >
        {loading ? (
          <option value="" disabled>
            Loading organizations...
          </option>
        ) : (
          <>
            <option value="" disabled>
              Select an organization
            </option>
            {orgsWithoutApp && orgsWithApp.length > 0 && (
              <optgroup label="Organizations with GitGrader Installed">
                {orgsWithApp.map((org) => (
                  <option
                    key={org.id}
                    value={org.login}
                    title="This organization has the GitGrader installed"
                  >
                    {org.login} ✔️
                  </option>
                ))}
              </optgroup>
            )}
            {orgsWithoutApp && orgsWithoutApp.length > 0 && (
              <optgroup label="Organizations without GitGrader Installed">
                {orgsWithoutApp.map((org) => (
                  <option
                    key={org.id}
                    value={org.login}
                    title="GitGrader not installed on this organization"
                  >
                    {org.login} ❌
                  </option>
                ))}
              </optgroup>
            )}
          </>
        )}
        <option value="create_new_org">Create a New Organization ➕</option>
      </select>
      {selectedOrg &&
        orgsWithoutApp.some((org) => org.login === selectedOrg.login) && (
          <a
            href={selectedOrg.html_url}
            target="_blank"
            rel="noopener noreferrer"
          >
            <button>Install GitGrader on {selectedOrg.login}</button>
          </a>
        )}
    </div>
  );
};

export default OrganizationDropdown;
