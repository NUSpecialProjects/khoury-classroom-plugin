import React, { useEffect, useState } from "react";
import "./styles.css";

interface IRubricItemProps {
    name: string;
    points: string;
    deduction?: boolean;
    onChange: (updatedFields: Partial<{ explanation: string; point_value: number; deduction: boolean }>) => void;
}

enum FeedbackType {
    Addition = "A",
    Deduction = "D",
    Neutral = "N"
}

const RubricItem: React.FC<IRubricItemProps> = ({ name, points, deduction, onChange,}) => {
    const [selection, setSelection] = useState<FeedbackType>(FeedbackType.Neutral);
    const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        onChange({ explanation: e.target.value });
    };


 
    const updatePointsBasedOnDeduction = (pointValue: number, feedbackType: FeedbackType) => {
        if (!isNaN(pointValue)) {
            const adjustedValue = feedbackType === FeedbackType.Deduction ? -1*Math.abs(pointValue) : Math.abs(pointValue)
            onChange({ point_value: adjustedValue });
        }
    }

    const handlePointsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const pointValue = parseInt(e.target.value, 10);
        updatePointsBasedOnDeduction(pointValue, selection)
        
    };


    const makeAddition = () => {
        if(selection !== FeedbackType.Addition) {
            setSelection(FeedbackType.Addition);
            updatePointsBasedOnDeduction(parseInt(points,10), FeedbackType.Addition)
        } else {
            setSelection(FeedbackType.Neutral)
        }
    };

    const makeDeduction = () => {
        if(selection !== FeedbackType.Deduction) {
            setSelection(FeedbackType.Deduction);
            updatePointsBasedOnDeduction(parseInt(points,10), FeedbackType.Deduction)
        } else {
            setSelection(FeedbackType.Neutral)
        }
    };


    // on startup
    useEffect(() => {
        console.log("rubric item explatioanio: ", name)
        if (deduction !== null || deduction !== undefined) {
            setSelection(FeedbackType.Neutral)
        } else {
            if (deduction) {
                setSelection(FeedbackType.Deduction)
            } else {
                setSelection(FeedbackType.Addition)
            }
        } 
    }, [])


    return (
        <div className="RubricItem__wrapper">
            <input
                className="RubricItem__itemName"
                id={name}
                name={name}
                value={name}
                placeholder="Add a rubric item..."
                onChange={handleNameChange}
            />

            <input
                className="RubricItem"
                id={points}
                name={points}
                value={points}
                placeholder="0"
                onChange={handlePointsChange}
            />

            <div>
                <button onClick={makeAddition} className={`RubricItem__button${selection === FeedbackType.Addition ? "AdditionActive" : ""}`}
                >+</button>

                <button onClick={makeDeduction} className={`RubricItem__button${selection === FeedbackType.Deduction ? "DeductionActive" : ""}`}
                >-</button>

            </div>

        </div>
    );
};

export default RubricItem;
