import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";

const Dashboard: React.FC = () => {
    return(
        <div className="Dashboard">
            <div className="Dashboard__classroomDetailsWrapper">
                <UserGroupCard label="Professors" number={1} />
                <UserGroupCard label="TAs" number={12} />
                <UserGroupCard label="Students" number={38} />
            </div>
            <div className="Dashboard__assignmentsWrapper">
                <div className="Dashboard__assignmentsHeader">
                    Active Assignments (1)
                    <div className="Dashboard__assignmentsTable"></div>
                </div>
                <div className="Dashboard__assignmentsHeader">
                    Inactive Assignments (3)
                    <div className="Dashboard__assignmentsTable"></div>
                </div>
            </div>
        </div>
    )
} 

export default Dashboard;