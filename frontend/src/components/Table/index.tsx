import React from "react";

import { ITable, ITableRow, ITableCell } from "./types";

import "./styles.css";

const Table: React.FC<ITable> = (props) => {
  return (
    <div
      {...props}
      className={"Table" + (props.className ? " " + props.className : "")}
    >
      {props.children}
    </div>
  );
};

const TableRow: React.FC<ITableRow> = (props) => {
  return (
    <div className="TableRow__group">
      <div
        {...props}
        className={
          "TableRow" +
          (props.labelRow ? " TableRow--labelRow" : "") +
          (props.className ? " " + props.className : "")
        }
      >
        {props.children}
      </div>
    </div>
  );
};

const TableCell: React.FC<ITableCell> = (props) => {
  return (
    <div
      {...props}
      className={
        "TableCell" +
        (props.primary ? " TableCell--primary" : "") +
        (props.className ? " " + props.className : "")
      }
    >
      {props.children}
    </div>
  );
};

export { Table, TableRow, TableCell };
