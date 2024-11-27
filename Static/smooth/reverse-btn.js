// Lấy các thành phần cần sử dụng
const reverseButton = document.getElementById("reverse-btn");
const chapterList = document.getElementById("chapter-list");

// Thêm sự kiện khi nhấn nút "Đảo ngược danh sách"
reverseButton.addEventListener("click", function () {
  const chapters = Array.from(chapterList.children); // Lấy danh sách các dòng (tr)
  chapters.reverse(); // Đảo ngược thứ tự
  chapterList.innerHTML = ""; // Xóa nội dung cũ
  chapters.forEach((chapter) => chapterList.appendChild(chapter)); // Thêm các dòng đã đảo ngược
});
