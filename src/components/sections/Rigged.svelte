<script>
  import LazyImage from "../LazyImage.svelte";
</script>

<!-- HTML -->
<div
  class="w-full ex-col justify-center py-12 md:py-20 relative md:my-20
  overflow-x-hidden"
>
  <!-- Background Illusration Right -->
  <img
    class="hidden md:block w-1/3 lg:w-1/4 h-auto absolute right-0 top-0"
    src="images/illustrations-background/bg-rigged-right.svg"
    alt="Rigged Right Illustration"
  />
  <!-- Background Illusration Left -->
  <img
    class="hidden md:block w-1/3 lg:w-1/4 h-auto absolute left-0 bottom-0"
    src="images/illustrations-background/bg-rigged-left.svg"
    alt="Rigged Left Illustration"
  />
  <div class="y-container relative">
    <h2
      class="font-bold text-4xl md:text-5xl mx-4 text-center leading-tight
      tracking-tight mt-5"
    >
      Masterfully Rigged
    </h2>
    <p class="text-center mx-4 md:mx-12 mt-3 text-2xl">
      We masterfully rig the system so that our friends can have special
      MonKeys.
      <br class="hidden lg:block" />
      But don’t worry, we can become friends.
    </p>
    <div
      class="flex flex-row flex-wrap justify-center -mx-16 lg:-mx-64 mt-2
      md:mt-4 lg:px-12"
    >
      <!-- Vader -->
      <LazyImage classes="w-1/2 lg:w-1/4 h-auto -mx-8 md:-mx-16">
        <img
          slot="content"
          class="w-full h-auto"
          src="images/illustrations-foreground/vader.svg"
          alt="Vader Illustration"
        />
        <img
          slot="placeholder"
          class="w-full h-auto"
          src="images/illustrations-foreground/vader-placeholder.svg"
          alt="Vader Illustration Placeholder"
        />
      </LazyImage>
      <!-- Iron Man -->
      <LazyImage classes="w-1/2 lg:w-1/4 h-auto -mx-8 md:-mx-16">
        <img
          slot="content"
          class="w-full h-auto"
          src="images/illustrations-foreground/ironman.svg"
          alt="Iron Man Illustration"
        />
        <img
          slot="placeholder"
          class="w-full h-auto"
          src="images/illustrations-foreground/ironman-placeholder.svg"
          alt="Iron Man Illustration Placeholder"
        />
      </LazyImage>
      <!-- Arthas -->
      <LazyImage classes="w-1/2 lg:w-1/4 h-auto -mx-8 md:-mx-16">
        <img
          slot="content"
          class="w-full h-auto"
          src="images/illustrations-foreground/arthas.svg"
          alt="Arthas Illustration"
        />
        <img
          slot="placeholder"
          class="w-full h-auto"
          src="images/illustrations-foreground/arthas-placeholder.svg"
          alt="Arthas Illustration Placeholder"
        />
      </LazyImage>
      <!-- Heisenberg -->
      <LazyImage classes="w-1/2 lg:w-1/4 h-auto -mx-8 md:-mx-16">
        <img
          slot="content"
          class="w-full h-auto"
          src="images/illustrations-foreground/heisenberg.svg"
          alt="Heisenberg Illustration"
        />
        <img
          slot="placeholder"
          class="w-full h-auto"
          src="images/illustrations-foreground/heisenberg-placeholder.svg"
          alt="Heisenberg Illustration Placeholder"
        />
      </LazyImage>
    </div>
  </div>
</div>
