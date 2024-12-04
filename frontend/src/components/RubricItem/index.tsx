import React, { useEffect, useState } from "react";
import "./styles.css";

interface IRubricItemProps {
    name: string;
    points: string;
    impact: ItemFeedbackType;
    onChange: (updatedFields: Partial<{ explanation: string; point_value: number | null; impact: ItemFeedbackType }>) => void;
}

export enum ItemFeedbackType {
    Addition = "A",
    Deduction = "D",
    Neutral = "N"
}

const RubricItem: React.FC<IRubricItemProps> = ({ name, points, impact, onChange,}) => {
    const [selection, setSelection] = useState<ItemFeedbackType>(ItemFeedbackType.Neutral)
    const [displayPoints, setDisplayPoints] = useState(points.toString())

    const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        onChange({ explanation: e.target.value });
    };


 
    const updatePointsBasedOnDeduction = (pointValue: number, feedbackType: ItemFeedbackType) => {
        if (!isNaN(pointValue)) {
            const adjustedValue = feedbackType === ItemFeedbackType.Deduction ? -1*Math.abs(pointValue) : Math.abs(pointValue)
            onChange({ point_value: adjustedValue, impact: feedbackType });
        }
    }

    const handlePointsChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value;
        if (value === "") {
            setDisplayPoints("")
            onChange({ point_value: null, impact: ItemFeedbackType.Neutral})
            return;
        }
        
        if (value === "0") {
            makeNeutral()
        }
    
        const pointValue = Math.abs(parseInt(value, 10));
        if (!isNaN(pointValue)) {
            setDisplayPoints(pointValue.toString())
            updatePointsBasedOnDeduction(pointValue, selection);
        }
    };
    
    const makeNeutral = () => {
        setSelection(ItemFeedbackType.Neutral)
    }


    const toggleAddition = () => {
        if(selection !== ItemFeedbackType.Addition) {
            setSelection(ItemFeedbackType.Addition);
            updatePointsBasedOnDeduction(parseInt(points, 10), ItemFeedbackType.Addition)
        } else {
            setSelection(ItemFeedbackType.Neutral)
            onChange({ impact: ItemFeedbackType.Neutral })
        }
    };

    const toggleDeduction = () => {
        if(selection !== ItemFeedbackType.Deduction) {
            setSelection(ItemFeedbackType.Deduction);
            updatePointsBasedOnDeduction(parseInt(points, 10), ItemFeedbackType.Deduction)
        } else {
            setSelection(ItemFeedbackType.Neutral)
            onChange({ impact: ItemFeedbackType.Neutral })
        }
    };


    // on startup
    useEffect(() => {
        setSelection(impact)
    }, [])


    return (
        <div className="RubricItem__wrapper">
            <input
                className="RubricItem__itemName"
                id={name}
                name={name}
                value={name}
                maxLength={250}
                placeholder="Add a rubric item..."
                onChange={handleNameChange}
            />

            <input
                className="RubricItem__itemPoints"
                id={points}
                name={"pointValue"}
                value={displayPoints}
                placeholder="Enter point value"
                maxLength={6}
                onChange={handlePointsChange}
            />

            <div className="RubricItem__buttonWrap">
                <button onClick={toggleAddition} className={`RubricItem__button${selection === ItemFeedbackType.Addition ? "AdditionActive" : ""}`}
                >+</button>

                <button onClick={toggleDeduction} className={`RubricItem__button${selection === ItemFeedbackType.Deduction ? "DeductionActive" : ""}`}
                >-</button>

            </div>

        </div>
    );
};

export default RubricItem;
