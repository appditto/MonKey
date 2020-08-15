<script>
  import LazyImage from "../LazyImage.svelte";
</script>

<!-- HTML -->
<div class="w-full flex-col justify-center py-12 md:my-20">
  <div class="y-container flex flex-col items-center">
    <h2
      class="font-bold text-4xl md:text-5xl mx-4 text-center leading-tight
      tracking-tight mt-5"
    >
      Wanna Integrate MonKey?
    </h2>
    <p class="w-full max-w-2xl text-center md:mx-12 mt-3 text-2xl px-4">
      Let’s get you started, it will be a breeze.
    </p>
    <div class="flex flex-row justify-center mt-6">
      <a
        href="/documentation"
        class="bg-primary btn-primary text-white text-xl font-bold rounded-xl
        border-2 border-black px-8 md:px-12 py-2"
      >
        Start Integrating
      </a>
    </div>
    <!-- Desktop Illustration -->
    <LazyImage classes="w-full hidden md:block px-12 lg:px-20 mt-12">
      <img
        slot="content"
        class="w-full h-auto"
        src="images/illustrations-foreground/integrate.svg"
        alt="Integrate Illustration"
      />
      <img
        slot="placeholder"
        class="w-full h-auto"
        src="images/illustrations-foreground/integrate-placeholder.svg"
        alt="Integrate Illustration Placeholder"
      />
    </LazyImage>
    <!-- Mobile Illustration-->
    <LazyImage classes="w-full max-w-sm md:hidden mx-auto mt-12 px-6">
      <img
        slot="content"
        class="w-full h-auto"
        src="images/illustrations-foreground/integrate-mobile.svg"
        alt="Integrate Illustration"
      />
      <img
        slot="placeholder"
        class="w-full h-auto"
        src="images/illustrations-foreground/integrate-mobile-placeholder.svg"
        alt="Integrate Illustration Placeholder"
      />
    </LazyImage>
  </div>
</div>
