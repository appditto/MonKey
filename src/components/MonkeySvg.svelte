<script>
  export let monkeySvg;
  let svgElement;
  $: if (svgElement && monkeySvg) {
    svgElement.innerHTML = monkeySvg;
  }
</script>

{#if monkeySvg}
  <div bind:this={svgElement} class="w-full h-full" />
{/if}
