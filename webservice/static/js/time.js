/**
 * getTimeNorwegian
 * Returns a date in norwegian time
 *
 * @return date in norwegian time
 * */
function getTimeNorwegian() {
    let date = new Date();
    let norwegianDate = Date.UTC(date.getUTCFullYear(), date.getUTCMonth(), date.getUTCDate(), date.getUTCHours() + 1, date.getUTCMinutes(), date.getUTCSeconds());
    return new Date(norwegianDate);
}