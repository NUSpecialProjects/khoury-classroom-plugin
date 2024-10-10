import React from "react";

import { ITable, ITableRow, ITableCell } from "./types";

import "./styles.css";

const Table: React.FC<ITable> = ({ children, primaryCol = null }) => {
  return (
    <div className="Table">
      {React.Children.map(children, (child) => {
        if (!React.isValidElement(child) || child.type !== TableRow) {
          return child;
        }
        return React.cloneElement(child as React.ReactElement<ITableRow>, {
          primaryCol,
        });
      })}
    </div>
  );
};

const TableRow: React.FC<ITableRow> = ({
  children,
  primaryCol = null,
  labelRow = false,
}) => {
  return (
    <div className="TableRow__group">
      <div className={"TableRow" + (labelRow ? " TableRow--labelRow" : "")}>
        {React.Children.map(children, (child, index) => {
          if (!React.isValidElement(child) || child.type !== TableCell) {
            return child;
          }
          return React.cloneElement(child as React.ReactElement<ITableCell>, {
            primary: index === primaryCol,
          });
        })}
      </div>
    </div>
  );
};

const TableCell: React.FC<ITableCell> = ({ children, primary = false }) => {
  return (
    <div className={"TableCell" + (primary ? " TableCell--primary" : "")}>
      {children}
    </div>
  );
};

export { Table, TableRow, TableCell };
