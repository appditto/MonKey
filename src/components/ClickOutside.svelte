<script>
  import { createEventDispatcher } from "svelte";
  let child;
  const dispatch = createEventDispatcher();
  function isChildClicked(target) {
    var clicked = target;
    while (clicked) {
      if (clicked === child) {
        return true;
      }
      clicked = clicked.parentNode;
    }
    return false;
  }
  function onClickOutside(event) {
    if (!isChildClicked(event.target)) {
      dispatch("clickoutside");
    }
  }
</script>

<svelte:body on:click={onClickOutside} />
<div bind:this={child}>
  <slot />
</div>
