const work = document.getElementById("work");
const movement = document.getElementById("movement");

(async () => {
    const data = await fetch("results.json");
    const results = await data.json();
    work.innerText = results.work;
    movement.innerText = results.movement;
})();
