const STROKE_WIDTH = 40;

// A reusable component that builds a horizontal box shape using svg
// Note: numbers here are px, not the real values in the dataset.

type HorizontalBoxProps = {
  min: number;
  q1: number;
  median: number;
  q3: number;
  max: number;
  height: number;
  outliers: number[]
  stroke: string;
  fill: string;
};

export const HorizontalBox = ({
  min,
  q1,
  median,
  q3,
  max,
  height,
  outliers,
  stroke,
  fill,
}: HorizontalBoxProps) => {
  const adjustedHeight = height / 2; 
  const verticalShift = height / 4;

  return (
    <>
      <line
        y1={adjustedHeight / 2  + verticalShift}
        y2={adjustedHeight / 2  + verticalShift}
        x1={min}
        x2={max}
        stroke={stroke}
        width={STROKE_WIDTH}
      />
      <rect
        x={q1}
        y={verticalShift}
        width={q3 - q1}
        height={adjustedHeight}
        stroke={stroke}
        fill={fill}
      />
      <line
        y1={verticalShift}
        y2={adjustedHeight + verticalShift}
        x1={median}
        x2={median}
        stroke={stroke}
        width={STROKE_WIDTH}
      />
      {outliers.map((outlier, index) => (
        <circle
          key={index}
          cx={outlier}
          cy={height / 2}
          r={STROKE_WIDTH / 12} 
          fill={stroke} 
        />
      ))}
     <line
        y1={verticalShift + (verticalShift * 2 / 3)}
        y2={adjustedHeight + (verticalShift/3)}
        x1={max}
        x2={max}
        stroke={stroke}
        width={STROKE_WIDTH}
      />
      <line
        y1={verticalShift + (verticalShift * 2 / 3)}
        y2={adjustedHeight + (verticalShift/3)}
        x1={min}
        x2={min}
        stroke={stroke}
        width={STROKE_WIDTH}
      />
    </>
  );
};
