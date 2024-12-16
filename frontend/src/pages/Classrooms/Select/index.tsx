import React, { useContext, useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Link } from "react-router-dom";
import { getClassroomsInOrg } from "@/api/classrooms";
import { Table, TableRow, TableCell } from "@/components/Table";
import EmptyDataBanner from "@/components/EmptyDataBanner";
import Button from "@/components/Button";
import { MdAdd } from "react-icons/md";
import { OrgRole } from "@/types/users";
import { useQuery } from "@tanstack/react-query";
import LoadingSpinner from "@/components/LoadingSpinner";

const ClassroomSelection: React.FC = () => {
  const location = useLocation();
  const { setSelectedClassroom } = useContext(SelectedClassroomContext);
  const navigate = useNavigate();
  const orgID = location.state?.orgID;

  useEffect(() => {
    if (!orgID) {
      console.log("No organization ID provided. Redirecting to organization selection.");
      navigate("/app/organization/select");
    }
  }, [orgID, navigate]);

  const { data, isLoading, error } = useQuery({
    queryKey: ['classrooms', orgID],
    queryFn: async () => {
      if (!orgID || isNaN(Number(orgID))) {
        throw new Error("Invalid organization ID");
      }
      return getClassroomsInOrg(Number(orgID));
    },
    enabled: !!orgID && !isNaN(Number(orgID)),
    retry: false
  });

  const classrooms = data?.classroom_users || [];
  const orgRole = data?.org_role || OrgRole.MEMBER;

  const handleClassroomSelect = (classroomUser: IClassroomUser) => {
    const classroom: IClassroom = {
      id: classroomUser.classroom_id,
      name: classroomUser.classroom_name,
      org_id: classroomUser.org_id,
      org_name: classroomUser.org_name,
    }
    setSelectedClassroom(classroom);
    navigate(`/app/dashboard`);
  };

  const hasClassrooms = classrooms.length > 0;

  return (
    <div className="Selection">
      <h1 className="Selection__title">Your Classrooms</h1>
      <div className="Selection__tableWrapper">
        {isLoading ? (
          <LoadingSpinner />
        ) : error ? (
          <p>Error loading classrooms: {error instanceof Error ? error.message : "Unknown error"}</p>
        ) : (
          <>
          <Table cols={2}>
            <TableRow>
              <TableCell>
                <div className="Selection__tableHeaderText">Classroom Name</div>
              </TableCell>
              <TableCell>
                {orgRole === OrgRole.ADMIN &&
                  (<div className="Selection__tableHeaderButton">
                    <Button 
                      size="small" 
                      onClick={() => navigate('/app/classroom/create', { state: { orgID } })}
                    >
                        <MdAdd /> New Classroom
                    </Button>
                  </div>
                )}
              </TableCell>
            </TableRow>
            {hasClassrooms && (
              classrooms.map((classroomUser, i) => (
                <TableRow key={i} className="Selection__tableRow">
                  <TableCell>
                    <div key={classroomUser.classroom_id} onClick={() => handleClassroomSelect(classroomUser)}>
                      {classroomUser.classroom_name}
                    </div>
                  </TableCell>
                  <TableCell>
                    <div>{classroomUser.classroom_role}</div>
                  </TableCell>
                </TableRow>
              ))
            )}
          </Table>
          {!hasClassrooms && (
            orgRole === OrgRole.ADMIN ? 
            (
              <TableRow className="Selection__tableRow">
                 <EmptyDataBanner>
                   <div className="emptyDataBannerMessage">
                      You have no classes in this organization.
                      <br></br>
                      Please create a new classroom to get started.
               </div>
               <Button variant="secondary"
                  onClick={() => navigate('/app/classroom/create', { state: { orgID } })}>
                   <MdAdd /> New Classroom
                 </Button>
                 </EmptyDataBanner>
       
              </TableRow>
            ) : (
              <TableRow className="Selection__tableRow">
                  <EmptyDataBanner>
                    You have no classes in this organization.
                    Your professor will need to invite you to a classroom.
                  </EmptyDataBanner>
              </TableRow>
            )
          )}
          <br></br>
          </>
        )}
      </div>
      <div className="Selection__linkWrapper">
        <Link to={`/app/organization/select`}>
          {" "}
          Choose a different organization →
        </Link>
      </div>
    </div>
  );
};

export default ClassroomSelection;
