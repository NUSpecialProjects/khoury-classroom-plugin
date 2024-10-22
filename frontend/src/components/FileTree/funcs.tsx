import { FileTreeDirectory, FileTreeFile } from ".";

export const buildTree = (tree1D: IGitTreeNode[]) => {
  let treeDepth = 0;
  const root: IFileTreeNode = {
    type: "tree",
    sha: "",
    name: "",
    path: "",
    childNodes: {},
  };
  tree1D.forEach((node) => {
    const fullPath = node.path.split("/");
    let level: IFileTreeNode = root;
    treeDepth = Math.max(treeDepth, fullPath.length);

    let path = "";
    fullPath.forEach((seg, i) => {
      path += "/" + seg;
      if (!(seg in level.childNodes)) {
        level.childNodes[seg] = {
          type: i === fullPath.length - 1 ? node.type : "tree",
          sha: i === fullPath.length - 1 ? node.sha : "",
          name: seg,
          path: path.substring(1),
          childNodes: {},
        };
      }

      level = level.childNodes[seg];
    });
  });

  return { root, treeDepth };
};

export const sortTreeNode = (node: IFileTreeNode) => {
  return Object.entries(node.childNodes).sort(
    ([nameA, nodeA], [nameB, nodeB]) => {
      // directories before file
      if (nodeA.type == "tree" && nodeB.type == "blob") return -1;
      // files after directories
      if (nodeA.type == "blob" && nodeB.type == "tree") return 1;
      // sort by alphabetical order afterwards
      return nameA.localeCompare(nameB);
    }
  );
};

// iterate through a tree and render appropriate components
export const renderTree = (
  node: IFileTreeNode,
  depth: number,
  treeDepth: number,
  selectedFile: string,
  selectFileCallback: (node: IFileTreeNode) => void
) => {
  if (node.type === "blob") {
    return (
      <FileTreeFile
        className={selectedFile == node.path ? "FileTreeFile--selected" : ""}
        key={node.path}
        depth={depth}
        name={node.name}
        path={node.path}
        onClick={() => {
          selectFileCallback(node);
        }}
      />
    );
  }

  // if not a blob (file), must be a tree (directory)
  return (
    <FileTreeDirectory
      key={node.name}
      name={node.name}
      depth={depth}
      treeDepth={treeDepth}
    >
      {sortTreeNode(node).map(([_, childNode]) =>
        renderTree(
          childNode,
          depth + 1,
          treeDepth,
          selectedFile,
          selectFileCallback
        )
      )}
    </FileTreeDirectory>
  );
};
