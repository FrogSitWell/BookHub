// common.js
async function fetchUserInfo() {
  const token = localStorage.getItem("token");

  if (token) {
    // Giải mã token để lấy thông tin user
    const decodedToken = jwt_decode(token);
    const userID = decodedToken.sub; // Giả sử userID lưu trong trường `sub` của token

    console.log("User ID:", userID);

    // Tiến hành gửi dữ liệu với userID
    fetch("http://localhost:8081/protected/profile", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + token,
      },
      body: formData,
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Upload successful:", data);
      })
      .catch((error) => {
        console.error("Error uploading:", error);
      });
  } else {
    console.log("User not logged in");
  }
}
