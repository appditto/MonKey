<script>
  import SocialIcon from "./SocialIcon.svelte";
  let icons = [
    { name: "Discord", href: "https://chat.banano.cc" },
    { name: "Reddit", href: "https://www.reddit.com/r/banano" },
    { name: "Twitter", href: "https://twitter.com/bananocoin" },
    { name: "Medium", href: "https://medium.com/banano" },
    { name: "Telegram", href: "https://t.me/banano_official" },
    { name: "Facebook", href: "https://www.facebook.com/bananocurrency" },
    { name: "Instagram", href: "https://www.instagram.com/bananocurrency" },
    { name: "GitHub", href: "https://github.com/bananocoin" },
  ];
</script>

<!-- HTML -->
<div class="w-full pt-4 pb-12 relative z-0">
  <!-- Desktop Background Illustration -->
  <img
    class="hidden lg:block w-full h-auto absolute bottom-0"
    src="images/illustrations-background/bg-footer.svg"
    alt="Footer Background Illustration"
  />
  <!-- Mobile Background Illustration -->
  <img
    class="lg:hidden w-full h-auto absolute bottom-0"
    src="images/illustrations-background/bg-footer-mobile.svg"
    alt="Footer Background Illustration"
  />
  <!-- Foreground -->
  <div class="y-container flex flex-col justify-between items-center relative px-2 md:px-4">
    <!-- Social Icons -->
    <div class="flex flex-row flex-wrap justify-center items-center">
      {#each icons as icon}
        <SocialIcon {icon} />
      {/each}
    </div>
    <!-- Footer Text -->
    <div class="flex flex-row mt-6">
      <span>©2020,&nbsp;</span>
      <a
        href="https://banano.cc"
        rel="noopener"
        target="_blank"
        class="font-bold hover:underline focus:underline"
      >
        Banano
      </a>
      <span>.</span>
    </div>
    <!-- Appditto Logo -->
    <div class="w-full flex flex-row justify-center items-center mt-5">
      <p>made by</p>
      <a class="logo-button p-1 ml-1" rel="noopener" target="_blank" href="https://appditto.com">
        <div class="w-28 h-auto">
          <img class="w-full h-auto" src="images/logos/logo-appditto-horizontal.svg" alt="Appditto" />
        </div>
      </a>
    </div>
  </div>
</div>
