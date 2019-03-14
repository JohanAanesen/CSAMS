/**
 * getTimeInNorwegian
 * Returns a date in norwegian time
 *
 * @return date in norwegian time
 * */
function getTimeInNorwegian() {
    let norTime = new Date().toLocaleString("no-no", {timeZone: "Europe/Oslo"});
    return new Date(norTime);
}

/**
 *
 * pad
 * Pads the numbers
 *
 * @param number
 * @return number padded
 * @link https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString
 * */
function pad(number) {
    if (number < 10) {
        return '0' + number;
    }
    return number;
}

/**
 *
 * toNORString
 *
 * Converts the date as it is time, it's like toISOString
 * only, it just converts the time to string without changing the time :)
 *
 * Inspired by:
 * @Link https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString
 * */
Date.prototype.toNORString = function () {
    return this.getFullYear() +
        '-' + pad((this.getMonth()) + 1) +
        '-' + pad(this.getDate()) +
        'T' + pad(this.getHours()) +
        ':' + pad(this.getMinutes());
};