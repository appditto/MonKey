<script>
  export let offset = 200;
  export let classes;
  import { onMount } from "svelte";
  let observer = null;
  let showContent = false;
  let item;
  function onIntersection(entries) {
    if (entries[0].isIntersecting && !showContent) {
      showContent = true;
    }
  }
  function lazyLoad(node) {
    observer && observer.observe(node);
    return {
      destroy() {
        observer && observer.unobserve(node);
      },
    };
  }
  onMount(async () => {
    let hasAPI = "IntersectionObserver" in window;
    if (hasAPI) {
      observer = new IntersectionObserver(onIntersection, { rootMargin: `${offset}px` });
      lazyLoad(item);
    }
  });
</script>

<div bind:this={item} class={classes}>
  <div class="w-full h-full relative">
    <div
      class="{showContent ? 'opacity-0' : 'opacity-100'} w-full h-full transition-opacity
      duration-500"
    >
      <slot name="placeholder" />
    </div>
    <div
      class="{showContent ? 'opacity-100' : 'opacity-0'} w-full h-full transition-opacity
      duration-300 ease-out absolute top-0 left-0"
    >
      {#if showContent}
        <slot name="content" />
        <slot />
      {/if}
    </div>
  </div>
</div>
