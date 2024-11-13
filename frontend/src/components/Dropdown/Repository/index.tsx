import React, { ChangeEvent } from 'react';

interface IRepositoryDropdownProps {
    repositories: IRepository[];
    onChange?: (selectedTemplate: IRepository) => void;
    selectedRepo: IRepository | null;
    loading: boolean;
}

const PLACEHOLDER_OPTION = 'Select a repository';
const LOADING_OPTION = 'Loading repositories...';
const NO_REPOSITORIES_OPTION = 'No repositories available';

const RepositoryDropdown: React.FC<IRepositoryDropdownProps> = ({
    onChange,
    repositories,
    selectedRepo,
    loading,
}) => {
    const handleChange = (event: ChangeEvent<HTMLSelectElement>) => {
        const selectedId = parseInt(event.target.value, 10);
        const selectedRepo = repositories.find((repo) => repo.id === selectedId);

        if (onChange && selectedRepo) {
            onChange(selectedRepo);
        }
    };

    const renderOptions = () => {
        if (loading) {
            return (
                <option value="" disabled>
                    {LOADING_OPTION}
                </option>
            );
        } else if (repositories.length === 0) {
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
            value={selectedRepo?.id ?? ''}
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
