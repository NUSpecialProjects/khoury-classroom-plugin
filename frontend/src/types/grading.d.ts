/******************************
 * Grading page component types
 ******************************/
interface IGradingSubmissionRow {
  name: string;
  score: number | null;
  maxScore: number;
}

interface IGradingAssignmentRow extends React.HTMLProps<HTMLDivElement> {
  submissions: IGradingSubmissionRow[];
}

/******************************
 * GitHub response object types
 ******************************/
interface IGitTreeNode {
  type: string;
  path: string;
  sha: string;
}

interface IGraderFile {
  name: string;
  content: string;
}
