function scrollToSection(sectionId) {
  // Lấy phần tử có id tương ứng
  const section = document.getElementById(sectionId);

  // Cuộn trang đến phần tử đó với hiệu ứng mượt mà
  section.scrollIntoView({
    behavior: "smooth", // Cuộn mượt mà
  });
}
