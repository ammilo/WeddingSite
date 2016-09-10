$( document ).ready(
  function() {
    //window.alert("Hiya!");
  }
)

window.onblur = function(e) {
  document.title = "where did you go!?"
}

window.onfocus = function(e) {
  var n = Math.floor(Math.random() * 3 + 1);
  if (n == 1) {
    document.title = "i missed you!";
  } else if (n == 2) {
    document.title = "you came back!";
  } else {
    document.title = "why did you leave me!?";
  }
  setTimeout(function() {
    document.title = "home!";
  },
  3000)
}
