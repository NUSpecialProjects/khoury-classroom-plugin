import { FaChevronLeft, FaChevronRight } from "react-icons/fa";
import { Light as CodeViewer } from "react-syntax-highlighter";
import { hybrid } from "react-syntax-highlighter/dist/esm/styles/hljs";
import { useEffect, useState } from "react";

import {
  buildTree,
  sortTreeNode,
  FileTree,
  FileTreeDirectory,
  FileTreeFile,
} from "@/components/FileTree";
import Button from "@/components/Button";

import "./styles.css";

const Grader: React.FC = () => {
  // states
  const [fileTree, setFileTree] = useState<IFileTreeNode>({
    type: "tree",
    sha: "",
    childNodes: {},
  });
  const [cachedContents, setCachedContents] = useState<Record<string, string>>(
    {}
  );
  const [currentContent, setCurrentContent] = useState<string | null>(null);

  // iterate through a tree and render appropriate components
  const renderTree = (node: IFileTreeNode, name: string) => {
    if (node.type === "blob") {
      return (
        <FileTreeFile
          key={name}
          name={name}
          onClick={() => {
            openFile(node);
          }}
        />
      );
    }

    // if not a blob (file), must be a tree (directory)
    return (
      <FileTreeDirectory key={name} name={name}>
        {sortTreeNode(node).map(([childName, childNode]) =>
          renderTree(childNode, childName)
        )}
      </FileTreeDirectory>
    );
  };

  useEffect(() => {
    fetch(
      "http://localhost:8080/file-tree/org/NUSpecialProjects/assignment/1/student/92pLytz-SgW~mKeuxDyuJg"
    )
      .then((response) => response.json())
      .then((data: IGitTreeNode[]) => {
        setFileTree(buildTree(data));
        console.log(buildTree(data));
      });
  }, []);

  const openDir = (dir: IGitTreeNode) => {};

  const openFile = (file: IFileTreeNode) => {
    // Check if the content is already cached
    if (cachedContents[file.sha]) {
      setCurrentContent(cachedContents[file.sha]);
      return;
    }

    fetch(
      "http://localhost:8080/file-tree/org/NUSpecialProjects/assignment/1/student/92pLytz-SgW~mKeuxDyuJg/blob/" +
        file.sha
    )
      .then((response) => response.text())
      .then((content) => {
        setCurrentContent(content);
        // Cache the content
        setCachedContents((prev) => ({
          ...prev,
          [file.sha]: content,
        }));
      });
  };

  return (
    <div className="Grader">
      <div className="Grader__head">
        <div className="Grader__title">
          <FaChevronLeft />
          <div>
            <h2>Assignment 3</h2>
            <span>Jane Doe</span>
          </div>
        </div>
        <div className="Grader__nav">
          <span>Submission 2/74</span>
          <div>
            <Button>
              <FaChevronLeft />
              Previous
            </Button>
            <Button>
              Next
              <FaChevronRight />
            </Button>
          </div>
        </div>
      </div>
      <div className="Grader__body">
        <FileTree className="Grader__files">
          {Object.entries(fileTree.childNodes).map(([name, node]) =>
            renderTree(node, name)
          )}
        </FileTree>
        <CodeViewer
          className="Grader__code"
          showLineNumbers
          lineNumberStyle={{ color: "#999", margin: "0 5px" }}
          style={hybrid}
        >
          {currentContent ?? ""}
        </CodeViewer>
      </div>
    </div>
  );
};

export default Grader;
