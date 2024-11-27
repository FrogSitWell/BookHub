// Hiển thị popup
document
  .querySelector(".action-buttons")
  .addEventListener("click", function () {
    document.getElementById("configPopup").style.display = "block";
  });

// Đóng popup
function closePopup() {
  document.getElementById("configPopup").style.display = "none";
}

// Áp dụng cấu hình ngoài popup
function applySettings() {
  const bgColor = document.getElementById("bgColor").value;
  const textColor = document.getElementById("textColor").value;
  const fontFamily = document.getElementById("fontFamily").value;
  const fontSize = document.getElementById("fontSize").value;

  // Thay đổi màu nền trang
  document.body.style.backgroundColor = bgColor;

  // Thay đổi màu chữ trang
  document.body.style.color = textColor;

  // Thay đổi kiểu chữ của trang
  document.body.style.fontFamily = fontFamily;

  // Thay đổi cỡ chữ của trang
  document.body.style.fontSize = fontSize + "px";
}
