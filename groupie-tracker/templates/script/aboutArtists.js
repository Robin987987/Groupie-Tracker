window.onload = function() {
createModalWindow();
}

function createModalWindow() {
    document.querySelector(".artistBox").addEventListener("click", openModal);
}

function openModal() {
    document.querySelector("#modal").showModal();
    document.querySelector("#modal").addEventListener("click", closeModal);
}

function closeModal() {
    document.querySelector("#modal").close();
}