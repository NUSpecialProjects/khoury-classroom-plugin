import "./styles.css"; // Assuming you will style this with CSS
import React from "react";

interface ITable extends React.HTMLProps<HTMLDivElement> {
  primaryCol?: number;
  cols: number;
}

const Table: React.FC<ITable> = ({
  primaryCol = 0,
  cols,
  className,
  style,
  ...props
}) => {
  const columns = Array(cols).fill("auto");
  columns[primaryCol] = "1fr"; // Make the specified column stretch

  const gridTemplateColumns = columns.join(" ");
  return (
    <div
      className={"Table" + (className ? " " + className : "")}
      style={{ gridTemplateColumns, ...style }}
      {...props}
    >
      {props.children}
    </div>
  );
};

const TableRow: React.FC<React.HTMLProps<HTMLDivElement>> = ({
  className,
  children,
  ...props
}) => {
  return (
    <div className={"TableRow" + (className ? " " + className : "")}>
      {React.Children.map(children, (child) =>
        React.isValidElement(child)
          ? React.cloneElement(child, { ...props })
          : child
      )}
    </div>
  );
};
const TableCell: React.FC<React.HTMLProps<HTMLDivElement>> = ({
  className,
  ...props
}) => {
  return (
    <div
      className={"TableCell" + (className ? " " + className : "")}
      {...props}
    >
      {props.children}
    </div>
  );
};
const TableDiv: React.FC<React.HTMLProps<HTMLDivElement>> = ({
  className,
  ...props
}) => {
  return (
    <div className={"TableDiv" + (className ? " " + className : "")} {...props}>
      {props.children}
    </div>
  );
};

export { Table, TableRow, TableCell, TableDiv };
