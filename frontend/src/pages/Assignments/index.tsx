import { getUserGradeByAssID, getUserGrades } from "@/api/assignment_requests";
import { GradeEntry } from "@/components/Viz/BoxPlot/Boxplot";
import { GradeDistBoxPlot } from "@/components/Viz/BoxPlot/GradeDistBoxPlot";
import { useEffect, useState } from "react";


const Assignments: React.FC = () => {
    const [userGrade, setUserGrade] = useState<number>(0)
    const [userGradeArray, setUserGradeArray] = useState<GradeEntry[]>([])

    useEffect(() => {
        const fetchSingleGrade= async () => {
            try {
                const tempData: GradeEntry[] = await getUserGrades("CamPlume1");
                setUserGradeArray(tempData);
            } catch (error) {
                console.error("Error fetching organizations:", error);
            }
        };
        const fetchUserGrades= async () => {
            try {
                const tempData: GradeEntry = await getUserGradeByAssID(1, "CamPlume1");
                setUserGrade(tempData.value);
            } catch (error) {
                console.error("Error fetching organizations:", error);
            } finally {
                ;
            }
        };
        fetchUserGrades();
        fetchSingleGrade();
    }, []);



    return(
        <>
        <div>
            {userGrade}
            <ul>
                    {userGradeArray.map((entry, index) => (
                        <li key={index}>
                            Name: {entry.name}, Value: {entry.value}
                        </li>
                    ))}
                </ul>
        <GradeDistBoxPlot />
        </div>
        </>
    )
} 

export default Assignments;