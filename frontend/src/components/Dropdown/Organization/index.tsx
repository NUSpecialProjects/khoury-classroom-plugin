import React from "react";
import GenericDropdown, { IDropdownOption } from "..";

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
  const getOptions = (): IDropdownOption[] => {
    const options: IDropdownOption[] = [];

    if (orgsWithApp.length > 0) {
      options.push({ value: "header-1", label: "Organizations with GitGrader Installed", disabled: true });
      orgsWithApp.forEach(org => {
        options.push({
          value: org.login,
          label: `${org.login} ✔️`,
        });
      });
    }

    if (orgsWithoutApp.length > 0) {
      options.push({ value: "header-2", label: "Organizations without GitGrader Installed", disabled: true });
      orgsWithoutApp.forEach(org => {
        options.push({
          value: org.login,
          label: `${org.login} ❌`,
        });
      });
    }

    options.push({
      value: "create_new_org",
      label: "Create a New Organization ➕"
    });

    return options;
  };

  const handleChange = async (selected: string) => {
    if (selected === "create_new_org") {
      window.open("https://github.com/organizations/plan", "_blank");
    } else {
      const selectedOrg = [...orgsWithApp, ...orgsWithoutApp].find(
        (org) => org.login === selected
      );
      if (selectedOrg) {
        await onSelect(selectedOrg);
      }
    }
  };

  return (
    <GenericDropdown
      options={getOptions()}
      onChange={handleChange}
      selectedOption={selectedOrg ? selectedOrg.login : null}
      loading={loading}
      labelText="Select an Organization"
      loadingText="Loading organizations..."
      placeholder="Select an organization"
    />
  );
};

export default OrganizationDropdown;

