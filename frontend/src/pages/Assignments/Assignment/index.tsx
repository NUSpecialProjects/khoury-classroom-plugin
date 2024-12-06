import Button from "@/components/Button";

import "./styles.css";
import { Link, useLocation, useParams } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Table, TableCell, TableRow } from "@/components/Table";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import { getAssignmentIndirectNav, postAssignmentToken } from "@/api/assignments";
import { getStudentWorkCommitsPerDay, getStudentWorks } from "@/api/student_works";
import { formatDateTime } from "@/utils/date";
import CopyLink from "@/components/CopyLink";
import MetricPanel from "@/components/Metrics/MetricPanel";
import SimpleMetric from "@/components/Metrics/SimpleMetric";

import { MdEditDocument } from "react-icons/md";
import { FaGithub } from "react-icons/fa";

import { ChartData, Chart as ChartJS, ChartOptions, Point, registerables } from "chart.js";
import { Line } from 'react-chartjs-2'
import ChartDataLabels from "chartjs-plugin-datalabels";

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
  const base_url: string = import.meta.env.VITE_PUBLIC_FRONTEND_DOMAIN as string;


  // useEffect for line chart 
  useEffect(() => {
    if (commitsPerDay) {
      const sortedDates = Array.from(commitsPerDay.keys()).sort((a, b) => a.valueOf() - b.valueOf()) 
      const sortedCounts: number[] = sortedDates.map((date) => commitsPerDay.get(date)!)

      const sortedDatesStrings = sortedDates.map((date) => `${date.getMonth()}/${date.getDate()}`)

      const lineData = {
        labels: sortedDatesStrings,
        datasets: [
          {
            data: sortedCounts,
            borderColor: 'rgba(244, 63, 94, 1)',
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
              const cPD = await getStudentWorkCommitsPerDay(selectedClassroom.id, a.id, 1)
              if (cPD !== null && cPD !== undefined) {
                setCommitsPerDay(cPD)
                console.log("ASDAUHDBSLKHJSDBSA: ", cPD instanceof Map)
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

          <div className="Assignment__externalButtons">
            <Button href="" variant="secondary" newTab>
              <FaGithub className="icon" /> View Template Repository
            </Button>
            <Button href="" variant="secondary" newTab>
              <MdEditDocument className="icon" />  View Rubric
            </Button>
            <Link to={`/app/assignments/${assignment.id}/rubric`} state={{ assignment }}>
              <Button href="" variant="secondary">
                View Rubric
              </Button>
            </Link>

          </div>

          <div className="Assignment__subSectionWrapper">
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
              <SimpleMetric metricTitle="First Commit Date" metricValue="6 Sep"></SimpleMetric>
              <SimpleMetric metricTitle="Total Commits" metricValue="941"></SimpleMetric>
              <SimpleMetric metricTitle="Extension  Requests" metricValue="0"></SimpleMetric>
              <SimpleMetric metricTitle="Regrade  Requests" metricValue="0"></SimpleMetric>
            </MetricPanel>
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
