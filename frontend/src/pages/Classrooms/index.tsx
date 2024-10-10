import TopNav from "@/components/Layout/TopNav";

import "./styles.css";
import { useEffect, useState } from "react";

type Classroom = {
    id: number;
    name: string;
    archived: string;
    url: string;
}


const Classrooms: React.FC = () => {
    const [classrooms, setClassrooms] = useState<Classroom[]>([]);

    useEffect(() => {
        // Replace with your API endpoint
        // fetch('https://api.example.com/items')
        //   .then(response => {
        //     if (!response.ok) {
        //       throw new Error('Network response was not ok');
        //     }
        //     return response.json();
        //   })
        //   .then(data => {
        //     setClassrooms(data);
        //   })
        //   .catch(error => {
        //     console.error('Error fetching data:', error);
        //   });
        //Fake data go brrrrrr

        const c: Classroom[] = [
            {
                id: 13232,
                name: "CS3200 Database Desgin",
                archived: ", false",
                url: "yes.com"
            },
            {
                id: 13232,
                name: "CS3200 Database Desgin",
                archived: ", false",
                url: "yes.com"
            }
        ]
        setClassrooms(c)
    }, []);

    return (
        <div className="Classrooms">
            <div>
                <TopNav />
            </div>
            <div className="Classrooms__bottom">
                <div className="Classrooms__content">
                    <div className="Classrooms__title">Classrooms</div>

                    {classrooms.length === 0 ? (
                        <div className="Classrooms__none"> No Classrooms Found...</div>
                    ) : (
                        <div className="Classrooms__list">
                                {classrooms.map(item => (
                                    <div className="Classrooms__classLink">
                                        <div>{item.name}</div>
                                        <div className="Classrooms__linkProf">PRofessor</div>
                                    </div>
                                ))}
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}

export default Classrooms;