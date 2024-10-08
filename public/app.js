const work = document.getElementById("work");
const movement = document.getElementById("movement");

async function renderFocus() {
    try {
        const response = await fetch("results.json");
        if (!response.ok) {
            throw new Error("no response");
        }
        const data = await response.json();
        work.innerText = await data.work;
        movement.innerText = await data.movement;
    } catch (err) {
        console.error("Could not fetch data:", err);
    }
}

document.addEventListener("DOMContentLoaded", async () => {
    await renderFocus();
    document.querySelector("main.content").style.display = "block";
});
