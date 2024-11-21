import { ResizableBox, ResizeHandle } from "react-resizable";
import { useState } from "react";

import "./styles.css";

interface ITreePanel extends React.HTMLProps<HTMLDivElement> {
  border: "left" | "right" | "both";
  hideable?: boolean;
  zIndex?: number;
}

const border2dir = {
  left: ["w"],
  right: ["e"],
  both: ["e, w"],
};

const ResizablePanel: React.FC<ITreePanel> = ({
  children,
  className,
  border,
  hideable = false,
  zIndex = 0,
}) => {
  const [largeStep, setLargeStep] = useState(false);

  return (
    <ResizableBox
      className={`TreePanel ${largeStep ? "TreePanel--largeStep" : ""}${className ?? ""}`}
      width={230}
      height={Infinity}
      resizeHandles={border2dir[border] as ResizeHandle[]}
      style={{ zIndex }}
      minConstraints={[20, 0]}
      draggableOpts={hideable && largeStep ? { grid: [211, 0] } : undefined}
      onResize={(_, { size }) => {
        console.log(size.width);
        setLargeStep(size.width < 230);
      }}
    >
      <>{children}</>
    </ResizableBox>
  );
};

export default ResizablePanel;
