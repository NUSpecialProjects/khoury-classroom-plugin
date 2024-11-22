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
    setThreshold(bound.x + (border == "right" ? bound.width : 0));
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
          const mouseX = (e as unknown as MouseEvent).clientX;
          const w = size.width;
          if (
            (border2dir[border].includes("w") && mouseX >= threshold) ||
            (border2dir[border].includes("e") && mouseX <= threshold)
          ) {
            console.log("test");
            if (
              (border2dir[border].includes("e") &&
                mouseX <= threshold - minWidth / 2) ||
              (border2dir[border].includes("w") &&
                mouseX >= threshold + minWidth / 2)
            ) {
              self.current?.setState({ width: 4 });
            } else {
              self.current?.setState({ width: minWidth });
            }
          } else {
            self.current?.setState({ width: w });
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
