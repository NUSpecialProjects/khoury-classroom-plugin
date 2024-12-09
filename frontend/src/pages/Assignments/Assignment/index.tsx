import { useLocation, useParams } from "react-router-dom";
import { MdEdit, MdEditDocument } from "react-icons/md";
import { FaGithub } from "react-icons/fa";
import { useContext, useEffect, useState } from "react";
import { Chart as ChartJS, registerables } from "chart.js";
import { Bar, Doughnut } from "react-chartjs-2";
import ChartDataLabels from "chartjs-plugin-datalabels";

import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import {
  getAssignmentIndirectNav,
  postAssignmentToken,
} from "@/api/assignments";
import { getStudentWorks } from "@/api/student_works";
import { formatDateTime } from "@/utils/date";

import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import CopyLink from "@/components/CopyLink";
import { Table, TableCell, TableRow } from "@/components/Table";
import Button from "@/components/Button";
import MetricPanel from "@/components/Metrics/MetricPanel";
import SimpleMetric from "@/components/Metrics/SimpleMetric";

import "./styles.css";

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
  const base_url: string = import.meta.env
    .VITE_PUBLIC_FRONTEND_DOMAIN as string;

  const data1 = {
    labels: ["Committed", "Accepted", "Not Accepted"],
    datasets: [
      {
        backgroundColor: ["#219386", "#f69c0e", "#f83b5c"],
        data: [29, 5, 4],
      },
    ],
  };

  const data2 = {
    labels: ["Graded", "Ungraded"],
    datasets: [
      {
        backgroundColor: ["#219386", "#e5e7eb"],
        data: [3000, 700],
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
        const tokenData = await postAssignmentToken(
          selectedClassroom.id,
          assignment.id
        );
        const url = `${base_url}/app/token/assignment/accept?token=${tokenData.token}`;
        setInviteLink(url);
      } catch (_) {
        setLinkError("Failed to generate assignment invite link");
      }
    };

    generateInviteLink();
  }, [assignment]);

  return (
    assignment && (
      <>
        <SubPageHeader
          pageTitle={assignment.name}
          chevronLink={"/app/dashboard"}
        >
          <div className="Assignment__dates">
            <div className="Assignment__date">
              <div className="Assignment__date--title"> {"Released on:"}</div>
              {assignment.created_at
                ? formatDateTime(new Date(assignment.created_at))
                : "N/A"}
            </div>
            <div className="Assignment__date">
              <div className="Assignment__date--title"> {"Due Date:"}</div>
              {assignment.main_due_date
                ? formatDateTime(new Date(assignment.main_due_date))
                : "N/A"}
            </div>
          </div>
        </SubPageHeader>

        <div className="Assignment">
          <div className="Assignment__externalButtons">
            <Button href="#" variant="secondary" newTab>
              <FaGithub className="icon" /> View Template Repository
            </Button>
            <Button
              href={`/app/assignments/${assignment.id}/rubric`}
              variant="secondary"
              state={{ assignment }}
            >
              <MdEditDocument className="icon" /> View Rubric
            </Button>
            <Button href="#" variant="secondary" newTab>
              <MdEdit className="icon" /> Edit Assignment
            </Button>
          </div>

          <div className="Assignment__link">
            <h2>Assignment Link</h2>
            <CopyLink link={inviteLink} name="invite-assignment" />
            {linkError && <p className="error">{linkError}</p>}
          </div>
          <div className="Assignment__link">
            <h2>Assignment Link</h2>
            <CopyLink link={inviteLink} name="invite-assignment" />
            {linkError && <p className="error">{linkError}</p>}
          </div>

          <div className="Assignment__metrics">
            <h2>Metrics</h2>
            <MetricPanel>
              <SimpleMetric
                metricTitle="First Commit Date"
                metricValue="6 Sep"
              ></SimpleMetric>
              <SimpleMetric
                metricTitle="Total Commits"
                metricValue="941"
              ></SimpleMetric>
              <SimpleMetric
                metricTitle="Extension  Requests"
                metricValue="0"
              ></SimpleMetric>
              <SimpleMetric
                metricTitle="Regrade  Requests"
                metricValue="0"
              ></SimpleMetric>
            </MetricPanel>
            <div className="Assignment__metricsCharts">
              <div className="Assignment__metricsChart">
                <Bar
                  data={data1}
                  options={{
                    maintainAspectRatio: false,
                    indexAxis: "y",
                    layout: {
                      padding: {
                        right: 50,
                      },
                    },
                    scales: {
                      x: {
                        display: false,
                      },
                      y: {
                        grid: {
                          display: false,
                        },
                        ticks: {
                          font: {
                            size: 10,
                          },
                        },
                      },
                    },
                    plugins: {
                      legend: {
                        display: false,
                      },
                      datalabels: {
                        align: "end",
                        anchor: "end",
                        color: "#000",
                        font: {
                          size: 10,
                        },
                      },
                      tooltip: {
                        enabled: false,
                      },
                    },
                  }}
                />
              </div>
              <div className="Assignment__metricsChart">
                <Doughnut
                  data={data2}
                  options={{
                    maintainAspectRatio: false,
                    plugins: {
                      legend: {
                        onClick: () => {},
                        display: true,
                        position: "bottom",
                        labels: {
                          usePointStyle: true,
                          font: {
                            size: 10,
                          },
                        },
                      },
                      datalabels: {
                        color: ["#fff", "#000"],
                        font: {
                          size: 10,
                        },
                      },
                      tooltip: {
                        enabled: false,
                      },
                    },
                    cutout: "50%",
                  }}
                />
              </div>
            </div>
          </div>

          <div>
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
        </div>
      </>
    )
  );
};

export default Assignment;
