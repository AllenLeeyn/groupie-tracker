const button = document.getElementById('changeStyleButton');
const stylesheet = document.getElementById('style');

(localStorage.getItem("cssFile") !== null)? stylesheet.href = localStorage.getItem("cssFile"): null;

button.addEventListener('click', () => {
    if (stylesheet.href.substring(stylesheet.href.length-"style-card-view.css".length) === 'style-card-view.css') {
        localStorage.setItem("cssFile", '/static/css/style-list-view.css')
    }
    else {
        localStorage.setItem("cssFile", '/static/css/style-card-view.css')
    }
    
    stylesheet.href = localStorage.getItem("cssFile")
});

/*------------------------------------------------------------------------------------------------*/

var yearsRangeInput0 = document.getElementById("yearsrange0")
var startDateOutput0 = document.getElementById("start-date0")
var endDateOutput0 = document.getElementById("end-date0")

var slider0 = document.getElementById("myRange0")

startDateOutput0.innerHTML = slider0.value; // Display the default slider value

yearsRangeInput0.oninput = function() {
    date = parseInt(startDateOutput0.innerHTML) + parseInt(this.value)
    if (!isNaN(date)) {
        latterDate = date.toString()
        endDateOutput0.innerHTML = latterDate
    }

    if (parseInt(endDateOutput0.innerHTML) < parseInt(startDateOutput0.innerHTML)) {
        [startDateOutput0.innerHTML, endDateOutput0.innerHTML] = [endDateOutput0.innerHTML, startDateOutput0.innerHTML] 
    }
    startDateOutput0.innerHTML = sanitize(startDateOutput0.innerHTML, 1963, 2018)
    endDateOutput0.innerHTML = sanitize(endDateOutput0.innerHTML, 1963, 2018)

    if (endDateOutput0.innerHTML[0] !== ' ')
        endDateOutput0.innerHTML = " " + endDateOutput0.innerHTML
    if (startDateOutput0.innerHTML[0] === ' ')
        startDateOutput0.innerHTML = startDateOutput0.innerHTML.substring(1)
}

slider0.oninput = function() {
    startDateOutput0.innerHTML = this.value;
    end = (parseInt(startDateOutput0.innerHTML) + parseInt(yearsRangeInput0.innerHTML)).toString()
    if (!isNaN(end)) {
        endDateOutput0.innerHTML = end
    }

    if (parseInt(endDateOutput0.innerHTML) < parseInt(startDateOutput0.innerHTML)) {
        [startDateOutput0.innerHTML, endDateOutput0.innerHTML] = [endDateOutput0.innerHTML, startDateOutput0.innerHTML] 
    }
    startDateOutput0.innerHTML = sanitize(startDateOutput0.innerHTML, 1963, 2018)
    endDateOutput0.innerHTML = sanitize(endDateOutput0.innerHTML, 1963, 2018)

    if (endDateOutput0.innerHTML[0] !== ' ')
        endDateOutput0.innerHTML = " " + endDateOutput0.innerHTML
    if (startDateOutput0.innerHTML[0] === ' ')
        startDateOutput0.innerHTML = startDateOutput0.innerHTML.substring(1)
}

/*---*/

var yearsRangeInput = document.getElementById("years range")
var startDateOutput = document.getElementById("start-date")
var endDateOutput = document.getElementById("end-date")

var slider = document.getElementById("myRange")

startDateOutput.innerHTML = slider.value; // Display the default slider value

yearsRangeInput.oninput = function() {
    date = parseInt(startDateOutput.innerHTML) + parseInt(this.value)
    if (!isNaN(date)) {
        latterDate = date.toString()
        endDateOutput.innerHTML = latterDate
    }

    if (parseInt(endDateOutput.innerHTML) < parseInt(startDateOutput.innerHTML)) {
        [startDateOutput.innerHTML, endDateOutput.innerHTML] = [endDateOutput.innerHTML, startDateOutput.innerHTML] 
    }
    startDateOutput.innerHTML = sanitize(startDateOutput.innerHTML, 1958, 2015)
    endDateOutput.innerHTML = sanitize(endDateOutput.innerHTML, 1958, 2015)

    if (endDateOutput.innerHTML[0] !== ' ')
        endDateOutput.innerHTML = " " + endDateOutput.innerHTML
    if (startDateOutput.innerHTML[0] === ' ')
        startDateOutput.innerHTML = startDateOutput.innerHTML.substring(1)
}

