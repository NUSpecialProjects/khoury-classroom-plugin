import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table";
import { Link, useNavigate } from "react-router-dom";
import { SelectedSemesterContext } from "@/contexts/selectedClassroom";
import AlertBanner from "@/components/Banner/AlertBanner";
import { useEffect, useState, useContext } from "react";
import { getAssignments } from "@/api/assignments";
import { formatDate } from "@/utils/date";

const Dashboard: React.FC = () => {
  const [assignments, setAssignments] = useState<IAssignment[]>([]);
  const { selectedClassroom: selectedSemester, setSelectedSemester } = useContext(
    SelectedSemesterContext
  );
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAssignments = async (semester: ISemester) => {
      if (semester) {
        getAssignments(semester.classroom_id)
          .then((assignments) => {
            setAssignments(assignments);
          })
          .catch((err: unknown) => {
            console.error("Error fetching assignments:", err);
          });
      }
    };

    const SyncWithClassroom = async (semester: ISemester) => {
      try {
        console.log("Using mocked API call for semester: ", semester);
        // const base_url: string = import.meta.env
        //   .VITE_PUBLIC_API_DOMAIN as string;
        // const result = await fetch(`${base_url}/github/sync`, {
        //   method: "POST",
        //   credentials: "include",
        //   headers: {
        //     "Content-Type": "application/json",
        //   },
        //   body: JSON.stringify({ classroom_id: semester.classroom_id }),
        // });
        const result = await Promise.resolve({ ok: true });

        if (!result.ok) {
          throw new Error("Network response was not ok");
        }
      } catch (error: unknown) {
        console.error("Error making API call:", error);
      }
    };

    if (selectedSemester !== null && selectedSemester !== undefined) {
      SyncWithClassroom(selectedSemester)
        .then(() => {
          fetchAssignments(selectedSemester).catch((error: unknown) => {
            console.log("Error fetching:", error);
          });
        })
        .catch((error: unknown) => {
          console.error("Error syncing:", error);
        });
    }
  }, [selectedSemester]);

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
      {selectedSemester && (
        <>
          <h1>
            {selectedSemester.org_name +
              " - " +
              selectedSemester.classroom_name}
          </h1>
          <AlertBanner
            semester={selectedSemester}
            onActivate={setSelectedSemester}
          />
          <div className="Dashboard__classroomDetailsWrapper">
            <UserGroupCard
              label="Professors"
              role_type="Professor"
              semester={selectedSemester}
              onClick={() => handleUserGroupClick("Professor")}
            />

            <UserGroupCard
              label="TAs"
              role_type="TA"
              semester={selectedSemester}
              onClick={() => handleUserGroupClick("TA")}
            />

            <UserGroupCard
              label="Students"
              role_type="Student"
              semester={selectedSemester}
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
    </div>
  );
};

export default Dashboard;
