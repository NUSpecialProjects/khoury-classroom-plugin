import { FaChevronRight, FaChevronDown } from "react-icons/fa";
import { useState } from "react";

import "./styles.css";

export const FileTree: React.FC<React.HTMLProps<HTMLDivElement>> = ({
  children,
  className,
  ...props
}) => {
  return (
    <div className={"FileTree" + (className ? " " + className : "")} {...props}>
      {children}
    </div>
  );
};

export const FileTreeDirectory: React.FC<IFileTreeDirectory> = ({
  name,
  children,
  className,
  ...props
}) => {
  const [collapsed, setCollapsed] = useState(false);
  return (
    <div
      className={"FileTreeDirectory" + (className ? " " + className : "")}
      {...props}
    >
      <div
        className="FileTreeDirectory__name"
        onClick={() => setCollapsed(!collapsed)}
      >
        {collapsed ? <FaChevronRight /> : <FaChevronDown />} {name}
      </div>
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

export const FileTreeFile: React.FC<IFileTreeFile> = ({
  name,
  className,
  ...props
}) => {
  return (
    <div
      className={"FileTreeFile" + (className ? " " + className : "")}
      {...props}
    >
      {name}
    </div>
  );
};

export const buildTree = (tree1D: IGitTreeNode[]) => {
  let root: IFileTreeNode = {
    type: "tree",
    sha: "",
    childNodes: {},
  };
  tree1D.forEach((node) => {
    const path = node.path.split("/");
    let level: IFileTreeNode = root;

    path.forEach((dir, i) => {
      if (!level.childNodes[dir]) {
        level.childNodes[dir] = {
          type: i === path.length - 1 ? node.type : "tree",
          sha: i === path.length - 1 ? node.sha : "",
          childNodes: {},
        };
      }

      level = level.childNodes[dir];
    });
  });

  return root;
};

export const sortTreeNode = (node: IFileTreeNode) => {
  return Object.entries(node.childNodes).sort(
    ([nameA, nodeA], [nameB, nodeB]) => {
      console.log(nodeA.type + ":" + nodeB.type);
      // Sort directories first and then by name
      if (nodeA.type === "tree" && nodeB.type === "blob") {
        console.log("test");
        return -1;
      }
      // dir before file
      if (nodeA.type === "blob" && nodeB.type === "tree") return 1; // file after dir
      return nameA.localeCompare(nameB); // alphabetical order
    }
  );
};

export default FileTree;
