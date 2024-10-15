import "./styles.css";

const FileTree: React.FC<React.HTMLProps<HTMLDivElement>> = ({ children }) => {
  return <div className="FileTree">{children}</div>;
};

const Directory: React.FC<React.HTMLProps<HTMLDivElement>> = ({ children }) => {
  return <div className="FileTreeDirectory">{children}</div>;
};

const File: React.FC<React.HTMLProps<HTMLDivElement>> = ({ children }) => {
  return <div className="FileTreeFile">{children}</div>;
};

export default FileTree;
