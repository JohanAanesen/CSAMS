/**
 * getTimeNorwegian
 * Returns a date in norwegian time
 *
 * @return date in norwegian time
 * */
function getTimeNorwegian() {
    return new Date().getTime() + (1000 * 60 * 60);
}