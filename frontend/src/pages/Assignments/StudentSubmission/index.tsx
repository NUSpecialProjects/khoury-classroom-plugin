import "./styles.css";
import { useLocation, useParams, Link } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import MetricPanel from "@/components/Metrics/MetricPanel";
import SimpleMetric from "@/components/Metrics/SimpleMetric";
import { getPaginatedStudentWork } from "@/api/student_works";


const StudentSubmission: React.FC = () => {
    const location = useLocation();
    const [submission, setSubmission] = useState<IStudentWork[]>([]);
    const { id } = useParams();
    // const { selectedClassroom } = useContext(SelectedClassroomContext);

    useEffect(() => {
        // check if assignment has been passed through
        if (location.state) {
            setSubmission(location.state.submission);
        }
        console.log(submission, id);
    }, []);

    return (
        <div className="StudentWork">
        </div>
    );
};

export default StudentSubmission;