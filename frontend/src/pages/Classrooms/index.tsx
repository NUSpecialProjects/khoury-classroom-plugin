import TopNav from "@/components/Layout/TopNav";

import "./styles.css";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

type Classroom = {
    id: number;
    name: string;
    archived: string;
    url: string;
}


const Classrooms: React.FC = () => {
    const [classrooms, setClassrooms] = useState<Classroom[]>([]);
    const navigate = useNavigate()

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

    const handleNavigation = () => {
        // Navigate to a specific classroom page or any other route
        navigate(`/app/dashboard`);
      };

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
                                    <button className="Classrooms__classLink"
                                            onClick={() => handleNavigation()}>
                                        <div>{item.name}</div>
                                        <div className="Classrooms__linkProf">PRofessor</div>
                                    </button>
                                ))}
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}

export default Classrooms;