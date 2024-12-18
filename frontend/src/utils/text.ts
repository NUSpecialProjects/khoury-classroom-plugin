//removes any underscores from a string
export const removeUnderscores = (text: string) => {
    return text
    .split('_')
    .map((str) => str.charAt(0).toUpperCase() + str.slice(1))
    .join(' ');
};