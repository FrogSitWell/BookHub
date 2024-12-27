// Fetch book information using JavaScript
// Lấy `bookID` từ URL
const urlParams = new URLSearchParams(window.location.search);
const bookID = urlParams.get("id");

// Lấy thông tin chi tiết sách
fetch(`http://localhost:8081/api/book/${bookID}`)
  .then((response) => {
    if (!response.ok) {
      throw new Error("Book not found");
    }
    return response.json();
  })
  .then((book) => {
    // Hiển thị thông tin sách
    document.getElementById("book-avatar").src = book.AvatarURL;
    document.getElementById("book-avatar").alt = book.Name;
    document.getElementById("book-name").textContent = book.Name;
    document.getElementById(
      "book-author"
    ).textContent = `Tác giả: ${book.Author}`;
    document.getElementById(
      "book-genre"
    ).textContent = `Thể loại: ${book.Genre.Name}`;
    document.getElementById(
      "book-total-chapters"
    ).textContent = `Tổng chương: ${book.TotalChapters}`;
    document.getElementById(
      "book-status"
    ).textContent = `Trạng thái: ${book.Status.StatusName}`;
    document.getElementById(
      "book-rating"
    ).textContent = `Đánh giá: ${book.AverageRate}/5`;

    // Gọi hàm loadRelatedBooks để hiển thị danh sách truyện cùng tác giả
    loadRelatedBooks(bookID);
  })
  .catch((error) => {
    console.error("Error fetching book data:", error);
  });

// Hàm lấy danh sách truyện cùng tác giả
function loadRelatedBooks(bookId) {
  // Gửi yêu cầu GET tới API
  fetch(`http://localhost:8081/api/author/${bookId}`)
    .then((response) => {
      if (!response.ok) {
        throw new Error("No related books found");
      }
      return response.json();
    })
    .then((data) => {
      if (data.books && data.books.length > 0) {
        const relatedBooksDiv = document.getElementById("related-books");
        relatedBooksDiv.innerHTML = ""; // Xóa nội dung cũ

        // Biến đếm để đánh số thứ tự
        let index = 1;

        // Thêm danh sách truyện vào giao diện
        data.books.forEach((book) => {
          const bookLink = document.createElement("p");
          bookLink.innerHTML = `<strong>${index++}.</strong> <a href="/books/${
            book.id
          }" target="_blank">${book.Name}</a>`;
          relatedBooksDiv.appendChild(bookLink);
        });
      } else {
        console.log("No related books available.");
      }
    })
    .catch((error) => console.error("Error loading related books:", error));
}

function fetchBookDescription(bookId) {
  // Gửi yêu cầu GET đến API để lấy mô tả sách
  fetch(`http://localhost:8081/api/description/${bookId}`)
    .then((response) => {
      if (!response.ok) {
        // Nếu API trả về lỗi, ném lỗi ra ngoài
        throw new Error("Không thể tải mô tả sách. Lỗi từ server.");
      }
      return response.json(); // Chuyển đổi phản hồi thành JSON
    })
    .then((data) => {
      // Kiểm tra toàn bộ dữ liệu nhận được từ API
      console.log("Dữ liệu nhận được từ API:", data);

      // Kiểm tra xem trường 'description' có tồn tại không
      if (data && data.description) {
        const postContent = document.getElementById("post-content");
        postContent.textContent = data.description; // Hiển thị mô tả sách
      } else {
        throw new Error("Mô tả sách không tồn tại.");
      }
    })
    .catch((error) => {
      // Hiển thị lỗi nếu có
      console.error("Lỗi khi gọi API mô tả sách:", error.message);
      const postContent = document.getElementById("post-content");
      postContent.textContent = `Lỗi: ${error.message}`;
    });
}
fetchBookDescription(bookID);
// Hàm lấy 3 chương mới nhất
// Hàm lấy 3 chương mới nhất
function loadTopChapters(bookID) {
  if (!bookID) {
    console.error("Invalid bookID provided");
    return;
  }

  // Gửi yêu cầu GET tới API để lấy 3 chương mới nhất
  fetch(`http://localhost:8081/api/newchapter/${bookID}`)
    .then((response) => {
      if (!response.ok) {
        throw new Error("No top chapters found");
      }
      return response.json();
    })
    .then((data) => {
      const chaptersDiv = document.getElementById("top-chapters-row");
      chaptersDiv.innerHTML = ""; // Xóa nội dung cũ

      if (data.chapters && data.chapters.length > 0) {
        data.chapters.sort((a, b) => a.SortOrder - b.SortOrder);
        // Duyệt qua các chương và thêm vào bảng
        data.chapters.forEach((chapter) => {
          const row = document.createElement("tr");
          const cell = document.createElement("td");
          cell.innerHTML = `<a href="#">${chapter.Title}</a>`;
          row.appendChild(cell);
          chaptersDiv.appendChild(row);
        });
      } else {
        // Nếu không có chương nào, hiển thị thông báo
        chaptersDiv.innerHTML = "<tr><td>No chapters available.</td></tr>";
      }
    })
    .catch((error) => {
      console.error("Error loading top chapters:", error);
      const chaptersDiv = document.getElementById("top-chapters-row");
      chaptersDiv.innerHTML =
        "<tr><td>Error loading chapters. Please try again later.</td></tr>";
    });
}

// Gọi hàm loadTopChapters với bookID
loadTopChapters(bookID);
function loadAllChapters(bookID) {
  if (!bookID) {
    console.error("Invalid bookID provided");
    return;
  }

  fetch(`http://localhost:8081/api/allchapter/${bookID}`)
    .then((response) => {
      if (!response.ok) {
        throw new Error("No chapters found");
      }
      return response.json();
    })
    .then((data) => {
      const chaptersDiv = document.getElementById("chapter-list");
      chaptersDiv.innerHTML = ""; // Xóa nội dung cũ

      if (data.chapters && data.chapters.length > 0) {
        data.chapters.forEach((chapter) => {
          const row = document.createElement("tr");
          const cell = document.createElement("td");
          cell.innerHTML = `<a href="#">${chapter.Title}</a>`;
          row.appendChild(cell);
          chaptersDiv.appendChild(row);
        });
      } else {
        chaptersDiv.innerHTML = "<tr><td>No chapters available.</td></tr>";
      }
    })
    .catch((error) => {
      console.error("Error loading chapters:", error);
      const chaptersDiv = document.getElementById("chapter-list");
      chaptersDiv.innerHTML =
        "<tr><td>Error loading chapters. Please try again later.</td></tr>";
    });
}
loadAllChapters(bookID);
