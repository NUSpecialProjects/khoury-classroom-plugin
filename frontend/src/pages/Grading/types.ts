export interface IGradingSubmissionRow {
  name: string;
  score: number | null;
  maxScore: number;
}

export interface IGradingAssignmentRow extends React.HTMLProps<HTMLDivElement> {
  submissions: IGradingSubmissionRow[];
}
