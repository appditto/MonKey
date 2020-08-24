<script>
  export let icon = {};
</script>

<!-- HTML -->
<a class="m-2 icon-button p-1" rel="noopener" target="_blank" href={icon.href}>
  <img class="w-12 h-12" src="images/icons/icon-{icon.name.toLowerCase()}.svg" alt="{icon.name} Icon" />
</a>

<style>
  .icon-button {
    transform: scale(1);
    transition: all 0.2s ease-out;
    border-radius: 0.35rem;
  }

  .icon-button:hover {
    transform: scale(1.15);
    transition: all 0.2s ease-out;
  }

  .icon-button:focus {
    background-color: rgba(0, 0, 0, 0.15);
  }
</style>
