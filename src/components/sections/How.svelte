<script>
  import LazyImage from "../LazyImage.svelte";
</script>

<!-- HTML -->
<div
  class="w-full flex flex-col justify-center py-16 md:pb-32 relative md:my-20">
  <!-- Background Illusration Right -->
  <img
    class="hidden md:block w-1/4 lg:w-1/5 h-auto absolute right-0 top-0"
    src="images/illustrations-background/bg-glasses-right.svg"
    alt="Glasses Illustration" />
  <!-- Background Illusration Left -->
  <img
    class="hidden md:block w-1/4 lg:w-1/5 h-auto absolute left-0 bottom-0"
    src="images/illustrations-background/bg-hat-left.svg"
    alt="Hats Illustration" />
  <div class="y-container relative">
    <h2
      class="font-bold text-5xl mx-4 text-center leading-tight tracking-tight
      mt-5">
      How Does It Work?
    </h2>
    <p class="text-center mx-4 md:mx-12 mt-3 text-2xl">
      We take your Banano address, apply some extremely
      <br class="hidden lg:block" />
      advanced science to it, and it turns into MonKey.
    </p>
    <!-- Desktop Illustration -->
    <LazyImage classes="w-full px-12 lg:px-20 mt-8 hidden md:block">
      <img
        slot="content"
        class="w-full h-auto"
        src="images/illustrations-foreground/formula.svg"
        alt="Formula Illustration" />
      <img
        slot="placeholder"
        class="w-full h-auto"
        src="images/illustrations-foreground/formula-placeholder.svg"
        alt="Formula Illustration Placeholder" />
    </LazyImage>
    <!-- Mobile Illustration-->
    <LazyImage classes="w-full max-w-xxs mt-8 md:hidden mx-auto px-4">
      <img
        slot="content"
        class="w-full h-auto"
        src="images/illustrations-foreground/formula-mobile.svg"
        alt="Formula Illustration" />
      <img
        slot="placeholder"
        class="w-full h-auto"
        src="images/illustrations-foreground/formula-mobile-placeholder.svg"
        alt="Formula Illustration Placeholder" />
    </LazyImage>
  </div>
</div>
