import "./styles.css";
import { useLocation, useParams } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import MetricPanel from "@/components/Metrics/MetricPanel";
import Metric from "@/components/Metrics";
import Button from "@/components/Button";
import { getStudentWorkById } from "@/api/student_works";
import { getFirstCommit, getTotalCommits } from "@/api/student_works";
import { formatDate } from "@/utils/date";

import { MdEditDocument } from "react-icons/md";
import { FaGithub } from "react-icons/fa";

const StudentSubmission: React.FC = () => {
  const location = useLocation();
  const [submission, setSubmission] = useState<IStudentWork>();
  const { id } = useParams();
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const assignmentID = location.state.assignmentId;
  const [firstCommit, setFirstCommit] = useState<string>("");
  const [totalCommits, setTotalCommits] = useState<string>();

  console.log(location.state);

  useEffect(() => {
    if (location.state && location.state.submission) {
      setSubmission(location.state.submission); // Use submission from state
    } else if (id && assignmentID && selectedClassroom) {
      // If state fails, use submission data as a fallback
      (async () => {
        try {
          const fetchedSubmission = await getStudentWorkById(
            selectedClassroom.id,
            assignmentID, // Use assignmentID from state
            +id // Student submission ID
          );
          if (fetchedSubmission) {
            setSubmission(fetchedSubmission);
          }
        } catch (error) {
          console.error("Failed to fetch submission:", error);
        }
      })();
    }
  }, [location.state, id, selectedClassroom]);

  useEffect(() => {
    if (
      submission !== null &&
      submission !== undefined &&
      selectedClassroom !== null &&
      selectedClassroom !== undefined
    ) {
      (async () => {
        try {
          const commitDate = await getFirstCommit(
            selectedClassroom.id,
            assignmentID,
            submission.student_work_id
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
  }, [selectedClassroom, submission]);

  useEffect(() => {
    if (
      submission !== null &&
      submission !== undefined &&
      selectedClassroom !== null &&
      selectedClassroom !== undefined
    ) {
      (async () => {
        try {
          const total = await getTotalCommits(
            selectedClassroom.id,
            assignmentID,
            submission.student_work_id
          );

          console.log(total);
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
  }, [selectedClassroom, submission]);

  return (
    <div className="StudentWork">
      <SubPageHeader
        pageTitle={submission?.contributors.map((contributor) => contributor.first_name + " " + contributor.last_name).join(", ")}
        pageSubTitle={submission?.assignment_name}
        chevronLink={`/app/assignments/${assignmentID}`}
      ></SubPageHeader>

      <div className="StudentSubmission__externalButtons">
        <Button
          href={`https://github.com/${submission?.org_name}/${submission?.repo_name}`}
          variant="secondary"
          newTab
        >
          <FaGithub className="icon" /> View Student Repository
        </Button>
        <Button
          href={`/app/grading/assignment/${assignmentID}/student/${submission?.student_work_id}`}
          variant="secondary"
        >
          <MdEditDocument className="icon" /> Grade Submission
        </Button>
      </div>

      <div className="StudentSubmission__subSectionWrapper">
        <h2 style={{ marginBottom: 10 }}>Metrics</h2>
        <MetricPanel>
          <Metric title="First Commit Date">{firstCommit}</Metric>
          <Metric title="Total Commits">{totalCommits ?? "N/A"}</Metric>
          <Metric title="First Commit Date">Put Chart Here</Metric>
        </MetricPanel>
      </div>
    </div>
  );
};

export default StudentSubmission;
