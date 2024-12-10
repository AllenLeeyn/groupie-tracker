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

var slider = document.getElementById("myRange")
    var output = document.getElementById("range-current-value");
    output.innerHTML = slider.value; // Display the default slider value

    // Update the current slider value (each time you drag the slider handle)
    slider.oninput = function() {
        output.innerHTML = this.value;
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
        } else if (sortCriteria === "name") {
            sortedCards = [...cards].sort(compareName);
        } else if (sortCriteria === "membersCount") {
            sortedCards = [...cards].sort(compareMembersCnt);
        } else if (sortCriteria === "firstAlbum") {
            sortedCards = [...cards].sort(compareFirstAlbum);
        } else {
            sortedCards = cards; // Default order (no sorting)
        }
        // Clear the container and append sorted cards
        container.innerHTML = '';
        sortedCards.forEach(card => container.appendChild(card));
    }
    // Attach the event listener to the dropdown
    sort.addEventListener("change", sortCards);
})

function revCards(){
    const container = document.querySelector(".container");
    let cards = Array.from(container.getElementsByClassName("card"));
    cards.reverse();
    container.innerHTML = '';
    cards.forEach(card => container.appendChild(card));
    const switchOrder = document.getElementById("switch-order");
    if (switchOrder) {
        switchOrder.value = (switchOrder.value === "▲") ? "▼" : "▲";
        switchOrder.textContent = switchOrder.value;
    }
}
