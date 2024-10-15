import { data } from './data';
import { Boxplot } from './Boxplot';
import "./styles.css"

export const BoxplotHorizontalDemo = ({ width = 700, height = 400 }) => (
  <div className="Chart">
  <Boxplot data={data} width={width} height={height} />
  {data[0].value}
  Width: {width}
  Height: {height}
  </div>
);
