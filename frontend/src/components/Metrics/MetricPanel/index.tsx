import React from "react";
import './styles.css';

interface IMetricPanel {
    children: React.ReactNode;
}

const MetricPanel: React.FC<IMetricPanel> = ({children}) => {
   return (
    <div className="metricPanel">
        {children}
    </div>
   )
}

export default MetricPanel;