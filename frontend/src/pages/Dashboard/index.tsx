import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table";
import { MdAdd } from "react-icons/md";
import { Link, useNavigate } from "react-router-dom";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useEffect, useState, useContext } from "react";
import { getAssignments } from "@/api/assignments";
import { formatDateTime, formatDate } from "@/utils/date";
import { useClassroomUser } from "@/hooks/useClassroomUser";
import { useClassroomUsersList } from "@/hooks/useClassroomUsersList";
import BreadcrumbPageHeader from "@/components/PageHeader/BreadcrumbPageHeader";
import Button from "@/components/Button";
import ClipLoader from "react-spinners/ClipLoader";
import MetricPanel from "@/components/Metrics/MetricPanel";
import Metric from "@/components/Metrics";
import { ClassroomRole } from "@/types/users";

const Dashboard: React.FC = () => {
  const [assignments, setAssignments] = useState<IAssignmentOutline[]>([]);
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const {
    classroomUser,
    error: classroomUserError,
    loading: loadingCurrentClassroomUser,
  } = useClassroomUser(selectedClassroom?.id, ClassroomRole.TA, "/app/organization/select");
  const { classroomUsers: classroomUsersList } = useClassroomUsersList(
    selectedClassroom?.id
  );
  const navigate = useNavigate();

  const getGCD = (a: number, b: number): number => {
    while (b !== 0) {
      const temp = b;
      b = a % b;
      a = temp;
    }
    return a;
  };

  const getTaToStudentRatio = (users: IClassroomUser[]): string => {
    if (!users || users.length === 0) {
      return "N/A";
    }

    const tas = users.filter((user) => user.classroom_role === "TA");

    const students = users.filter((user) => user.classroom_role === "STUDENT");

    if (tas.length === 0 || students.length === 0) {
      return "N/A";
    }

    const taCount = tas.length;
    const studentCount = students.length;
    const gcd = getGCD(taCount, studentCount);

    const reducedTaCount = taCount / gcd;
    const reducedStudentCount = studentCount / gcd;

    return `${reducedTaCount} : ${reducedStudentCount}`;
  };

  useEffect(() => {
    const fetchAssignments = async (classroom: IClassroom) => {
      if (classroom) {
        getAssignments(classroom.id)
          .then((assignments) => {
            setAssignments(assignments);
          })
          .catch((_: unknown) => {
            // do nothing
          });
      }
    };

    if (selectedClassroom !== null && selectedClassroom !== undefined) {
      fetchAssignments(selectedClassroom).catch((_: unknown) => {
        // do nothing
      });
    }
  }, [selectedClassroom]);

  useEffect(() => {
    if (
      !loadingCurrentClassroomUser &&
      (classroomUserError || !classroomUser)
    ) {
      console.log(
        "Attempted to view a classroom without access. Redirecting to classroom select."
      );
      navigate(`/app/organization/select`);
    }
  }, [
    loadingCurrentClassroomUser,
    classroomUserError,
    classroomUser,
    selectedClassroom?.org_id,
    navigate,
  ]);

  const handleUserGroupClick = (group: string, users: IClassroomUser[]) => {
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

  if (loadingCurrentClassroomUser) {
    return (
      <div className="Dashboard__loading">
        <ClipLoader size={50} color={"#123abc"} loading={true} />
        <p>Loading classroom details...</p>
      </div>
    );
  }

  return (
    <div className="Dashboard">
      {selectedClassroom && (
        <>
          <BreadcrumbPageHeader
            pageTitle={selectedClassroom?.org_name}
            breadcrumbItems={[selectedClassroom?.name]}
          />

          <div className="Dashboard__sectionWrapper">
            <MetricPanel>
              <div className="Dashboard__classroomDetailsWrapper">
                <UserGroupCard
                  label="Students"
                  givenUsersList={classroomUsersList.filter(
                    (user) => user.classroom_role === ClassroomRole.STUDENT
                  )}
                  onClick={() =>
                    handleUserGroupClick(
                      "Student",
                      classroomUsersList.filter(
                        (user) => user.classroom_role === ClassroomRole.STUDENT
                      )
                    )
                  }
                />

                <UserGroupCard
                  label="TAs"
                  givenUsersList={classroomUsersList.filter(
                    (user) => user.classroom_role === ClassroomRole.TA
                  )}
                  onClick={() =>
                    handleUserGroupClick(
                      "TA",
                      classroomUsersList.filter(
                        (user) => user.classroom_role === ClassroomRole.TA
                      )
                    )
                  }
                />

                <UserGroupCard
                  label="Professors"
                  givenUsersList={classroomUsersList.filter(
                    (user) => user.classroom_role === ClassroomRole.PROFESSOR
                  )}
                  onClick={() =>
                    handleUserGroupClick(
                      "Professor",
                      classroomUsersList.filter(
                        (user) => user.classroom_role === ClassroomRole.PROFESSOR
                      )
                    )
                  }
                />
              </div>

              <Metric title="Created on">
                {formatDate(selectedClassroom.created_at ?? null)}
              </Metric>
              <Metric title="Assignments">
                {assignments.length.toString()}
              </Metric>
              <Metric title="TA to Student Ratio">
                {getTaToStudentRatio(classroomUsersList)}
              </Metric>
            </MetricPanel>
          </div>

          <div className="Dashboard__sectionWrapper">
            <div className="Dashboard__assignmentsHeader">
              <h2 style={{ marginBottom: 0 }}>Assignments</h2>
              <div className="Dashboard__createAssignmentButton">
                <Button
                  variant="secondary"
                  size="small"
                  href={`/app/assignments/create?org_name=${selectedClassroom?.org_name}`}
                >
                  <MdAdd className="icon" /> Create Assignment
                </Button>
              </div>
            </div>
            <Table cols={2}>
              <TableRow style={{ borderTop: "none" }}>
                <TableCell>Assignment Name</TableCell>
                <TableCell>Created Date</TableCell>
              </TableRow>
              {assignments.map((assignment, i: number) => (
                <TableRow key={i} className="Assignment__submission">
                  <TableCell>
                    <Link
                      to={`/app/assignments/${assignment.id}`}
                      state={{ assignment }}
                      className="Dashboard__assignmentLink"
                    >
                      {assignment.name}
                    </Link>
                  </TableCell>
                  <TableCell>{formatDateTime(assignment.created_at)}</TableCell>
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
