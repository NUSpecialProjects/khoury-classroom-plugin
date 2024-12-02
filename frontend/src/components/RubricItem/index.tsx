import React, { useEffect, useState } from "react";
import "./styles.css";

interface IRubricItemProps {
    name: string;
    points: string;
    impact?: boolean;
    onChange: (updatedFields: Partial<{ explanation: string; point_value: number; impact: boolean }>) => void;
}

enum FeedbackType {
    Addition = "A",
    Deduction = "D",
    Neutral = "N"
}

const RubricItem: React.FC<IRubricItemProps> = ({ name, points, impact, onChange,}) => {
    const [selection, setSelection] = useState<FeedbackType>(FeedbackType.Neutral);
    const [displayPoints, setDisplayPoints] = useState(points.toString())

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
        const value = e.target.value;
        if (value === "") {
            setDisplayPoints("")
            onChange({ point_value: null })
            return;
        }
    
        const pointValue = parseInt(value, 10);
        if (!isNaN(pointValue)) {
            setDisplayPoints(pointValue.toString())
            updatePointsBasedOnDeduction(pointValue, selection);
        }
    };
    


    const makeAddition = () => {
        if(selection !== FeedbackType.Addition) {
            setSelection(FeedbackType.Addition);
            updatePointsBasedOnDeduction(parseInt(points, 10), FeedbackType.Addition)
        } else {
            setSelection(FeedbackType.Neutral)
        }
    };

    const makeDeduction = () => {
        if(selection !== FeedbackType.Deduction) {
            setSelection(FeedbackType.Deduction);
            updatePointsBasedOnDeduction(parseInt(points, 10), FeedbackType.Deduction)
        } else {
            setSelection(FeedbackType.Neutral)
        }
    };


    // on startup
    useEffect(() => {
        if (impact !== null && impact === undefined) {
            setSelection(FeedbackType.Neutral)
        } else {
            if (impact) {
                setSelection(FeedbackType.Addition)
            } else {
                setSelection(FeedbackType.Deduction)
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
                name={"pointValue"}
                value={displayPoints}
                placeholder="Enter point value"
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
