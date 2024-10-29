const options: Intl.DateTimeFormatOptions = {
  year: "numeric",
  month: "short",
  day: "numeric",
  hour: "numeric",
  minute: "2-digit",
};

export const formatDate = (date: Date | null) => {
  return date ? new Date(date).toLocaleDateString("en-US", options) : "N/A";
};
