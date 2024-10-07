const work = document.getElementById("work");
const movement = document.getElementById("movement");

(async () => {
    try {
        const response = await fetch("results.json");
        const data = await response.json();
        work.innerText = data.work;
        movement.innerText = data.movement;
    } catch (err) {
        console.error("Could not fetch data:", err);
    }
})();
