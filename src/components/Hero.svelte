<style>
  .bg-hero {
    background-image: url("/images/illustrations-background/bg-hero.svg");
    background-size: auto 100%;
    background-position: 100% 50%;
    background-repeat: no-repeat;
  }
</style>

<!-- HTML -->
<div class="w-full flex flex-col justify-center pt-20 bg-hero">
  <h1
    class="font-bold text-5xl mx-5 text-center leading-none tracking-tight mt-5">
    Welcome to the jungle.
  </h1>
  <p class="text-center mx-5 mt-5 text-2xl">
    In here, Banano addresses are MonKeys.
  </p>
  <button
    class="bg-primary btn-primary text-white text-xl font-bold rounded-xl
    border-2 border-black px-6 py-2 mx-auto mt-6">
    Show My MonKey
  </button>
  <!-- Desktop Illustration -->
  <img
    class="w-full h-auto hidden md:block"
    src="images/illustrations-foreground/hero.svg"
    alt="Hero Illustration" />
  <!-- Mobile Illustration-->
  <img
    class="w-full h-auto md:hidden"
    src="images/illustrations-foreground/hero-mobile.svg"
    alt="Hero Illustration" />
</div>
