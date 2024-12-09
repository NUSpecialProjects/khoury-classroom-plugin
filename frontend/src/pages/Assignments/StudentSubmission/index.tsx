import "./styles.css";
import { useLocation, useParams, Link } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import MetricPanel from "@/components/Metrics/MetricPanel";
import SimpleMetric from "@/components/Metrics/SimpleMetric";
import { getStudentWorkById } from "@/api/student_works";


const StudentSubmission: React.FC = () => {
    const location = useLocation();
    const [submission, setSubmission] = useState<IStudentWork>();
    const { id } = useParams();
    const { selectedClassroom } = useContext(SelectedClassroomContext);
    const assignmentID = location.state.assignmentId;

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
    
    return (
        <div className="StudentWork">
            <SubPageHeader
                pageTitle={submission?.contributors.join(", ")}
                pageSubTitle={submission?.assignment_name}
                chevronLink={`/app/assignments/${assignmentID}`}
            >
            </SubPageHeader>
        </div>
    );
};

export default StudentSubmission;