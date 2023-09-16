const postContainer = document.querySelector("#posts");
const layoutSwitchBtn = document.querySelector("#layoutToggle");
const indexHeader = document.querySelector(".index-list-header");

function getCookie(cookieName) {
  const cookies = document.cookie.split("; ");
  for (let i = 0; i < cookies.length; i++) {
    const cookie = cookies[i].split("=");
    const name = cookie[0];
    const value = cookie[1];
    if (name === cookieName) {
      return decodeURIComponent(value);
    }
  }
  return null;
}

const layoutCookie = getCookie("layout");

if (layoutCookie == "compact") {
  postContainer.classList.add("compact");
  layoutSwitchBtn.classList.add("compact");
  indexHeader.classList.add("compact");
} else {
  layoutSwitchBtn.classList.add("full");
}

layoutSwitchBtn.addEventListener("click", () => {
  const innerCookie = getCookie("layout");

  switch (innerCookie) {
    case "full":
      layoutSwitchBtn.classList.add("compact");
      layoutSwitchBtn.classList.remove("full");
      postContainer.classList.add("compact");
      indexHeader.classList.add("compact");
      break;
    case "compact":
      layoutSwitchBtn.classList.add("full");
      layoutSwitchBtn.classList.remove("compact");
      postContainer.classList.remove("compact");
      indexHeader.classList.remove("compact");
      break;
  }
});
