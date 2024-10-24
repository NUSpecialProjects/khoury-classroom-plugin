import { useState } from "react";
import { FaChevronRight, FaChevronDown } from "react-icons/fa";
import { ResizableBox } from "react-resizable";

import { buildTree, renderTree, sortTreeNode } from "./funcs";

import "./styles.css";

/****************
 * TREE COMPONENT
 ****************/
export const FileTree: React.FC<IFileTree> = ({
  gitTree,
  selectFileCallback,
  children,
  className,
  ...props
}) => {
  const [selectedFile, setSelectedFile] = useState<string>("");
  const { root, treeDepth } = buildTree(gitTree);

  return (
    <ResizableBox
      className={"FileTree" + (className ? " " + className : "")}
      style={{ zIndex: treeDepth * 2 }}
      width={230}
      height={Infinity}
      resizeHandles={["e"]}
      handle={
        <div className="ResizeHandle" style={{ zIndex: treeDepth * 2 + 1 }} />
      }
    >
      <>
        <div className="FileTree__head">Files</div>
        <div className="FileTree__body" {...props}>
          {sortTreeNode(root).map((node) =>
            renderTree(node, 0, treeDepth, selectedFile, (n) => {
              setSelectedFile(n.path);
              selectFileCallback(n);
            })
          )}
          {children}
        </div>
      </>
    </ResizableBox>
  );
};

/*********************
 * DIRECTORY COMPONENT
 *********************/
export const FileTreeDirectory: React.FC<IFileTreeDirectory> = ({
  name,
  depth,
  treeDepth,
  children,
  className,
  ...props
}) => {
  const [collapsed, setCollapsed] = useState(true);
  return (
    <div
      className={"FileTreeDirectory" + (className ? " " + className : "")}
      {...props}
    >
      <div
        className="FileTree__nodeName"
        style={{
          paddingLeft: (depth * 15 + 10).toString() + "px",
          top: (depth * 24).toString() + "px",
          zIndex: (treeDepth - depth) * 2,
        }}
        onClick={() => {
          setCollapsed(!collapsed);
        }}
      >
        {collapsed ? <FaChevronRight /> : <FaChevronDown />} <span>{name}</span>
      </div>
      <div
        className="FileTreeDirectory__bars"
        style={{
          marginLeft: (depth * 15 + 15).toString() + "px",
          zIndex: (treeDepth - depth) * 2 - 1,
        }}
      />
      <div
        className={
          "FileTreeDirectory__children" +
          (collapsed ? " FileTreeDirectory--collapsed" : "")
        }
      >
        {children}
      </div>
    </div>
  );
};

/****************
 * FILE COMPONENT
 ****************/
export const FileTreeFile: React.FC<IFileTreeFile> = ({
  name,
  path,
  depth,
  className,
  ...props
}) => {
  return (
    <div
      className={"FileTreeFile" + (className ? " " + className : "")}
      style={{ paddingLeft: (depth * 15 + 10).toString() + "px" }}
      {...props}
    >
      <span className="FileTree__nodeName" data-path={path}>
        {name}
      </span>
    </div>
  );
};

/*const ResizeHandle = forwardRef<HTMLDivElement, IResizeHandle>(
  ({ zIndex }, ref) => {
    return <div ref={ref} className="ResizeHandle" style={{ zIndex }} />;
  }
);*/

export default FileTree;
