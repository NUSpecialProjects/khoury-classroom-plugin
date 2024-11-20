import { ResizableBox, ResizeHandle } from "react-resizable";
import SimpleBar from "simplebar-react";

import "./styles.css";

interface ITreePanel extends React.HTMLProps<HTMLDivElement> {
  panelName: string;
  border: "left" | "right" | "both";
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
  panelName,
  border,
  zIndex = 0,
}) => {
  return (
    <ResizableBox
      className={`TreePanel ${className ?? ""}`}
      width={230}
      height={Infinity}
      resizeHandles={border2dir[border] as ResizeHandle[]}
      style={{ zIndex }}
    >
      <>
        <div className="TreePanel__head">{panelName}</div>
        <SimpleBar className="TreePanel__body">{children}</SimpleBar>
      </>
    </ResizableBox>
  );
};

export default ResizablePanel;
