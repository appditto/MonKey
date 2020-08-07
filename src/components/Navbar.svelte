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
<div class="w-full px-2 py-2 md:px-4 md:py-3">
  <div class="y-container flex flex-row justify-between items-center">
    <a href="/">
      <img
        class="w-32 md:w-40 h-auto m-2"
        alt="Logo"
        src="images/logos/logo-monkey-horizontal.svg" />
    </a>
    <div class="flex flex-row justify-end items-center w-1/2 font-bold">
      <a
        href="/documentation"
        class="m-2 md:ml-4 line {segment === 'documentation' ? 'line-active' : ''}">
        API
      </a>
      <a
        href="https://github.com/appditto/monkey"
        target="_blank"
        class="m-2 md:ml-4 line">
        GitHub
      </a>
      <a href="https://banano.cc" target="_blank" class="m-2 md:ml-4 line">
        Banano
      </a>
    </div>
  </div>
</div>
