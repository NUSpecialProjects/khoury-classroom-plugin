import { useContext, useEffect, useState } from "react";

import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import PageHeader from "@/components/PageHeader";
import { getRubricsInClassroom } from "@/api/rubrics";
import { Table, TableCell, TableDiv, TableRow } from "@/components/Table";
import { FaChevronDown, FaChevronRight } from "react-icons/fa6";


interface IRubricRowData extends React.HTMLProps<HTMLDivElement> {
    rubricItems: IRubricItem[];
    rubricName: string
}



const RubricRow: React.FC<IRubricRowData> = ({
    rubricItems,
    rubricName
}) => {
    const [collapsed, setCollapsed] = useState(true);


    return (
        <>
            <TableRow
                className={!collapsed ? "TableRow--expanded" : undefined}
                onClick={() => {
                    setCollapsed(!collapsed);
                }}
            >
                <TableCell>
                    {collapsed ? <FaChevronRight /> : <FaChevronDown />}
                </TableCell>
                <TableCell>
                    {rubricName}

                </TableCell>
            </TableRow>
            {!collapsed && (
                <TableDiv>
                    <Table cols={2} className="ItemTable">
                        {rubricItems &&
                            rubricItems.map((item, i: number) => (
                                <TableRow
                                    key={i}>
                                    <TableCell>
                                        {item.explanation}
                                    </TableCell>
                                    <TableCell>{item.point_value}</TableCell>
                                </TableRow>
                            ))}
                    </Table>
                </TableDiv>
            )}
        </>
    );

}



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

            <Table cols={2} primaryCol={1} className="RubricsTable">
                <TableRow style={{ borderTop: "none" }}>
                    <TableCell></TableCell>
                    <TableCell>Rubric Name</TableCell>
                </TableRow>
                {rubrics &&
                    rubrics.length > 0 &&
                    rubrics.map((rubric, i) => (
                        <RubricRow
                            key={i}
                            rubricItems={rubric.rubric_items}
                            rubricName={rubric.rubric.name}>
                        </RubricRow>
                    ))}
            </Table>
        </div>
    );
};

export default Rubrics;