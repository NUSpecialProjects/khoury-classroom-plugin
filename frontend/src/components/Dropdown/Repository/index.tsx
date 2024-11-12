<<<<<<< HEAD
import React, { useState, useEffect } from 'react';

interface RepositoryDropdownProps {
    repositories: IRepository[];
    onChange?: (selectedRepoId: number) => void;
    selectedRepoId?: number;
    loading: boolean;
}

const RepositoryDropdown: React.FC<RepositoryDropdownProps> = ({
    repositories,
    onChange,
    selectedRepoId,
    loading,
}) => {
    const [value, setValue] = useState<string>(selectedRepoId?.toString() || '');

    useEffect(() => {
        setValue(selectedRepoId?.toString() || '');
    }, [selectedRepoId]);

    const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedId = parseInt(event.target.value, 10);
        setValue(event.target.value);
        if (onChange) {
            onChange(selectedId);
        }
    };

    return (
        <select value={value} onChange={handleChange}>
            <option value="" disabled>
                {value
                    ? repositories.find(repo => repo.id === parseInt(value, 10))?.name || 'Select a repository'
                    : 'Select a repository'}
            </option>
            {loading ? (
                <option className="Dropdown__option" value="" disabled>
                    Loading templates...
                </option>
            ) : (
                repositories.length > 0 ? (
                    repositories.map((repo) => (
                        <option key={repo.id} value={repo.id.toString()}>
                            {repo.name}
                        </option>
                    ))
                ) : (
                    <option value="" disabled>
                        No repositories available
                    </option>
                )
            )}
        </select>
    );
=======
import React, { useState, useEffect } from "react";

interface RepositoryDropdownProps {
  repositories: IRepository[];
  onChange?: (selectedRepoId: number) => void;
  selectedRepoId?: number;
  loading: boolean;
}

const RepositoryDropdown: React.FC<RepositoryDropdownProps> = ({
  repositories,
  onChange,
  selectedRepoId,
  loading,
}) => {
  const [value, setValue] = useState<string>(selectedRepoId?.toString() || "");

  useEffect(() => {
    setValue(selectedRepoId?.toString() || "");
  }, [selectedRepoId]);

  const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedId = parseInt(event.target.value, 10);
    setValue(event.target.value);
    if (onChange) {
      onChange(selectedId);
    }
  };

  return (
    <select value={value} onChange={handleChange}>
      <option value="" disabled>
        {value
          ? repositories.find((repo) => repo.id === parseInt(value, 10))
              ?.name || "Select a repository"
          : "Select a repository"}
      </option>
      {loading ? (
        <option className="Dropdown__option" value="" disabled>
          Loading templates...
        </option>
      ) : repositories.length > 0 ? (
        repositories.map((repo) => (
          <option key={repo.id} value={repo.id.toString()}>
            {repo.name}
          </option>
        ))
      ) : (
        <option value="" disabled>
          No repositories available
        </option>
      )}
    </select>
  );
>>>>>>> main
};

export default RepositoryDropdown;
