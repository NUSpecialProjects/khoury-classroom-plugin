interface IAssignmentAcceptanceMetricsResponse {
  assignment_id: number;
  status: IAssignmentAcceptanceMetrics;
}

interface IAssignmentAcceptanceMetrics {
  accepted: number;
  not_accepted: number;
  in_grading: number;
  started: number;
  submitted: number;
}

interface IAssignmentGradedMetricsResponse {
  assignment_id: number;
  status: IAssignmentGradedMetrics;
}

interface IAssignmentGradedMetrics {
  graded: number;
  ungraded: number;
}

interface IChartJSData {
  labels: string[];
  datasets: [
    {
      backgroundColor: string | string[];
      data: number[];
    },
  ];
}

interface IMetric extends React.HTMLProps<HTMLDivElement> {
  title: string;
}
