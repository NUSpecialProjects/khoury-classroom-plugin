interface IFileTreeNode {
  type: string;
  name: string;
  path: string;
  sha: string;
  diff: IGitDiff[] | null;
  status: string;
  childNodes: {
    [name: string]: IFileTreeNode;
  };
}

interface IFileTree extends React.HTMLProps<HTMLDivElement> {
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
