let manche1 = false;

const cardList = document.getElementsByClassName("card-list")[0];
const shuffle = document.getElementsByClassName("btn-shuffle")[0];
const reset = document.getElementsByClassName("btn-reset")[0];

const card2List = document.getElementsByClassName("card2-list")[0];

// equipe 1 shuffle ane take card to wirte answer...

reset.disabled = true;
if (manche1 == false) {
    shuffle.addEventListener("click", () => {
        cardList.classList.add("is-animated");
      
        shuffle.disabled = true;
        reset.disabled = false;
      });
    reset.addEventListener("click", () => {
        cardList.classList.remove("is-animated");
        card2List.classList.add("is-animated");
      
        reset.disabled = true;
        shuffle.disabled = false;
        manche1 = true;
      });
}