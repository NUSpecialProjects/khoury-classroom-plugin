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
  const [selectedSha, setSelectedSha] = useState<string>("");
  const root = buildTree(gitTree);

  return (
    <ResizableBox
      className={"FileTree" + (className ? " " + className : "")}
      width={230}
      height={Infinity}
      resizeHandles={["e"]}
      handle={<div className="ResizeHandle"></div>}
    >
      <div className="FileTree__head">Files</div>
      <div className="FileTree__body" {...props}>
        {sortTreeNode(root).map(([name, node]) =>
          renderTree(node, name, 0, selectedSha, (n) => {
            setSelectedSha(n.sha);
            selectFileCallback(n);
          })
        )}
        {children}
      </div>
    </ResizableBox>
  );
};

/*********************
 * DIRECTORY COMPONENT
 *********************/
export const FileTreeDirectory: React.FC<IFileTreeDirectory> = ({
  name,
  depth,
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
        className="FileTreeDirectory__name"
        style={{ paddingLeft: (depth * 15).toString() + "px" }}
        onClick={() => {
          setCollapsed(!collapsed);
        }}
      >
        {collapsed ? <FaChevronRight /> : <FaChevronDown />} <span>{name}</span>
      </div>
      <div
        className="FileTreeDirectory__bars"
        style={{ marginLeft: (depth * 15 + 5).toString() + "px" }}
      ></div>
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
  depth,
  className,
  ...props
}) => {
  return (
    <div
      className={"FileTreeFile" + (className ? " " + className : "")}
      style={{ paddingLeft: (depth * 15).toString() + "px" }}
      {...props}
    >
      {name}
    </div>
  );
};

export default FileTree;
