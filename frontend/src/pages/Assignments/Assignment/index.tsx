import { useLocation, useParams, Link } from "react-router-dom";
import { MdEdit, MdEditDocument } from "react-icons/md";
import { FaGithub } from "react-icons/fa";
import { useContext, useEffect, useState } from "react";
import { Bar, Doughnut } from "react-chartjs-2";
import { ChartData, Chart as ChartJS, ChartOptions, Point, registerables } from "chart.js";
import { Line } from 'react-chartjs-2'
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
import { getStudentWorkCommitsPerDay, getStudentWorks } from "@/api/student_works";
import { formatDate } from "@/utils/date";

import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import CopyLink from "@/components/CopyLink";
import { Table, TableCell, TableRow } from "@/components/Table";
import Button from "@/components/Button";
import MetricPanel from "@/components/Metrics/MetricPanel";
import Metric from "@/components/Metrics";

import "./styles.css";

ChartJS.register(...registerables);
ChartJS.register(ChartDataLabels);

const Assignment: React.FC = () => {
  const location = useLocation();
  const [assignment, setAssignment] = useState<IAssignmentOutline>();
  const [studentWorks, setStudentAssignment] = useState<IStudentWork[]>([]);
  const [commitsPerDay, setCommitsPerDay] = useState<Map<Date, number>>(new Map());
  const [lineData, setLineData] = useState<ChartData<"line", (number | Point | null)[], unknown>>()
  const [lineOptions, setLineOptions] = useState<ChartOptions<"line">>()
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const { id } = useParams();
  const [inviteLink, setInviteLink] = useState<string>("");
  const [linkError, setLinkError] = useState<string | null>(null);

  const [firstCommit, setFirstCommit] = useState<string>("");
  const [totalCommits, setTotalCommits] = useState<string>();

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

  const base_url: string = import.meta.env
    .VITE_PUBLIC_FRONTEND_DOMAIN as string;

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
              const cPD = await getStudentWorkCommitsPerDay(selectedClassroom.id, a.id, 3)
              if (cPD !== null && cPD !== undefined) {
                setCommitsPerDay(cPD)
              }
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

  // useEffect for line chart 
  useEffect(() => {
    if (commitsPerDay) {
      const sortedDates = Array.from(commitsPerDay.keys()).sort((a, b) => a.valueOf() - b.valueOf()) 
      // end dates at today or due date, whichever is sooner
      if (assignment) {
        if (assignment.main_due_date) {
          sortedDates.push(assignment.main_due_date.valueOf() - Date.now() ? assignment.main_due_date : new Date())
        } else {
          const today = new Date()
          today.setUTCHours(0)
          today.setUTCMinutes(0)
          today.setUTCSeconds(0)
          if (sortedDates[sortedDates.length-1].toDateString() !== (today.toDateString())) {
            sortedDates.push(new Date())
          }
          
        }
      }

      console.log(sortedDates)
      const sortedCounts: number[] = (sortedDates.map((date) => commitsPerDay.get(date) ?? 0))
      const sortedDatesStrings = sortedDates.map((date) => `${date.getUTCMonth()}/${date.getUTCDate()}`)
      console.log(sortedDatesStrings)
      console.log(sortedCounts)


      //add in days with 0 commits
      const sortedDatesStringsCopy = [...sortedDatesStrings]
      for (let i = 0; i < sortedDatesStringsCopy.length-1; i++) {
        const firstMonth = Number(sortedDatesStringsCopy[i].split("/")[0])
        const firstDay = Number(sortedDatesStringsCopy[i].split("/")[1])
        const secondDay = Number(sortedDatesStrings[i+1].split("/")[1])


        const difference = firstDay - secondDay

        const adjacent = (difference === -1) 
        const adjacentWrapped = ((difference === 30 || difference === 29 || difference === 27) && (secondDay === 1))

        if (!adjacent && !adjacentWrapped) {
          for (let j = 1; j < Math.abs(difference); j++) {
            if (firstMonth === 2 && firstDay === 29 ) {
              sortedDatesStrings.splice(i+j, 0, `${3}/${1}`);

            } else if (firstDay === 30 && (firstMonth === 10 || firstMonth === 4 || firstMonth === 5 || firstMonth === 11)){
              sortedDatesStrings.splice(i+j, 0, `${firstMonth+1}/${1}`);

            } else if (firstDay === 31 && !(firstMonth === 10 || firstMonth === 4 || firstMonth === 5 || firstMonth === 11)){
              if (firstMonth === 12) {
                sortedDatesStrings.splice(i+j, 0, `${firstMonth+1}/${1}`);
              } else {
                sortedDatesStrings.splice(i+j, 0, `${11}/${1}`);
              }
            } else {
              sortedDatesStrings.splice(i+j, 0, `${firstMonth}/${firstDay+j}`);
            }
            sortedCounts.splice(i+j, 0, 0)
          }
          
        }
      } 




      const lineData = {
        labels: sortedDatesStrings,
        datasets: [
          {
            data: sortedCounts,
            borderColor: 'rgba(13, 148, 136, 1)',
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            tension: 0.05,
          },
        ],
      }
      setLineData(lineData)
  
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
      }
      setLineOptions(lineOptions)
    }

  }, [commitsPerDay])

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
  if (assignment !== null && assignment !== undefined && selectedClassroom !== null && selectedClassroom !== undefined) {
    (async () => {
      try {
        const total = await getAssignmentTotalCommits (
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
            {linkError && <p className="error">{linkError}</p>}
          </div>

          <div className="Assignment__subSectionWrapper">
            <h2>Metrics</h2>
            <p>Metrics go here</p>

            <div className="Assignment__metricsWrapper">
              {lineData && lineOptions && (
                <Line
                  options={lineOptions}
                  data={lineData}
                />
              )}
            </div>



            <h2 style={{ marginBottom: 10 }}>Metrics</h2>
            <MetricPanel>
              <SimpleMetric metricTitle="First Commit Date" metricValue={firstCommit}></SimpleMetric>
              <SimpleMetric metricTitle="Total Commits" metricValue={totalCommits ?? "N/A"}></SimpleMetric>
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
                studentWorks.map((sa, i) => (
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
