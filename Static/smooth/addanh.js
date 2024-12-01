document
  .getElementById("upload-image")
  .addEventListener("change", function (event) {
    const file = event.target.files[0];
    const preview = document.getElementById("image-preview");

    if (file) {
      const reader = new FileReader();
      reader.onload = function (e) {
        preview.innerHTML = `<img src="${e.target.result}" alt="Ảnh bìa" />`;
      };
      reader.readAsDataURL(file);
    } else {
      preview.innerHTML = "Chưa có ảnh nào được chọn";
    }
  });
