const loginForm = document.getElementById("adminLoginForm");

loginForm.addEventListener("submit", async (event) => {
  event.preventDefault();

  const formData = new FormData(loginForm);

  try {
    const response = await fetch("/admin/api/login", {
      method: "POST",
      body: formData,
    });

    if (!response.ok) {
      throw new Error(response.status);
    }

    location.href = "/admin";
  } catch (error) {}
});
