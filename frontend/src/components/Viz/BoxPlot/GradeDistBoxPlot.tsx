import { useEffect, useState } from 'react';
import { Boxplot, GradeEntry } from './Boxplot';
import './styles.css';
import { getGrades } from '@/api/assignment_requests';

export const GradeDistBoxPlot = ({ width = 700, height = 400, assID=1}) => {
  const [grades, setGrades] = useState<GradeEntry[]>([]);

  useEffect(() => {
    const fetchGrades = async () => {
      try {
        const tempData: GradeEntry[] = await getGrades(assID);
        setGrades(tempData);
      } catch (err) {
      }
    };

    fetchGrades(); 
  }, []);


  return (
    <div className="Chart">
      <p className="ChartTitle"><b>Current Grade Distribution</b></p>
      <Boxplot data={grades} width={width} height={height} />
    </div>
  );
};
