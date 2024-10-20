/* eslint-disable */

import React, { useEffect, useState } from "react";
import "./styles.css";
import { useNavigate } from "react-router-dom";
import OrganizationDropdown from "@/components/Dropdown/Organization";
import { Organization, OrganizationsResponse } from "@/types/organization";
import ClassroomDropdown from "@/components/Dropdown/Classroom";
import { Classroom, ClassroomResponse } from "@/types/classroom";
import { getClassrooms, getOrganizationDetails, getOrganizations, postSemester } from "@/api/requests";

enum SemesterCreationStatus {
    NONE = "NONE",
    CREATING = "CREATING",
    ERRORED = "ERRORED",
    CREATED = "CREATED"
}

const SemesterCreation: React.FC = () => {
    const [orgsWithApp, setOrgsWithApp] = useState<Organization[]>([]);
    const [orgsWithoutApp, setOrgsWithoutApp] = useState<Organization[]>([]);
    const [selectedOrg, setSelectedOrg] = useState<Organization | null>(null);

    const [availableClassrooms, setAvailableClassrooms] = useState<Classroom[]>([]);
    const [unavailableClassrooms, setUnavailableClassrooms] = useState<Classroom[]>([]);
    const [selectedClassroom, setSelectedClassroom] = useState<Classroom | null>(null);

    const [semesterCreationStatus, setSemesterCreationStatus] = useState(SemesterCreationStatus.NONE)

    const navigate = useNavigate();

    useEffect(() => {
        const fetchOrganizations = async () => {
            try {
                const data: OrganizationsResponse = await getOrganizations();
                setOrgsWithApp(data.orgs_with_app);
                setOrgsWithoutApp(data.orgs_without_app);
            } catch (error) {
                setOrgsWithApp([]);
                setOrgsWithoutApp([]);
                console.error("Error fetching organizations:", error);
            }
        };

        fetchOrganizations();
    }, []);

    useEffect(() => {
        if (selectedOrg) {
            const fetchClassrooms = async () => {
                try {
                    const data: ClassroomResponse = await getClassrooms(selectedOrg.id);
                    setAvailableClassrooms(data.available_classrooms);
                    setUnavailableClassrooms(data.unavailable_classrooms);
                } catch (error) {
                    setAvailableClassrooms([]);
                    setUnavailableClassrooms([]);
                    console.error("Error fetching classrooms:", error);
                }
            };

            fetchClassrooms();
        }
    }, [selectedOrg]);

    const handleOrganizationSelect = async (org: Organization) => {
        try {
            const data: Organization = await getOrganizationDetails(org.login);
            setSelectedOrg(data);
            setSelectedClassroom(null);
        } catch (error) {
            console.error("Error fetching organization details:", error);
        }
    };

    const handleClassroomSelect = async (classroom: Classroom) => {
        setSemesterCreationStatus(SemesterCreationStatus.NONE)
        setSelectedClassroom(classroom);
    };

    const handleCreateSemester = async () => {
        if (selectedOrg && selectedClassroom) {
            setSemesterCreationStatus(SemesterCreationStatus.CREATING)
            try {
                await postSemester(selectedOrg.id, selectedClassroom.id, `${selectedOrg.login}:${selectedClassroom.name}`);
                setSemesterCreationStatus(SemesterCreationStatus.CREATED)
                // Move selectedClassroom to unavailableClassrooms
                setAvailableClassrooms(prevAvailableClassrooms =>
                    prevAvailableClassrooms.filter(classroom => classroom.id !== selectedClassroom.id)
                );
                setUnavailableClassrooms(prevUnavailableClassrooms => [
                    ...prevUnavailableClassrooms,
                    selectedClassroom,
                ]);
            } catch (error) {
                setSemesterCreationStatus(SemesterCreationStatus.ERRORED)
                console.error("Error creating semester:", error);
            }
        }
    };

    return (
        <div className="SemesterCreation">
            <h1>Create a New Semester</h1>
            {semesterCreationStatus == SemesterCreationStatus.NONE && (
                <>
                    <OrganizationDropdown
                        orgsWithApp={orgsWithApp || []}
                        orgsWithoutApp={orgsWithoutApp || []}
                        selectedOrg={selectedOrg}
                        onSelect={handleOrganizationSelect}
                    />
                    {selectedOrg && orgsWithApp.find(org => org.id === selectedOrg.id) &&
                        (<ClassroomDropdown
                            availableClassrooms={availableClassrooms || []}
                            unavailableClassrooms={unavailableClassrooms || []}
                            selectedClassroom={selectedClassroom}
                            onSelect={handleClassroomSelect}
                        />)}

                </>
            )}
            {semesterCreationStatus == SemesterCreationStatus.CREATED && (
                <div>Semester successfully created!</div>
            )}
            <div>
                {selectedClassroom && selectedOrg &&
                    availableClassrooms.find(
                        classroom => classroom.id === selectedClassroom.id
                    ) && (
                        <>
                            {semesterCreationStatus === SemesterCreationStatus.CREATING && (
                                <button onClick={handleCreateSemester} disabled={true}>
                                    `Creating ${selectedClassroom.name}...`
                                </button>
                            )}
                            {(
                                semesterCreationStatus === SemesterCreationStatus.NONE ||
                                semesterCreationStatus === SemesterCreationStatus.ERRORED
                            ) && (
                                    <button onClick={handleCreateSemester}>
                                        {`Create Semester: "${selectedOrg.login}:${selectedClassroom.name}"`}
                                    </button>
                                )}
                            {semesterCreationStatus === SemesterCreationStatus.ERRORED && (
                                <div>
                                    Error creating semester. Please try again.
                                </div>
                            )}
                        </>
                    )}
            </div>
            <button onClick={() => { navigate("/semester-selection") }}> Go to Select Semester Page</button>
            {semesterCreationStatus !== SemesterCreationStatus.NONE && (
                <button onClick={() => { setSemesterCreationStatus(SemesterCreationStatus.NONE) }}> Create another semester</button>
            )}
        </div>
    );
};

export default SemesterCreation;