<script>
  import NavbarItems from "./NavbarItems.svelte";
  export let segment;
  let navItems = [
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
<div class="w-full px-2 py-2 md:px-4 md:py-3 absolute z-20">
  <div class="y-container flex flex-row justify-between items-center">
    <!-- Logo -->
    <a class="w-32 md:w-36 p-1 logo-button m-2" href="/">
      <img
        class="w-full h-auto"
        alt="Logo"
        src="images/logos/logo-monkey-horizontal.svg"
      />
    </a>
    <!-- Menu Options -->
    <NavbarItems {navItems} {segment} />
  </div>
</div>
