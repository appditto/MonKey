<script>
  export let segment;
</script>

<style>
  .line::after {
    content: "";
    height: 4px;
    width: calc(100% + 8px);
    margin-left: -4px;
    display: block;
    border-radius: 2px;
    transition: transform 0.2s ease-out;
    transform-origin: center;
    transform: scaleX(0);
  }
  .line-active::after {
    background-color: black;
    transform: scaleX(1);
  }
</style>

<!-- HTML -->
<div class="w-full px-2 py-2 md:px-4 md:py-3 absolute z-20">
  <div class="y-container flex flex-row justify-between items-center">
    <!-- Logo -->
    <a class="w-32 md:w-36 p-1 logo-button m-2" href="/">
      <img
        class="w-full h-auto"
        alt="Logo"
        src="images/logos/logo-monkey-horizontal.svg" />
    </a>
    <!-- Menu Options -->
    <div
      class="flex flex-row flex-wrap justify-end items-center w-2/3 font-medium
      pr-1">
      <a
        href="/documentation"
        class="menu-button hover:text-brownLight px-2 pt-1 mx-1 md:ml-4 line {segment === 'documentation' ? 'line-active' : ''}">
        API
      </a>
      <a
        href="https://github.com/appditto/monkey"
        class="menu-button hover:text-brownLight px-2 pt-1 mx-1 md:ml-4 line">
        GitHub
      </a>
      <a
        href="https://banano.cc"
        class="menu-button hover:text-brownLight px-2 pt-1 mx-1 md:ml-4 line">
        Banano
      </a>
    </div>
  </div>
</div>
