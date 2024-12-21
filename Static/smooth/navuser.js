document.addEventListener("DOMContentLoaded", function () {
  // Hàm để lấy danh sách thể loại từ backend
  function loadGenres() {
    fetch("http://localhost:8081/api/genres") // Gọi API backend để lấy danh sách thể loại
      .then((response) => response.json())
      .then((genres) => {
        displayGenres(genres.slice(0, 5)); // Hiển thị 5 thể loại đầu tiên
        addSeeMoreLink(); // Thêm liên kết "Xem thêm"
      })
      .catch((error) => {
        console.error("Lỗi khi lấy thể loại:", error);
      });
  }

  // Hàm hiển thị thể loại lên dropdown
  function displayGenres(genres) {
    const genreList = document.getElementById("genre-list");
    genreList.innerHTML = ""; // Xóa danh sách hiện tại

    // Hiển thị 5 thể loại đầu tiên
    genres.forEach((genre) => {
      const li = document.createElement("li");
      const a = document.createElement("a");
      a.href = "/Views/homepage/file.html"; // có thể yêu cầu quyery
      a.textContent = genre.Name;
      li.appendChild(a);
      genreList.appendChild(li);
    });
  }

  // Hàm thêm nút "Xem thêm" với link dẫn đến trang khác
  function addSeeMoreLink() {
    const genreList = document.getElementById("genre-list");
    const seeMoreLi = document.createElement("li");
    const seeMoreA = document.createElement("a");
    seeMoreA.href = "/Views/homepage/file.html"; // Liên kết dẫn đến trang khác
    seeMoreA.textContent = "Xem thêm>>";
    seeMoreLi.appendChild(seeMoreA);
    genreList.appendChild(seeMoreLi);
  }

  // Bắt đầu tải danh sách thể loại
  loadGenres();

  const searchInput = document.getElementById("searchInput");
  const resultsContainer = document.getElementById("results");

  // Hàm tìm kiếm
  function searchBooks() {
    const query = searchInput.value;

    if (query.length >= 2) {
      // Bắt đầu tìm kiếm khi người dùng nhập ít nhất 2 ký tự
      fetch(`http://localhost:8081/api/search?query=${query}`)
        .then((response) => response.json())
        .then((data) => {
          displayResults(data.books); // Hiển thị kết quả tìm kiếm
        })
        .catch((error) => console.error("Error:", error));
    } else {
      clearResults(); // Nếu không có ký tự tìm kiếm, xóa kết quả hiện tại
    }
  }

  // Hàm hiển thị kết quả tìm kiếm dưới dạng danh sách gợi ý
  function displayResults(books) {
    resultsContainer.innerHTML = ""; // Xóa kết quả cũ
    resultsContainer.style.display = "block"; // Hiển thị danh sách gợi ý

    if (books.length > 0) {
      books.forEach((book) => {
        const resultItem = document.createElement("div");
        resultItem.classList.add("result-item");
        resultItem.innerHTML = `
            <strong>${book.Name}</strong><br>
          `;
        resultItem.addEventListener("click", function () {
          searchInput.value = book.Name; // Gán tên sách vào ô tìm kiếm khi người dùng chọn kết quả
          clearResults(); // Ẩn kết quả khi người dùng chọn một kết quả
        });
        resultsContainer.appendChild(resultItem);
      });
    } else {
      resultsContainer.innerHTML =
        '<div class="result-item">Không có kết quả phù hợp</div>';
    }
  }

  // Hàm xóa kết quả
  function clearResults() {
    resultsContainer.innerHTML = "";
    resultsContainer.style.display = "none"; // Ẩn danh sách gợi ý
  }

  // Gắn sự kiện cho ô tìm kiếm
  searchInput.addEventListener("input", searchBooks);

  // Ẩn kết quả khi người dùng click ra ngoài
  document.addEventListener("click", function (event) {
    if (
      !resultsContainer.contains(event.target) &&
      event.target !== searchInput
    ) {
      clearResults();
    }
  });
  let banners = [];
  let currentBannerIndex = 0;

  // Fetch banners from the backend
  fetch("http://localhost:8081/api/banners")
    .then((response) => response.json())
    .then((data) => {
      banners = data;
      displayBanner();
      setInterval(nextBanner, 5000); // Chuyển banner mỗi 5 giây
    });

  function displayBanner() {
    if (banners.length > 0) {
      const banner = banners[currentBannerIndex];
      document.getElementById("banner1").src = banner.ImageURL;
    }
  }

  function nextBanner() {
    currentBannerIndex = (currentBannerIndex + 1) % banners.length;
    displayBanner();
  }
});
