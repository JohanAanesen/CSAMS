/**
 * copyLinks
 * Copies the join course url to clipboard for login and register
 *
 * @link https://hackernoon.com/copying-text-to-clipboard-with-javascript-df4d4988697f
 * @augments hash string
 * */
function copyLinks(hash) {

    // Create an element and attach to document body
    const element = document.createElement('textarea');

    // TODO brede : dynamically get root url
    element.value = 'http://localhost:8088/login?courseid=' + hash;
    document.body.appendChild(element);

    // select element and copy to clipboard
    element.select();
    document.execCommand("copy");

    // Remove element again
    document.body.removeChild(element);
}