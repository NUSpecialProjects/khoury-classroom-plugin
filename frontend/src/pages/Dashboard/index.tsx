import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table";
import { Link } from "react-router-dom";
import useSelectedSemester from "@/contexts/useSelectedSemester";
import AlertBanner from "@/components/Banner/AlertBanner";
import { activateSemester, deactivateSemester } from "@/api/semesters";
import { useEffect, useState } from "react";
import ErrorMessage from "@/components/Error";




const Dashboard: React.FC = () => {
  const [assignments, setAssignments] = useState<IAssignment[]>([]);
  const { selectedSemester, setSelectedSemester } = useSelectedSemester();
  const [error, setError] = useState<string | null>(null);

  const options: Intl.DateTimeFormatOptions = {
    weekday: 'short', year: 'numeric', month: 'short', day: 'numeric',
    hour: '2-digit', minute: '2-digit', timeZoneName: 'short'
  };

  const handleActivate = async (newSemester: ISemester) => {
    setSelectedSemester(newSemester);
  };

  const handleActivateClick = async () => {
    if (selectedSemester) {
      try {
        const newSemester = await activateSemester(
          selectedSemester.classroom_id
        );
        handleActivate(newSemester);
        setError(null);
      } catch (err) {
        console.log(err);
        setError("Failed to activate the class. Please try again.");
      }
    }
  };

  const handleDeactivateClick = async () => {
    if (selectedSemester) {
      try {
        const newSemester = await deactivateSemester(
          selectedSemester.classroom_id
        );
        handleActivate(newSemester);
        setError(null);
      } catch (err) {
        console.log(err);
        setError("Failed to deactivate the class. Please try again.");
      }
    }
  };

  useEffect(() => {
    const fetchAssignments = async (semester: ISemester) => {
      try {
        if (semester.classroom_id) {
          const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
          const result = await fetch(`${base_url}/assignments/${semester.classroom_id}`, {
            method: 'GET',
            credentials: 'include',
            headers: {
              'Content-Type': 'application/json',
            },
          });

          if (!result.ok) {
            throw new Error('Network response was not ok');
          }

          const data: IAssignment[] = (await result.json() as IAssignment[])
          const assignmentGoodDate = data.map((assignment: IAssignment) => ({
            ...assignment,
            main_due_date: assignment.main_due_date ? new Date(assignment.main_due_date) : null,
          }))
          console.log("Setting Assignment data: ", assignmentGoodDate)
          setAssignments(assignmentGoodDate);
        }

      } catch (error: unknown) {
        console.error('Error fetching assignments:', error);
      }
    };

    const SyncWithClassroom = async (semester: ISemester) => {
      try {
        const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
        const result = await fetch(`${base_url}/github/sync`, {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ classroom_id: semester.classroom_id }),
        })

        if (!result.ok) {
          throw new Error('Network response was not ok');
        }

      } catch (error: unknown) {
        console.error('Error making API call:', error);
      }
    };

    console.log("We in dashboard: ", selectedSemester)
    if (selectedSemester !== null && selectedSemester !== undefined) {
      SyncWithClassroom(selectedSemester).then(() => {
        fetchAssignments(selectedSemester).catch((error: unknown) => {
          console.log("Error fetching:", error)
        })
      }).catch((error: unknown) => {
        console.error('Error syncing:', error);
      });
    }

  }, [selectedSemester]);

  return (<div className="Dashboard">
    {selectedSemester && (
      <>
        <h1>
          {selectedSemester.org_name +
            " - " +
            selectedSemester.classroom_name}
        </h1>
        <AlertBanner
          semester={selectedSemester}
          onActivate={handleActivate}
        />
        <div className="Dashboard__classroomDetailsWrapper">
          <UserGroupCard label="Professors" number={1} />
          <UserGroupCard label="TAs" number={12} />
          <UserGroupCard label="Students" number={38} />
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
                <TableCell> <Link to={`/app/assignments/?id=${assignment.id}`} className="Dashboard__assignmentLink">{assignment.name}</Link></TableCell>
                <TableCell> {assignment.main_due_date ? assignment.main_due_date.toLocaleDateString("en-US", options) : "N/A"}</TableCell>
              </TableRow>
            ))}
          </Table>


        </div>
        <div>
          <p>Temporary Classroom Settings</p>
          {selectedSemester && !selectedSemester.active && (
            <button onClick={handleActivateClick}>Activate Class</button>
          )}
          {selectedSemester && selectedSemester.active && (
            <button onClick={handleDeactivateClick}>Deactivate Class</button>
          )}
        </div>
        {error && <ErrorMessage message={error} />}
      </>
    )}
  </div>
  );

};


export default Dashboard;