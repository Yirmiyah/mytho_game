/*timer*/
const startButton = document.getElementById("start-button");
const timerDisplay = document.getElementById("timer");
let fakeAns1 = document.getElementById("bluff1")
let fakeAns2 = document.getElementById("bluff2")



let count = 60;

startButton.addEventListener("click", function () {
  const countdown = setInterval(function () {
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
function fadeInImage() {
  let image = document.getElementById("image");
  let txt = document.getElementById("txt");
  let txt2 = document.getElementById("txt2");
  let div = document.querySelector(".fr")
  let field = document.querySelector(".field")
  let bluff = document.querySelector(".bluff")

  image.style.opacity = 1;
  div.style.opacity = 1;
  field.style.opacity = 1;
  txt.style.opacity = 1;
  txt2.style.opacity = 1;
  bluff.style.opacity = 1;
}
/* End of flip card*/


/* Shuffle card*/
let manche1 = false;

const cardList = document.getElementsByClassName("card-list")[0];
const shuffle = document.getElementsByClassName("btn-shuffle")[0];
const reset = document.getElementsByClassName("btn-reset")[0];
const card2List = document.getElementsByClassName("card2-list")[0];

// equipe 1 shuffle ane take card to write answer...

reset.disabled = true;

shuffle.addEventListener("click", () => {
  cardList.classList.add("is-animated");
  shuffle.disabled = true;
  reset.disabled = false;
  setTimeout(function () {
    fadeInImage();
  }, 3000);
});

reset.addEventListener("click", () => {
  cardList.classList.remove("is-animated");
  card2List.classList.add("is-animated");

  reset.disabled = true;
  shuffle.disabled = false;
  manche1 = true;
});

/*Randomiser les Questions/Réponses */


function PickQRL() {

  fetch('quizz.json')
    .then(response => response.json())
    .then(data => {
      // console.log(data);
      let Questions = data.Question
      let Reponses = data.Response
      let Level = data.Level

      numberRand = Math.floor(Math.random() * 93);
      // Questions[numberRand]
      // Reponses[numberRand]
      // Level[numberRand]

      console.log("Questions: " + Questions[numberRand])
      console.log("Responses: " + Reponses[numberRand])
      console.log("Level: " + Level[numberRand])

      const images = [
        "../../data/frontcard.png",
        "../../data/frontcard2.png",
        "../../data/frontcard3.png",
      ];

      if (Level[numberRand] == "Débutant") {
        document.getElementById("image").src = images[0];
      } else if (Level[numberRand] == "Intermédiaire") {
        document.getElementById("image").src = images[1];
      } else if (Level[numberRand] == "Expert") {
        document.getElementById("image").src = images[2];
      }


      // const Q = `${Questions}`;
      // const R = `${Reponses}`;
      // document.getElementById("Question").innerHTML = Q;
      // document.getElementById("Reponse").innerHTML = R;
      console.log("Questions: " + Questions[numberRand])
      console.log("Reponses: " + Reponses[numberRand])
      document.querySelector(".front").innerHTML = `${Questions[numberRand]}`;
      document.querySelector(".good-rep").innerHTML = `${Reponses[numberRand]}`;



    })
    .catch(error => {
      console.error('Une erreur est survenue', error);
    });
}

const StartTeam1 = document.getElementById("start-button")

StartTeam1.addEventListener("click", (e) => {
  if (e) {
    PickQRL()
  }
})



// const form = document.querySelector('.avatar');

// form.addEventListener('submit', function(event) {
//   event.preventDefault();

//   // Récupère la valeur de l'élément input radio sélectionné
//   const radios = document.getElementsByName('position');
//   let selectedValue;
//   for (const radio of radios) {
//     if (radio.checked) {
//       selectedValue = radio.value;
//       break;
//     }
//   }

//   console.log(selectedValue); // Affiche la valeur de l'élément sélectionné
// });




