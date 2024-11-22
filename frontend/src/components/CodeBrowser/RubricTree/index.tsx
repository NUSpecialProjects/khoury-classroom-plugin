import { useContext } from "react";
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
          {rubric.rubric_items.map((rubricItem, i) => (
            <div className="RubricTree__item" key={i}>
              {rubricItem.explanation}
            </div>
          ))}
        </SimpleBar>
        <div className="RubricTree__foot">
          <Button onClick={postFeedback}>Submit Grade</Button>
        </div>
      </ResizablePanel>
    )
  );
};

export default RubricTree;
