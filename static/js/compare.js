// Standard Cookie Helpers
function setCookie(name, value, days = 7) {
    const date = new Date();
    date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
    const expires = "; expires=" + date.toUTCString();
    document.cookie = name + "=" + (value || "") + expires + "; path=/; SameSite=Strict";
}

function getCookie(name) {
    const nameEQ = name + "=";
    const ca = document.cookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') c = c.substring(1, c.length);
        if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
    }
    return null;
}

function toggleCompare(button) {
    const carId = button.getAttribute('data-car-id');
    const isSelected = button.getAttribute('data-selected') === 'true';

    const nextState = !isSelected;

    let savedIdsCookie = getCookie("compare_cars");
    let comparedIds = savedIdsCookie ? savedIdsCookie.split(",") : [];

    if (nextState) {
        if (!comparedIds.includes(carId)) {
            comparedIds.push(carId);
        }
    } else {
        comparedIds = comparedIds.filter(id => id !== carId);
    }

    if (comparedIds.length > 0) {
        setCookie("compare_cars", comparedIds.join(","), 7);
    } else {
        setCookie("compare_cars", "", -1);
    }

    button.setAttribute('data-selected', nextState);

    const icon = button.querySelector('.compare-icon');
    if (nextState) {
        icon.classList.remove('fa-regular', 'fa-square');
        icon.classList.add('fa-solid', 'fa-square-check');
    } else {
        icon.classList.remove('fa-solid', 'fa-square-check');
        icon.classList.add('fa-regular', 'fa-square');
    }
}