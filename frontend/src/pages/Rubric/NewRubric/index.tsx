import { useContext, useEffect, useState } from "react";

import "./styles.css";
import Button from "@/components/Button";
import Input from "@/components/Input";
import RubricItem from "@/components/RubricItem";
import { createRubric, updateRubric } from "@/api/rubrics";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { setAssignmentRubric } from "@/api/assignments";

interface NewRubricProps {
    givenRubricData?: IFullRubric; // Accept rubricData as a prop
    assignment?: IAssignmentOutline;
}

const NewRubric: React.FC<NewRubricProps> = ({ givenRubricData, assignment }) => {
    const { selectedClassroom } = useContext(SelectedClassroomContext)

    // front end only id for each rubric item, kept track of using a counter
    const [itemCount, setitemCount] = useState(0);
    const incrementCount = () => setitemCount(itemCount + 1);

    // default item for adding new items
    const newRubricItem: IRubricItem = {
        id: itemCount,
        point_value: 0,
        explanation: "",
        rubric_id: null,
        created_at: null
    }

    //data for the rubric
    const [rubricData, setRubricData] = useState<IFullRubric>()
    const [rubricItemData, setRubricItemData] = useState<IRubricItem[]>([])
    const [rubricName, setRubricName] = useState<string>("")

    // if there has been any changes since last save
    const [rubricEdited, setRubricEdited] = useState<boolean>(false)

    // saving the rubric, creates a rubric in the backend
    const saveRubric = async () => {
        if (selectedClassroom !== null && selectedClassroom !== undefined && rubricEdited) {
            const rubricItems = (rubricItemData.map(item => ({ ...item, id: null })));

            const fullRubric: IFullRubric = {
                rubric: {
                    id: null,
                    name: rubricName,
                    org_id: selectedClassroom.org_id,
                    classroom_id: selectedClassroom.id,
                    reusable: true,
                    created_at: null
                },
                rubric_items: rubricItems
            }
            
            // update existing rubric
            if (rubricData) {
                console.log("updating")
                await updateRubric(rubricData.rubric.id!, fullRubric)
                .then((updatedRubric) => {
                    setRubricEdited(false)
                    setRubricData(updatedRubric)
                    if (assignment !== null && assignment !== undefined) {
                        setAssignmentRubric(updatedRubric.rubric.id!, selectedClassroom.id, assignment.id)
                    }
                })
                .catch((error) => {
                    console.error("Error creating rubric:", error);
                });

            // create new rubric
            } else {
                console.log("creating")
                console.log(fullRubric)
                await createRubric(fullRubric)
                .then((createdRubric) => {
                    setRubricEdited(false)
                    setRubricData(createdRubric)
                    if (assignment !== null && assignment !== undefined) {
                        setAssignmentRubric(createdRubric.rubric.id!, selectedClassroom.id, assignment.id)
                    }
                })
                .catch((error) => {
                    console.error("Error creating rubric:", error);
                });
            }
        }
    };

    // handles when any rubric item is updated
    const handleItemChange = (id: number, updatedFields: Partial<IRubricItem>) => {
        setRubricEdited(true)
        setRubricItemData((prevItems) =>
            prevItems.map((item) =>
                item.id === id ? { ...item, ...updatedFields } : item
            )
        );
    };

    // handles when the rubric's name is changed
    const handleNameChange = (newName: string) => {
        setRubricName(newName)
        setRubricEdited(true)
    }

    // handles adding another rubric item
    const addNewItem = () => {
        setRubricItemData([...rubricItemData, newRubricItem]);
        incrementCount()
    };

    // on startup, store an assignment if we have one 
    // Also make sure there is atleast one editable rubric item already on the screen
    useEffect(() => {
        if (assignment !== null && assignment !== undefined) {
            if (givenRubricData) {
                console.log("recieved data in NewRurbic ", givenRubricData)
                setRubricData(givenRubricData)
                setRubricName(givenRubricData.rubric.name)
                setRubricItemData(givenRubricData.rubric_items)
                console.log("Rubric Items: ", givenRubricData.rubric_items)
            } else {
                setRubricName(`${assignment.name} Rubric`)
            }
        } else {
            console.log("WHERE IS TEH DAMN ASSIGNMENT")
            setRubricName("New Rubric")
        }

        if (rubricItemData.length === 0) {
            addNewItem()
        }

        console.log("rubric Item data", rubricItemData)
    }, [assignment, givenRubricData, rubricItemData]);



    return (
        <div className="NewRubric__body">
            {!rubricData &&
                <div className="NewRubric__title">  
                    {assignment !== null && assignment !== undefined ? `${assignment.name} > ` : ""}
                    {rubricEdited ? "New Rubric *" : "New Rubric"}
                </div>
            }

            <Input
                label="Rubric name"
                name="rubric-name"
                placeholder="Enter a name for your classroom..."
                required
                value={rubricName}
                onChange={(n) => { handleNameChange(n.target.value) }}
            />

            <div className="NewRubric__itemsTitle"> Rubric Items </div>

            {rubricItemData && rubricItemData.length > 0 &&
                rubricItemData.map((item) => (
                    <RubricItem
                        key={item.id}
                        name={item.explanation}
                        points={Math.abs(item.point_value).toString()}
                        deduction={item.point_value ? item.point_value > 0 : undefined}
                        onChange={(newItem) => handleItemChange(item.id ? item.id : 0, newItem)}
                    />
                ))
            }

            <Button href="" variant="secondary" onClick={addNewItem}> + Add a new item </Button>


            <div className="NewRubric__decisionButtons">
                <Button href="" variant="secondary"> Cancel </Button>
                <Button href="" onClick={saveRubric}> Save rubric </Button>
            </div>
        </div>
    );
};

export default NewRubric;