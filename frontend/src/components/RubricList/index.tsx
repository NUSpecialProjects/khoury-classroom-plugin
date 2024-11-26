import { Link } from "react-router-dom";
import { Table, TableCell, TableDiv, TableRow } from "../Table";
import "./styles.css";
import { FaChevronDown, FaChevronRight } from "react-icons/fa6";
import Button from "../Button";
import { useState } from "react";


interface IRubricRowData extends React.HTMLProps<HTMLDivElement> {
    rubricData: IFullRubric,
}

const RubricRow: React.FC<IRubricRowData> = ({
    rubricData,
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
                    {rubricData.rubric.name}
                </TableCell>
                <TableCell>
                <Link to={`/app/rubrics/new`} state={{ rubricData }}>
                    <Button href="" variant="secondary" >
                        Edit this Rubric
                    </Button>
                </Link>
                </TableCell>
            </TableRow>
            {!collapsed && (
                <TableDiv>
                    <Table cols={3} className="ItemTable">
                        {rubricData.rubric_items &&
                            rubricData.rubric_items.map((item, i: number) => (
                                <TableRow
                                    key={i}>

                                    <TableCell>
                                        {item.explanation}
                                    </TableCell>

                                    <TableCell> 
                                        {item.point_value > 0 ? "+" + item.point_value : item.point_value}
                                    </TableCell>

                                    <TableCell/>
                                    
                                </TableRow>
                            ))}
                    </Table>
                </TableDiv>
            )}
        </>
    );

}


interface IRubricListData extends React.HTMLProps<HTMLDivElement> {
    rubrics: IFullRubric[]
}
const RubricList: React.FC<IRubricListData> = ({
    rubrics,
}) => {

    return (
        <Table cols={3} primaryCol={1} className="RubricsTable">
                <TableRow style={{ borderTop: "none" }}>
                    <TableCell></TableCell>
                    <TableCell>Rubric Name</TableCell>
                    <TableCell></TableCell>
                </TableRow>
                {rubrics &&
                    rubrics.length > 0 &&
                    rubrics.map((rubric, i) => (
                        <RubricRow
                            key={i}
                            rubricData={rubric}>
                        </RubricRow>
                    ))}
            </Table>
    );
}


export default RubricList;