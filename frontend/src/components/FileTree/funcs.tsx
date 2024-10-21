import { FileTreeDirectory, FileTreeFile } from ".";

export const buildTree = (tree1D: IGitTreeNode[]) => {
  const root: IFileTreeNode = {
    type: "tree",
    sha: "",
    name: "",
    childNodes: {},
  };
  tree1D.forEach((node) => {
    const path = node.path.split("/");
    let level: IFileTreeNode = root;

    path.forEach((seg, i) => {
      if (!(seg in level.childNodes)) {
        level.childNodes[seg] = {
          type: i === path.length - 1 ? node.type : "tree",
          sha: i === path.length - 1 ? node.sha : "",
          name: seg,
          childNodes: {},
        };
      }

      level = level.childNodes[seg];
    });
  });

  return root;
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
  name: string,
  depth: number,
  selectedSha: string,
  selectFileCallback: (node: IFileTreeNode) => void
) => {
  if (node.type === "blob") {
    return (
      <FileTreeFile
        className={selectedSha == node.sha ? "FileTreeFile--selected" : ""}
        key={name}
        depth={depth}
        name={name}
        onClick={() => {
          selectFileCallback(node);
        }}
      />
    );
  }

  // if not a blob (file), must be a tree (directory)
  return (
    <FileTreeDirectory key={name} name={name} depth={depth}>
      {sortTreeNode(node).map(([childName, childNode]) =>
        renderTree(
          childNode,
          childName,
          depth + 1,
          selectedSha,
          selectFileCallback
        )
      )}
    </FileTreeDirectory>
  );
};
