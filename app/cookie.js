function getCookie() {
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    var ca1 = ca[0].split('=');
    return ca1;
}

//checks cookie and updates it
function checkCookie() {
    var cookie = getCookie();
    var cookieCheck = true;
    if (cookie == "") {
        cookieCheck = false;

    } else {
        document.cookie =
            cookie[0] + "=" + cookie[1] + ";" + "max-age=" + 20 * 60 + ";path=/";
        cookieCheck = true;
    }

    // console.log(cookieCheck);

    return cookieCheck
}

function removeCookie(name) {
    document.cookie = name + '=; Max-Age=-99999999;';
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
    console.log("date ", dateString);
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