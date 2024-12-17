import React from "react";
import "./styles.css";

const Metric: React.FC<IMetric> = ({ title, className, children }) => {
  return (
    <div className={`Metric ${className ?? ""}`}>
      <div className="Metric__title">{title}</div>
      <div className="Metric__content">{children}</div>
    </div>
  );
};

export default Metric;
