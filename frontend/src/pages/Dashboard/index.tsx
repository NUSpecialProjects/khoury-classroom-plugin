import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table";
import { Link, useNavigate } from "react-router-dom";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useEffect, useState, useContext } from "react";
import { getAssignments } from "@/api/assignments";
import { formatDate } from "@/utils/date";

const Dashboard: React.FC = () => {
  const [assignments, setAssignments] = useState<IAssignment[]>([]);
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAssignments = async (classroom: IClassroom) => {
      if (classroom) {
        getAssignments(classroom.id)
          .then((assignments) => {
            setAssignments(assignments);
          })
          .catch((err: unknown) => {
            console.error("Error fetching assignments:", err);
          });
      }
    };

    const SyncWithClassroom = async (classroom: IClassroom) => {
      try {
        console.log("Using mocked API call for classroom: ", classroom);
        // const base_url: string = import.meta.env
        //   .VITE_PUBLIC_API_DOMAIN as string;
        // const result = await fetch(`${base_url}/github/sync`, {
        //   method: "POST",
        //   credentials: "include",
        //   headers: {
        //     "Content-Type": "application/json",
        //   },
        //   body: JSON.stringify({ classroom_id: classroom.classroom_id }),
        // });
        const result = await Promise.resolve({ ok: true });

        if (!result.ok) {
          throw new Error("Network response was not ok");
        }
      } catch (error: unknown) {
        console.error("Error making API call:", error);
      }
    };

    if (selectedClassroom !== null && selectedClassroom !== undefined) {
      SyncWithClassroom(selectedClassroom)
        .then(() => {
          fetchAssignments(selectedClassroom).catch((error: unknown) => {
            console.log("Error fetching:", error);
          });
        })
        .catch((error: unknown) => {
          console.error("Error syncing:", error);
        });
    }
  }, [selectedClassroom]);

  const handleUserGroupClick = (group: string) => {
    console.log(`Clicked on ${group}`);
    if (group === "Professor") {
      navigate("/app/professors");
    }
    if (group === "TA") {
      navigate("/app/tas");
    }
    if (group === "Student") {
      navigate("/app/students");
    }
  };

  return (
    <div className="Dashboard">
      {selectedClassroom && (
        <>
          <h1>{selectedClassroom.org_name + " - " + selectedClassroom.name}</h1>
          <div className="Dashboard__classroomDetailsWrapper">
            <UserGroupCard
              label="Professors"
              role_type="Professor"
              classroom={selectedClassroom}
              onClick={() => handleUserGroupClick("Professor")}
            />

            <UserGroupCard
              label="TAs"
              role_type="TA"
              classroom={selectedClassroom}
              onClick={() => handleUserGroupClick("TA")}
            />

            <UserGroupCard
              label="Students"
              role_type="Student"
              classroom={selectedClassroom}
              onClick={() => handleUserGroupClick("Student")}
            />
          </div>
          <div className="Dashboard__assignmentsWrapper">
            <h2 style={{ marginBottom: 0 }}>Assignments</h2>
            <Table cols={2}>
              <TableRow style={{ borderTop: "none" }}>
                <TableCell>Assignment Name</TableCell>
                <TableCell>Due Date</TableCell>
              </TableRow>
              {assignments.map((assignment, i: number) => (
                <TableRow key={i} className="Assignment__submission">
                  <TableCell>
                    {" "}
                    <Link
                      to={`/app/assignments/${i + 1}`}
                      className="Dashboard__assignmentLink"
                    >
                      {assignment.name}
                    </Link>
                  </TableCell>
                  <TableCell>{formatDate(assignment.main_due_date)}</TableCell>
                </TableRow>
              ))}
            </Table>
          </div>
        </>
      )}
      <div className="Dashboard__linkWrapper">
        <Link to={`/app/classroom/select?org_id=${selectedClassroom?.id}`}>View other classrooms</Link>
        </div>
    </div>
  );
};

export default Dashboard;
