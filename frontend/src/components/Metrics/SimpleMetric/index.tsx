import React from "react";
import './styles.css';
import '../styles.css';

const SimpleMetric : React.FC<ISimpleMetric> = ({metricTitle, metricValue}) => {
    return (
        <div className="metric">
            <div className="metricTitle">{metricTitle}</div>
            <div className="simpleMetric__value">
                {metricValue}
            </div>
        </div>
    )
}

export default SimpleMetric;