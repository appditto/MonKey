<script>
  import LazyImage from "../LazyImage.svelte";
</script>

<style>
  .bg-banano {
    background-image: url("/images/illustrations-background/bg-banano.svg");
    background-size: 100% auto;
    background-position: 100% 0%;
    background-repeat: no-repeat;
  }
</style>

<!-- HTML -->
<div class="w-full py-12 md:my-20">
  <div class="w-full flex flex-col justify-center relative bg-banano">
    <div class="y-container relative">
      <h2
        class="font-bold text-5xl mx-4 text-center leading-none tracking-tight
        mt-5">
        Available on Lots of
        <br class="hidden md:block" />
        Banano Services
      </h2>
      <p class="text-center mx-4 md:mx-12 mt-3 text-2xl">
        MonKeys are available in lots of places in the Banano ecosystem.
        <br class="hidden lg:block" />
        Most notable two are
        <a href="https://kalium.banano.cc" target="_blank" class="font-bold">
          Kalium
        </a>
        and
        <a href="https://creeper.banano.cc" target="_blank" class="font-bold">
          Creeper
        </a>
        .
      </p>
      <div class="w-full flex flex-row flex-wrap justify-center px-4 mt-2">
        <a
          href="https://kalium.banano.cc"
          target="_blank"
          class="w-full md:w-64 bg-primary btn-primary text-white text-xl
          font-bold rounded-xl border-2 border-black px-6 py-2 mx-2 mt-6
          text-center">
          Download Kalium
        </a>
        <a
          href="https://creeper.banano.cc"
          target="_blank"
          class="w-full md:w-64 bg-primary btn-primary text-white text-xl
          font-bold rounded-xl border-2 border-black px-6 py-2 mx-2 mt-6
          text-center">
          Visit Creeper
        </a>
      </div>
      <div
        class="w-full flex flex-row flex-wrap justify-center items-center mt-6
        px-4">
        <LazyImage classes="w-full max-w-xs px-2 md:px-4 my-6">
          <img
            slot="content"
            class="w-full h-auto"
            src="images/illustrations-foreground/kalium.svg"
            alt="Kalium Illustration" />
          <img
            slot="placeholder"
            class="w-full h-auto"
            src="images/illustrations-foreground/kalium-placeholder.svg"
            alt="Kalium Illustration Placeholder" />
        </LazyImage>
        <div class="flex flex-col items-center justify-center md:mt-4">
          <LazyImage classes="w-full max-w-md mx-auto px-2 my-2">
            <img
              slot="content"
              class="w-full h-auto"
              src="images/illustrations-foreground/creeper-1.svg"
              alt="Creeper 1 Illustration" />
            <img
              slot="placeholder"
              class="w-full h-auto "
              src="images/illustrations-foreground/creeper-1-placeholder.svg"
              alt="Creeper 1 Illustration Placeholder" />
          </LazyImage>
          <LazyImage classes="w-full max-w-md mx-auto px-6 md:px-6 my-2">
            <img
              slot="content"
              class="w-full h-auto"
              src="images/illustrations-foreground/creeper-2.svg"
              alt="Creeper 2 Illustration" />
            <img
              slot="placeholder"
              class="w-full h-auto "
              src="images/illustrations-foreground/creeper-2-placeholder.svg"
              alt="Creeper 2 Illustration Placeholder" />
          </LazyImage>
        </div>
      </div>
    </div>
  </div>
</div>
