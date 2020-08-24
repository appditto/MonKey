<script>
  import { modalIn, modalOut } from "../plugins/transitions.js";
  import ClickOutside from "./ClickOutside.svelte";
  export let navItems = [];
  export let segment;
  let isModalOpen = false;
  function toggleModal() {
    isModalOpen = !isModalOpen;
  }
  function closeModal() {
    isModalOpen = false;
  }
  function isTabOrEscPressed(event) {
    if (event.keyCode == "9" || event.keyCode == "27") {
      closeModal();
    }
  }
  function isEscPressed(event) {
    if (event.keyCode == "27") {
      closeModal();
    }
  }
</script>

<!-- If Mobile -->
<ClickOutside on:clickoutside={isModalOpen ? closeModal : null}>
  <div class="flex md:hidden flex-col justify-end items-center font-medium">
    <button on:keydown={isEscPressed} on:click={toggleModal} class="menu-button">
      <img
        class="{isModalOpen ? '-rotate-90' : ''} w-11 h-11 transition-all transform ease-out duration-300"
        src="images/icons/icon-menu.svg"
        alt="Menu Icon"
      />
    </button>
    <!-- Mobile Menu -->
    {#if isModalOpen}
      <div in:modalIn out:modalOut class="w-full relative z-40 origin-top-right">
        <div
          class="modal absolute w-56 flex flex-col justify-center bg-white rounded-lg shadow-2xl
          border-offWhite border-2 top-0 right-0 mt-2 p-2"
        >
          {#each navItems as item, i}
            {#if item.text && item.href}
              <a
                on:keydown={navItems.length - 1 == i ? isTabOrEscPressed : isEscPressed}
                on:click={closeModal}
                href={item.href}
                rel="noopener"
                class="{item.href == '/' + segment || (item.href == '/' && !segment) ? 'text-brownLight' : ''}
                menu-item w-full px-6 py-2 my-1 text-xl text-center transition-all duration-300 ease-out
                rounded-md"
              >
                <span
                  class="inline-block line {item.href == '/' + segment || (item.href == '/' && !segment) ? 'line-brownLight-active' : ''}"
                >
                  {item.text}
                </span>
              </a>
            {/if}
          {/each}
        </div>
      </div>
    {/if}
  </div>
</ClickOutside>
<!-- If >= Mobile -->
<div class="hidden md:flex flex-row flex-wrap justify-end items-center font-medium">
  {#each navItems as item}
    {#if item.text && item.href}
      <a
        href={item.href}
        rel="noopener"
        class="menu-button hover:text-brownLight px-2 pt-1 mx-1 md:ml-4 line {item.href == '/' + segment || (item.href == '/' && !segment) ? 'line-active' : ''}"
      >
        {item.text}
      </a>
    {/if}
  {/each}
</div>

<style>
  .line::after {
    content: "";
    height: 3px;
    width: calc(100% + 8px);
    margin-left: -4px;
    margin-top: 0.15rem;
    display: block;
    border-radius: 1.5px;
    transition: transform 0.2s ease-out;
    transform-origin: center;
    transform: scaleX(0);
  }
  .line-active::after {
    background-color: black;
    transform: scaleX(1);
  }
  .line-active:hover::after {
    background-color: #cd9e6c;
  }
  .menu-item:hover,
  .menu-item:focus {
    background-color: rgba(0, 0, 0, 0.1);
  }
  .menu-button {
    transition: all 0.2s ease-out;
    border-radius: 0.2rem;
  }
  .menu-button:focus {
    background-color: rgba(0, 0, 0, 0.15);
  }
  .line-brownLight-active::after {
    background-color: #cd9e6c;
    transform: scaleX(1);
  }
</style>
