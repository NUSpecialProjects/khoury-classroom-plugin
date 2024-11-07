import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table";
import { Link, useNavigate } from "react-router-dom";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useEffect, useState, useContext } from "react";
import { getAssignments } from "@/api/assignments";
import { formatDate } from "@/utils/date";
import { useClassroomUser } from "@/hooks/useClassroomUser";
import { useClassroomUsersList } from "@/hooks/useClassroomUsersList";

const Dashboard: React.FC = () => {
  const [assignments, setAssignments] = useState<IAssignmentOutline[]>([]);
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const { classroomUser, loading: loadingCurrentClassroomUser } = useClassroomUser(selectedClassroom?.id);
  const { classroomUsers: classroomUsersList, loading: loadingClassroomUsersList } = useClassroomUsersList(selectedClassroom?.id);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAssignments = async (classroom: IClassroom) => {
      if (classroom) {
        getAssignments(classroom.id)
          .then((assignments) => {
            console.log("Assignments:", assignments);
            setAssignments(assignments);
          })
          .catch((err: unknown) => {
            console.error("Error fetching assignments:", err);
          });
      }
    };


    if (selectedClassroom !== null && selectedClassroom !== undefined) {
      fetchAssignments(selectedClassroom).catch((error: unknown) => {
        console.log("Error fetching:", error);
      });
    }
  }, [selectedClassroom]);

  const handleUserGroupClick = (group: string, users: IClassroomUser[]) => {
    console.log(`Clicked on ${group}`);
    if (group === "Professor") {
      navigate("/app/professors", { state: { users } });
    }
    if (group === "TA") {
      navigate("/app/tas", { state: { users } });
    }
    if (group === "Student") {
      navigate("/app/students", { state: { users } });
    }
  };

  return (
    <div className="Dashboard">
      {selectedClassroom && (
        <>
          <h1>{selectedClassroom.org_name + " - " + selectedClassroom.name}</h1>
          {!loadingCurrentClassroomUser && classroomUser && (
            <p>{"Viewing as a " + classroomUser.classroom_role}</p>
          )}
          {!loadingCurrentClassroomUser && !classroomUser && (
            <p>{"Viewing classroom you aren't in!! (Eventually, this should be impossible)"}</p>
          )}
          <div className="Dashboard__classroomDetailsWrapper">
            <UserGroupCard
              label="Professors"
              role_type="PROFESSOR"
              classroom={selectedClassroom}
              givenUsersList={classroomUsersList.filter((user) => user.classroom_role === "PROFESSOR")}
              onClick={() => handleUserGroupClick("Professor", classroomUsersList.filter((user) => user.classroom_role === "PROFESSOR"))}
            />

            <UserGroupCard
              label="TAs"
              role_type="TA"
              classroom={selectedClassroom}
              givenUsersList={classroomUsersList.filter((user) => user.classroom_role === "TA")}
              onClick={() => handleUserGroupClick("TA", classroomUsersList.filter((user) => user.classroom_role === "TA"))}
            />

            <UserGroupCard
              label="Students"
              role_type="STUDENT"
              classroom={selectedClassroom}
              givenUsersList={classroomUsersList.filter((user) => user.classroom_role === "STUDENT")}
              onClick={() => handleUserGroupClick("Student", classroomUsersList.filter((user) => user.classroom_role === "STUDENT"))}
            />
          </div>
          <div className="Dashboard__assignmentsWrapper">
            <h2 style={{ marginBottom: 0 }}>Assignments</h2>
            <Table cols={2}>
              <TableRow style={{ borderTop: "none" }}>
                <TableCell>Assignment Name</TableCell>
                <TableCell>Created Date</TableCell>
              </TableRow>
              {assignments.map((assignment, i: number) => (
                <TableRow key={i} className="Assignment__submission">
                  <TableCell> <Link to={`/app/assignments/${assignment.id}`} state={{assignment}} className="Dashboard__assignmentLink">{assignment.name}</Link></TableCell>
                  <TableCell>{formatDate(assignment.created_at)}</TableCell>
                </TableRow>
              ))}
            </Table>
          </div>
        </>
      )}
      <div className="Dashboard__linkWrapper">
        <Link to={`/app/classroom/select?org_id=${selectedClassroom?.org_id}`}>
          View other classrooms
        </Link>
      </div>
    </div>
  );
};

export default Dashboard;
