// Tạo và cấu hình CKEditor
ClassicEditor.create(document.querySelector("#gioi-thieu"), {
  toolbar: [
    "undo",
    "redo",
    "|",
    "bold",
    "italic",
    "underline",
    "strike",
    "|",
    "subscript",
    "superscript",
    "|",
    "fontSize",
    "fontFamily",
    "textColor",
    "bgColor",
    "|",
  ],
  placeholder: "Viết nội dung ở đây...",
  height: 300,
})
  .then((editor) => {
    console.log("CKEditor đã được khởi tạo!");
  })
  .catch((error) => {
    console.error("Có lỗi khi khởi tạo CKEditor:", error);
  });
