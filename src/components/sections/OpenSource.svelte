<script>
  import LazyImage from "../LazyImage.svelte";
</script>

<!-- HTML -->
<div class="w-full flex-col justify-center py-12 relative md:my-20">
  <!-- Background Illusration Left -->
  <img
    class="hidden md:block w-1/4 lg:w-1/5 h-auto absolute left-0 top-0"
    src="images/illustrations-background/bg-code-left.svg"
    alt="Code Illustration" />
  <!-- Background Illusration Right -->
  <img
    class="hidden md:block w-1/4 lg:w-1/5 h-auto absolute right-0 bottom-0"
    src="images/illustrations-background/bg-heart-right.svg"
    alt="Heart Illustration" />
  <div class="y-container flex flex-col items-center relative">
    <h2
      class="font-bold text-5xl mx-4 text-center leading-tight tracking-tight
      mt-5">
      Open Source
    </h2>
    <p class="w-full max-w-2xl text-center md:mx-12 mt-3 text-2xl px-4">
      MonKey is open source. So, you can take our code and build something else
      with it or you can contribute to it.
    </p>
    <div class="flex flex-row justify-center mt-6">
      <a
        href="https://github.com/appditto/monkey"
        target="_blank"
        class="bg-primary btn-primary text-white text-xl font-bold rounded-xl
        border-2 border-black px-8 md:px-12 py-2">
        Visit the Repo
      </a>
    </div>
    <!-- Desktop Illustration -->
    <LazyImage classes="w-full hidden md:block px-12 lg:px-20 mt-12">
      <img
        slot="content"
        class="w-full h-auto"
        src="images/illustrations-foreground/open-source.svg"
        alt="Open Source Illustration" />
      <img
        slot="placeholder"
        class="w-full h-auto"
        src="images/illustrations-foreground/open-source-placeholder.svg"
        alt="Open Source Illustration Placeholder" />
    </LazyImage>
    <!-- Mobile Illustration-->
    <LazyImage classes="w-full md:hidden mx-auto mt-12 px-6">
      <img
        slot="content"
        class="w-full h-auto"
        src="images/illustrations-foreground/open-source-mobile.svg"
        alt="Open Source Illustration" />
      <img
        slot="placeholder"
        class="w-full h-auto"
        src="images/illustrations-foreground/open-source-mobile-placeholder.svg"
        alt="Open Source Illustration Placeholder" />
    </LazyImage>
  </div>
</div>
