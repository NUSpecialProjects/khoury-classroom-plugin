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
        </div>
    )
} 

export default Dashboard;