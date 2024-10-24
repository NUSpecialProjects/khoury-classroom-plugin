import React, { useEffect, useState } from "react";
import "./styles.css";
import { useNavigate } from "react-router-dom";
import OrganizationDropdown from "@/components/Dropdown/Organization";
import ClassroomDropdown from "@/components/Dropdown/Classroom";
import Panel from "@/components/Panel";
import Button from "@/components/Button";
import {
  getUserSemesters,
  getClassrooms,
  getOrganizationDetails,
  getOrganizations,
  postSemester,
} from "@/api/semesters";

enum SemesterCreationStatus {
  NONE = "NONE",
  CREATING = "CREATING",
  ERRORED = "ERRORED",
  CREATED = "CREATED",
}

const SemesterCreation: React.FC = () => {
  const [orgsWithApp, setOrgsWithApp] = useState<IOrganization[]>([]);
  const [orgsWithoutApp, setOrgsWithoutApp] = useState<IOrganization[]>([]);
  const [loadingOrganizations, setLoadingOrganizations] = useState(true);
  const [selectedOrg, setSelectedOrg] = useState<IOrganization | null>(null);

  const [availableClassrooms, setAvailableClassrooms] = useState<IClassroom[]>(
    []
  );
  const [unavailableClassrooms, setUnavailableClassrooms] = useState<
    IClassroom[]
  >([]);
  const [loadingClassrooms, setLoadingClassrooms] = useState(true);
  const [selectedClassroom, setSelectedClassroom] = useState<IClassroom | null>(
    null
  );

  const [semesterCreationStatus, setSemesterCreationStatus] = useState(
    SemesterCreationStatus.NONE
  );
 const [hasActiveSemester, setHasActiveSemesters] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchOrganizations = async () => {
      try {
        const data: IOrganizationsResponse = await getOrganizations();

        //checking if data exists before populating and setting lists
        if (data.orgs_with_app) {
          setOrgsWithApp(data.orgs_with_app);
        }
        if (data.orgs_without_app) {
          setOrgsWithoutApp(data.orgs_without_app);
        }
      } catch (error) {
        console.error("Error fetching organizations:", error);
      } finally {
        setLoadingOrganizations(false);
      }
    };

    void fetchOrganizations();
  }, []);

  useEffect(() => {
    if (selectedOrg) {
      const fetchClassrooms = async () => {
        try {
          const data: IClassroomResponse = await getClassrooms(selectedOrg.id);
          setAvailableClassrooms(data.available_classrooms);
          setUnavailableClassrooms(data.unavailable_classrooms);
        } catch (error) {
          setAvailableClassrooms([]);
          setUnavailableClassrooms([]);
          console.error("Error fetching classrooms:", error);
        } finally {
          setLoadingClassrooms(false);
        }
      };

      void fetchClassrooms();
    }
  }, [selectedOrg]);

  const handleOrganizationSelect = async (org: IOrganization) => {
    try {
      const data: IOrganization = await getOrganizationDetails(org.login);
      setSelectedOrg(data);
      setSelectedClassroom(null);
    } catch (error) {
      console.error("Error fetching organization details:", error);
    }
  };

  const handleClassroomSelect = async (classroom: IClassroom) => {
    setSemesterCreationStatus(SemesterCreationStatus.NONE);
    setSelectedClassroom(classroom);
  };

  const handleCreateSemester = async () => {
    if (selectedOrg && selectedClassroom) {
      setSemesterCreationStatus(SemesterCreationStatus.CREATING);
      try {
        await postSemester(
          selectedOrg.id,
          selectedClassroom.id,
          selectedOrg.login,
          selectedClassroom.name
        );
        setSemesterCreationStatus(SemesterCreationStatus.CREATED);
        // Move selectedClassroom to unavailableClassrooms
        setAvailableClassrooms((prevAvailableClassrooms) =>
          prevAvailableClassrooms.filter(
            (classroom) => classroom.id !== selectedClassroom.id
          )
        );
        setUnavailableClassrooms((prevUnavailableClassrooms) => [
          ...prevUnavailableClassrooms,
          selectedClassroom,
        ]);
      } catch (error) {
        setSemesterCreationStatus(SemesterCreationStatus.ERRORED);
        console.error("Error creating class:", error);
      }
    }
  };

  useEffect(() => {
    const fetchSemesters = async () => {
      try {
        const data: IUserSemestersResponse = await getUserSemesters();
        if(data.active_semesters.length !== 0) {
          setHasActiveSemesters(true);
        }
      } catch (error) {
        console.error("Error fetching semesters:", error);
      } 
    };

    void fetchSemesters();
  }, []);

  return (
    <Panel title="Create a New Classroom">
      {semesterCreationStatus == SemesterCreationStatus.NONE && (
        <>
          <OrganizationDropdown
            orgsWithApp={orgsWithApp}
            orgsWithoutApp={orgsWithoutApp}
            selectedOrg={selectedOrg}
            loading={loadingOrganizations}
            onSelect={handleOrganizationSelect}
          />
          {selectedOrg &&
            orgsWithApp.find((org) => org.id === selectedOrg.id) && (
              <ClassroomDropdown
                availableClassrooms={availableClassrooms}
                unavailableClassrooms={unavailableClassrooms}
                selectedClassroom={selectedClassroom}
                loading={loadingClassrooms}
                onSelect={handleClassroomSelect}
              />
            )}
        </>
      )}
      {semesterCreationStatus == SemesterCreationStatus.CREATED && (
        <div className="Creation__message">Class successfully created!</div>
      )}
      <div className="Creation__buttonWrapper">
        {selectedClassroom &&
          selectedOrg &&
          availableClassrooms.find(
            (classroom) => classroom.id === selectedClassroom.id
          ) && (
            <>
              {semesterCreationStatus === SemesterCreationStatus.CREATING && (

                <Button variant="primary" onClick={handleCreateSemester} disabled={true}>Creating classroom...</Button>
                // <button onClick={handleCreateSemester} disabled={true}>
                //   `Creating ${selectedClassroom.name}...`
                // </button>
              )}
              {(semesterCreationStatus === SemesterCreationStatus.NONE ||
                semesterCreationStatus === SemesterCreationStatus.ERRORED) && (
                  <Button variant="primary" onClick={handleCreateSemester}>Create classroom</Button>
                  // <button onClick={handleCreateSemester}>
                  //   {`Create Class: "${selectedOrg.login}:${selectedClassroom.name}"`}
                  // </button>
                )}
              {semesterCreationStatus === SemesterCreationStatus.ERRORED && (
                <div>Error creating classroom. Please try again.</div>
              )}
            </>
          )
        }

        {hasActiveSemester === true &&
          <Button variant="secondary" onClick={() => {
            navigate("/class-selection");
          }}>View existing classrooms</Button>
        }

        {semesterCreationStatus !== SemesterCreationStatus.NONE && (
          <Button variant="secondary" onClick={() => {
            setSemesterCreationStatus(SemesterCreationStatus.NONE);
          }}>Create another classroom</Button>
        )}
      </div>
    </Panel >
  );
};

export default SemesterCreation;
