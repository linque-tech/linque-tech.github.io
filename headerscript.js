window.addEventListener("scroll", function() {
    const navbar = document.getElementById("header");
    const homeSection = document.getElementById("home");

    // Calculate the height of the home section to know when to apply the background color
    const homeHeight = homeSection.offsetHeight;

    // Add or remove the 'scrolled' class based on the scroll position
    if (window.scrollY > homeHeight - 50) {
        navbar.classList.add("scrolled");
    } else {
        navbar.classList.remove("scrolled");
    }
});
