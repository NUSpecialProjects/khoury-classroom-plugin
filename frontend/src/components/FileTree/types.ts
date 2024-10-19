interface IFileTreeNode {
  type: string;
  sha: string;
  childNodes: {
    [name: string]: IFileTreeNode;
  };
}

interface IFileTree extends React.HTMLProps<HTMLDivElement> {
  gitTree: IGitTreeNode[];
  selectFileCallback: Function;
}

interface IFileTreeDirectory extends React.HTMLProps<HTMLDivElement> {
  name: string;
}

interface IFileTreeFile extends React.HTMLProps<HTMLDivElement> {
  name: string;
}
