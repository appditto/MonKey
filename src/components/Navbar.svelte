<script>
  import NavbarItems from "./NavbarItems.svelte";
  export let segment;
  let navItems = [
    {
      text: "Home",
      href: "/",
    },
    {
      text: "API",
      href: "/documentation",
    },
    {
      text: "Banano",
      href: "https://banano.cc",
    },
    {
      text: "GitHub",
      href: "https://github.com/appditto/monkey",
    },
  ];
</script>

<!-- HTML -->
<div class="w-full px-3 py-3 md:px-5 md:py-4 absolute z-20">
  <div class="y-container flex flex-row justify-between items-center">
    <!-- Logo -->
    <a class="w-36 h-auto p-1 logo-button mr-2" href="/">
      <img class="w-full h-auto" alt="Logo" src="images/logos/logo-monkey-horizontal.svg" />
    </a>
    <!-- Menu Options -->
    <NavbarItems {navItems} {segment} />
  </div>
</div>
