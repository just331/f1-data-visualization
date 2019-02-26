document.addEventListener("keydown", function(e) {
  if (e.key === "ArrowUp" || e.key === "ArrowDown") {
    var ele = document.getElementById('year');
    e.key === "ArrowUp" ? ele.innerHTML++ : ele.innerHTML--;
  } else if (e.key === "ArrowLeft" || e.key === "ArrowRight") {
    var ele = document.getElementById('circuit');
    if (e.key === "ArrowLeft") {
      ele.innerHTML = "prev"
    } else {
      ele.innerHTML = "next"
    }
  }
});
