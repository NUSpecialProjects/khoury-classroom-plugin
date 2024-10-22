import * as d3 from "d3";
import { getSummaryStats } from "./summary-stats";

import { HorizontalBox } from "./HorizontalBox";
import { AxisBottom } from "./AxisBottom";

const MARGIN = { top: 30, right: 30, bottom: 30, left: 50 };

type BoxplotProps = {
  width: number;
  height: number;
  data: { name: string; value: number }[];
};

export const Boxplot = ({ width, height, data }: BoxplotProps) => {
  // The bounds (= area inside the axis) is calculated by substracting the margins from total width / height
  const boundsWidth = width - MARGIN.right - MARGIN.left;
  const boundsHeight = height - MARGIN.top - MARGIN.bottom;

  // Compute scales
  const xScale = d3.scaleLinear().domain([0, 100]).range([0, boundsWidth]);

  const yScale = d3.scaleBand().range([0, boundsHeight]);
  //.padding(0.5);

  // Build the box shapes
  const groupData = data.map((d) => d.value);
  const sumStats = getSummaryStats(groupData);

  if (!sumStats) {
    return null;
  }

  const { min, q1, median, q3, max, outliers } = sumStats;
  console.log(min);

  return (
    <div>
      <svg width={width} height={height}>
        <g
          width={boundsWidth}
          height={boundsHeight}
          transform={`translate(${[MARGIN.left, MARGIN.top].join(",")})`}
        >
          {/*<AxisLeft yScale={yScale} />*/}

          {/* X axis uses an additional translation to appear at the bottom */}
          <g
            className="BoxPlotAxes"
            transform={`translate(0, ${boundsHeight.toString()})`}
          >
            <AxisBottom
              xScale={xScale}
              height={boundsHeight}
              pixelsPerTick={80}
            />
          </g>
          <g>
            <HorizontalBox
              height={yScale.bandwidth()}
              q1={xScale(q1)}
              median={xScale(median)}
              q3={xScale(q3)}
              min={xScale(min)}
              max={xScale(max)}
              outliers={outliers.map((outlier) => xScale(outlier))} //Apply xScale to every point in this outliers array
              stroke="black"
              fill={"orange"}
            />
          </g>
        </g>
      </svg>
    </div>
  );
};
