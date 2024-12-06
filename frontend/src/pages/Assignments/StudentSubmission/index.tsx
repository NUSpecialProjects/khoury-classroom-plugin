import "./styles.css";
import { useLocation, useParams, Link } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import MetricPanel from "@/components/Metrics/MetricPanel";
import SimpleMetric from "@/components/Metrics/SimpleMetric";
import { getPaginatedStudentWork } from "@/api/student_works";


const StudentSubmission: React.FC = () => {
    const location = useLocation();
    const [submission, setSubmission] = useState<IStudentWork>();
    const { id } = useParams();
    const { selectedClassroom } = useContext(SelectedClassroomContext);

    console.log(location.state);

    useEffect(() => {
        if (location.state && location.state.submission) {
          setSubmission(location.state.submission); // Use submission from state
        } else if (id && location.state.assignmentId && selectedClassroom) {
          // Fetch submission data as a fallback
          (async () => {
            try {
              const fetchedSubmission = await getPaginatedStudentWork(
                selectedClassroom.id,
                location.state.assignmentId, // Use assignmentId from state
                id // Student submission ID
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
                pageTitle={"Submission"}
                chevronLink={"/app/dashboard"}
            ></SubPageHeader>
        </div>
    );
};

export default StudentSubmission;