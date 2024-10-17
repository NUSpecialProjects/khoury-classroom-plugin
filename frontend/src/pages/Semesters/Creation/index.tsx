import React, { useEffect, useState } from "react";
import "./styles.css";
import { useNavigate } from "react-router-dom";

interface Organization {
    login: string;
    id: number;
    html_url: string;
    name: string;
    avatar_url: string;
}

interface Classroom {
    id: number;
    name: string;
    url: string;
}

interface OrganizationsResponse {
    orgs_with_app: Organization[];
    orgs_without_app: Organization[];
}

interface ClassroomResponse {
    available_classrooms: Classroom[];
    unavailable_classrooms: Classroom[];
}

const SemesterCreation: React.FC = () => {
    const [orgsWithApp, setOrgsWithApp] = useState<Organization[]>([]);
    const [orgsWithoutApp, setOrgsWithoutApp] = useState<Organization[]>([]);
    const [selectedOrg, setSelectedOrg] = useState<Organization | null>(null);

    const [availableClassrooms, setAvailableClassrooms] = useState<Classroom[]>([]);
    const [unavailableClassrooms, setUnavailableClassrooms] = useState<Classroom[]>([]);
    const [selectedClassroom, setSelectedClassroom] = useState<Classroom | null>(null);

    const [isCreatingSemester, setIsCreatingSemester] = useState(false);
    const [semesterCreated, setSemesterCreated] = useState(false);

    const navigate = useNavigate();

    useEffect(() => {
        const fetchOrganizations = async () => {
            try {
                const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                const response = await fetch(`${base_url}/github/user/orgs`, {
                    method: "GET",
                    credentials: 'include',
                    headers: {
                        "Content-Type": "application/json",
                    },
                });
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                const data: OrganizationsResponse = await response.json() as OrganizationsResponse;
                setOrgsWithApp(data.orgs_with_app);
                setOrgsWithoutApp(data.orgs_without_app);
            } catch (error) {
                console.error("Error fetching organizations:", error);
            }
        };

        void fetchOrganizations();
    }, []);

    useEffect(() => {
        if (selectedOrg) {
            const fetchClassrooms = async () => {
                try {
                    const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                    const response = await fetch(`${base_url}/github/user/orgs/${selectedOrg.id}/classrooms`, {
                        method: "GET",
                        credentials: 'include',
                        headers: {
                            "Content-Type": "application/json",
                        },
                    });
                    if (!response.ok) {
                        throw new Error("Network response was not ok");
                    }
                    const data: ClassroomResponse = await response.json() as ClassroomResponse;
                    setAvailableClassrooms(data.available_classrooms);
                    setUnavailableClassrooms(data.unavailable_classrooms);
                } catch (error) {
                    setAvailableClassrooms([]);
                    setUnavailableClassrooms([]);
                    console.error("Error fetching classrooms:", error);
                }
            };

            void fetchClassrooms();
        }
    }, [selectedOrg]);

    const handleOrganizationSelect = async (org: Organization) => {
        try {
            const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
            const response = await fetch(`${base_url}/github/orgs/${org.login}`, {
                method: "GET",
                credentials: 'include',
                headers: {
                    "Content-Type": "application/json",
                },
            });
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const resp = await response.json();
            const data: Organization = resp.org as Organization;
            setSelectedOrg(data);
            setSelectedClassroom(null);
        } catch (error) {
            console.error("Error fetching organization details:", error);
        }
    };

    const handleClassroomSelect = (classroom: Classroom) => {
        setSelectedClassroom(classroom);
    };

    const handleCreateSemester = async () => {
        if (selectedOrg && selectedClassroom) {
            setIsCreatingSemester(true);
            try {
                const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                const response = await fetch(`${base_url}/github/semesters`, {
                    method: "POST",
                    credentials: 'include',
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        org_id: selectedOrg.id,
                        classroom_id: selectedClassroom.id,
                        name: selectedClassroom.name,
                    }),
                });
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                setSemesterCreated(true);
            } catch (error) {
                console.error("Error creating semester:", error);
            } finally {
                setIsCreatingSemester(false);
            }
        }
    };

    return (
        <div className="SemesterCreation">
            <h1>Create a New Semester</h1>
            <div className="form-group">
                <label htmlFor="organization">Select Organization:</label>
                <select
                    id="organization"
                    value={selectedOrg ? selectedOrg.login : ""}
                    onChange={(e) => {
                        const selectedLogin = e.target.value;
                        if (selectedLogin === "create_new_org") {
                            window.open("https://github.com/organizations/plan", "_blank");
                        } else {
                            const selected = [...orgsWithApp, ...orgsWithoutApp].find(org => org.login === selectedLogin);
                            if (selected) {
                                handleOrganizationSelect(selected);
                            }
                        }
                    }}
                >
                    <option value="" disabled>Select an organization</option>
                    {orgsWithApp.length > 0 && (
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
                    {orgsWithoutApp.length > 0 && (
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
                    <option value="create_new_org">Create a New Organization ➕</option>
                </select>
                {selectedOrg && orgsWithoutApp.some(org => org.login === selectedOrg.login) && (
                    <a href={selectedOrg.html_url} target="_blank" rel="noopener noreferrer">
                        <button>Install GitGrader on {selectedOrg.login}</button>
                    </a>
                )}
            </div>
            {selectedOrg && orgsWithApp.find(org => org.login === selectedOrg.login) && (
                <div className="form-group">
                    <label htmlFor="classroom">Select Classroom:</label>
                    <select
                        id="classroom"
                        value={selectedClassroom ? selectedClassroom.id : ""}
                        onChange={(e) => {
                            const selectedId = Number(e.target.value);
                            if (selectedId === -1) {
                                window.open("https://classroom.github.com/classrooms/new", "_blank");
                            } else {
                                const selected = [...availableClassrooms, ...unavailableClassrooms].find(classroom => classroom.id === selectedId);
                                if (selected) {
                                    handleClassroomSelect(selected);
                                }
                            }
                        }}
                    >
                        <option value="" disabled>Select a classroom</option>
                        {availableClassrooms.length > 0 && (
                            <optgroup label="Available Classrooms">
                                {availableClassrooms.map((classroom) => (
                                    <option key={classroom.id} value={classroom.id}>
                                        {classroom.name} ✔️
                                    </option>
                                ))}
                            </optgroup>
                        )}
                        {unavailableClassrooms.length > 0 && (
                            <optgroup label="Unavailable Classrooms">
                                {unavailableClassrooms.map((classroom) => (
                                    <option key={classroom.id} value={classroom.id}>
                                        {classroom.name} ❌
                                    </option>
                                ))}
                            </optgroup>
                        )}
                        <option value="-1">Create New Classroom ➕</option>
                    </select>
                </div>
            )}
            <div>
                {selectedClassroom && selectedOrg && !semesterCreated &&
                    availableClassrooms.find(
                        classroom => classroom.id === selectedClassroom.id
                    ) && (
                        <button onClick={handleCreateSemester} disabled={isCreatingSemester}>
                            {isCreatingSemester ? `Creating ${selectedClassroom.name}...` : `Create Semester: \"${selectedOrg.login}:${selectedClassroom.name}\"`}
                        </button>
                    )}
                {semesterCreated && (
                    <button onClick={() => console.log("Navigate to select semester screen")}>
                        Go to Select Semester
                    </button>
                )}
            </div>
            <button onClick={() => navigate("/semester-selection")}> Go to Select Semester Page</button>
        </div>
    );
};

export default SemesterCreation;