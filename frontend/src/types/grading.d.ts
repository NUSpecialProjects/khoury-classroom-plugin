/********************
 * Grading page types
 ********************/
interface IGradingCommentMap {
  [path: string]: { [line: number]: { [commentID: number]: IGradingComment } };
}
interface IGradingComment {
  path: string;
  line: number;
  body: string;
  points: number;
}

/******************************
 * GitHub response object types
 ******************************/
interface IGitDiff {
  Start: number;
  End: number;
}
interface IGitTreeNode {
  Status: {
    Status: string;
    Diff: IGitDiff[] | null;
  };
  Entry: {
    type: string;
    path: string;
    sha: string;
    status: string;
  };
}

interface IGitTreeResponse {
  tree: IGitTreeNode[];
}
