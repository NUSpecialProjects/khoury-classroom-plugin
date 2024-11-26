import { useContext, useEffect, useState } from "react";

import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import PageHeader from "@/components/PageHeader";
import { getRubricsInClassroom } from "@/api/rubrics";
import Button from "@/components/Button";
import { Link } from "react-router-dom";
import RubricList from "@/components/RubricList";

const Rubrics: React.FC = () => {
    const { selectedClassroom } = useContext(SelectedClassroomContext)
    const [rubrics, setRubricsData] = useState<IFullRubric[]>([])

    useEffect(() => {
        if (selectedClassroom) {
            (async () => {
                try {
                    const retrievedRubrics = await getRubricsInClassroom(selectedClassroom.id)
                    if (retrievedRubrics !== null) {
                        console.log("Assignment rubric retrieved rubric data, ", retrievedRubrics)
                        setRubricsData(retrievedRubrics)
                    }
                } catch (_) {
                    //do nothing
                }

            })();
        }
        
        console.log("rubric Item data")
    }, []);



    return (
        <div>
            <PageHeader pageTitle="Rubrics"></PageHeader>

            <RubricList rubrics={rubrics} />

            <Link to={`/app/rubrics/new`}>
                <Button href=""> Create New Rubric </Button>
            </Link>
        </div>
    );
};

export default Rubrics;