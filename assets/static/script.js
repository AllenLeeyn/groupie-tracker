document.addEventListener("DOMContentLoaded", () => {
    const sort = document.getElementById("sort");

    const container = document.querySelector(".container");
    const cards = Array.from(container.getElementsByClassName("card"));

    // Sorting function for names
    function compareName(a, b) {
        const nameA = a.querySelector("p").textContent.toLowerCase();
        const nameB = b.querySelector("p").textContent.toLowerCase();
        return nameA.localeCompare(nameB);
    }

    function compareDate(a, b) {
        const dateElementA = a.querySelector("p:nth-child(2)");
        const dateElementB = b.querySelector("p:nth-child(2)");
    
        const yearA = parseInt(dateElementA.textContent.split('Started ')[1], 10);
        const yearB = parseInt(dateElementB.textContent.split('Started ')[1], 10);

        return yearA - yearB; // Sort in ascending order
    }
    
    // Function to sort and display the cards
    function sortCards() {
        const sortCriteria = sort.value;
        let sortedCards;

        if (sortCriteria === "creation_date") {
            sortedCards = [...cards].sort(compareDate);
        } else if (sortCriteria === "name") {
            sortedCards = [...cards].sort(compareName);
        } else {
            sortedCards = cards; // Default order (no sorting)
        }

        // Clear the container and append sorted cards
        container.innerHTML = '';
        sortedCards.forEach(card => container.appendChild(card));
    }

    // Attach the event listener to the dropdown
    sort.addEventListener("change", sortCards);
});

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