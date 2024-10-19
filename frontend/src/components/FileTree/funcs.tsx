import { FileTreeDirectory, FileTreeFile } from ".";

export const buildTree = (tree1D: IGitTreeNode[]) => {
  let root: IFileTreeNode = {
    type: "tree",
    sha: "",
    childNodes: {},
  };
  tree1D.forEach((node) => {
    const path = node.path.split("/");
    let level: IFileTreeNode = root;

    path.forEach((dir, i) => {
      if (!level.childNodes[dir]) {
        level.childNodes[dir] = {
          type: i === path.length - 1 ? node.type : "tree",
          sha: i === path.length - 1 ? node.sha : "",
          childNodes: {},
        };
      }

      level = level.childNodes[dir];
    });
  });

  return root;
};

export const sortTreeNode = (node: IFileTreeNode) => {
  return Object.entries(node.childNodes).sort(
    ([nameA, nodeA], [nameB, nodeB]) => {
      console.log(nodeA.type + ":" + nodeB.type);
      // Sort directories first and then by name
      if (nodeA.type === "tree" && nodeB.type === "blob") {
        console.log("test");
        return -1;
      }
      // dir before file
      if (nodeA.type === "blob" && nodeB.type === "tree") return 1; // file after dir
      return nameA.localeCompare(nameB); // alphabetical order
    }
  );
};

// iterate through a tree and render appropriate components
export const renderTree = (
  node: IFileTreeNode,
  name: string,
  selectFileCallback: Function
) => {
  if (node.type === "blob") {
    return (
      <FileTreeFile
        key={name}
        name={name}
        onClick={() => {
          selectFileCallback(node.sha);
        }}
      />
    );
  }

  // if not a blob (file), must be a tree (directory)
  return (
    <FileTreeDirectory key={name} name={name}>
      {sortTreeNode(node).map(([childName, childNode]) =>
        renderTree(childNode, childName, selectFileCallback)
      )}
    </FileTreeDirectory>
  );
};
