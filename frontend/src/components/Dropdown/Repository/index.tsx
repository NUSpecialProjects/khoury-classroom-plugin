import React, { ChangeEvent } from 'react';

interface RepositoryDropdownProps {
  repositories: IRepository[];
  onChange?: (selectedRepoId: number) => void;
  selectedRepoId?: number;
  loading: boolean;
}

const PLACEHOLDER_OPTION = 'Select a repository';
const LOADING_OPTION = 'Loading repositories...';
const NO_REPOSITORIES_OPTION = 'No repositories available';

const RepositoryDropdown: React.FC<RepositoryDropdownProps> = ({
    onChange,
    repositories,
    selectedRepoId,
    loading,
}) => {
    const handleChange = (event: ChangeEvent<HTMLSelectElement>) => {
        const selectedId = parseInt(event.target.value, 10);
        if (onChange && !isNaN(selectedId)) {
            onChange(selectedId);
        }
    };

    const renderOptions = () => {
        if (loading) {
            return (
                <option value="" disabled>
                    {LOADING_OPTION}
                </option>
            );
        }

        if (repositories.length === 0) {
            return (
                <option value="" disabled>
                    {NO_REPOSITORIES_OPTION}
                </option>
            );
        }

        return repositories.map((repo) => (
            <option key={repo.id} value={repo.id}>
                {repo.name}
            </option>
        ));
    };

    return (
        <select
            value={selectedRepoId ?? ''}
            onChange={handleChange}
        >
            <option value="" disabled>
                {PLACEHOLDER_OPTION}
            </option>
            {renderOptions()}
        </select>
    );
};

export default RepositoryDropdown;
