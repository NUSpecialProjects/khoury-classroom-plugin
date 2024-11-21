import { useContext, useEffect, useState } from "react";

import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import PageHeader from "@/components/PageHeader";
import { getRubricsInClassroom } from "@/api/rubrics";
import { Table, TableCell, TableRow } from "@/components/Table";


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
                } catch (error) {
                    //do nothing
                }

            })();
        }


        console.log("rubric Item data")
    }, []);



    return (
        <div>
            <PageHeader pageTitle="Rubrics"></PageHeader>

            <Table cols={1}>
              <TableRow style={{ borderTop: "none" }}>
                <TableCell>Rubric Name</TableCell>
              </TableRow>
              {rubrics &&
                rubrics.length > 0 &&
                rubrics.map((rubric, i) => (
                  <TableRow key={i} className="Assignment__submission">
                    <TableCell>{rubric.rubric.name}</TableCell>
                  </TableRow>
                ))}
            </Table>

            


        </div>
    );
};

export default Rubrics;