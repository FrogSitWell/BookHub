document.addEventListener("DOMContentLoaded", function () {
  const dropdownButton = document.querySelector(".dropdown-button");
  const dropdownMenu = document.querySelector(".dropdown-menu");

  // Hiện/ẩn menu khi click vào button
  dropdownButton.addEventListener("click", function () {
    dropdownMenu.style.display =
      dropdownMenu.style.display === "block" ? "none" : "block";
  });

  // Đóng menu khi click bên ngoài dropdown
  document.addEventListener("click", function (event) {
    if (!event.target.closest(".dropdown")) {
      dropdownMenu.style.display = "none";
    }
  });
});
