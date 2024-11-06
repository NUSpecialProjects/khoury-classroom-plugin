import React from 'react';

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
    const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedId = parseInt(event.target.value, 10);
        if (onChange) {
            onChange(selectedId);
        }
    };

    return (
        <select value={selectedRepoId ?? ''} onChange={handleChange}>
            <option value="" disabled>
                Select a repository
            </option>
            {loading ? (
                <option className="Dropdown__option" value="" disabled>
                    Loading templates...
                </option>
            ) : (
                repositories.length > 0 ? (
                    repositories.map((repo) => (
                        <option key={repo.id} value={repo.id}>
                            {repo.name}
                        </option>
                    ))
                ) : (
                    <option value="" disabled>
                        No repositories available
                    </option>
                ))}
        </select>
    );
};

export default RepositoryDropdown;
