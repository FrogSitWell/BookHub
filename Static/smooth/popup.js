const popup = document.querySelector(".popup-content");
const overlay = document.getElementById("overlay");
const closePopupButton = document.getElementById("closePopup");

// Mở popup
document.getElementById("openPopup").addEventListener("click", () => {
  popup.style.display = "block";
  overlay.style.display = "block";
});

// Đóng popup
closePopupButton.addEventListener("click", () => {
  popup.style.display = "none";
  overlay.style.display = "none";
});

// Đóng popup khi nhấn vào overlay
overlay.addEventListener("click", () => {
  popup.style.display = "none";
  overlay.style.display = "none";
});
// Đóng popup nếu người dùng nhấn ra ngoài popup
window.onclick = function (event) {
  const popup = document.getElementById("popup");
  if (event.target === popup) {
    popup.style.display = "none";
  }
};
