import Button from "@/components/Button";

import "./styles.css";
import { useLocation, useParams } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Table, TableCell, TableRow } from "@/components/Table";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import { getAssignmentIndirectNav, postAssignmentToken } from "@/api/assignments";
import { getStudentWorks } from "@/api/student_works";
import { formatDate } from "@/utils/date";
import CopyLink from "@/components/CopyLink";

import { Chart as ChartJS, registerables, Tooltip } from "chart.js";
import { Line } from 'react-chartjs-2'
import ChartDataLabels from "chartjs-plugin-datalabels";

ChartJS.register(...registerables);
ChartJS.register(ChartDataLabels);

const Assignment: React.FC = () => {
  const location = useLocation();
  const [assignment, setAssignment] = useState<IAssignmentOutline>();
  const [studentWorks, setStudentAssignment] = useState<IStudentWork[]>([]);
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const { id } = useParams();
  const [inviteLink, setInviteLink] = useState<string>("");
  const [linkError, setLinkError] = useState<string | null>(null);
  const base_url: string = import.meta.env.VITE_PUBLIC_FRONTEND_DOMAIN as string;

  const lineOptions = {
    responsive: true,
    plugins: {
      legend: {
        display: false,
      },
      title: {
        display: true,
        text: 'Commits over time',
      },
      datalabels: {
        display: false,
      },
    },
    scales: {
      x: {
        grid: {
          display: false,
        },
      },
      y: {
        grid: {
          display: false, 
        },
        ticks: {
          maxTicksLimit: 5, 
        },
      },
    },
    elements: {
      point: {
        radius: 1,
      },
      labels: {
        display: false
      }
    },
  };

  const lineTempData = {
    labels: ['5/10', '5/11', '5/12', '5/13', '5/14', '5/15', '5/16', '5/17', '5/18', '5/19', '5/20',],
    datasets: [
      {
        data: [0, 0, 1, 2, 7, 2, 121, 2, 0, 0, 2],
        borderColor: 'rgba(244, 63, 94, 1)',
        backgroundColor: 'rgba(75, 192, 192, 0.2)',
        tension: 0.05,
      },
    ],
  };

  useEffect(() => {
    // check if assignment has been passed through
    if (location.state) {
      setAssignment(location.state.assignment);
      const a: IAssignmentOutline = location.state.assignment;

      // sync student assignments
      if (selectedClassroom !== null && selectedClassroom !== undefined) {
        (async () => {
          try {
            const studentWorks = await getStudentWorks(
              selectedClassroom.id,
              a.id
            );
            if (studentWorks !== null && studentWorks !== undefined) {
              setStudentAssignment(studentWorks);
            }
          } catch (_) {
            // do nothing
          }
        })();
      }
    } else {
      // fetch the assignment from backend
      if (id && selectedClassroom !== null && selectedClassroom !== undefined) {
        (async () => {
          try {
            const fetchedAssignment = await getAssignmentIndirectNav(
              selectedClassroom.id,
              +id
            );
            if (fetchedAssignment !== null && fetchedAssignment !== undefined) {
              setAssignment(fetchedAssignment);
              const studentWorks = await getStudentWorks(
                selectedClassroom.id,
                fetchedAssignment.id
              );
              if (studentWorks !== null && studentWorks !== undefined) {
                setStudentAssignment(studentWorks);
              }
            }
          } catch (_) {
            // do nothing
          }
        })();
      }
    }
  }, [selectedClassroom]);

  useEffect(() => {
    const generateInviteLink = async () => {
      if (!assignment) return;

      try {
        if (!selectedClassroom) return;
        const tokenData = await postAssignmentToken(selectedClassroom.id, assignment.id);
        const url = `${base_url}/app/token/assignment/accept?token=${tokenData.token}`;
        setInviteLink(url);
      } catch (_) {
        setLinkError("Failed to generate assignment invite link");
      }
    };


    generateInviteLink();
  }, [assignment]);

  return (
    <div className="Assignment">
      {assignment && (
        <>
          <SubPageHeader
            pageTitle={assignment.name}
            chevronLink={"/app/dashboard"}
          >
            <div className="Assignment__dates">
              <div className="Assignment__date">
                <div className="Assignment__date--title"> {"Released on:"}</div>
                {assignment.created_at
                  ? formatDate(new Date(assignment.created_at))
                  : "N/A"}
              </div>
              <div className="Assignment__date">
                <div className="Assignment__date--title"> {"Due Date:"}</div>
                {assignment.main_due_date
                  ? formatDate(new Date(assignment.main_due_date))
                  : "N/A"}
              </div>
            </div>
          </SubPageHeader>

          <div className="Assignment__externalButtons">
            <Button href="" variant="secondary" newTab>
              View in Github Classroom
            </Button>
            <Button href="" variant="secondary" newTab>
              View Starter Code
            </Button>
            <Button href="" variant="secondary" newTab>
              View Rubric
            </Button>
          </div>

          <h2>Assignment Link</h2>
          <CopyLink link={inviteLink} name="invite-assignment" />
          {linkError && <p className="error">{linkError}</p>}

          <div className="Assignment__subSectionWrapper">
            <h2>Metrics</h2>
            <p>Metrics go here</p>

            <div className="Assignment__metricsWrapper">
              <Line
                options={lineOptions}
                data={lineTempData}
              />
            </div>



          </div>

          <div className="Assignment__subSectionWrapper">
            <h2 style={{ marginBottom: 0 }}>Student Assignments</h2>
            <Table cols={3}>
              <TableRow style={{ borderTop: "none" }}>
                <TableCell>Student Name</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Last Commit</TableCell>
              </TableRow>
              {studentWorks &&
                studentWorks.length > 0 &&
                studentWorks.map((sa, i) => (
                  <TableRow key={i} className="Assignment__submission">
                    <TableCell>{sa.contributors.join(", ")}</TableCell>
                    <TableCell>Passing</TableCell>
                    <TableCell>12 Sep, 11:34pm</TableCell>
                  </TableRow>
                ))}
            </Table>
          </div>
        </>
      )}
    </div>
  );
};

export default Assignment;