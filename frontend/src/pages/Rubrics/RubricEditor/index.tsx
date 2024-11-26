import { useContext, useEffect, useState } from "react";

import "./styles.css";
import Button from "@/components/Button";
import Input from "@/components/Input";
import RubricItem from "@/components/RubricItem";
import { createRubric, updateRubric } from "@/api/rubrics";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { setAssignmentRubric } from "@/api/assignments";
import { useLocation, useNavigate } from "react-router-dom";

interface IEditableItem {
    frontFacingIndex: number;
    rubricItem: IRubricItem;
}


const RubricEditor: React.FC = () => {
    const location = useLocation()
    const navigate = useNavigate();

    const { selectedClassroom } = useContext(SelectedClassroomContext)

    // potential data for the assignment
    const [assignmentData, setAssignmentData] = useState<IAssignmentOutline>()
    //data for the rubric
    const [rubricData, setRubricData] = useState<IFullRubric>()
    const [rubricItemData, setRubricItemData] = useState<IEditableItem[]>([])
    const [rubricName, setRubricName] = useState<string>("")

    // if there has been any changes since last save
    const [rubricEdited, setRubricEdited] = useState<boolean>(false)

    // front end id for each rubric item, kept track of using a counter
    const [itemCount, setitemCount] = useState(0);
    const incrementCount = () => setitemCount(itemCount + 1);

    // default item for adding new items
    const newRubricItem: IEditableItem = {
        frontFacingIndex: itemCount,
        rubricItem: {
            id: null,
            point_value: 0,
            explanation: "",
            rubric_id: null,
            created_at: null
        }
    }

    const backButton = () => {
        navigate(-1);
    }
    // saving the rubric, creates a rubric in the backend
    const saveRubric = async () => {
        if (selectedClassroom !== null && selectedClassroom !== undefined && rubricEdited) {
            const rubricItems = (rubricItemData.map(item => item.rubricItem));

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
                        console.log(updatedRubric)
                        setRubricEdited(false)
                        setRubricData(updatedRubric)
                        if (assignmentData !== null && assignmentData !== undefined) {
                            setAssignmentRubric(updatedRubric.rubric.id!, selectedClassroom.id, assignmentData.id)
                        }
                    })
                    .catch((_) => {
                    });

                // create new rubric
            } else {
                await createRubric(fullRubric)
                    .then((createdRubric) => {
                        setRubricEdited(false)
                        setRubricData(createdRubric)
                        if (assignmentData !== null && assignmentData !== undefined) {
                            setAssignmentRubric(createdRubric.rubric.id!, selectedClassroom.id, assignmentData.id)
                        }
                    })
                    .catch((_) => {
                    });
            }
        }
    };

    // handles when any rubric item is updated
    const handleItemChange = (id: number, updatedFields: Partial<IRubricItem>) => {
        setRubricEdited(true);
        setRubricItemData((prevItems) =>
            prevItems.map((item) =>
                item.frontFacingIndex === id
                    ? {
                          ...item,
                          rubricItem: {
                            ...item.rubricItem,
                            ...updatedFields,
                        },
                      }
                    : item
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
        if (location.state) {
            if (location.state.assignment) {
                const assignment = location.state.assignment
                setAssignmentData(assignment)
                setRubricName(`${assignment.name} Rubric`)
            }
            if (location.state.rubricData) {
                const rubric = location.state.rubricData
                setRubricData(rubric)
                setRubricName(rubric.rubric.name)
            }
        } else {
            setRubricName("New Rubric")
        }

        if (rubricItemData.length === 0) {
            addNewItem()
        }

    }, [location.state]);


    useEffect(() => {
        if (rubricData) {
            let localCount = itemCount
            const editableItems = rubricData.rubric_items.map((item) => {
                const editableItem: IEditableItem = {
                    rubricItem: item,
                    frontFacingIndex: localCount,
                };
                localCount++; // Increment itemCount for each item
                return editableItem;
            });
            setitemCount(localCount)
            setRubricItemData(editableItems)
        }
    }, [rubricData])



    return (
        <div className="NewRubric__body">
            <div className="NewRubric__title">
                {assignmentData !== null && assignmentData !== undefined ? `${assignmentData.name} > ` : ""}
                {rubricData !== null && rubricData !== undefined ? "Edit Rubric" : "New Rubric"}
                {rubricEdited ? "*" : ""}
            </div>


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
                        key={item.frontFacingIndex}
                        name={item.rubricItem.explanation}
                        points={Math.abs(item.rubricItem.point_value).toString()}
                        impact={item.rubricItem.point_value === 0 ? undefined : item.rubricItem.point_value > 0}
                        onChange={(newItem) => handleItemChange(item.frontFacingIndex, newItem)}
                    />
                ))
            }

            <Button href="" variant="secondary" onClick={addNewItem}> + Add a new item </Button>


            <div className="NewRubric__decisionButtons">
                <Button href="" variant="secondary" onClick={backButton}> Cancel </Button>
                <Button href="" onClick={saveRubric}> Save rubric </Button>
            </div>
        </div>
    );
};

export default RubricEditor;