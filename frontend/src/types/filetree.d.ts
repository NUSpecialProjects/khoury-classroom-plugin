interface IFileTreeNode {
  type: string;
  name: string;
  path: string;
  sha: string;
  status: string;
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
  path: string;
  depth: number;
  status: string;
  treeDepth: number;
}

interface IFileTreeFile extends React.HTMLProps<HTMLDivElement> {
  name: string;
  path: string;
  depth: number;
  status: string;
}
