import { ITableRow } from "@/components/Table/types";

export interface IGradingSubmissionRow {
  name: string;
  score: number | null;
  maxScore: number;
}

export interface IGradingAssignmentRow extends ITableRow {
  submissions: IGradingSubmissionRow[];
}
