import { useContext, useState } from "react";
import SimpleBar from "simplebar-react";

import ResizablePanel from "../ResizablePanel";
import Button from "@/components/Button";
import { GraderContext } from "@/contexts/grader";

import "./styles.css";

const RubricTree: React.FC = () => {
  const { rubric, postFeedback } = useContext(GraderContext);

  return (
    rubric && (
      <ResizablePanel border="left">
        <div className="RubricTree__head">Rubric</div>

        <SimpleBar className="RubricTree__body scrollable">
          <div className="RubricTree__items">
            {rubric.rubric_items.map((rubricItem, i) => (
              <RubricItem key={i} {...rubricItem} />
            ))}
          </div>
        </SimpleBar>
        <div className="RubricTree__foot">
          <Button onClick={postFeedback}>Submit Grade</Button>
        </div>
      </ResizablePanel>
    )
  );
};

const RubricItem: React.FC<IRubricItem> = ({
  id,
  point_value,
  explanation,
}) => {
  const { selectRubricItem, deselectRubricItem } = useContext(GraderContext);

  const [selected, setSelected] = useState(false);

  return (
    <div
      className={`RubricItem${selected ? " RubricItem--selected" : ""}`}
      onClick={() => {
        if (selected) {
          deselectRubricItem(id);
        } else {
          selectRubricItem(id);
        }
        setSelected(!selected);
      }}
    >
      <div
        className={`RubricItem__points RubricItem__points--${point_value > 0 ? "positive" : point_value < 0 ? "negative" : "neutral"}`}
      >
        {point_value == 0
          ? "Comment"
          : point_value > 0
            ? `+${point_value}`
            : point_value}
      </div>
      <div className="RubricItem__explanation">{explanation}</div>
    </div>
  );
};

export default RubricTree;
