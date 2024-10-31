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

enum ClassroomCreationStatus {
  NONE = "NONE",
  CREATING = "CREATING",
  ERRORED = "ERRORED",
  CREATED = "CREATED",
}

const ClassroomCreation: React.FC = () => {
  const [orgsWithApp, setOrgsWithApp] = useState<IOrganization[]>([]);
  const [orgsWithoutApp, setOrgsWithoutApp] = useState<IOrganization[]>([]);
  const [loadingOrganizations, setLoadingOrganizations] = useState(true);
  const [selectedOrg, setSelectedOrg] = useState<IOrganization | null>(null);

  const [classroomCreationStatus, setClassroomCreationStatus] = useState(
    ClassroomCreationStatus.NONE
  );
  const [hasClassroom, setHasClassroom] = useState(false);
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
    setClassroomCreationStatus(ClassroomCreationStatus.NONE);
    setSelectedClassroom(classroom);
  };

  const handleCreateSemester = async () => {
    if (selectedOrg && selectedClassroom) {
      setClassroomCreationStatus(ClassroomCreationStatus.CREATING);
      try {
        await postSemester(
          selectedOrg.id,
          selectedClassroom.id,
          selectedOrg.login,
          selectedClassroom.name
        );
        setClassroomCreationStatus(ClassroomCreationStatus.CREATED);
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

        //Since a new semester has been created, flag this as true to render button
        setHasClassroom(true);
      } catch (error) {
        setClassroomCreationStatus(ClassroomCreationStatus.ERRORED);
        console.error("Error creating class:", error);
      }
    }
  };

  useEffect(() => {
    const fetchSemesters = async () => {
      try {
        const data: IUserSemestersResponse = await getUserSemesters();

        if (
          data.active_semesters.length > 0 ||
          data.inactive_semesters.length > 0
        ) {
          setHasClassroom(true);
          console.log(hasClassroom);
        }
      } catch (error) {
        console.error("Error fetching semesters:", error);
      }
    };

    void fetchSemesters();
  }, []);

  return (
    <Panel title="Create a New Classroom" logo={true}>
      {classroomCreationStatus == ClassroomCreationStatus.NONE && (
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
      {classroomCreationStatus == ClassroomCreationStatus.CREATED && (
        <div className="Creation__message">Class successfully created!</div>
      )}
      <div className="Creation__buttonWrapper">
        {selectedClassroom &&
          selectedOrg &&
          availableClassrooms.find(
            (classroom) => classroom.id === selectedClassroom.id
          ) && (
            <>
              {classroomCreationStatus === ClassroomCreationStatus.CREATING && (
                <Button
                  variant="primary"
                  onClick={handleCreateSemester}
                  disabled={true}
                >
                  Creating classroom...
                </Button>
              )}
              {(classroomCreationStatus === ClassroomCreationStatus.NONE ||
                classroomCreationStatus === ClassroomCreationStatus.ERRORED) && (
                <Button variant="primary" onClick={handleCreateSemester}>
                  Create classroom
                </Button>
              )}
              {classroomCreationStatus === ClassroomCreationStatus.ERRORED && (
                <div>Error creating classroom. Please try again.</div>
              )}
            </>
          )}

        {selectedOrg &&
          orgsWithoutApp.some((org) => org.login === selectedOrg.login) && (
            <Button variant="primary" href={selectedOrg.html_url}>
              Install GitGrader for {selectedOrg.login}
            </Button>
          )}

        {(hasClassroom ||
          classroomCreationStatus === ClassroomCreationStatus.CREATED) && (
          <Button
            variant="secondary"
            onClick={() => {
              navigate("/app/classroom/select");
            }}
          >
            View existing classrooms
          </Button>
        )}

        {classroomCreationStatus !== ClassroomCreationStatus.NONE && (
          <Button
            variant="secondary"
            onClick={() => {
              setClassroomCreationStatus(ClassroomCreationStatus.NONE);
            }}
          >
            Create another classroom
          </Button>
        )}
      </div>
    </Panel>
  );
};

export default ClassroomCreation;
