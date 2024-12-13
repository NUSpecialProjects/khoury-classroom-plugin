import { useLocation, useParams, Link } from "react-router-dom";
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
  getAssignmentFirstCommit,
  getAssignmentTotalCommits,
} from "@/api/assignments";
import {
  getAssignmentAcceptanceMetrics,
  getAssignmentGradedMetrics,
} from "@/api/metrics";
import { getStudentWorks } from "@/api/student_works";
import { formatDate } from "@/utils/date";

import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import CopyLink from "@/components/CopyLink";
import { Table, TableCell, TableRow } from "@/components/Table";
import Button from "@/components/Button";
import MetricPanel from "@/components/Metrics/MetricPanel";
import Metric from "@/components/Metrics";
import { useQuery } from "@tanstack/react-query";

import "./styles.css";

ChartJS.register(...registerables);
ChartJS.register(ChartDataLabels);

const Assignment: React.FC = () => {
  const location = useLocation();
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const { id } = useParams();
  const base_url: string = import.meta.env.VITE_PUBLIC_FRONTEND_DOMAIN as string;
  const [firstCommit, setFirstCommit] = useState<string>("");
  const [totalCommits, setTotalCommits] = useState<string>();

  const { data: assignment } = useQuery({
    queryKey: ['assignment', selectedClassroom?.id, id],
    queryFn: async () => {
      if (!selectedClassroom?.id || !id) return null;
      if (location.state?.assignment) {
        return location.state.assignment;
      }
      return await getAssignmentIndirectNav(selectedClassroom.id, +id);
    },
    enabled: !!selectedClassroom?.id && !!id
  });

  const { data: studentWorks = [] } = useQuery({
    queryKey: ['studentWorks', selectedClassroom?.id, assignment?.id],
    queryFn: async () => {
      if (!selectedClassroom?.id || !assignment?.id) return [];
      return await getStudentWorks(selectedClassroom.id, assignment.id);
    },
    enabled: !!selectedClassroom?.id && !!assignment?.id
  });

  const { data: inviteLink = "", error: linkError } = useQuery({
    queryKey: ['assignmentToken', selectedClassroom?.id, assignment?.id],
    queryFn: async () => {
      if (!selectedClassroom?.id || !assignment?.id) return "";
      const tokenData = await postAssignmentToken(selectedClassroom.id, assignment.id);
      return `${base_url}/app/token/assignment/accept?token=${tokenData.token}`;
    },
    enabled: !!selectedClassroom?.id && !!assignment?.id
  });

  const [acceptanceMetrics, setAcceptanceMetrics] = useState<IChartJSData>({
    labels: ["Not Accepted", "Accepted", "Started", "Submitted", "In Grading"],
    datasets: [
      {
        backgroundColor: [
          "#f83b5c",
          "#50c878",
          "#fece5a",
          "#7895cb",
          "#219386",
        ],
        data: [],
      },
    ],
  });
  const [gradedMetrics, setGradedMetrics] = useState<IChartJSData>({
    labels: ["Graded", "Ungraded"],
    datasets: [
      {
        backgroundColor: ["#219386", "#e5e7eb"],
        data: [],
      },
    ],
  });

  useEffect(() => {
    if (!selectedClassroom || !id) return;

    // populate acceptance metrics
    getAssignmentAcceptanceMetrics(selectedClassroom.id, Number(id)).then(
      (metrics) => {
        acceptanceMetrics.datasets[0].data = [
          metrics.not_accepted,
          metrics.accepted,
          metrics.started,
          metrics.submitted,
          metrics.in_grading,
        ];
        setAcceptanceMetrics(acceptanceMetrics);
      }
    );

    // populate graded status metrics
    getAssignmentGradedMetrics(selectedClassroom.id, Number(id)).then(
      (metrics) => {
        gradedMetrics.datasets[0].data = [metrics.graded, metrics.ungraded];
        setGradedMetrics(gradedMetrics);
      }
    );
  }, [selectedClassroom]);

  useEffect(() => {
    if (
      assignment !== null &&
      assignment !== undefined &&
      selectedClassroom !== null &&
      selectedClassroom !== undefined
    ) {
      (async () => {
        try {
          const commitDate = await getAssignmentFirstCommit(
            selectedClassroom.id,
            assignment.id
          );
          if (commitDate !== null && commitDate !== undefined) {
            setFirstCommit(formatDate(commitDate));
          } else {
            setFirstCommit("N/A");
          }
        } catch (_) {
          // do nothing
        }
      })();
    }
  }, [selectedClassroom, assignment]);

  useEffect(() => {
    if (
      assignment !== null &&
      assignment !== undefined &&
      selectedClassroom !== null &&
      selectedClassroom !== undefined
    ) {
      (async () => {
        try {
          const total = await getAssignmentTotalCommits(
            selectedClassroom.id,
            assignment.id
          );
          if (totalCommits !== null && totalCommits !== undefined) {
            setTotalCommits(total.toString());
          } else {
            setTotalCommits("N/A");
          }
        } catch (_) {
          // do nothing
        }
      })();
    }
  }, [selectedClassroom, assignment]);

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
                ? formatDate(assignment.created_at)
                : "N/A"}
            </div>
            <div className="Assignment__date">
              <div className="Assignment__date--title"> {"Due Date:"}</div>
              {assignment.main_due_date
                ? formatDate(assignment.main_due_date)
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
            {linkError && <p className="error">Failed to generate assignment invite link</p>}
          </div>

          <div className="Assignment__metrics">
            <h2>Metrics</h2>
            <MetricPanel>
              <Metric title="First Commit Date">{firstCommit}</Metric>
              <Metric title="Total Commits">{totalCommits ?? "N/A"}</Metric>
            </MetricPanel>

            <div className="Assignment__metricsCharts">
              <Metric
                title="Grading Status"
                className="Assignment__metricsChart Assignment__metricsChart--graded"
              >
                <Doughnut
                  redraw={true}
                  data={gradedMetrics}
                  options={{
                    maintainAspectRatio: true,
                    plugins: {
                      legend: {
                        onClick: () => {},
                        display: true,
                        position: "bottom",
                        labels: {
                          usePointStyle: true,
                          font: {
                            size: 12,
                          },
                        },
                      },
                      datalabels: {
                        color: ["#fff", "#000"],
                        font: {
                          size: 12,
                        },
                      },
                      tooltip: {
                        enabled: false,
                      },
                    },
                    cutout: "50%",
                    borderColor: "transparent",
                  }}
                />
              </Metric>

              <Metric
                title="Repository Status"
                className="Assignment__metricsChart Assignment__metricsChart--acceptance"
              >
                <Bar
                  redraw={true}
                  data={acceptanceMetrics}
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
                            size: 12,
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
                          size: 12,
                        },
                      },
                      tooltip: {
                        enabled: false,
                      },
                    },
                  }}
                />
              </Metric>
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
                studentWorks.map((sa: IStudentWork, i: number) => (
                  <TableRow key={i} className="Assignment__submission">
                    <TableCell>
                      <Link
                        to={`/app/submissions/${sa.student_work_id}`}
                        state={{ submission: sa, assignmentId: assignment.id }}
                        className="Dashboard__assignmentLink"
                      >
                        {sa.contributors.join(", ")}
                      </Link>
                    </TableCell>
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
