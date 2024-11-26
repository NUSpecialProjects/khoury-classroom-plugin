import { ResizableBox, ResizeHandle } from "react-resizable";
import { useState, useRef, useEffect } from "react";

import "./styles.css";

interface ITreePanel extends React.HTMLProps<HTMLDivElement> {
  border: "left" | "right" | "both";
  minWidth?: number;
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
  minWidth = 230,
}) => {
  const [collapsed, setCollapsed] = useState(false);
  const [threshold, setThreshold] = useState(0);

  const wrapper = useRef<HTMLDivElement>(null);
  const self = useRef<ResizableBox>(null);

  useEffect(() => {
    if (!wrapper.current) return;
    const bound = wrapper.current?.children[0].getBoundingClientRect();
    setThreshold(bound.x + (border == "right" ? minWidth : 0));
  }, [wrapper]);

  return (
    <div className="TreePanel__wrapper" ref={wrapper}>
      <ResizableBox
        ref={self}
        className={`TreePanel ${collapsed ? "TreePanel--collapsed" : ""}${className ?? ""}`}
        width={minWidth}
        height={Infinity}
        resizeHandles={border2dir[border] as ResizeHandle[]}
        minConstraints={[4, 0]}
        onResize={(e, { size }) => {
          if (!self.current) return;

          // check if mouse is halfway between the panel edge-to-edge
          const mouseX = (e as unknown as MouseEvent).clientX;
          const w = size.width;
          if (
            (border2dir[border].includes("w") && mouseX >= threshold) ||
            (border2dir[border].includes("e") && mouseX <= threshold)
          ) {
            // override default resizing behavior if so
            if (
              (border2dir[border].includes("e") &&
                mouseX <= threshold - minWidth / 2) ||
              (border2dir[border].includes("w") &&
                mouseX >= threshold + minWidth / 2)
            ) {
              // collapse panel if mouse under halfway
              self.current.setState({ width: 4 });
            } else {
              // expand panel if mouse over halfway
              self.current.setState({ width: minWidth });
            }
          }
          setCollapsed(w == 4);
        }}
      >
        <>{children}</>
      </ResizableBox>
    </div>
  );
};

export default ResizablePanel;
