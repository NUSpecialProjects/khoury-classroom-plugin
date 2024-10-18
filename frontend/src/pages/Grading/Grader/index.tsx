import { FaChevronLeft, FaChevronRight } from "react-icons/fa";
import { Light as CodeViewer } from "react-syntax-highlighter";
import { hybrid } from "react-syntax-highlighter/dist/esm/styles/hljs";
import { useEffect, useState } from "react";

import Button from "@/components/Button";

import "./styles.css";

const Grader: React.FC = () => {
  const [cachedContents, setCachedContents] = useState<Record<string, string>>(
    {}
  );
  const [currentFile, setCurrentFile] = useState<IRepoObject | null>(null);
  const [currentContent, setCurrentContent] = useState<string | null>(null);
  const [files, setFiles] = useState<IRepoObject[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    fetch(
      "http://localhost:8080/files/org/NUSpecialProjects/assignment/1/student/92pLytz-SgW~mKeuxDyuJg"
    )
      .then((response) => response.json())
      .then((data: IRepoObject[]) => {
        setFiles(data);
      });
  }, []);

  const openObject = (obj: IRepoObject) => {
    if (obj.type == "dir") {
      return openDir(obj);
    }
    return openFile(obj);
  };

  const openDir = (dir: IRepoObject) => {};

  const openFile = (file: IRepoObject) => {
    if (file.type == "dir") {
      return openDir;
    }

    setLoading(true);
    setCurrentFile(file);

    // Check if the content is already cached
    if (cachedContents[file.download_url]) {
      setCurrentContent(cachedContents[file.download_url]);
      setLoading(false);
      return;
    }

    fetch(file.download_url)
      .then((response) => response.text())
      .then((content) => {
        setCurrentContent(content);
        // Cache the content
        setCachedContents((prev) => ({
          ...prev,
          [file.download_url]: content,
        }));
      })
      .finally(() => {
        setLoading(false);
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
        <div className="Grader__files">
          {files.map((obj, i) => {
            return (
              <div
                className="Grader__file"
                key={i}
                onClick={() => {
                  openObject(obj);
                }}
              >
                {obj.name}
              </div>
            );
          })}
        </div>
        <CodeViewer
          className="Grader__code"
          showLineNumbers
          lineNumberStyle={{ color: "#999", margin: "0 5px" }}
          language="python"
          style={hybrid}
        >
          {currentContent ?? ""}
        </CodeViewer>
      </div>
    </div>
  );
};

export default Grader;
