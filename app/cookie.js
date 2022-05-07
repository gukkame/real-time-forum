function getCookie() {
    var decodedCookie = decodeURIComponent(document.cookie);
    if (decodedCookie != "") {
        var ca = decodedCookie.split(';');
        var ca1 = ca[0].split('=');
        var ca2 = ca1[1].split(':');
        ca2.reverse()
        return ca2;
    }
    return []
}

//checks cookie and updates it
function checkCookie() {
    var cookie = getCookie();
    var cookieCheck = true;
    if (cookie == "") {
        cookieCheck = false;

    } else {
        document.cookie =
            "session=" + cookie[1] + ":" + cookie[0] + ";" + "max-age=" + 20 * 60 + ";path=/";
        cookieCheck = true;
    }
    return cookieCheck
}

function removeCookie(value) {
    document.cookie = 'session=' + value + '; expires=Thu, 01 Jan 1970 00:00:00 UTC;';
    console.log("donee");
}

function dateToString(date) {
    console.log("date ", date);
    if (typeof date === "undefined") {
        date = Date.parse(date)
    }
    console.log("date222 ", date);
    return (
        date.getFullYear() +
        "-" +
        pad(date.getMonth() + 1) +
        "-" +
        pad(date.getDate())
    );
}

function pad(n, s = String(n)) {
    return s.length < 2 ? `0${s}` : s;
}

function getAge(dateString) {
    dateString = dateString.toString();
    var today = new Date();
    var birthDate = new Date(dateString);
    var age = today.getFullYear() - birthDate.getFullYear();
    var m = today.getMonth() - birthDate.getMonth();
    if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
        age--;
    }
    return age;
}
export { getCookie, removeCookie, dateToString, checkCookie, getAge }