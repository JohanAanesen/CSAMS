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

    // Dynamically get root url
    const rootURL = window.location.origin;

    // Set value to element
    element.value = rootURL + '/login?courseid=' + hash;

    // Append to body
    document.body.appendChild(element);

    // select element and copy to clipboard
    element.select();
    document.execCommand("copy");

    // Remove element again
    document.body.removeChild(element);
}