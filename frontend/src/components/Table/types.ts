export interface ITable extends React.HTMLProps<HTMLDivElement> {
  primaryCol?: number;
  cols: number;
}

export interface ITableRow extends React.HTMLProps<HTMLDivElement> {}

export interface ITableCell extends React.HTMLProps<HTMLDivElement> {}
