const optionsDate: Intl.DateTimeFormatOptions = {
  year: "numeric",
  month: "short",
  day: "numeric",
};

const optionsDateTime: Intl.DateTimeFormatOptions = {
  year: "numeric",
  month: "short",
  day: "numeric",
  hour: "numeric",
  minute: "2-digit",
};

//formats a date with only the day, month and year
export const formatDate = (date: Date | undefined) => {
  return date ? new Date(date).toLocaleDateString("en-US", optionsDate) : "N/A";
};

//formats a date with the date and timestamp
export const formatDateTime = (date: Date | undefined) => {
  return date ? new Date(date).toLocaleDateString("en-US", optionsDateTime) : "N/A";
};
