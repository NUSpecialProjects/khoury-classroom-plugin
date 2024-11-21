import { useContext } from "react";

import ResizablePanel from "../ResizablePanel";
import { GraderContext } from "@/contexts/grader";

const RubricTree: React.FC = () => {
  const { rubric } = useContext(GraderContext);
  return (
    rubric && (
      <ResizablePanel panelName="Rubric" border="left">
        {rubric.rubric_items.map((rubricItem) => (
          <>{rubricItem.explanation}</>
        ))}
      </ResizablePanel>
    )
  );
};

export default RubricTree;
