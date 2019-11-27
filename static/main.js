let encodedUrl = "";
let encodedFilter = "";

function rot13(str) {
    var input     = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
    var output    = 'NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm';
    var index     = x => input.indexOf(x);
    var translate = x => index(x) > -1 ? output[index(x)] : x;
    return str.split('').map(translate).join('');
}

function updateLink() {
    var button = document.getElementById("go");
    var elem = document.getElementById("link");

    let finalLink = "http://127.0.0.1:8080/proxy?u=" + encodedUrl + "&c=" + encodedFilter;

    button.setAttribute("href", finalLink);
    elem.setAttribute("href", finalLink);
    elem.innerText = finalLink;
}

function urlChanged() {
    var url = document.getElementById("url").value;
    encodedUrl = encodeURIComponent(rot13(url));

    updateLink();
}

function filterChanged() {
    var filters = document.getElementById("filtered").value;
    var expressions = filters.split(/[\r\n]+/);
    var golangRegex = "(?i)(" + expressions.join("|") + ")";
    encodedFilter = encodeURIComponent(rot13(golangRegex));

    updateLink();
}

urlChanged();
filterChanged();