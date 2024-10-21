interface IFileTreeNode {
  type: string;
  name: string;
  sha: string;
  childNodes: {
    [name: string]: IFileTreeNode;
  };
}

interface IFileTree extends React.HTMLProps<HTMLDivElement> {
  gitTree: IGitTreeNode[];
  selectFileCallback: (node: IFileTreeNode) => void;
}

interface IFileTreeDirectory extends React.HTMLProps<HTMLDivElement> {
  name: string;
  depth: number;
  treeDepth: number;
}

interface IFileTreeFile extends React.HTMLProps<HTMLDivElement> {
  name: string;
  depth: number;
}