slider.oninput = function() {
    startDateOutput.innerHTML = this.value;
    end = (parseInt(startDateOutput.innerHTML) + parseInt(yearsRangeInput.innerHTML)).toString()
    if (!isNaN(end)) {
        endDateOutput.innerHTML = end
    }

    if (parseInt(endDateOutput.innerHTML) < parseInt(startDateOutput.innerHTML)) {
        [startDateOutput.innerHTML, endDateOutput.innerHTML] = [endDateOutput.innerHTML, startDateOutput.innerHTML] 
    }
    startDateOutput.innerHTML = sanitize(startDateOutput.innerHTML, 1958, 2015)
    endDateOutput.innerHTML = sanitize(endDateOutput.innerHTML, 1958, 2015)

    if (endDateOutput.innerHTML[0] !== ' ')
        endDateOutput.innerHTML = " " + endDateOutput.innerHTML
    if (startDateOutput.innerHTML[0] === ' ')
        startDateOutput.innerHTML = startDateOutput.innerHTML.substring(1)
}
/*---*/
function sanitize(x, minDate, maxDate) {
    if (parseInt(x) > maxDate) {
        return maxDate.toString()
    }

    if (parseInt(x) < minDate) {
        return minDate.toString()
    }
    return x
}


/*------------------------------------------------------------------------------------------------*/

const filterVisibility = document.getElementById('toggle');
(localStorage.getItem("filterToggle") === "true")? filterVisibility.checked = true: filterVisibility.checked = false;

filterVisibility.addEventListener('change', () => {
    if (filterVisibility.checked){
        localStorage.setItem("filterToggle", "true");
        filterVisibility.checked = true;
    }  else {
        localStorage.setItem("filterToggle", "false");
        filterVisibility.checked = false;
    }
});

/*------------------------------------------------------------------------------------------------*/

document.addEventListener("DOMContentLoaded", () =>{
    const sort = document.getElementById("sort");
    (localStorage.getItem("sortBy") !== null)? sort.value=localStorage.getItem("sortBy"): null;

    const container = document.querySelector(".container");
    const cards = Array.from(container.getElementsByClassName("card"));

    function compareName(a, b) {
        const nameA = a.querySelector("p").textContent.toLowerCase();
        const nameB = b.querySelector("p").textContent.toLowerCase();
        return nameA.localeCompare(nameB);
    }
    function compareDate(a, b) {
        const dateElementA = a.querySelector("p:nth-child(2)").textContent;
        const dateElementB = b.querySelector("p:nth-child(2)").textContent;
        return dateElementA.localeCompare(dateElementB);
    }
    function compareMembersCnt(a, b) {
        const membersCntA = a.querySelector("p:nth-child(3)").textContent;
        const membersCntB = b.querySelector("p:nth-child(3)").textContent;
        return membersCntA.localeCompare(membersCntB);
    }
    function compareFirstAlbum(a, b) {
        const firstAlbumA = a.querySelector("p:nth-child(4)").textContent.slice(-4);
        const firstAlbumB = b.querySelector("p:nth-child(4)").textContent.slice(-4);
        return firstAlbumA.localeCompare(firstAlbumB);
    }
    
    // Function to sort and display the cards
    function sortCards() {
        const sortCriteria = sort.value;
        let sortedCards;
        if (sortCriteria === "creation_date") {
            sortedCards = [...cards].sort(compareDate);
            localStorage.setItem("sortBy", "creation_date");
        } else if (sortCriteria === "name") {
            sortedCards = [...cards].sort(compareName);
            localStorage.setItem("sortBy", "name");
        } else if (sortCriteria === "membersCount") {
            sortedCards = [...cards].sort(compareMembersCnt);
            localStorage.setItem("sortBy", "membersCount");
        } else if (sortCriteria === "firstAlbum") {
            sortedCards = [...cards].sort(compareFirstAlbum);
            localStorage.setItem("sortBy", "firstAlbum");
        } else {
            sortedCards = cards; // Default order (no sorting)
            localStorage.setItem("sortBy", "default");
        }
        // Clear the container and append sorted cards
        container.innerHTML = '';
        sortedCards.forEach(card => container.appendChild(card));
    }

    sortCards();
    let sortedCards = Array.from(container.getElementsByClassName("card"));

    const switchOrder = document.getElementById("switch-order");
    (localStorage.getItem("order") !== null)? switchOrder.textContent=localStorage.getItem("order"): null;

    if (switchOrder.textContent === "▲") {
        sortedCards.reverse();
    }
    container.innerHTML = '';
    sortedCards.forEach(card => container.appendChild(card));

    const resultCount = document.getElementById('resultCount');
    resultCount.textContent = "Showing " + sortedCards.length + " artist(s)"

    // Attach the event listener to the dropdown
    sort.addEventListener("change", sortCards);
})

/*------------------------------------------------------------------------------------------------*/

function revCards(){
    const container = document.querySelector(".container");
    
    let cards = Array.from(container.getElementsByClassName("card"));
    cards.reverse();
    container.innerHTML = '';
    cards.forEach(card => container.appendChild(card));
    
    const switchOrder = document.getElementById("switch-order");

    switchOrder.textContent = (switchOrder.textContent === "▲") ? "▼" : "▲";
    localStorage.setItem("order", switchOrder.textContent);
}
