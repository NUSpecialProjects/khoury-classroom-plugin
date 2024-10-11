import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";

const Dashboard: React.FC = () => {
    return (
        <div className="Dashboard">
            <div className="Dashboard__classroomDetailsWrapper">
                <UserGroupCard label="Professors" number={1} />
                <UserGroupCard label="TAs" number={12} />
                <UserGroupCard label="Students" number={38} />
            </div>
            <div className="Dashboard__assignmentsWrapper">
                <div className="Dashboard__assignmentsHeader">
                    Active Assignments (1)</div>
                <div className="Dashboard__assignments">
                    <div className="Dashboard__assignment">
                        <div>Assignment Name</div>
                        <div>Released</div>
                        <div>Due Date</div>
                    </div>
                    {Array.from({ length: 1 }).map((_, i: number) => (
                        <div key={i} className="Dashboard__assignment">
                            <div><a href="#" className="Dashboard__assignmentLink">Assignment 3</a></div>
                            <div>5 Sep, 10:00AM</div>
                            <div>15 Sep, 11:59pm</div>
                        </div>
                    ))}
                </div>
                <div className="Dashboard__assignmentsHeader">
                    Inactive Assignments (3)</div>
                <div className="Dashboard__assignments">
                    <div className="Dashboard__assignment">
                        <div>Assignment Name</div>
                        <div>Released</div>
                        <div>Due Date</div>
                    </div>
                    {Array.from({ length: 3 }).map((_, i: number) => (
                        <div key={i} className="Dashboard__assignment">
                            <div><a href="#" className="Dashboard__assignmentLink">Assignment 1</a></div>
                            <div>5 Sep, 10:00AM</div>
                            <div>15 Sep, 11:59pm</div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}

export default Dashboard;