import { useState } from "react";
import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";

const Dashboard: React.FC = () => {
    const [response, setResponse] = useState<string | null>(null);


    const handleButtonClick = async () => {
    
        try {
            const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
            const result = await fetch(`${base_url}/github/sync`, {
                method: 'POST', 
                credentials: 'include',
                headers: {
                  'Content-Type': 'application/json',
                },
                body: JSON.stringify({classroom_id : 237209}), //237210
              })
      
            if (!result.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await result.json();
            setResponse(JSON.stringify(data));
        } catch (error) {
            console.error('Error making API call:', error);
            setResponse('Error fetching data');
        } finally {
            console.log("Done")
        }
      };

    return (
        <div className="Dashboard">
            {/* Header group cards */}
            <div className="Dashboard__classroomDetailsWrapper">
                <UserGroupCard label="Professors" number={1} />
                <UserGroupCard label="TAs" number={12} />
                <UserGroupCard label="Students" number={38} />
            </div>
            
            {/* Assignments */}
            <div className="Dashboard__assignmentsWrapper">
                {/* Active Assignments */}

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
                <button onClick={handleButtonClick}>
                    ASSINGMENT DATA SYNC
                </button>
                {/* Inactive Assignments */}
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