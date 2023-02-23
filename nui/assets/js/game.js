/*timer*/
const startButton = document.getElementById("start-button");
const timerDisplay = document.getElementById("timer");

let count = 60;

startButton.addEventListener("click", function() {
  const countdown = setInterval(function() {
    count--;
    timerDisplay.innerHTML = count < 10 ? "0" + count.toString() : count;

    if (count <= 10) {
      timerDisplay.style.color = "red";
    }

    if (count === 0) {
      clearInterval(countdown);
      resetTimer();
    }
  }, 1000);
});

function resetTimer() {
  count = 60;
  timerDisplay.innerHTML = count;
  timerDisplay.style.color = "black";
}
/* End of timer*/
/* flip card*/
var card = document.querySelector('.card');
/* End of flip card*/





/* Shuffle card*/
let manche1 = false;

const cardList = document.getElementsByClassName("card-list")[0];
const shuffle = document.getElementsByClassName("btn-shuffle")[0];
const reset = document.getElementsByClassName("btn-reset")[0];

const card2List = document.getElementsByClassName("card2-list")[0];

// equipe 1 shuffle ane take card to wirte answer...

reset.disabled = true;

shuffle.addEventListener("click", () => {
  cardList.classList.add("is-animated");
  shuffle.disabled = true;
  reset.disabled = false;
  setTimeout(function() {
    card.classList.toggle('is-flipped');
  }, 4000);
  });

reset.addEventListener("click", () => {
  cardList.classList.remove("is-animated");
  card2List.classList.add("is-animated");
      
  reset.disabled = true;
  shuffle.disabled = false;
  manche1 = true;
  });