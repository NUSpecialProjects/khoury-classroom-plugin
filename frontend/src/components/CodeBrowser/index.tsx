import { useState, useEffect, useContext } from "react";
import Prism from "prismjs";

import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import {
  ext2lang,
  extractExtension,
  ext2langLoader,
  dependencies,
} from "@/utils/prism-lang-loader";
import { getFileBlob } from "@/api/grading";
import CodeLine from "./CodeLine";

import "@/assets/prism-vs-dark.css";
import "./styles.css";

interface ICodeBrowser extends React.HTMLProps<HTMLDivElement> {
  assignmentID: string | undefined;
  studentWorkID: string | undefined;
  file: IFileTreeNode | null;
}

const CodeBrowser: React.FC<ICodeBrowser> = ({
  assignmentID,
  studentWorkID,
  file,
  className,
  ...props
}) => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  const [fileContents, setFileContents] = useState<React.ReactNode>();
  const [cachedFileContents, setCachedFileContents] = useState<
    Record<string, React.ReactNode>
  >({});

  // get file contents
  useEffect(() => {
    if (!selectedClassroom || !assignmentID || !studentWorkID || !file) return;

    // Check if the content is already cached
    if (file.sha in cachedFileContents) {
      setFileContents(cachedFileContents[file.sha]);
      return;
    }

    (async () => {
      if (!selectedClassroom || !assignmentID || !studentWorkID) return;
      let wrapped: React.ReactNode;
      try {
        const blob = await getFileBlob(
          selectedClassroom.id,
          Number(assignmentID),
          Number(studentWorkID),
          file.sha
        );

        const highlighted = await highlightCode(blob);
        wrapped = await wrapCode(highlighted);
      } catch {
        wrapped = <></>;
      } finally {
        setFileContents(wrapped);
        setCachedFileContents((prev) => ({
          ...prev,
          [file.sha]: wrapped,
        }));
      }
    })();
  }, [assignmentID, studentWorkID, file]);

  // when a new file is selected, import any necessary
  // prismjs language syntax files and trigger a rehighlight
  const highlightCode = async (code: string) => {
    if (!file) return "";

    const lang = ext2lang[extractExtension(file.name)];
    try {
      const deps: string | string[] = dependencies[lang];
      if (deps) {
        if (typeof deps === "string") {
          await ext2langLoader[deps]();
        }
        if (Array.isArray(deps)) {
          for (const dep of deps) {
            await ext2langLoader[dep]();
          }
        }
      }
      await ext2langLoader[lang]();
    } catch {
      // Prism does not support language or mapping does not exist
      return code;
    }
    return Prism.highlight(code, Prism.languages[lang], lang);
  };

  const wrapCode = async (code: string) => {
    if (!file) return <></>;

    const lines = code.split("\n");

    // MEMOIZATION :D
    let memo = [];
    if (file.diff) {
      memo = Array(lines.length).fill(0);
      for (const diff of file.diff) {
        memo[diff.Start - 1]++;
        memo[diff.End - 1]--;
      }
    }

    const wrappedLines: React.ReactNode[] = [];
    for (let i = 0; i < lines.length; i++) {
      if (file.diff && memo && i > 0) memo[i] += memo[i - 1];
      wrappedLines.push(
        <CodeLine
          key={i}
          path={file.path}
          line={i + 1}
          isDiff={(file.diff && memo && memo[i] > 0) ?? false}
          code={lines[i]}
        />
      );
    }

    return <>{wrappedLines}</>;
  };

  return (
    <div
      className={"CodeBrowser" + (className ? " " + className : "")}
      {...props}
    >
      <pre>
        <code
          data-diff={JSON.stringify(file?.diff ?? "")}
          className={
            file
              ? "language-" + ext2lang[extractExtension(file.name)]
              : "language-undefined"
          }
        >
          {file ? fileContents : "Select a file to view its contents."}
        </code>
      </pre>
    </div>
  );
};

export default CodeBrowser;
