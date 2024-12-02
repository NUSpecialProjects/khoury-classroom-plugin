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
    const [failedRurbicRetrival, setfailedRurbicRetrival] = useState(false)

    useEffect(() => {
        if (selectedClassroom) {
            (async () => {
                try {
                    const retrievedRubrics = await getRubricsInClassroom(selectedClassroom.id)
                    if (retrievedRubrics !== null) {
                        setRubricsData(retrievedRubrics)
                    }

                } catch (_) {
                    setfailedRurbicRetrival(true)
                }

            })();
        }
    }, []);



    return (
        <div>
            <PageHeader pageTitle="Rubrics"></PageHeader>

            {!failedRurbicRetrival ?
                <div>
                    {rubrics.length > 0 ?
                        <RubricList rubrics={rubrics} />
                        :
                        <div> No Rubrics Found </div>
                    }

                    <Link to={`/app/rubrics/new`}>
                        <Button href=""> Create New Rubric </Button>
                    </Link>
                </div>

                :
                
                <div>
                    <div> Failed to get existing rubrics </div>
                    <Button href="" onClick={() => window.location.reload()}>Click to retry</Button>
                </div>
            }

        </div>
    );
};

export default Rubrics;