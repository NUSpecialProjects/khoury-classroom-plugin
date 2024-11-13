import React, { useState } from "react";
import "./styles.css";

interface IRubricItemProps {
    name: string;
    points: string;
    deduction: boolean;
    onChange: (updatedFields: Partial<{ explanation: string; point_value: number; deduction: boolean }>) => void;
}

enum FeedbackType {
    Addition = "A",
    Deduction = "D",
    Neutral = "N"
}

const RubricItem: React.FC<IRubricItemProps> = ({
    name,
    points,
    deduction, 
    onChange,
}) => {

    const [selection, setSelection] = useState<FeedbackType>(FeedbackType.Neutral);
    const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        onChange({ explanation: e.target.value });
    };


    const updatePointsBasedOnDeduction = (pointValue: number, feedbackType: FeedbackType) => {
        if (!isNaN(pointValue)) {
            console.log("point value pre update, ", pointValue)
            console.log("feedback type ", feedbackType)
            const adjustedValue = feedbackType === FeedbackType.Deduction ? -1*Math.abs(pointValue) : Math.abs(pointValue)
            console.log("adjuasted value pre update, ", adjustedValue)
            onChange({ point_value: adjustedValue });
        }
    }

    const handlePointsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const pointValue = parseInt(e.target.value, 10);
        updatePointsBasedOnDeduction(pointValue, selection)
        
    };


    const makeAddition = () => {
        console.log("made addutuin")

        if(selection !== FeedbackType.Addition) {
            setSelection(FeedbackType.Addition);
            updatePointsBasedOnDeduction(parseInt(points,10), FeedbackType.Addition)

        } else {
            setSelection(FeedbackType.Neutral)
        }
    };

    const makeDeduction = () => {
        console.log("made deduciton")
        if(selection !== FeedbackType.Deduction) {
            setSelection(FeedbackType.Deduction);
            updatePointsBasedOnDeduction(parseInt(points,10), FeedbackType.Deduction)
        } else {
            setSelection(FeedbackType.Neutral)
        }
    };


    return (
        <div className="RubricItem__wrapper">
            <input
                className="RubricItem__itemName"
                id={name}
                name={name}
                placeholder="Add a rubric item..."
                onChange={handleNameChange}
            />

            <input
                className="RubricItem"
                id={points}
                name={points}
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
