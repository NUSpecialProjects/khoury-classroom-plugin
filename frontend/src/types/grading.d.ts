/******************************
 * Grading page component types
 ******************************/
interface IGradingSubmissionRow {
  name: string;
  score: number | null;
  maxScore: number;
}

interface IGradingAssignmentRow extends React.HTMLProps<HTMLDivElement> {
  assignmentId: number;
}

/********************
 * Grading page types
 ********************/
interface IGradingComment {
  path: string;
  line: number;
  body: string;
}

/******************************
 * GitHub response object types
 ******************************/
interface IGitTreeNode {
  Status: string;
  Entry: {
    type: string;
    path: string;
    sha: string;
    status: string;
  };
}

interface IGraderFile {
  name: string;
  content: string;
}
