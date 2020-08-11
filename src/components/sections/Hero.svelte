<script>
  import MonkeyGenerator from "../MonkeyGenerator.svelte";
  import { onMount } from "svelte";
  let generatorVisibility = false;
  const toggleGenerator = () => {
    generatorVisibility = !generatorVisibility;
  };
  let showGenerator = false;
  onMount(async () => {
    showGenerator = true;
  });
</script>

<style>
  .bg-hero {
    background-image: url("/images/illustrations-background/bg-hero.svg");
    background-size: auto 100%;
    background-position: 100% 50%;
    background-repeat: no-repeat;
  }
</style>

<!-- HTML -->
<div class="w-full flex flex-col items-center pt-20 bg-hero">
  <div class="w-full flex flex-col items-center px-4">
    <h1
      class="font-bold text-4xl md:text-5xl mx-5 text-center leading-tight
      tracking-tight mt-5 md:mt-16">
      Welcome to the jungle.
    </h1>
    <p class="text-center mx-5 mt-4 text-2xl">
      In here, Banano addresses are MonKeys.
    </p>
    <button
      on:click={() => {
        toggleGenerator();
      }}
      class="{generatorVisibility ? 'bg-danger btn-danger border-dangerDark' : 'bg-primary btn-primary border-black'}
      w-64 max-w-full text-white text-xl font-bold rounded-xl border-2 px-3
      md:px-6 py-2 mx-auto mt-6">
      {generatorVisibility ? 'Close' : 'Show My MonKey'}
    </button>
    <div class="w-full flex flex-row justify-center relative z-50">
      {#if showGenerator}
        <MonkeyGenerator {generatorVisibility} />
      {/if}
    </div>
  </div>
  <!-- Desktop Illustration -->
  <img
    class="w-full h-auto hidden md:block"
    src="images/illustrations-foreground/hero.svg"
    alt="Hero Illustration" />
  <!-- Mobile Illustration-->
  <img
    class="w-full h-auto md:hidden mt-4"
    src="images/illustrations-foreground/hero-mobile.svg"
    alt="Hero Illustration" />
</div>
