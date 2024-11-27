// script.js

// Xử lý khi nhấn nút "Trước" hoặc "Tiếp"
const chapters = [
  "Chương 01: Chỉ Có Người Chết",
  "Chương 02: Hiện Thực! Võ Đạo Thất Trọng!",
  "Chương 03: Hệ Thống Không Dễ Xơi!",
  "Chương 04: Hạ Gia! Hạ Uyển!",
  "Chương 05: Nhuyễn Khí Tán",
  "Chương 06: Không Chế Không Tốt!",
];

let currentIndex = 0;
const chapterDisplay = document.querySelector(".current-chapter");
const prevButton = document.querySelector(".prev");
const nextButton = document.querySelector(".next");
const extraPrevButton = document.getElementById("prev-btn");
const extraNextButton = document.getElementById("next-btn");

// Cập nhật chương
function updateChapter() {
  chapterDisplay.textContent = chapters[currentIndex];
  prevButton.disabled = currentIndex === 0;
  nextButton.disabled = currentIndex === chapters.length - 1;
  extraPrevButton.disabled = currentIndex === 0;
  extraNextButton.disabled = currentIndex === chapters.length - 1;
}

// Sự kiện cho nút "Trước"
prevButton.addEventListener("click", () => {
  if (currentIndex > 0) {
    currentIndex--;
    updateChapter();
  }
});

// Sự kiện cho nút "Tiếp"
nextButton.addEventListener("click", () => {
  if (currentIndex < chapters.length - 1) {
    currentIndex++;
    updateChapter();
  }
});

// Sự kiện cho nút bổ sung
extraPrevButton.addEventListener("click", () => {
  if (currentIndex > 0) {
    currentIndex--;
    updateChapter();
  }
});

extraNextButton.addEventListener("click", () => {
  if (currentIndex < chapters.length - 1) {
    currentIndex++;
    updateChapter();
  }
});

// Khởi tạo
updateChapter();
