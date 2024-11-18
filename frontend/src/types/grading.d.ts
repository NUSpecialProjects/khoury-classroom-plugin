/********************
 * Grading page types
 ********************/
interface IGradingFeedbackMap {
  [commentID: number]: IGradingFeedback;
}
interface IGradingFeedback {
  path: string;
  line: number;
  body: string;
  points: number;
}

/******************************
 * GitHub response object types
 ******************************/
interface IGitDiff {
  start: number;
  end: number;
}
interface IGitTreeNode {
  status: {
    status: string;
    diff: IGitDiff[] | null;
  };
  entry: {
    type: string;
    path: string;
    sha: string;
    status: string;
  };
}

interface IGitTreeResponse {
  tree: IGitTreeNode[];
}
