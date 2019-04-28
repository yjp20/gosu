document.addEventListener("DOMContentLoaded", () => {
  let fm = document.getElementsByClassName("flash-message")
  for (let i of fm) {
    i.addEventListener("click", (e) => {
      i.setAttribute("remove", "");
      setTimeout(() => {
        i.parentNode.removeChild(i);
      }, 200);
    });
  }
});
