import { data } from './data';
import { Boxplot } from './Boxplot';
import "./styles.css"

export const GradeDistBoxPlot = ({ width = 700, height = 400 }) => (
  <div className="Chart">
    <p className="ChartTitle"><b>Current Grade Distribution</b></p>
    <Boxplot data={data} width={width} height={height} />
  </div>
);
