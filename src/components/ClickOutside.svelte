<script>
  import { createEventDispatcher } from "svelte";
  let subject;
  const dispatch = createEventDispatcher();
  function isChildClicked(target) {
    if (target === subject || subject.contains(target)) {
      return true;
    } else {
      return false;
    }
  }
  function onClickOutside(event) {
    if (!isChildClicked(event.target)) {
      dispatch("clickoutside");
    }
  }
</script>

<svelte:body on:click={onClickOutside} />
<div bind:this={subject}>
  <slot />
</div>
