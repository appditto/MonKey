<script>
  import MonkeyGenerator from "../MonkeyGenerator.svelte";
  let showGenerator = false;
  function toggleGenerator() {
    showGenerator = !showGenerator;
  }
  let windowWidth;
</script>

<svelte:window bind:innerWidth={windowWidth} />
<!-- HTML -->
<div class="w-full flex flex-col items-center pt-20 md:pt-16 bg-hero">
  <div class="w-full flex flex-col items-center px-4">
    <h1
      class="font-bold text-5xl md:text-6xl mx-5 md:mx-12 text-center leading-tight tracking-tight
      mt-5 md:mt-16 lg:mt-20"
    >
      Welcome to the jungle.
    </h1>
    <p class="text-center mx-5 mt-4 text-2xl">In here, Banano addresses are MonKeys.</p>
    <button
      on:click={toggleGenerator}
      class="{showGenerator ? 'bg-danger btn-danger border-dangerDark' : 'bg-primary btn-primary border-black'}
      w-64 max-w-full text-white text-xl font-bold rounded-xl border-2 px-3 md:px-6 py-2 mx-auto
      mt-6"
    >
      {showGenerator ? 'Close' : 'Show My MonKey'}
    </button>
    <div class="w-full flex flex-row justify-center relative z-50">
      <MonkeyGenerator {showGenerator} />
    </div>
  </div>
  <div class="w-full relative">
    {#if windowWidth >= 1024}
      <!-- Desktop Illustration -->
      <img
        class="absolute top-0 left-0 w-full h-auto hidden lg:block"
        src="images/illustrations-foreground/hero.svg"
        alt="Hero Illustration"
      />
    {/if}
    {#if windowWidth >= 768 && windowWidth < 1024}
      <!-- Tablet Illustration -->
      <img
        class="absolute top-0 left-0 w-full h-auto hidden md:block lg:hidden"
        src="images/illustrations-foreground/hero-tablet.svg"
        alt="Hero Illustration"
      />
    {/if}
    {#if windowWidth < 768}
      <!-- Mobile Illustration-->
      <img
        class="absolute top-0 left-0 w-full h-auto md:hidden mt-4"
        src="images/illustrations-foreground/hero-mobile.svg"
        alt="Hero Illustration"
      />
    {/if}
  </div>
  <!-- Desktop BG -->
  <img
    class="w-full h-auto hidden lg:block"
    src="images/illustrations-foreground/hero-bg.svg"
    alt="Hero Illustration Background"
  />
  <!-- Tablet Illustration -->
  <img
    class="w-full h-auto hidden md:block lg:hidden"
    src="images/illustrations-foreground/hero-bg-tablet.svg"
    alt="Hero Illustration Background"
  />
  <!-- Mobile Illustration-->
  <img
    class="w-full h-auto md:hidden mt-4"
    src="images/illustrations-foreground/hero-bg-mobile.svg"
    alt="Hero Illustration Background"
  />
</div>

<style>
  .bg-hero {
    background-image: url("/images/illustrations-background/bg-hero.svg");
    background-size: auto 100%;
    background-position: 100% 50%;
    background-repeat: no-repeat;
  }
</style>
