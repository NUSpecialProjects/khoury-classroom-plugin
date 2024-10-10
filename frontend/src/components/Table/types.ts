export interface ITable extends React.HTMLProps<HTMLDivElement> {
  primaryCol?: number | null;
}

export interface ITableRow extends ITable {
  labelRow?: boolean;
}

export interface ITableCell extends React.HTMLProps<HTMLDivElement> {
  primary?: boolean;
}
