<script>
  export let navItems = [];
  export let segment;
</script>

<div
  class="flex flex-row flex-wrap justify-end items-center w-2/3 font-medium pr-1"
>
  {#each navItems as item}
    {#if item.text && item.href}
      <a
        href={item.href}
        rel="noopener"
        class="menu-button hover:text-brownLight px-2 pt-1 mx-1 md:ml-4 line {item.href == '/documentation' && segment === 'documentation' ? 'line-active' : ''}"
      >
        {item.text}
      </a>
    {/if}
  {/each}
</div>

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
  .line-active:hover::after {
    background-color: #cd9e6c;
  }
</style>
