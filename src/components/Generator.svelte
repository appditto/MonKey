<script>
  export let generatorVisibility = false;
</script>

<style>
  .generator {
    width: calc(100vw - 1rem);
    height: calc(100vw - 1rem);
    border-radius: 1rem;
    border-width: 0.25rem;
    box-shadow: 0rem 0.6rem 0rem 0rem #404040;
  }
  @media screen and (min-width: 768px) {
    .generator {
      width: 50vw;
      height: 50vw;
    }
  }
  @media screen and (min-width: 768px) {
    .generator {
      min-width: 24rem;
      min-height: 24rem;
      width: 20vw;
      height: 20vw;
    }
  }
</style>

{#if generatorVisibility}
  <div
    class="generator border-primary flex flex-col bg-white absolute top-0 mt-5">
    <div class="w-full h-full flex flex-col px-4 md:px-8 py-4 md:py-6">
      <div class="flex flex-col items-center my-auto">
        <input
          class="w-full text-xl font-bold px-4 py-2 border-2 border-black
          rounded-xl"
          type="text"
          placeholder="enter your address" />
        <button
          class="w-full bg-primary btn-primary text-white text-xl font-bold
          rounded-xl border-2 border-black px-6 py-2 mx-auto mt-3">
          Show Me
        </button>
      </div>
      <button
        class="bg-primary btn-primary text-white text-lg font-bold rounded-lg
        border-2 border-black px-6 md:px-8 py-1 mx-auto">
        Randomize
      </button>
    </div>

  </div>
{/if}
